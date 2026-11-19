# Frontend and API foundation (T02 / T03)

Status: Initial foundation; backend resource endpoints are **contract-only**, not implemented.

## Local usage

Node 22+ is required. Run `cd frontend && npm install && npm run dev` (the Vite development server proxies `/api` to the Go server at localhost:8080). Run `npm run typecheck && npm test && npm run build` to validate. Until a committed lockfile is generated, the pipeline generates a temporary npm lock and then uses `npm ci`; commit the generated `frontend/package-lock.json` to obtain fully reproducible installs.

## Architecture

- Vite + React + TypeScript, React Router with an SPA fallback in `web/default.conf`.
- Tailwind CSS v4, a locally owned shadcn/ui-style Button (Radix Slot + CVA), TanStack Query for backend health polling.
- Browser calls `/api/*`; Nginx strips `/api/` and forwards requests to Go `/*`. Go currently implements only `/health/live` and `/health/ready`.
- `docs/openapi.yaml` defines **planned** auth, websites, incidents, checks, renewals, consistent UTC timestamps, HTTP codes, error envelopes and pagination. The generated-style typed API adapter in `frontend/src/lib/api` follows the contract. Do not assume planned endpoints are live yet.
- The empty dashboard is deliberate; no mock websites or misleading counts are displayed. Website mutations require a later backend persistence and authentication implementation.

## Verification

- Frontend: `npm ci && npm run typecheck && npm test && npm run build` after lockfile generation.
- API contract: `npx --yes @redocly/cli@1.34.5 lint docs/openapi.yaml` (CI).
- Go: `cd backend && go test ./...`.
- Docker: `docker compose config --quiet`, `docker compose up -d --build`, then `curl -f http://localhost:PORT/healthz` through the web container or reverse proxy.
- SPA direct-refresh: navigate to a nonexistent frontend route and refresh; the React 404 page should appear, not an Nginx error. API 404 responses must not be rewritten to HTML.

## Security and operating notes

The future authenticated API uses HttpOnly Secure SameSite cookies and must enforce CSRF protection for mutating requests, rate limiting, authorization and SSRF protections on monitored URLs. The TypeScript input checks are UX validation only; the Go backend must independently validate every field. Never store secrets in frontend code or expose Go directly to the internet. Use RFC3339 UTC timestamps with the trailing `Z`.
