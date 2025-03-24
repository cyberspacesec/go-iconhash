package hasher

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"hash"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/twmb/murmur3"
)

// HashOptions defines options for icon hashing
type HashOptions struct {
	UseUint32          bool
	RequestTimeout     time.Duration
	InsecureSkipVerify bool
	UserAgent          string
}

// DefaultOptions returns a HashOptions with sensible defaults
func DefaultOptions() *HashOptions {
	return &HashOptions{
		UseUint32:          false,
		RequestTimeout:     10 * time.Second,
		InsecureSkipVerify: true,
		UserAgent:          "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_0) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.56 Safari/535.11",
	}
}

// IconHasher provides methods to calculate MMH3 hash of favicons
type IconHasher struct {
	options    *HashOptions
	httpClient *http.Client
}

// New creates a new IconHasher with the given options
func New(options *HashOptions) *IconHasher {
	if options == nil {
		options = DefaultOptions()
	}

	transport := http.DefaultTransport.(*http.Transport).Clone()
	if transport.TLSClientConfig != nil {
		transport.TLSClientConfig.InsecureSkipVerify = options.InsecureSkipVerify
	}

	return &IconHasher{
		options: options,
		httpClient: &http.Client{
			Timeout:   options.RequestTimeout,
			Transport: transport,
		},
	}
}

// HashFromURL downloads and calculates the hash of an icon from a URL
func (h *IconHasher) HashFromURL(url string) (string, error) {
	data, err := h.getContentFromURL(url)
	if err != nil {
		return "", fmt.Errorf("failed to get content from URL: %w", err)
	}

	return h.HashFromBytes(data)
}

// HashFromFile calculates the hash of an icon from a file
func (h *IconHasher) HashFromFile(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}

	return h.HashFromBytes(data)
}

// HashFromBase64 calculates the hash of an icon from base64 encoded data
func (h *IconHasher) HashFromBase64(base64Data string) (string, error) {
	// Strip prefix if exists
	if strings.HasPrefix(base64Data, "data:image/vnd.microsoft.icon;base64,") {
		base64Data = base64Data[37:]
	}

	// Format with newlines for every 76 characters
	formattedBytes := h.formatBase64WithNewlines([]byte(base64Data))

	return h.calculateHash(formattedBytes)
}

// HashFromBytes calculates the hash of an icon from bytes
func (h *IconHasher) HashFromBytes(data []byte) (string, error) {
	encodedBytes := h.standardBase64Encode(data)
	return h.calculateHash(encodedBytes)
}

// getContentFromURL fetches content from a URL
func (h *IconHasher) getContentFromURL(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", h.options.UserAgent)

	resp, err := h.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP request failed with status code: %d", resp.StatusCode)
	}

	return io.ReadAll(resp.Body)
}

// standardBase64Encode encodes bytes to base64 and formats with newlines
func (h *IconHasher) standardBase64Encode(data []byte) []byte {
	encodedStr := base64.StdEncoding.EncodeToString(data)
	return h.formatBase64WithNewlines([]byte(encodedStr))
}

// formatBase64WithNewlines formats base64 data with newlines every 76 characters
func (h *IconHasher) formatBase64WithNewlines(data []byte) []byte {
	var buffer bytes.Buffer

	for i := 0; i < len(data); i++ {
		buffer.WriteByte(data[i])
		if (i+1)%76 == 0 {
			buffer.WriteByte('\n')
		}
	}

	// Add final newline if not already present
	if len(data)%76 != 0 {
		buffer.WriteByte('\n')
	}

	return buffer.Bytes()
}

// calculateHash computes the MMH3 hash
func (h *IconHasher) calculateHash(data []byte) (string, error) {
	var h32 hash.Hash32 = murmur3.New32()
	_, err := h32.Write(data)
	if err != nil {
		return "", fmt.Errorf("failed to calculate hash: %w", err)
	}

	if h.options.UseUint32 {
		return fmt.Sprintf("%d", h32.Sum32()), nil
	}
	return fmt.Sprintf("%d", int32(h32.Sum32())), nil
}
