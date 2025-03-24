package util

import (
	"fmt"

	"github.com/fatih/color"
)

// Result represents a hash calculation result interface
type Result interface {
	Format(cfg interface{}) string
}

// PrintResult prints the result based on the configuration
func PrintResult(result Result, cfg interface{}) {
	// Get formatted result
	formatted := result.Format(cfg)

	// Print the formatted result
	fmt.Println(color.GreenString(formatted))
}
