package api

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/cyberspacesec/go-iconhash/pkg/hasher"
	"github.com/cyberspacesec/go-iconhash/pkg/mcp"
	"github.com/cyberspacesec/go-iconhash/pkg/util"
)

// Server represents the HTTP API server
type Server struct {
	config     *Config
	iconHasher *hasher.IconHasher
	logger     *util.Logger
	mcpHandler *mcp.Handler
	debug      bool
}

// Config holds the server configuration
type Config struct {
	Host               string
	Port               int
	ReadTimeout        time.Duration
	WriteTimeout       time.Duration
	AuthToken          string
	EnableDebug        bool
	InsecureSkipVerify bool
	RequestTimeout     time.Duration
}

// DefaultConfig returns a default server configuration
func DefaultConfig() *Config {
	return &Config{
		Host:               "127.0.0.1",
		Port:               8080,
		ReadTimeout:        30 * time.Second,
		WriteTimeout:       30 * time.Second,
		AuthToken:          "",
		EnableDebug:        false,
		InsecureSkipVerify: true,
		RequestTimeout:     10 * time.Second,
	}
}

// NewServer creates a new API server
func NewServer(config *Config) *Server {
	if config == nil {
		config = DefaultConfig()
	}

	// Create options for hasher
	options := &hasher.HashOptions{
		UseUint32:          false,
		RequestTimeout:     config.RequestTimeout,
		InsecureSkipVerify: config.InsecureSkipVerify,
		UserAgent:          "IconHash API Server",
	}

	// Create the server
	logger := util.NewLogger(config.EnableDebug)

	// Create standard icon hasher
	h := hasher.New(options)

	return &Server{
		config:     config,
		iconHasher: h,
		logger:     logger,
		mcpHandler: mcp.NewHandler(config.EnableDebug),
		debug:      config.EnableDebug,
	}
}

// Start starts the HTTP server
func (s *Server) Start() error {
	// Create router
	mux := http.NewServeMux()

	// Setup routes
	mux.HandleFunc("/health", s.handleHealth)
	mux.HandleFunc("/hash/url", s.handleHashURL)
	mux.HandleFunc("/hash/file", s.handleHashFile)
	mux.HandleFunc("/hash/base64", s.handleHashBase64)
	mux.HandleFunc("/mcp", s.handleMCP)

	// Wrap with auth middleware if token is set
	var handler http.Handler = mux
	if s.config.AuthToken != "" {
		handler = s.authMiddleware(mux)
	}

	// Create server
	addr := s.config.Host + ":" + strconv.Itoa(s.config.Port)
	server := &http.Server{
		Addr:         addr,
		Handler:      handler,
		ReadTimeout:  s.config.ReadTimeout,
		WriteTimeout: s.config.WriteTimeout,
	}

	// Log startup
	if s.debug {
		s.logger.Debugf("Starting server on %s", addr)
		s.logger.Debugf("Auth token: %v", s.config.AuthToken != "")
		s.logger.Debugf("Debug enabled: %v", s.config.EnableDebug)
	}

	// Start server
	return server.ListenAndServe()
}

// authMiddleware adds authentication to routes
func (s *Server) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Skip auth for health endpoint
		if r.URL.Path == "/health" {
			next.ServeHTTP(w, r)
			return
		}

		// Check for token in header
		authHeader := r.Header.Get("Authorization")
		if authHeader != "" {
			// Check if it's a Bearer token
			if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
				token := authHeader[7:]
				if token == s.config.AuthToken {
					next.ServeHTTP(w, r)
					return
				}
			}
		}

		// Check for token in query parameter
		token := r.URL.Query().Get("token")
		if token == s.config.AuthToken {
			next.ServeHTTP(w, r)
			return
		}

		// Unauthorized
		sendErrorResponse(w, "Unauthorized: Invalid or missing authentication token", http.StatusUnauthorized)
	})
}

// handleHealth handles the health check endpoint
func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		sendErrorResponse(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "ok",
		"version": "v1",
		"time":    time.Now().Format(time.RFC3339),
	})
}

// handleMCP handles the Model Context Protocol endpoint
func (s *Server) handleMCP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		sendErrorResponse(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Read request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		sendErrorResponse(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	// Parse MCP request
	var req mcp.Request
	if err := json.Unmarshal(body, &req); err != nil {
		sendErrorResponse(w, "Invalid MCP request: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Process the request
	resp, err := s.mcpHandler.Process(&req)
	if err != nil {
		sendErrorResponse(w, "Error processing MCP request: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Serialize response
	respData, err := json.Marshal(resp)
	if err != nil {
		sendErrorResponse(w, "Error serializing response", http.StatusInternalServerError)
		return
	}

	// Send response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respData)
}

// HashResponse is the response format for hash endpoints
type HashResponse struct {
	Hash      string `json:"hash"`
	Format    string `json:"format,omitempty"`
	Formatted string `json:"formatted,omitempty"`
	Error     string `json:"error,omitempty"`
}

// handleHashURL handles the hash from URL endpoint
func (s *Server) handleHashURL(w http.ResponseWriter, r *http.Request) {
	// Set content type
	w.Header().Set("Content-Type", "application/json")

	// Get URL from query or form
	var urlStr string
	if r.Method == http.MethodGet {
		urlStr = r.URL.Query().Get("url")
	} else if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			sendErrorResponse(w, "Error parsing form", http.StatusBadRequest)
			return
		}
		urlStr = r.FormValue("url")
	} else {
		sendErrorResponse(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Validate URL
	if urlStr == "" {
		sendErrorResponse(w, "URL parameter is required", http.StatusBadRequest)
		return
	}

	// Check for format parameter
	formatStr := r.URL.Query().Get("format")
	format := parseFormatParam(formatStr)

	// Update hasher options if uint32 is specified
	uint32Param := r.URL.Query().Get("uint32")
	useUint32 := uint32Param == "true" || uint32Param == "1"
	s.updateHasherOptions(useUint32)

	// Debug output
	if s.debug {
		s.logger.Debugf("URL hash request: %s", urlStr)
		s.logger.Debugf("Format: %s", getFormatName(format))
		s.logger.Debugf("UseUint32: %v", useUint32)
	}

	// Calculate hash
	hash, err := s.iconHasher.HashFromURL(urlStr)
	if err != nil {
		sendErrorResponse(w, "Error calculating hash: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Format hash based on requested format
	formatted := util.FormatHash(hash, format)
	sendHashResponse(w, hash, getFormatName(format), formatted)
}

// handleHashFile handles the hash from file upload endpoint
func (s *Server) handleHashFile(w http.ResponseWriter, r *http.Request) {
	// Set content type
	w.Header().Set("Content-Type", "application/json")

	// Validate method
	if r.Method != http.MethodPost {
		sendErrorResponse(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse multipart form
	err := r.ParseMultipartForm(10 << 20) // Max 10MB
	if err != nil {
		sendErrorResponse(w, "Error parsing multipart form", http.StatusBadRequest)
		return
	}

	// Get file from form
	file, _, err := r.FormFile("file")
	if err != nil {
		sendErrorResponse(w, "Error getting file from form", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Read file
	fileData, err := io.ReadAll(file)
	if err != nil {
		sendErrorResponse(w, "Error reading file", http.StatusInternalServerError)
		return
	}

	// Check for format parameter
	formatStr := r.URL.Query().Get("format")
	format := parseFormatParam(formatStr)

	// Update hasher options if uint32 is specified
	uint32Param := r.URL.Query().Get("uint32")
	useUint32 := uint32Param == "true" || uint32Param == "1"
	s.updateHasherOptions(useUint32)

	// Debug output
	if s.debug {
		s.logger.Debugf("File hash request: %d bytes", len(fileData))
		s.logger.Debugf("Format: %s", getFormatName(format))
		s.logger.Debugf("UseUint32: %v", useUint32)
	}

	// Calculate hash
	hash, err := s.iconHasher.HashFromBytes(fileData)
	if err != nil {
		sendErrorResponse(w, "Error calculating hash: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Format hash based on requested format
	formatted := util.FormatHash(hash, format)
	sendHashResponse(w, hash, getFormatName(format), formatted)
}

// handleHashBase64 handles the hash from base64 endpoint
func (s *Server) handleHashBase64(w http.ResponseWriter, r *http.Request) {
	// Set content type
	w.Header().Set("Content-Type", "application/json")

	// Validate method
	if r.Method != http.MethodPost {
		sendErrorResponse(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse form
	if err := r.ParseForm(); err != nil {
		sendErrorResponse(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	// Get base64 data from form
	base64Data := r.FormValue("data")
	if base64Data == "" {
		sendErrorResponse(w, "Base64 data is required", http.StatusBadRequest)
		return
	}

	// Check for format parameter
	formatStr := r.URL.Query().Get("format")
	format := parseFormatParam(formatStr)

	// Update hasher options if uint32 is specified
	uint32Param := r.URL.Query().Get("uint32")
	useUint32 := uint32Param == "true" || uint32Param == "1"
	s.updateHasherOptions(useUint32)

	// Debug output
	if s.debug {
		s.logger.Debugf("Base64 hash request: %d bytes", len(base64Data))
		s.logger.Debugf("Format: %s", getFormatName(format))
		s.logger.Debugf("UseUint32: %v", useUint32)
	}

	// Calculate hash
	hash, err := s.iconHasher.HashFromBase64(base64Data)
	if err != nil {
		sendErrorResponse(w, "Error calculating hash: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Format hash based on requested format
	formatted := util.FormatHash(hash, format)
	sendHashResponse(w, hash, getFormatName(format), formatted)
}

// updateHasherOptions updates the hash options
func (s *Server) updateHasherOptions(useUint32 bool) {
	// Apply the standard options
	options := &hasher.HashOptions{
		UseUint32:          useUint32,
		RequestTimeout:     s.config.RequestTimeout,
		InsecureSkipVerify: s.config.InsecureSkipVerify,
		UserAgent:          "IconHash API Server",
	}

	// Create new hasher with updated options
	s.iconHasher = hasher.New(options)
}

// sendHashResponse sends a hash response
func sendHashResponse(w http.ResponseWriter, hash, formatName, formatted string) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(HashResponse{
		Hash:      hash,
		Format:    formatName,
		Formatted: formatted,
	})
}

// sendErrorResponse sends an error response
func sendErrorResponse(w http.ResponseWriter, errMessage string, statusCode int) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(HashResponse{
		Error: errMessage,
	})
}

// parseFormatParam parses the format parameter
func parseFormatParam(formatStr string) util.OutputFormat {
	switch formatStr {
	case "plain":
		return util.FormatPlain
	case "fofa":
		return util.FormatFofa
	case "shodan":
		return util.FormatShodan
	default:
		// Default to fofa
		return util.FormatFofa
	}
}

// getFormatName returns the name of the format
func getFormatName(format util.OutputFormat) string {
	switch format {
	case util.FormatPlain:
		return "plain"
	case util.FormatFofa:
		return "fofa"
	case util.FormatShodan:
		return "shodan"
	default:
		return "unknown"
	}
}
