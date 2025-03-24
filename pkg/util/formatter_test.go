package util

import "testing"

func TestFormatHash(t *testing.T) {
	tests := []struct {
		name     string
		hash     string
		format   OutputFormat
		expected string
	}{
		{"Plain format", "12345", FormatPlain, "12345"},
		{"Fofa format", "12345", FormatFofa, "icon_hash=\"12345\""},
		{"Shodan format", "12345", FormatShodan, "http.favicon.hash:12345"},
		{"Negative number with Plain format", "-12345", FormatPlain, "-12345"},
		{"Negative number with Fofa format", "-12345", FormatFofa, "icon_hash=\"-12345\""},
		{"Negative number with Shodan format", "-12345", FormatShodan, "http.favicon.hash:-12345"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := FormatHash(test.hash, test.format)
			if result != test.expected {
				t.Errorf("FormatHash(%q, %v) = %q, expected %q", test.hash, test.format, result, test.expected)
			}
		})
	}
}

func TestNewOutputOptions(t *testing.T) {
	options := NewOutputOptions()

	if options == nil {
		t.Fatal("NewOutputOptions() returned nil")
	}

	if options.Format != FormatFofa {
		t.Errorf("NewOutputOptions().Format = %v, expected %v", options.Format, FormatFofa)
	}

	if options.Debug != false {
		t.Errorf("NewOutputOptions().Debug = %v, expected false", options.Debug)
	}
}
