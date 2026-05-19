// In Go, every file starts with a package declaration.
// "main" is special: it's the entry point package for an executable program.
package main

// "import" brings in standard library packages (or external ones).
// These are Go's built-in packages — no installation needed.
import (
	"fmt"
	"log"      // for logging messages to stderr
	"net/http" // for building HTTP servers and clients
)

// setupRoutes creates and configures the router (called a ServeMux in Go).
// It returns a *http.ServeMux — the * means "pointer to a ServeMux value".
// Returning it (instead of using a global) makes the function easier to test.
func setupRoutes() *http.ServeMux {
	// http.NewServeMux() creates a new request multiplexer (router).
	// It matches incoming request URLs to handler functions.
	mux := http.NewServeMux()

	// HandleFunc registers a handler function for a URL pattern.
	// "/" matches everything that no other pattern matched (catch-all).
	// "/health" matches exactly that path.
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/health", healthHandler)

	// "return" sends the value back to the caller.
	return mux
}

// main() is the entry point — Go always starts execution here.
func main() {
	mux := setupRoutes()

	fmt.Printf("Server is running on %d", 8080)
	// log.Println writes a timestamped line to stderr.
	// log.Println("listening on ")

	// http.ListenAndServe starts the HTTP server.
	// It blocks forever (runs until the program is killed).
	// On error it returns a non-nil error; log.Fatal logs it and exits.
	log.Fatal(http.ListenAndServe(":8080", mux))
}
