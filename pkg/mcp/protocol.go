package mcp

import (
	"encoding/json"
	"fmt"
	"time"
)

// Protocol constants
const (
	ProtocolVersion = "1.0"
	ProtocolName    = "Model Context Protocol"
)

// Message represents a message in the MCP
type Message struct {
	Role    string                 `json:"role"`
	Content string                 `json:"content"`
	Name    string                 `json:"name,omitempty"`
	Meta    map[string]interface{} `json:"meta,omitempty"`
}

// Context represents the conversation context
type Context struct {
	Messages []Message              `json:"messages"`
	Meta     map[string]interface{} `json:"meta,omitempty"`
}

// Request represents an MCP request
type Request struct {
	Version  string                 `json:"version"`
	Protocol string                 `json:"protocol"`
	Context  Context                `json:"context"`
	Config   map[string]interface{} `json:"config,omitempty"`
	Meta     map[string]interface{} `json:"meta,omitempty"`
}

// Response represents an MCP response
type Response struct {
	Version  string                 `json:"version"`
	Protocol string                 `json:"protocol"`
	Message  Message                `json:"message"`
	Usage    *Usage                 `json:"usage,omitempty"`
	Meta     map[string]interface{} `json:"meta,omitempty"`
}

// Usage represents the resource usage information
type Usage struct {
	PromptTokens     int       `json:"prompt_tokens,omitempty"`
	CompletionTokens int       `json:"completion_tokens,omitempty"`
	TotalTokens      int       `json:"total_tokens,omitempty"`
	ProcessingTime   float64   `json:"processing_time,omitempty"` // in seconds
	StartedAt        time.Time `json:"started_at,omitempty"`
	CompletedAt      time.Time `json:"completed_at,omitempty"`
}

// NewRequest creates a new MCP request
func NewRequest() *Request {
	return &Request{
		Version:  ProtocolVersion,
		Protocol: ProtocolName,
		Context: Context{
			Messages: []Message{},
			Meta:     make(map[string]interface{}),
		},
		Config: make(map[string]interface{}),
		Meta:   make(map[string]interface{}),
	}
}

// NewResponse creates a new MCP response
func NewResponse() *Response {
	return &Response{
		Version:  ProtocolVersion,
		Protocol: ProtocolName,
		Message: Message{
			Role:    "assistant",
			Content: "",
			Meta:    make(map[string]interface{}),
		},
		Usage: &Usage{
			StartedAt: time.Now(),
		},
		Meta: make(map[string]interface{}),
	}
}

// AddMessage adds a message to the request context
func (r *Request) AddMessage(role, content string) {
	r.Context.Messages = append(r.Context.Messages, Message{
		Role:    role,
		Content: content,
	})
}

// Validate validates the request
func (r *Request) Validate() error {
	if r.Version == "" {
		return fmt.Errorf("version is required")
	}
	if r.Protocol == "" {
		return fmt.Errorf("protocol is required")
	}
	if len(r.Context.Messages) == 0 {
		return fmt.Errorf("at least one message is required")
	}
	return nil
}

// Complete completes the response with usage information
func (r *Response) Complete() {
	r.Usage.CompletedAt = time.Now()
	r.Usage.ProcessingTime = r.Usage.CompletedAt.Sub(r.Usage.StartedAt).Seconds()
	// We're not doing actual token counting here, just an example
	r.Usage.PromptTokens = len(r.Message.Content) / 4
	r.Usage.CompletionTokens = len(r.Message.Content) / 4
	r.Usage.TotalTokens = r.Usage.PromptTokens + r.Usage.CompletionTokens
}

// String returns a string representation of the request or response
func ObjectToString(obj interface{}) string {
	data, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		return fmt.Sprintf("Error marshaling to JSON: %v", err)
	}
	return string(data)
}
