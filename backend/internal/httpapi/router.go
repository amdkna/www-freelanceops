package httpapi

import (
	"encoding/json"
	"net/http"
	"sync/atomic"

	"github.com/go-chi/chi/v5"
)

type Health struct {
	ready atomic.Bool
}

func NewHealth() *Health {
	h := &Health{}
	h.ready.Store(true)
	return h
}

func (h *Health) SetReady(ready bool) {
	h.ready.Store(ready)
}

func Router(h *Health) http.Handler {
	r := chi.NewRouter()
	r.Get("/health/live", func(w http.ResponseWriter, _ *http.Request) {
		writeStatus(w, http.StatusOK, "ok")
	})
	r.Get("/health/ready", func(w http.ResponseWriter, _ *http.Request) {
		if !h.ready.Load() {
			writeStatus(w, http.StatusServiceUnavailable, "not_ready")
			return
		}
		writeStatus(w, http.StatusOK, "ok")
	})
	return r
}

func writeStatus(w http.ResponseWriter, code int, status string) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Cache-Control", "no-store")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(struct {
		Status string `json:"status"`
	}{Status: status})
}
