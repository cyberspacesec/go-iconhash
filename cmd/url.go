package cmd

import (
	"fmt"
	"os"

	"github.com/cyberspacesec/go-iconhash/pkg/hasher"
	"github.com/cyberspacesec/go-iconhash/pkg/util"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// NewURLCommand 创建URL命令
func NewURLCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "url [url]",
		Short: "Generate hash from a URL",
		Long: `Generate a favicon hash from a URL.
		
This command will fetch the favicon from the specified URL and calculate its hash.
The hash can be formatted for use with search engines like Fofa or Shodan.

Examples:
  iconhash url https://example.com
  iconhash url -u https://example.com/favicon.ico --shodan
  iconhash url https://example.com --uint32`,
		Run: runURL,
		Args: func(cmd *cobra.Command, args []string) error {
			// If URL is provided as positional arg, set it in the flags
			if len(args) > 0 {
				// If URL already set via flag, don't override
				if URL == "" {
					URL = args[0]
				}
			}

			// Validate we have a URL
			if URL == "" {
				return fmt.Errorf("URL is required. Provide it as an argument or with --url flag")
			}
			return nil
		},
	}

	return cmd
}

// runURL handles the URL command execution
func runURL(cmd *cobra.Command, args []string) {
	// Create a new hasher with the options from root command
	options := &hasher.HashOptions{
		UseUint32:          Uint32Flag,
		RequestTimeout:     Timeout,
		InsecureSkipVerify: SkipVerify,
		UserAgent:          UserAgent,
	}
	h := hasher.New(options)

	// Debug info if enabled
	if Debug {
		fmt.Fprintf(os.Stderr, "🔍 URL: %s\n", URL)
		fmt.Fprintf(os.Stderr, "🔧 Options: uint32=%v, timeout=%ds, skip-verify=%v\n",
			options.UseUint32, options.RequestTimeout.Seconds(), options.InsecureSkipVerify)
		if options.UserAgent != "" {
			fmt.Fprintf(os.Stderr, "🕵️ User-Agent: %s\n", options.UserAgent)
		}
	}

	// Calculate hash
	fmt.Fprintf(os.Stderr, "🌐 Fetching favicon from %s...\n", URL)
	hash, err := h.HashFromURL(URL)
	if err != nil {
		color.Red("❌ Error calculating hash: %v", err)
		os.Exit(1)
	}

	// Determine output format
	var format util.OutputFormat
	if ShodanFormat {
		format = util.FormatShodan
	} else if FofaFormat {
		format = util.FormatFofa
	} else {
		format = util.FormatPlain
	}

	// Format the hash
	formatted := util.FormatHash(hash, format)

	// Print hash with color
	boldGreen := color.New(color.FgGreen, color.Bold)
	boldGreen.Println("✅ Hash calculated successfully!")
	fmt.Println()

	boldCyan := color.New(color.FgCyan, color.Bold)
	boldCyan.Printf("Hash: ")
	fmt.Println(hash)

	if format != util.FormatPlain {
		boldCyan.Printf("Formatted: ")
		fmt.Println(formatted)
	}
}
