# go-sample

Minimal Go HTTP server for learning purposes.

## Commands

```bash
go run .          # start server on :8080
go test ./...     # run tests
go build ./...    # compile check
go vet ./...      # static analysis
```

## Structure

- `main.go` — entry point, route setup
- `handler.go` — HTTP handlers + `writeJSON` helper
- `handler_test.go` — tests using `net/http/httptest`
- `go.mod` — module definition (no external deps)

## Conventions

- All code in `package main` (flat layout, no subdirectories)
- Stdlib only — no third-party packages
- Every handler must be covered by a test
