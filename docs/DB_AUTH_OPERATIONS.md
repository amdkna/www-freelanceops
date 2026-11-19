# Database, authentication, and quality operations (issues #4–#6)

## SQLite / migrations
The Go API uses SQLite via `database/sql` with the pure-Go modernc driver, WAL mode, 5-second busy timeout, foreign keys, and a single open connection. Versioned migrations are embedded in `backend/internal/storage/migrations` and applied atomically with a `schema_migrations` ledger on startup. Schema tables cover websites, checks, incidents, expirations, notifications, renewals, admins and sessions.

Docker Compose mounts the named volume `freelanceops-data` at `/data`. Do not use `docker compose down -v` in production. Before migration changes, back up the database using a SQLite-consistent online backup / `VACUUM INTO` while the application is running (do not copy only the WAL-backed .db file). Restore into an offline container with correct ownership, then deploy an image using compatible migrations. Migrations are forward-only; rollbacks require a verified backup.

## First admin setup
On the **server**, create `/opt/apps/freelanceops/.env` from `.env.example` and set `ADMIN_PASSWORD` to a unique 12+ character password, along with the correct HTTPS `PUBLIC_ORIGIN`. Protect the file (`chmod 600 .env`). On first boot, a bcrypt password hash is stored; subsequent restarts do not reset it even if the environment variable remains set. Remove `ADMIN_PASSWORD` from the environment after successful bootstrap and redeploy. With no password and no admin present, the application starts locked (health checks pass, login returns 503 SETUP_REQUIRED). There is **no default password** and no anonymous admin access.

Production cookies are Secure, HttpOnly, SameSite=Strict, scoped to `/api`, valid for 24 hours. Session tokens are random and stored only as SHA256 hashes; logout revokes the session. Browser mutation requests need the same-origin `Origin` header and the session-specific `X-CSRF-Token` value. Login requires Origin and has a basic in-memory rate limiter; deploy behind the existing TLS reverse proxy and configure additional edge rate limits. If deployed to a different hostname, update PUBLIC_ORIGIN; HTTP development requires a local HTTPS reverse proxy to use Secure cookies.

## Validation
Run `cd backend && gofmt -w ./cmd ./internal && go mod tidy && go vet ./... && go test ./...` and `cd frontend && npm ci && npm run lint && npm run typecheck && npm test && npm run build`. The CI workflow validates format, vet, unit tests, frontend checks, OpenAPI, Compose and configuration. Secrets must never appear in logs, commits, or screenshots.

## Current API scope
Implemented: `GET /health/live`, `GET /health/ready`, `POST /v1/auth/login`, `GET /v1/auth/session`, and `POST /v1/auth/logout`. Protected resource endpoints currently return HTTP 501 after authentication and CSRF validation; website CRUD and monitoring are separate later tasks. Authenticated login and session responses include `csrf_token`, which is never put in the cookie.
