package main

import (
	"encoding/json" // for encoding Go values as JSON
	"net/http"      // for HTTP request/response types
)

// rootHandler handles GET / requests.
// Go HTTP handlers always have this exact signature:
//   - http.ResponseWriter: lets you write the response (headers + body)
//   - *http.Request: contains everything about the incoming request
func rootHandler(w http.ResponseWriter, r *http.Request) {
	// ServeMux routes "/" as a catch-all, so we must manually reject
	// any path that isn't exactly "/".
	// r.URL.Path is the path part of the request URL (e.g. "/foo").
	if r.URL.Path != "/" {
		// http.NotFound writes a 404 response with a plain-text body.
		http.NotFound(w, r)
		// "return" exits the function early so nothing else runs.
		return
	}

	// map[string]string is a Go map (like a dictionary/object in other languages).
	// We build the response body inline and pass it straight to writeJSON.
	writeJSON(w, http.StatusOK, map[string]string{
		"status":  "ok",
		"message": "hello",
	})
}

// healthHandler handles GET /health requests.
// It always returns 200 with a simple JSON body — useful for load balancers
// and uptime monitors to check that the server is alive.
func healthHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{
		"status": "healthy",
	})
}

// writeJSON is a shared helper that writes a JSON response.
// Extracting repeated logic into a helper keeps handlers short and readable.
//
// Parameters:
//   - w      : the response writer we write to
//   - status : HTTP status code (e.g. 200, 404)
//   - v      : any Go value — "any" is Go's way of saying "any type"
func writeJSON(w http.ResponseWriter, status int, v any) {
	// Set the Content-Type header so clients know the body is JSON.
	// Headers must be set BEFORE calling WriteHeader or writing the body.
	w.Header().Set("Content-Type", "application/json")

	// WriteHeader sends the HTTP status code (e.g. 200 OK).
	// After this call, you can no longer change headers.
	w.WriteHeader(status)

	// json.NewEncoder(w) creates a JSON encoder that writes directly to w.
	// .Encode(v) converts v to JSON and writes it, appending a newline.
	json.NewEncoder(w).Encode(v)
}
