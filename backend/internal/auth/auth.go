package auth

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"crypto/subtle"
	"database/sql"
	"encoding/hex"
	"errors"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"golang.org/x/crypto/bcrypt"
)

const CookieName = "freelanceops_session"

var ErrNotConfigured = errors.New("admin not configured")
var ErrInvalidCredentials = errors.New("invalid credentials")

type Store struct {
	DB         *sql.DB
	Origin     string
	attemptsMu sync.Mutex
	attempts   map[string]attempt
}
type attempt struct {
	Count int
	Until time.Time
}

func New(db *sql.DB, origin string) *Store {
	return &Store{DB: db, Origin: strings.TrimRight(origin, "/"), attempts: make(map[string]attempt)}
}
func (s *Store) Bootstrap(ctx context.Context, password string) error {
	var count int
	if err := s.DB.QueryRowContext(ctx, "SELECT count(*) FROM admins").Scan(&count); err != nil {
		return err
	}
	if count != 0 {
		return nil
	}
	if password == "" {
		return nil
	} // Locked until a secure bootstrap password is supplied.
	if len(password) < 12 {
		return errors.New("ADMIN_PASSWORD must be at least 12 characters")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	_, err = s.DB.ExecContext(ctx, "INSERT INTO admins(id,password_hash,created_at) VALUES(1,?,?)", string(hash), time.Now().UTC().Format(time.RFC3339Nano))
	return err
}
func randomToken() (string, error) {
	raw := make([]byte, 32)
	if _, err := rand.Read(raw); err != nil {
		return "", err
	}
	return hex.EncodeToString(raw), nil
}
func hashToken(token string) string {
	h := sha256.Sum256([]byte(token))
	return hex.EncodeToString(h[:])
}
func (s *Store) CheckLoginLimit(key string) bool {
	s.attemptsMu.Lock()
	defer s.attemptsMu.Unlock()
	now := time.Now()
	a := s.attempts[key]
	if now.After(a.Until) {
		a = attempt{Until: now.Add(15 * time.Minute)}
	}
	if a.Count >= 5 {
		return false
	}
	a.Count++
	s.attempts[key] = a
	return true
}
func (s *Store) ClearLoginLimit(key string) {
	s.attemptsMu.Lock()
	delete(s.attempts, key)
	s.attemptsMu.Unlock()
}
func (s *Store) Login(ctx context.Context, password string) (token, csrf string, err error) {
	var stored string
	if err = s.DB.QueryRowContext(ctx, "SELECT password_hash FROM admins WHERE id=1").Scan(&stored); errors.Is(err, sql.ErrNoRows) {
		return "", "", ErrNotConfigured
	} else if err != nil {
		return "", "", err
	}
	if bcrypt.CompareHashAndPassword([]byte(stored), []byte(password)) != nil {
		return "", "", ErrInvalidCredentials
	}
	token, err = randomToken()
	if err != nil {
		return
	}
	csrf, err = randomToken()
	if err != nil {
		return
	}
	now := time.Now().UTC()
	_, err = s.DB.ExecContext(ctx, "INSERT INTO sessions(token_hash,admin_id,csrf_token,expires_at,created_at) VALUES(?,1,?,?,?)", hashToken(token), csrf, now.Add(24*time.Hour).Format(time.RFC3339Nano), now.Format(time.RFC3339Nano))
	return
}
func (s *Store) Session(ctx context.Context, token string) (string, bool) {
	if len(token) != 64 {
		return "", false
	}
	var csrf, expires string
	if err := s.DB.QueryRowContext(ctx, "SELECT csrf_token,expires_at FROM sessions WHERE token_hash=?", hashToken(token)).Scan(&csrf, &expires); err != nil {
		return "", false
	}
	expiry, err := time.Parse(time.RFC3339Nano, expires)
	return csrf, err == nil && time.Now().Before(expiry)
}
func (s *Store) Logout(ctx context.Context, token string) error {
	if token == "" {
		return nil
	}
	_, err := s.DB.ExecContext(ctx, "DELETE FROM sessions WHERE token_hash=?", hashToken(token))
	return err
}
func (s *Store) ValidOrigin(r *http.Request) bool {
	origin := r.Header.Get("Origin")
	if origin == "" {
		return false
	}
	parsed, err := url.Parse(origin)
	if err != nil || parsed.Scheme == "" || parsed.Host == "" || parsed.User != nil || parsed.Path != "" || parsed.RawQuery != "" || parsed.Fragment != "" {
		return false
	}
	return subtle.ConstantTimeCompare([]byte(strings.TrimRight(origin, "/")), []byte(s.Origin)) == 1
}
func (s *Store) CSRFMatches(got, want string) bool {
	if got == "" || want == "" {
		return false
	}
	return subtle.ConstantTimeCompare([]byte(got), []byte(want)) == 1
}
