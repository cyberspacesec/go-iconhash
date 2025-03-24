package util

import (
	"fmt"
	"io"
	"os"
)

// Logger provides debugging functionality
type Logger struct {
	enabled bool
	output  io.Writer
}

// NewLogger creates a new Logger with the given debug flag
func NewLogger(debug bool) *Logger {
	return &Logger{
		enabled: debug,
		output:  os.Stdout,
	}
}

// SetOutput changes the output writer
func (l *Logger) SetOutput(w io.Writer) {
	l.output = w
}

// Debugf formats and prints a debug message if debugging is enabled
func (l *Logger) Debugf(format string, args ...interface{}) {
	if !l.enabled {
		return
	}
	msg := fmt.Sprintf(format, args...)
	fmt.Fprintln(l.output, msg)
}

// DebugSection prints a debug message with a section header and footer
func (l *Logger) DebugSection(section string, fn func()) {
	if !l.enabled {
		fn()
		return
	}

	l.Debugf("--------------------------- START %s ---------------------------", section)
	fn()
	l.Debugf("--------------------------- END %s ---------------------------", section)
}

// IsDebugEnabled returns whether debugging is enabled
func (l *Logger) IsDebugEnabled() bool {
	return l.enabled
}

// Enable enables debugging
func (l *Logger) Enable() {
	l.enabled = true
}

// Disable disables debugging
func (l *Logger) Disable() {
	l.enabled = false
}
