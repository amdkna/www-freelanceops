package httpapi

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthRoutes(t *testing.T) {
	health := NewHealth()
	handler := Router(health)

	for _, path := range []string{"/health/live", "/health/ready"} {
		w := httptest.NewRecorder()
		handler.ServeHTTP(w, httptest.NewRequest(http.MethodGet, path, nil))
		if w.Code != http.StatusOK {
			t.Fatalf("%s returned %d", path, w.Code)
		}
		if got := w.Header().Get("Content-Type"); got != "application/json; charset=utf-8" {
			t.Fatalf("unexpected content type: %q", got)
		}
		var response struct{ Status string }
		if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil || response.Status != "ok" {
			t.Fatalf("invalid response for %s: %s (%v)", path, w.Body.String(), err)
		}
	}

	health.SetReady(false)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, httptest.NewRequest(http.MethodGet, "/health/ready", nil))
	if w.Code != http.StatusServiceUnavailable {
		t.Fatalf("expected 503 during shutdown, got %d", w.Code)
	}
	w = httptest.NewRecorder()
	handler.ServeHTTP(w, httptest.NewRequest(http.MethodGet, "/health/live", nil))
	if w.Code != http.StatusOK {
		t.Fatalf("live endpoint should stay live, got %d", w.Code)
	}
}

func TestUnknownAndWrongMethod(t *testing.T) {
	handler := Router(NewHealth())
	tests := []struct {
		method, path string
		status       int
	}{
		{http.MethodGet, "/missing", http.StatusNotFound},
		{http.MethodPost, "/health/live", http.StatusMethodNotAllowed},
	}
	for _, tt := range tests {
		w := httptest.NewRecorder()
		handler.ServeHTTP(w, httptest.NewRequest(tt.method, tt.path, nil))
		if w.Code != tt.status {
			t.Errorf("%s %s = %d, want %d", tt.method, tt.path, w.Code, tt.status)
		}
	}
}
