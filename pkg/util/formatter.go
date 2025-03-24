package util

import "fmt"

// OutputFormat represents the format of the hash output
type OutputFormat int

const (
	// FormatPlain outputs the hash as is
	FormatPlain OutputFormat = iota
	// FormatFofa outputs the hash in Fofa search format
	FormatFofa
	// FormatShodan outputs the hash in Shodan search format
	FormatShodan
)

// FormatHash formats a hash value according to the specified output format
func FormatHash(hash string, format OutputFormat) string {
	switch format {
	case FormatFofa:
		return fmt.Sprintf("icon_hash=\"%s\"", hash)
	case FormatShodan:
		return fmt.Sprintf("http.favicon.hash:%s", hash)
	default:
		return hash
	}
}

// OutputOptions contains configuration for output
type OutputOptions struct {
	Format OutputFormat
	Debug  bool
}

// NewOutputOptions creates OutputOptions with default values
func NewOutputOptions() *OutputOptions {
	return &OutputOptions{
		Format: FormatFofa, // Default to Fofa format like the original
		Debug:  false,
	}
}
