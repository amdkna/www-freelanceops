package config

import (
	"testing"
	"time"
)

func TestLoad(t *testing.T) {
	tests := []struct {
		name, port, timeout string
		wantError           bool
	}{
		{"missing", "", "", true},
		{"invalid port", "abc", "", true},
		{"zero port", "0", "", true},
		{"port too large", "65536", "", true},
		{"bad timeout", "8080", "oops", true},
		{"negative timeout", "8080", "-1s", true},
		{"valid", "8080", "2s", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("PORT", tt.port)
			t.Setenv("SHUTDOWN_TIMEOUT", tt.timeout)
			cfg, err := Load()
			if (err != nil) != tt.wantError {
				t.Fatalf("Load() error = %v, wantError = %v", err, tt.wantError)
			}
			if !tt.wantError && (cfg.Addr != ":8080" || cfg.ShutdownTimeout != 2*time.Second) {
				t.Fatalf("unexpected config: %+v", cfg)
			}
		})
	}
}
