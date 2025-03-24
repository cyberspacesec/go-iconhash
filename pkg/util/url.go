package util

import (
	"strings"
)

// IsURL checks if a string looks like a URL
func IsURL(s string) bool {
	return strings.Contains(s, "://") || strings.HasPrefix(s, "www.")
}
