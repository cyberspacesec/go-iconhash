package hasher

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"
)

func TestDefaultOptions(t *testing.T) {
	options := DefaultOptions()

	if options == nil {
		t.Fatal("DefaultOptions() returned nil")
	}

	if options.UseUint32 != false {
		t.Errorf("DefaultOptions().UseUint32 = %v, expected false", options.UseUint32)
	}

	if options.RequestTimeout != 10*time.Second {
		t.Errorf("DefaultOptions().RequestTimeout = %v, expected 10s", options.RequestTimeout)
	}

	if options.InsecureSkipVerify != true {
		t.Errorf("DefaultOptions().InsecureSkipVerify = %v, expected true", options.InsecureSkipVerify)
	}

	if options.UserAgent == "" {
		t.Error("DefaultOptions().UserAgent is empty")
	}
}

func TestNew(t *testing.T) {
	// Test with default options
	hasher := New(nil)
	if hasher == nil {
		t.Fatal("New(nil) returned nil")
	}

	// Test with custom options
	customOptions := &HashOptions{
		UseUint32:      true,
		RequestTimeout: 20 * time.Second,
		UserAgent:      "CustomUserAgent",
	}

	hasher = New(customOptions)
	if hasher == nil {
		t.Fatal("New(customOptions) returned nil")
	}

	if hasher.options.UseUint32 != true {
		t.Errorf("hasher.options.UseUint32 = %v, expected true", hasher.options.UseUint32)
	}

	if hasher.options.RequestTimeout != 20*time.Second {
		t.Errorf("hasher.options.RequestTimeout = %v, expected 20s", hasher.options.RequestTimeout)
	}

	if hasher.options.UserAgent != "CustomUserAgent" {
		t.Errorf("hasher.options.UserAgent = %q, expected 'CustomUserAgent'", hasher.options.UserAgent)
	}
}

func TestHashFromURL(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Send a simple favicon byte array
		w.Header().Set("Content-Type", "image/x-icon")
		w.Write([]byte{0, 0, 1, 0, 1, 0, 16, 16})
	}))
	defer server.Close()

	hasher := New(nil)

	// Test successful URL hash
	hash, err := hasher.HashFromURL(server.URL)
	if err != nil {
		t.Errorf("HashFromURL(%q) returned error: %v", server.URL, err)
	}

	if hash == "" {
		t.Errorf("HashFromURL(%q) returned empty hash", server.URL)
	}

	// Test with non-existent URL
	_, err = hasher.HashFromURL("http://non-existent-url.example")
	if err == nil {
		t.Error("HashFromURL with non-existent URL did not return error")
	}
}

func TestHashFromFile(t *testing.T) {
	// Create a temporary test file
	tempFile, err := os.CreateTemp("", "favicon-test-*.ico")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	// Write some test data
	if _, err := tempFile.Write([]byte{0, 0, 1, 0, 1, 0, 16, 16}); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	tempFile.Close()

	hasher := New(nil)

	// Test successful file hash
	hash, err := hasher.HashFromFile(tempFile.Name())
	if err != nil {
		t.Errorf("HashFromFile(%q) returned error: %v", tempFile.Name(), err)
	}

	if hash == "" {
		t.Errorf("HashFromFile(%q) returned empty hash", tempFile.Name())
	}

	// Test with non-existent file
	_, err = hasher.HashFromFile("/non/existent/favicon.ico")
	if err == nil {
		t.Error("HashFromFile with non-existent file did not return error")
	}
}

func TestHashFromBase64(t *testing.T) {
	// Simple base64 encoded data
	base64Data := "AAEAAQAQEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAf"

	hasher := New(nil)

	// Test successful base64 hash
	hash, err := hasher.HashFromBase64(base64Data)
	if err != nil {
		t.Errorf("HashFromBase64() returned error: %v", err)
	}

	if hash == "" {
		t.Error("HashFromBase64() returned empty hash")
	}

	// Test with data URL prefix
	prefixedData := "data:image/vnd.microsoft.icon;base64," + base64Data
	prefixHash, err := hasher.HashFromBase64(prefixedData)
	if err != nil {
		t.Errorf("HashFromBase64() with prefix returned error: %v", err)
	}

	if prefixHash != hash {
		t.Errorf("HashFromBase64() with prefix = %q, expected %q", prefixHash, hash)
	}
}

func TestFormatBase64WithNewlines(t *testing.T) {
	hasher := New(nil)

	// Test with data shorter than 76 chars
	shortData := []byte("short data")
	formatted := hasher.formatBase64WithNewlines(shortData)
	if !strings.HasSuffix(string(formatted), "\n") {
		t.Error("formatBase64WithNewlines() didn't add trailing newline for short data")
	}

	// Test with data longer than 76 chars
	longData := []byte(strings.Repeat("A", 100))
	formatted = hasher.formatBase64WithNewlines(longData)

	// Should have a newline at position 76
	if formatted[76] != '\n' {
		t.Error("formatBase64WithNewlines() didn't add newline at position 76")
	}

	// Should have a final newline
	if formatted[len(formatted)-1] != '\n' {
		t.Error("formatBase64WithNewlines() didn't add trailing newline for long data")
	}
}

func TestCalculateHash(t *testing.T) {
	hasher := New(&HashOptions{UseUint32: false})

	// Test int32 hash
	int32Hash, err := hasher.calculateHash([]byte("test data"))
	if err != nil {
		t.Errorf("calculateHash() returned error: %v", err)
	}

	// Test uint32 hash
	hasher = New(&HashOptions{UseUint32: true})
	uint32Hash, err := hasher.calculateHash([]byte("test data"))
	if err != nil {
		t.Errorf("calculateHash() with uint32 returned error: %v", err)
	}

	// The hash values should be different representations of the same number
	if int32Hash[0] != '-' && int32Hash == uint32Hash {
		t.Errorf("int32Hash (%s) and uint32Hash (%s) should be different for negative numbers",
			int32Hash, uint32Hash)
	}
}
