# go-sample

Minimal Go HTTP server. Two endpoints, no external dependencies.

## Endpoints

| Method | Path      | Response                              |
|--------|-----------|---------------------------------------|
| GET    | `/`       | `{"status":"ok","message":"hello"}`   |
| GET    | `/health` | `{"status":"healthy"}`                |

## Run

```bash
go run .
# listening on :8080
```

## Test

```bash
go test ./...
```

## Build

```bash
go build -o bin/server .
./bin/server
```
