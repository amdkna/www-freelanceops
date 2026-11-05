# FreelanceOps backend

Minimal Go HTTP service using [Chi](https://github.com/go-chi/chi), tracked by [issue #1](https://github.com/amdkna/www-freelanceops/issues/1).

## Requirements

Go 1.23 or newer. No database or secrets are required for this foundation task.

## Run

From the repository root:

```sh
cd backend
PORT=8080 go run ./cmd/server
```

Windows PowerShell:

```powershell
cd backend
$env:PORT = "8080"
go run ./cmd/server
```

The service deliberately fails at startup without a valid `PORT` to avoid silently using an unexpected bind port. It binds all interfaces on the selected port; restrict exposure at the reverse proxy or firewall.

Optional `SHUTDOWN_TIMEOUT` accepts a positive Go duration, default `10s`.

## Routes

| Method | Path | Success | Failure |
| --- | --- | --- | --- |
| GET | `/health/live` | 200 `{"status":"ok"}` | — |
| GET | `/health/ready` | 200 `{"status":"ok"}` | 503 `{"status":"not_ready"}` during shutdown |

Unknown paths return 404; unsupported methods return 405. Health endpoints are unauthenticated and contain no private data. Readiness currently measures only the server's accepting state, not future database connections.

## Test and build

```sh
cd backend
go test ./...
go vet ./...
go build ./...
```

GitHub Actions automatically discovers `backend/go.mod` and runs formatting, vet, tests, and build on pull requests and main.

## Shutdown

SIGINT or SIGTERM marks readiness false, stops accepting new connections, and allows in-flight requests up to `SHUTDOWN_TIMEOUT`; an expired deadline forces connection closure. No credentials or full environment are logged.
