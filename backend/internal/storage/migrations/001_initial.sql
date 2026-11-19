CREATE TABLE IF NOT EXISTS admins (
 id INTEGER PRIMARY KEY CHECK (id = 1),
 password_hash TEXT NOT NULL,
 created_at TEXT NOT NULL
);
CREATE TABLE IF NOT EXISTS websites (
 id TEXT PRIMARY KEY,
 name TEXT NOT NULL CHECK (length(trim(name)) BETWEEN 1 AND 120),
 url TEXT NOT NULL UNIQUE,
 status TEXT NOT NULL DEFAULT 'unknown' CHECK (status IN ('up','down','unknown')),
 domain_expires_at TEXT,
 ssl_expires_at TEXT,
 created_at TEXT NOT NULL
);
CREATE TABLE IF NOT EXISTS checks (
 id TEXT PRIMARY KEY,
 website_id TEXT NOT NULL REFERENCES websites(id) ON DELETE CASCADE,
 checked_at TEXT NOT NULL,
 status_code INTEGER CHECK (status_code IS NULL OR status_code BETWEEN 100 AND 599),
 success INTEGER NOT NULL CHECK (success IN (0,1))
);
CREATE INDEX IF NOT EXISTS idx_checks_website_time ON checks (website_id, checked_at DESC);
CREATE TABLE IF NOT EXISTS incidents (
 id TEXT PRIMARY KEY,
 website_id TEXT NOT NULL REFERENCES websites(id) ON DELETE CASCADE,
 started_at TEXT NOT NULL,
 resolved_at TEXT,
 CHECK (resolved_at IS NULL OR resolved_at >= started_at)
);
CREATE INDEX IF NOT EXISTS idx_incidents_website ON incidents(website_id, started_at DESC);
CREATE TABLE IF NOT EXISTS expirations (
 id TEXT PRIMARY KEY,
 website_id TEXT NOT NULL REFERENCES websites(id) ON DELETE CASCADE,
 kind TEXT NOT NULL CHECK (kind IN ('domain','ssl')),
 expires_at TEXT NOT NULL,
 UNIQUE (website_id,kind)
);
CREATE TABLE IF NOT EXISTS notifications (
 id TEXT PRIMARY KEY,
 website_id TEXT NOT NULL REFERENCES websites(id) ON DELETE CASCADE,
 event_type TEXT NOT NULL,
 event_key TEXT NOT NULL UNIQUE,
 created_at TEXT NOT NULL,
 sent_at TEXT
);
CREATE TABLE IF NOT EXISTS renewals (
 id TEXT PRIMARY KEY,
 website_id TEXT NOT NULL REFERENCES websites(id) ON DELETE CASCADE,
 kind TEXT NOT NULL CHECK (kind IN ('domain','ssl')),
 expires_at TEXT NOT NULL,
 confirmed_at TEXT
);
CREATE TABLE IF NOT EXISTS sessions (
 token_hash TEXT PRIMARY KEY,
 admin_id INTEGER NOT NULL REFERENCES admins(id) ON DELETE CASCADE,
 csrf_token TEXT NOT NULL,
 expires_at TEXT NOT NULL,
 created_at TEXT NOT NULL
);
CREATE INDEX IF NOT EXISTS idx_sessions_expiry ON sessions(expires_at);
