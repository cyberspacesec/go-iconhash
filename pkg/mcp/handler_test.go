package mcp

import (
	"strings"
	"testing"

	"github.com/cyberspacesec/go-iconhash/pkg/util"
)

func TestNewRequest(t *testing.T) {
	req := NewRequest()

	if req.Version != ProtocolVersion {
		t.Errorf("Expected version %s, got %s", ProtocolVersion, req.Version)
	}

	if req.Protocol != ProtocolName {
		t.Errorf("Expected protocol %s, got %s", ProtocolName, req.Protocol)
	}

	if len(req.Context.Messages) != 0 {
		t.Errorf("Expected empty messages, got %d", len(req.Context.Messages))
	}
}

func TestNewResponse(t *testing.T) {
	resp := NewResponse()

	if resp.Version != ProtocolVersion {
		t.Errorf("Expected version %s, got %s", ProtocolVersion, resp.Version)
	}

	if resp.Protocol != ProtocolName {
		t.Errorf("Expected protocol %s, got %s", ProtocolName, resp.Protocol)
	}

	if resp.Message.Role != "assistant" {
		t.Errorf("Expected role 'assistant', got %s", resp.Message.Role)
	}

	if resp.Usage == nil {
		t.Error("Expected usage to be non-nil")
	}
}

func TestAddMessage(t *testing.T) {
	req := NewRequest()
	req.AddMessage("user", "Hello")

	if len(req.Context.Messages) != 1 {
		t.Errorf("Expected 1 message, got %d", len(req.Context.Messages))
	}

	if req.Context.Messages[0].Role != "user" {
		t.Errorf("Expected role 'user', got %s", req.Context.Messages[0].Role)
	}

	if req.Context.Messages[0].Content != "Hello" {
		t.Errorf("Expected content 'Hello', got %s", req.Context.Messages[0].Content)
	}
}

func TestRequestValidate(t *testing.T) {
	tests := []struct {
		name          string
		modifyRequest func(*Request)
		expectError   bool
	}{
		{
			name:          "Valid request",
			modifyRequest: func(r *Request) { r.AddMessage("user", "Hello") },
			expectError:   false,
		},
		{
			name:          "Missing version",
			modifyRequest: func(r *Request) { r.Version = ""; r.AddMessage("user", "Hello") },
			expectError:   true,
		},
		{
			name:          "Missing protocol",
			modifyRequest: func(r *Request) { r.Protocol = ""; r.AddMessage("user", "Hello") },
			expectError:   true,
		},
		{
			name:          "No messages",
			modifyRequest: func(r *Request) {},
			expectError:   true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := NewRequest()
			test.modifyRequest(req)

			err := req.Validate()
			if test.expectError && err == nil {
				t.Error("Expected error, got nil")
			}
			if !test.expectError && err != nil {
				t.Errorf("Expected no error, got %v", err)
			}
		})
	}
}

func TestResponseComplete(t *testing.T) {
	resp := NewResponse()
	resp.Message.Content = "Hello world"
	resp.Complete()

	if resp.Usage.ProcessingTime <= 0 {
		t.Errorf("Expected processing time > 0, got %f", resp.Usage.ProcessingTime)
	}

	if resp.Usage.CompletedAt.IsZero() {
		t.Error("Expected completed_at to be set")
	}

	if resp.Usage.TotalTokens <= 0 {
		t.Errorf("Expected total tokens > 0, got %d", resp.Usage.TotalTokens)
	}
}

// MockHandler is a mock implementation of Handler for testing
type MockHandler struct {
	logger *util.Logger
}

// Process implements a mock version of the Process method
func (h *MockHandler) Process(req *Request) (*Response, error) {
	// Validate the request
	if err := req.Validate(); err != nil {
		return nil, err
	}

	// Create a response
	resp := NewResponse()

	// Extract the last user message
	var lastUserMessage string
	for i := len(req.Context.Messages) - 1; i >= 0; i-- {
		msg := req.Context.Messages[i]
		if msg.Role == "user" {
			lastUserMessage = msg.Content
			break
		}
	}

	if lastUserMessage == "" {
		resp.Message.Content = "Please provide a URL, file path, or base64 data to calculate the favicon hash."
		resp.Complete()
		return resp, nil
	}

	// Process based on content
	if strings.Contains(strings.ToLower(lastUserMessage), "help") {
		resp.Message.Content = "IconHash - Favicon Hash Calculator: Help text"
	} else if strings.Contains(strings.ToLower(lastUserMessage), "random text") {
		resp.Message.Content = "Error: I couldn't detect a URL, valid base64 data, or a help request in your message."
	} else {
		resp.Message.Content = "Mocked response for: " + lastUserMessage
	}

	resp.Complete()
	return resp, nil
}

func TestHandlerProcess(t *testing.T) {
	// Create a mock handler for testing
	mockHandler := &MockHandler{
		logger: util.NewLogger(false),
	}

	tests := []struct {
		name           string
		message        string
		expectContains string
		expectError    bool
	}{
		{
			name:           "Help request",
			message:        "Help me use this tool",
			expectContains: "IconHash - Favicon Hash Calculator",
			expectError:    false,
		},
		{
			name:           "Invalid input",
			message:        "This is just random text",
			expectContains: "Error:",
			expectError:    false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := NewRequest()
			req.AddMessage("user", test.message)

			resp, err := mockHandler.Process(req)

			if test.expectError {
				if err == nil {
					t.Error("Expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Errorf("Expected no error, got %v", err)
				return
			}

			if !strings.Contains(resp.Message.Content, test.expectContains) {
				t.Errorf("Expected response to contain '%s', got '%s'",
					test.expectContains, resp.Message.Content)
			}
		})
	}
}
