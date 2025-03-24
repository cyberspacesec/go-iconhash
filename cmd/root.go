package cmd

import (
	"fmt"
	"time"

	"github.com/cyberspacesec/go-iconhash/pkg/util"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// RootCmd represents the base command
var RootCmd = &cobra.Command{
	Use:     "iconhash [flags] [url or file]",
	Short:   "Icon Hash Calculator - A tool for cybersecurity reconnaissance",
	Version: Version,
	Long: `Icon Hash Calculator - A tool for cybersecurity reconnaissance

Calculate the MMH3 hash of a favicon.ico file for use with search engines like
Fofa or Shodan. This tool can process favicons from URLs, files, or base64 data.

Examples:
  iconhash https://example.com/favicon.ico       # Hash from URL
  iconhash favicon.ico                           # Hash from file
  iconhash -b64 base64file.txt                   # Hash from base64 file
  iconhash server -p 8080                        # Start API server on port 8080`,
	Run: func(cmd *cobra.Command, args []string) {
		// If no args or subcommands specified, show help
		if len(args) == 0 && FilePath == "" && URL == "" && Base64Path == "" {
			cmd.Help()
			return
		}

		// If a positional argument is provided, try to determine if it's a URL or file
		if len(args) > 0 {
			arg := args[0]

			// Check if it looks like a URL (has ://)
			if util.IsURL(arg) {
				URL = arg
			} else {
				// Assume it's a file
				FilePath = arg
			}
		}

		// If URL is set, forward to URL command
		if URL != "" {
			runURL(cmd, args)
			return
		}

		// If FilePath is set, forward to File command
		if FilePath != "" {
			runFile(cmd, args)
			return
		}

		// If Base64Path is set, forward to Base64 command
		if Base64Path != "" {
			runBase64(cmd, args)
			return
		}
	},
}

// PrintLogo prints the ASCII art logo
func PrintLogo() {
	cyan := color.New(color.FgCyan).SprintFunc()
	blue := color.New(color.FgBlue).SprintFunc()

	logo := `
  ██▓ ▄████▄   ▒█████   ███▄    █     ██░ ██  ▄▄▄       ██████  ██░ ██ 
 ▓██▒▒██▀ ▀█  ▒██▒  ██▒ ██ ▀█   █    ▓██░ ██▒▒████▄   ▒██    ▒ ▓██░ ██▒
 ▒██▒▒▓█    ▄ ▒██░  ██▒▓██  ▀█ ██▒   ▒██▀▀██░▒██  ▀█▄ ░ ▓██▄   ▒██▀▀██░
 ░██░▒▓▓▄ ▄██▒▒██   ██░▓██▒  ▐▌██▒   ░▓█ ░██ ░██▄▄▄▄██  ▒   ██▒░▓█ ░██ 
 ░██░▒ ▓███▀ ░░ ████▓▒░▒██░   ▓██░   ░▓█▒░██▓ ▓█   ▓██▒▒██████▒▒░▓█▒░██▓
 ░▓  ░ ░▒ ▒  ░░ ▒░▒░▒░ ░ ▒░   ▒ ▒     ▒ ░░▒░▒ ▒▒   ▓▒█░▒ ▒▓▒ ▒ ░ ▒ ░░▒░▒
  ▒ ░  ░  ▒     ░ ▒ ▒░ ░ ░░   ░ ▒░    ▒ ░▒░ ░  ▒   ▒▒ ░░ ░▒  ░ ░ ▒ ░▒░ ░
  ▒ ░░        ░ ░ ░ ▒     ░   ░ ░     ░  ░░ ░  ░   ▒   ░  ░  ░   ░  ░░ ░
  ░  ░ ░          ░ ░           ░     ░  ░  ░      ░  ░      ░   ░  ░  ░
     ░                                                                   
`
	coloredLogo := cyan(logo)
	fmt.Println(coloredLogo)
	fmt.Printf("%s %s - Version %s\n",
		blue("IconHash Calculator"),
		cyan("by Cyberspace Security"),
		blue(Version))
	fmt.Printf("Build Date: %s | Hash: %s\n\n", BuildDate, BuildHash)
}

// Initialize function to set up all commands and flags
func Initialize() {
	// Add all commands that we've created
	RootCmd.AddCommand(NewURLCommand())
	RootCmd.AddCommand(NewFileCommand())
	RootCmd.AddCommand(NewBase64Command())
	RootCmd.AddCommand(NewServerCommand())

	// Define global flags
	RootCmd.PersistentFlags().BoolVarP(&Debug, "debug", "d", false, "Enable debug output")
	RootCmd.PersistentFlags().BoolVarP(&Uint32Flag, "uint32", "n", false, "Output hash as uint32 instead of int32")
	RootCmd.PersistentFlags().StringVarP(&URL, "url", "u", "", "URL to favicon")
	RootCmd.PersistentFlags().StringVarP(&FilePath, "file", "f", "", "Path to favicon file")
	RootCmd.PersistentFlags().StringVarP(&Base64Path, "b64", "b", "", "Path to file containing base64 encoded favicon")
	RootCmd.PersistentFlags().StringVarP(&UserAgent, "user-agent", "a", "", "User agent for HTTP requests")
	RootCmd.PersistentFlags().BoolVarP(&FofaFormat, "fofa", "o", false, "Format output for Fofa search")
	RootCmd.PersistentFlags().BoolVarP(&ShodanFormat, "shodan", "s", false, "Format output for Shodan search")
	RootCmd.PersistentFlags().BoolVarP(&SkipVerify, "insecure", "k", false, "Skip TLS certificate verification")
	RootCmd.PersistentFlags().DurationVarP(&Timeout, "timeout", "t", 30*time.Second, "HTTP request timeout")
	RootCmd.PersistentFlags().StringVarP(&OutputFormat, "format", "", "text", "Output format (text, json, csv)")
}
