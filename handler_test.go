// Tests live in the same package as the code they test.
// Go finds and runs them automatically with "go test ./..."
package main

import (
	"encoding/json"  // to decode the JSON response body
	"net/http"       // for HTTP status code constants (http.StatusOK, etc.)
	"net/http/httptest" // Go's built-in HTTP testing helpers — no live server needed
	"testing"        // Go's built-in test framework
)

// Test functions must start with "Test" and accept *testing.T.
// t is your handle for reporting failures and logging.
func TestRootHandler(t *testing.T) {
	// httptest.NewRequest builds a fake HTTP request.
	// We never touch a real network — it's all in-memory.
	req := httptest.NewRequest(http.MethodGet, "/", nil) // nil = no request body

	// httptest.NewRecorder is a fake ResponseWriter that records
	// whatever the handler writes, so we can inspect it afterwards.
	w := httptest.NewRecorder()

	// Call the handler directly, passing the fake request and recorder.
	rootHandler(w, req)

	// w.Code holds the HTTP status code the handler sent.
	if w.Code != http.StatusOK {
		// t.Fatalf logs the message and stops the test immediately.
		t.Fatalf("expected 200, got %d", w.Code)
	}

	// Decode the response body from JSON into a Go map.
	// The variable type map[string]string means: keys and values are both strings.
	var body map[string]string
	if err := json.NewDecoder(w.Body).Decode(&body); err != nil {
		// & takes the address of body — NewDecoder needs a pointer so it can fill it in.
		t.Fatalf("invalid JSON: %v", err)
	}

	// t.Errorf logs a failure but lets the test continue (unlike Fatalf).
	if body["status"] != "ok" {
		t.Errorf("expected status=ok, got %q", body["status"])
	}
	if body["message"] != "hello" {
		t.Errorf("expected message=hello, got %q", body["message"])
	}
}

func TestHealthHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()
	healthHandler(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	var body map[string]string
	if err := json.NewDecoder(w.Body).Decode(&body); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}
	if body["status"] != "healthy" {
		t.Errorf("expected status=healthy, got %q", body["status"])
	}
}

// TestUnknownRoute tests the 404 behaviour for an unregistered path.
// We use the full mux here (not rootHandler directly) because the ServeMux
// is what decides which handler runs for a given URL.
func TestUnknownRoute(t *testing.T) {
	// setupRoutes() returns the configured *http.ServeMux from main.go.
	mux := setupRoutes()

	req := httptest.NewRequest(http.MethodGet, "/unknown", nil)
	w := httptest.NewRecorder()

	// ServeHTTP runs the mux as if it received a real HTTP request.
	mux.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Fatalf("expected 404, got %d", w.Code)
	}
}
