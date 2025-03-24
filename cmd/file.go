package cmd

import (
	"fmt"
	"os"

	"github.com/cyberspacesec/go-iconhash/pkg/hasher"
	"github.com/cyberspacesec/go-iconhash/pkg/util"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// NewFileCommand 创建文件命令
func NewFileCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "file [filepath]",
		Short: "Generate hash from a file",
		Long: `Generate a favicon hash from a local file.
		
This command will read the specified favicon file and calculate its hash.
The hash can be formatted for use with search engines like Fofa or Shodan.

Examples:
  iconhash file favicon.ico
  iconhash file -f /path/to/favicon.ico --shodan
  iconhash file icon.png --uint32`,
		Run: runFile,
		Args: func(cmd *cobra.Command, args []string) error {
			// If filepath is provided as positional arg, set it in the flags
			if len(args) > 0 {
				// If filepath already set via flag, don't override
				if FilePath == "" {
					FilePath = args[0]
				}
			}

			// Validate we have a filepath
			if FilePath == "" {
				return fmt.Errorf("filepath is required. Provide it as an argument or with --file flag")
			}

			// Check if file exists
			_, err := os.Stat(FilePath)
			if err != nil {
				return fmt.Errorf("file not found: %s", FilePath)
			}

			return nil
		},
	}

	return cmd
}

// runFile handles the file command execution
func runFile(cmd *cobra.Command, args []string) {
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
		fmt.Fprintf(os.Stderr, "🔍 File: %s\n", FilePath)
		fmt.Fprintf(os.Stderr, "🔧 Options: uint32=%v\n", options.UseUint32)
	}

	// Read file data
	fmt.Fprintf(os.Stderr, "📂 Reading file %s...\n", FilePath)
	fileData, err := os.ReadFile(FilePath)
	if err != nil {
		color.Red("❌ Error reading file: %v", err)
		os.Exit(1)
	}

	// Calculate hash
	fmt.Fprintf(os.Stderr, "🧮 Calculating hash...\n")
	hash, err := h.HashFromBytes(fileData)
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
