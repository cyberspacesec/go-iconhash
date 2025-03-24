package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/cyberspacesec/go-iconhash/pkg/api"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// NewServerCommand 创建服务器命令
func NewServerCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "server",
		Short: "Start the HTTP API server",
		Long: `Start the HTTP API server for favicon hash calculation.
	
This command starts a REST API server that can calculate favicon hashes from URLs,
files, or base64 encoded data. The server provides endpoints for health checking,
hash calculation, and supports the Model Context Protocol (MCP).

Examples:
  iconhash server
  iconhash server -p 8080 --host 0.0.0.0
  iconhash server --auth-token secret123 --debug`,
		Run: runServer,
	}

	// 添加服务器特定的标志
	cmd.Flags().StringVarP(&Host, "host", "H", "127.0.0.1", "Host to bind server")
	cmd.Flags().IntVarP(&Port, "port", "p", 8000, "Port to bind server")
	cmd.Flags().StringVar(&AuthToken, "auth-token", "", "Authentication token for API requests")
	cmd.Flags().DurationVar(&ReadTimeout, "read-timeout", 30, "HTTP server read timeout in seconds")
	cmd.Flags().DurationVar(&WriteTimeout, "write-timeout", 30, "HTTP server write timeout in seconds")

	return cmd
}

// runServer handles the server command execution
func runServer(cmd *cobra.Command, args []string) {
	// Create server config from flags
	config := &api.Config{
		Host:               Host,
		Port:               Port,
		ReadTimeout:        ReadTimeout,
		WriteTimeout:       WriteTimeout,
		AuthToken:          AuthToken,
		EnableDebug:        Debug,
		InsecureSkipVerify: SkipVerify,
		RequestTimeout:     Timeout,
	}

	// Create and start the server
	server := api.NewServer(config)

	// Handle graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-stop
		fmt.Println("\n🛑 Shutting down server...")
		os.Exit(0)
	}()

	// Print startup message
	cyan := color.New(color.FgCyan).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()

	fmt.Println("🚀", cyan("Starting IconHash API Server"))
	fmt.Printf("⚙️  %s: %s\n", yellow("Configuration"), fmt.Sprintf("%s:%d", Host, Port))
	fmt.Printf("🔑 %s: %v\n", yellow("Authentication"), AuthToken != "")
	fmt.Printf("🐛 %s: %v\n", yellow("Debug Mode"), Debug)
	fmt.Printf("⏱️  %s: %v\n", yellow("Request Timeout"), Timeout)
	fmt.Printf("🔐 %s: %v\n", yellow("Insecure Skip Verify"), SkipVerify)

	// Construct the base URL for easy access
	scheme := "http"
	baseURL := fmt.Sprintf("%s://%s:%d", scheme, Host, Port)
	if Host == "0.0.0.0" {
		baseURL = fmt.Sprintf("%s://localhost:%d", scheme, Port)
	}

	fmt.Println("\n📋", cyan("API Endpoints:"))
	fmt.Printf("  %s: %s/health\n", yellow("Health Check"), baseURL)
	fmt.Printf("  %s: %s/hash/url?url=...\n", yellow("URL Hash"), baseURL)
	fmt.Printf("  %s: %s/hash/file\n", yellow("File Hash"), baseURL)
	fmt.Printf("  %s: %s/hash/base64\n", yellow("Base64 Hash"), baseURL)
	fmt.Printf("  %s: %s/mcp\n", yellow("Model Context Protocol"), baseURL)

	fmt.Println("\n🔍", cyan("Query Parameters:"))
	fmt.Printf("  %s: uint32=true|false - Use uint32 format\n", yellow("Optional"))
	fmt.Printf("  %s: format=fofa|shodan|plain - Output format\n", yellow("Optional"))

	if AuthToken != "" {
		fmt.Println("\n🔒", cyan("Authentication:"))
		fmt.Printf("  %s: \"Authorization: Bearer %s\"\n", yellow("Header"), AuthToken)
		fmt.Printf("  %s: \"?token=%s\"\n", yellow("Query"), AuthToken)
	}

	fmt.Println("\n📢", cyan("Press Ctrl+C to stop the server"))
	fmt.Println(yellow("--------------------------------------------------"))

	// Start the server
	err := server.Start()
	if err != nil {
		color.Red("❌ Server error: %v", err)
		os.Exit(1)
	}
}
