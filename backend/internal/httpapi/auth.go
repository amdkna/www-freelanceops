package httpapi

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/amdkna/www-freelanceops/backend/internal/auth"
	"github.com/go-chi/chi/v5"
)

func writeError(w http.ResponseWriter, status int, code, message string) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Cache-Control", "no-store")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(map[string]any{"error": map[string]string{"code": code, "message": message}})
}
func authRoutes(r chi.Router, s *auth.Store) {
	r.Post("/v1/auth/login", func(w http.ResponseWriter, req *http.Request) {
		if !s.ValidOrigin(req) {
			writeError(w, 403, "ORIGIN_REJECTED", "Invalid request origin")
			return
		}
		if !s.CheckLoginLimit(req.RemoteAddr) {
			writeError(w, 429, "RATE_LIMITED", "Too many login attempts")
			return
		}
		req.Body = http.MaxBytesReader(w, req.Body, 4096)
		var body struct {
			Password string `json:"password"`
		}
		if err := json.NewDecoder(req.Body).Decode(&body); err != nil || strings.TrimSpace(body.Password) == "" {
			writeError(w, 400, "VALIDATION_ERROR", "Password required")
			return
		}
		token, csrf, err := s.Login(req.Context(), body.Password)
		if errors.Is(err, auth.ErrNotConfigured) {
			writeError(w, 503, "SETUP_REQUIRED", "Admin account has not been configured")
			return
		}
		if errors.Is(err, auth.ErrInvalidCredentials) {
			writeError(w, 401, "INVALID_CREDENTIALS", "Invalid credentials")
			return
		}
		if err != nil {
			writeError(w, 500, "INTERNAL_ERROR", "Unexpected server error")
			return
		}
		s.ClearLoginLimit(req.RemoteAddr)
		http.SetCookie(w, &http.Cookie{Name: auth.CookieName, Value: token, Path: "/api", HttpOnly: true, Secure: true, SameSite: http.SameSiteStrictMode, MaxAge: 86400})
		writeJSON(w, 200, map[string]any{"authenticated": true, "csrf_token": csrf})
	})
	r.Group(func(private chi.Router) {
		private.Use(func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				cookie, err := req.Cookie(auth.CookieName)
				if err != nil {
					writeError(w, 401, "UNAUTHORIZED", "Authentication required")
					return
				}
				csrf, ok := s.Session(req.Context(), cookie.Value)
				if !ok {
					writeError(w, 401, "UNAUTHORIZED", "Authentication required")
					return
				}
				if req.Method != http.MethodGet && req.Method != http.MethodHead && req.Method != http.MethodOptions {
					if !s.ValidOrigin(req) || !s.CSRFMatches(req.Header.Get("X-CSRF-Token"), csrf) {
						writeError(w, 403, "CSRF_REJECTED", "CSRF token or origin invalid")
						return
					}
				}
				next.ServeHTTP(w, req)
			})
		})
		private.Get("/v1/auth/session", func(w http.ResponseWriter, req *http.Request) {
			c, _ := req.Cookie(auth.CookieName)
			csrf, _ := s.Session(req.Context(), c.Value)
			writeJSON(w, 200, map[string]any{"authenticated": true, "csrf_token": csrf})
		})
		private.Post("/v1/auth/logout", func(w http.ResponseWriter, req *http.Request) {
			c, _ := req.Cookie(auth.CookieName)
			if err := s.Logout(req.Context(), c.Value); err != nil {
				writeError(w, 500, "INTERNAL_ERROR", "Unexpected server error")
				return
			}
			http.SetCookie(w, &http.Cookie{Name: auth.CookieName, Path: "/api", MaxAge: -1, HttpOnly: true, Secure: true, SameSite: http.SameSiteStrictMode})
			w.WriteHeader(204)
		})
		private.Route("/v1", func(v chi.Router) {
			v.Get("/websites", notImplemented)
			v.Post("/websites", notImplemented)
			v.Get("/websites/{websiteId}", notImplemented)
			v.Get("/websites/{websiteId}/checks", notImplemented)
			v.Get("/websites/{websiteId}/incidents", notImplemented)
			v.Get("/websites/{websiteId}/renewals", notImplemented)
			v.Post("/renewals/{renewalId}/confirm", notImplemented)
		})
	})
}
func notImplemented(w http.ResponseWriter, _ *http.Request) {
	writeError(w, 501, "NOT_IMPLEMENTED", "This endpoint is planned for a later task")
}
func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Cache-Control", "no-store")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
