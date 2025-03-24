package api

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// ServerCommand returns a cobra command for running the API server
func ServerCommand() *cobra.Command {
	var (
		host            string
		port            int
		authToken       string
		debug           bool
		insecureSkipTLS bool
		timeout         int
	)

	cmd := &cobra.Command{
		Use:   "server",
		Short: "Start the HTTP API server",
		Long: `Start the HTTP API server for iconhash

The server provides HTTP endpoints for calculating favicon hashes:
- /hash/url - Hash a favicon from a URL
- /hash/file - Hash a favicon from an uploaded file
- /hash/base64 - Hash a favicon from base64 encoded data
- /health - Server health check
- /mcp - Model Context Protocol endpoint

Authentication can be enabled by setting an auth token. When set,
all requests must include the token in the Authorization header
or as a 'token' query parameter.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			// Create server config
			config := &Config{
				Host:               host,
				Port:               port,
				ReadTimeout:        30 * time.Second,
				WriteTimeout:       30 * time.Second,
				AuthToken:          authToken,
				EnableDebug:        debug,
				InsecureSkipVerify: insecureSkipTLS,
				RequestTimeout:     time.Duration(timeout) * time.Second,
			}

			// Create and start the server
			server := NewServer(config)

			// Handle graceful shutdown
			c := make(chan os.Signal, 1)
			signal.Notify(c, os.Interrupt, syscall.SIGTERM)

			// Start server in a goroutine
			errChan := make(chan error, 1)
			go func() {
				// 使用彩色输出来显示服务器信息
				boldGreen := color.New(color.FgGreen, color.Bold).SprintFunc()
				boldCyan := color.New(color.FgCyan, color.Bold).SprintFunc()

				fmt.Printf("\n%s\n", boldGreen("🚀 Starting IconHash API Server"))
				fmt.Printf("%s %s:%d\n", boldCyan("Server listening on:"), host, port)

				if authToken != "" {
					fmt.Println(color.YellowString("Authentication is enabled"))
				}

				// 打印API信息
				fmt.Printf("\n%s\n", boldGreen("Available API Endpoints:"))
				fmt.Printf("  %s %s\n", boldCyan("GET  /health"), "- Server health check")
				fmt.Printf("  %s %s\n", boldCyan("GET  /hash/url"), "- Hash from URL")
				fmt.Printf("  %s %s\n", boldCyan("POST /hash/file"), "- Hash from file upload")
				fmt.Printf("  %s %s\n", boldCyan("POST /hash/base64"), "- Hash from base64 data")
				fmt.Printf("  %s %s\n", boldCyan("POST /mcp"), "- Model Context Protocol endpoint")

				fmt.Printf("\n%s\n", boldGreen("Press Ctrl+C to stop the server"))

				errChan <- server.Start()
			}()

			// Wait for interrupt or error
			select {
			case err := <-errChan:
				return err
			case <-c:
				fmt.Println(color.GreenString("\n👋 Shutting down server..."))
				time.Sleep(250 * time.Millisecond) // Give time for connections to close
				return nil
			}
		},
	}

	// Add flags
	flags := cmd.Flags()
	flags.StringVarP(&host, "host", "H", "127.0.0.1", "Host address to bind to")
	flags.IntVarP(&port, "port", "p", 8080, "Port to listen on")
	flags.StringVarP(&authToken, "auth-token", "a", "", "Authentication token (empty for no auth)")
	flags.BoolVarP(&debug, "debug", "d", false, "Enable debug logging")
	flags.BoolVarP(&insecureSkipTLS, "insecure", "k", true, "Skip TLS verification for outbound requests")
	flags.IntVarP(&timeout, "timeout", "t", 10, "Timeout for outbound requests in seconds")

	return cmd
}

// GetServerInfo returns formatted information about the API server and endpoints
func GetServerInfo(host string, port int, authEnabled bool) string {
	scheme := "http"
	addr := host
	if addr == "0.0.0.0" {
		addr = "localhost"
	}

	baseURL := fmt.Sprintf("%s://%s:%d", scheme, addr, port)

	info := fmt.Sprintf(`
API Server running at %s

Endpoints:
  %s/health                  - Health check (GET)
  %s/hash/url?url=<url>      - Hash from URL (GET/POST)
  %s/hash/file               - Hash from file upload (POST)
  %s/hash/base64             - Hash from base64 data (POST)
  %s/mcp                     - Model Context Protocol (POST)

Parameters:
  format=plain|fofa|shodan   - Output format (default: fofa)
  uint32=true|false          - Use uint32 format (default: false)
`, baseURL, baseURL, baseURL, baseURL, baseURL, baseURL)

	if authEnabled {
		info += `
Authentication:
  Add token in "Authorization: Bearer <token>" header
  OR
  Add "?token=<token>" to the URL
`
	}

	info += `
Example curl commands:
  curl -X GET "${baseURL}/hash/url?url=https://example.com/favicon.ico"
  curl -X POST -F "file=@favicon.ico" ${baseURL}/hash/file
  curl -X POST -d "data=$(base64 -i favicon.ico)" ${baseURL}/hash/base64
  curl -X POST -H "Content-Type: application/json" -d '{"version":"1.0","protocol":"Model Context Protocol","context":{"messages":[{"role":"user","content":"Calculate the hash for https://example.com/favicon.ico"}]}}' ${baseURL}/mcp
`

	return info
}
