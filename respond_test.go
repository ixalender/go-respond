package respond_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ixalender/go-respond"
)

func TestOK(t *testing.T) {
	w := httptest.NewRecorder()
	respond.OK(w)

	if w.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, w.Code)
	}

	if body := w.Body.String(); body != "" {
		t.Errorf("expected empty body, got %q", body)
	}
}

func TestOkay(t *testing.T) {
	w := httptest.NewRecorder()
	testData := struct{ Name string }{Name: "test"}
	respond.Okay(w, testData)

	if w.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, w.Code)
	}

	expectedBody := `{"Name":"test"}` + "\n"
	if body := w.Body.String(); body != expectedBody {
		t.Errorf("expected body %q, got %q", expectedBody, body)
	}
}

func TestSuccess(t *testing.T) {
	w := httptest.NewRecorder()
	testData := struct{ ID int }{ID: 123}
	status := http.StatusCreated
	respond.Respond(w, status, testData)

	if w.Code != status {
		t.Errorf("expected status %d, got %d", status, w.Code)
	}

	expectedBody := `{"ID":123}` + "\n"
	if body := w.Body.String(); body != expectedBody {
		t.Errorf("expected body %q, got %q", expectedBody, body)
	}
}

func TestError(t *testing.T) {
	w := httptest.NewRecorder()
	status := http.StatusNotFound
	message := "Not found"
	respond.Error(w, status, message)

	if w.Code != status {
		t.Errorf("expected status %d, got %d", status, w.Code)
	}

	var resp respond.RestError
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	if resp.Message != message {
		t.Errorf("expected message %q, got %q", message, resp.Message)
	}

	if resp.Status != status {
		t.Errorf("expected status %d, got %d", status, resp.Status)
	}
}

func TestBadRequest(t *testing.T) {
	tests := []struct {
		name     string
		message  string
		expected string
	}{
		{
			name:     "custom message",
			message:  "Invalid input",
			expected: "Invalid input",
		},
		{
			name:     "default message",
			message:  "",
			expected: "Bad Request",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			respond.BadRequest(w, tt.message)

			if w.Code != http.StatusBadRequest {
				t.Errorf("expected status %d, got %d", http.StatusBadRequest, w.Code)
			}

			var resp respond.RestError
			if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
				t.Fatalf("failed to unmarshal response: %v", err)
			}

			if resp.Message != tt.expected {
				t.Errorf("expected message %q, got %q", tt.expected, resp.Message)
			}
		})
	}
}

func TestInternalError(t *testing.T) {
	tests := []struct {
		name     string
		message  string
		expected string
	}{
		{
			name:     "custom message",
			message:  "DB error",
			expected: "DB error",
		},
		{
			name:     "default message",
			message:  "",
			expected: "Internal Server Error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			respond.InternalError(w, tt.message)

			if w.Code != http.StatusInternalServerError {
				t.Errorf("expected status %d, got %d", http.StatusInternalServerError, w.Code)
			}

			var resp respond.RestError
			if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
				t.Fatalf("failed to unmarshal response: %v", err)
			}

			if resp.Message != tt.expected {
				t.Errorf("expected message %q, got %q", tt.expected, resp.Message)
			}
		})
	}
}

func TestNotFound(t *testing.T) {
	tests := []struct {
		name     string
		message  string
		expected string
	}{
		{
			name:     "custom message",
			message:  "User not found",
			expected: "User not found",
		},
		{
			name:     "default message",
			message:  "",
			expected: "Not Found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			respond.NotFound(w, tt.message)

			if w.Code != http.StatusNotFound {
				t.Errorf("expected status %d, got %d", http.StatusNotFound, w.Code)
			}

			var resp respond.RestError
			if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
				t.Fatalf("failed to unmarshal response: %v", err)
			}

			if resp.Message != tt.expected {
				t.Errorf("expected message %q, got %q", tt.expected, resp.Message)
			}
		})
	}
}
