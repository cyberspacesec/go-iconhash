package util

import (
	"bytes"
	"strings"
	"testing"
)

func TestNewLogger(t *testing.T) {
	tests := []struct {
		name     string
		debug    bool
		expected bool
	}{
		{"Debug enabled", true, true},
		{"Debug disabled", false, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			logger := NewLogger(test.debug)
			if logger.IsDebugEnabled() != test.expected {
				t.Errorf("NewLogger(%v).IsDebugEnabled() = %v, expected %v",
					test.debug, logger.IsDebugEnabled(), test.expected)
			}
		})
	}
}

func TestDebugf(t *testing.T) {
	var buf bytes.Buffer
	logger := NewLogger(true)
	logger.SetOutput(&buf)

	// Test with debug enabled
	logger.Debugf("Test message: %s", "hello")
	if !strings.Contains(buf.String(), "Test message: hello") {
		t.Errorf("Debug message not found in output: %q", buf.String())
	}

	// Test with debug disabled
	buf.Reset()
	logger.Disable()
	logger.Debugf("This should not appear")
	if buf.Len() > 0 {
		t.Errorf("Expected no output with debug disabled, got: %q", buf.String())
	}
}

func TestDebugSection(t *testing.T) {
	var buf bytes.Buffer
	logger := NewLogger(true)
	logger.SetOutput(&buf)

	// Test with debug enabled
	called := false
	logger.DebugSection("TEST SECTION", func() {
		called = true
		logger.Debugf("Inside section")
	})

	if !called {
		t.Error("Function was not called in DebugSection")
	}

	output := buf.String()
	if !strings.Contains(output, "START TEST SECTION") {
		t.Errorf("Section start message not found in output: %q", output)
	}
	if !strings.Contains(output, "Inside section") {
		t.Errorf("Function debug message not found in output: %q", output)
	}
	if !strings.Contains(output, "END TEST SECTION") {
		t.Errorf("Section end message not found in output: %q", output)
	}

	// Test with debug disabled
	buf.Reset()
	logger.Disable()
	called = false
	logger.DebugSection("INVISIBLE SECTION", func() {
		called = true
		logger.Debugf("This should not appear")
	})

	if !called {
		t.Error("Function was not called in DebugSection with debug disabled")
	}
	if buf.Len() > 0 {
		t.Errorf("Expected no output with debug disabled, got: %q", buf.String())
	}
}

func TestIsDebugEnabled(t *testing.T) {
	logger := NewLogger(false)
	if logger.IsDebugEnabled() {
		t.Error("IsDebugEnabled() = true, expected false")
	}

	logger.Enable()
	if !logger.IsDebugEnabled() {
		t.Error("IsDebugEnabled() = false after Enable(), expected true")
	}

	logger.Disable()
	if logger.IsDebugEnabled() {
		t.Error("IsDebugEnabled() = true after Disable(), expected false")
	}
}
