package mcp

import (
	"encoding/base64"
	"fmt"
	"net/url"
	"regexp"
	"strings"

	"github.com/cyberspacesec/go-iconhash/pkg/hasher"
	"github.com/cyberspacesec/go-iconhash/pkg/util"
)

// Handler processes MCP requests and generates responses
type Handler struct {
	iconHasher *hasher.IconHasher
	logger     *util.Logger
	options    *hasher.HashOptions
	debug      bool
}

// NewHandler creates a new MCP handler
func NewHandler(debug bool) *Handler {
	options := &hasher.HashOptions{
		UseUint32:          false,
		RequestTimeout:     hasher.DefaultOptions().RequestTimeout,
		InsecureSkipVerify: true,
		UserAgent:          hasher.DefaultOptions().UserAgent,
	}

	return &Handler{
		iconHasher: hasher.New(options),
		logger:     util.NewLogger(debug),
		options:    options,
		debug:      debug,
	}
}

// Process processes an MCP request and returns a response
func (h *Handler) Process(req *Request) (*Response, error) {
	// Validate the request
	if err := req.Validate(); err != nil {
		return nil, err
	}

	// Create a response
	resp := NewResponse()

	// Extract the last user message
	var lastUserMessage string
	for i := len(req.Context.Messages) - 1; i >= 0; i-- {
		msg := req.Context.Messages[i]
		if msg.Role == "user" {
			lastUserMessage = msg.Content
			break
		}
	}

	if lastUserMessage == "" {
		resp.Message.Content = "Please provide a URL, file path, or base64 data to calculate the favicon hash."
		resp.Complete()
		return resp, nil
	}

	// Debug logging
	if h.debug {
		h.logger.Debugf("Processing MCP message: %s", lastUserMessage)
	}

	// Process the message and generate a response
	result, err := h.processMessage(lastUserMessage)
	if err != nil {
		resp.Message.Content = fmt.Sprintf("Error: %v", err)
		resp.Message.Meta["error"] = err.Error()
		resp.Complete()
		return resp, nil
	}

	resp.Message.Content = result
	resp.Complete()
	return resp, nil
}

// processMessage processes a user message and returns a result
func (h *Handler) processMessage(message string) (string, error) {
	// Check if the message contains a URL
	urlPattern := regexp.MustCompile(`https?://[^\s]+`)
	if urlPattern.MatchString(message) {
		urls := urlPattern.FindAllString(message, -1)
		if h.debug {
			h.logger.Debugf("Found URL in message: %s", urls[0])
		}
		return h.processURL(urls[0])
	}

	// Check if the message contains base64 data
	base64Pattern := regexp.MustCompile(`(?:data:[^;]+;base64,)?([A-Za-z0-9+/=]+)`)
	if base64Pattern.MatchString(message) {
		matches := base64Pattern.FindStringSubmatch(message)
		if len(matches) > 1 {
			// Check if it's a valid base64 string
			if _, err := base64.StdEncoding.DecodeString(matches[1]); err == nil {
				if h.debug {
					h.logger.Debugf("Found base64 data in message (length: %d)", len(matches[1]))
				}
				return h.processBase64(matches[1])
			}
		}
	}

	// Assume the message contains commands or requests about the tool
	if strings.Contains(strings.ToLower(message), "help") ||
		strings.Contains(strings.ToLower(message), "how to") ||
		strings.Contains(strings.ToLower(message), "example") {
		if h.debug {
			h.logger.Debugf("Detected help request, returning help text")
		}
		return h.getHelpText(), nil
	}

	return "", fmt.Errorf("I couldn't detect a URL, valid base64 data, or a help request in your message. Please provide a favicon URL or base64 data to calculate the hash.")
}

// processURL processes a URL and returns the hash
func (h *Handler) processURL(urlStr string) (string, error) {
	// Validate URL
	_, err := url.ParseRequestURI(urlStr)
	if err != nil {
		return "", fmt.Errorf("invalid URL: %v", err)
	}

	if h.debug {
		h.logger.Debugf("Processing URL: %s", urlStr)
	}

	// Calculate hash
	hash, err := h.iconHasher.HashFromURL(urlStr)
	if err != nil {
		return "", fmt.Errorf("error calculating hash: %v", err)
	}

	if h.debug {
		h.logger.Debugf("Calculated hash: %s", hash)
	}

	// Format the result
	result := fmt.Sprintf("Favicon Hash for %s:\n\n", urlStr)
	result += fmt.Sprintf("Plain hash: %s\n", hash)
	result += fmt.Sprintf("Fofa format: %s\n", util.FormatHash(hash, util.FormatFofa))
	result += fmt.Sprintf("Shodan format: %s\n", util.FormatHash(hash, util.FormatShodan))

	return result, nil
}

// processBase64 processes base64 data and returns the hash
func (h *Handler) processBase64(data string) (string, error) {
	if h.debug {
		h.logger.Debugf("Processing base64 data (length: %d)", len(data))
	}

	// Calculate hash
	hash, err := h.iconHasher.HashFromBase64(data)
	if err != nil {
		return "", fmt.Errorf("error calculating hash: %v", err)
	}

	if h.debug {
		h.logger.Debugf("Calculated hash: %s", hash)
	}

	// Format the result
	result := "Favicon Hash for provided base64 data:\n\n"
	result += fmt.Sprintf("Plain hash: %s\n", hash)
	result += fmt.Sprintf("Fofa format: %s\n", util.FormatHash(hash, util.FormatFofa))
	result += fmt.Sprintf("Shodan format: %s\n", util.FormatHash(hash, util.FormatShodan))

	return result, nil
}

// getHelpText returns help text for the MCP
func (h *Handler) getHelpText() string {
	return `# IconHash - Favicon Hash Calculator

This tool calculates the MMH3 hash of favicons for use in cybersecurity reconnaissance.

## Examples:

1. Calculate hash from a URL:
   "Calculate the hash for https://example.com/favicon.ico"

2. Calculate hash from base64 data:
   "Calculate the hash for this base64: AAABAAEAEBAAAAEAIABoBAAAFgAAA..."

## How to use the results:

The hash can be used in search engines like:
- Fofa: icon_hash="123456789" 
- Shodan: http.favicon.hash:123456789

For more information, visit: https://github.com/cyberspacesec/go-iconhash
`
}
