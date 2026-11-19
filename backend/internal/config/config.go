package config

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	Addr            string
	ShutdownTimeout time.Duration
	DBPath          string
	PublicOrigin    string
	AdminPassword   string
	LogLevel        string
}

func Load() (Config, error) {
	port := strings.TrimSpace(os.Getenv("PORT"))
	if port == "" {
		return Config{}, errors.New("PORT is required")
	}
	n, err := strconv.Atoi(port)
	if err != nil || n < 1 || n > 65535 {
		return Config{}, errors.New("PORT must be an integer between 1 and 65535")
	}
	timeout := 10 * time.Second
	if raw := strings.TrimSpace(os.Getenv("SHUTDOWN_TIMEOUT")); raw != "" {
		timeout, err = time.ParseDuration(raw)
		if err != nil || timeout <= 0 {
			return Config{}, errors.New("SHUTDOWN_TIMEOUT must be a positive Go duration")
		}
	}
	dbPath := strings.TrimSpace(os.Getenv("DB_PATH"))
	if dbPath == "" {
		dbPath = "/data/freelanceops.db"
	}
	if !strings.HasPrefix(dbPath, "/") {
		return Config{}, errors.New("DB_PATH must be absolute")
	}
	origin := strings.TrimRight(strings.TrimSpace(os.Getenv("PUBLIC_ORIGIN")), "/")
	if origin == "" {
		return Config{}, errors.New("PUBLIC_ORIGIN is required")
	}
	parsed, err := url.Parse(origin)
	if err != nil || parsed.Host == "" || parsed.User != nil || parsed.RawQuery != "" || parsed.Fragment != "" || parsed.Path != "" || parsed.Scheme != "https" && parsed.Scheme != "http" {
		return Config{}, errors.New("PUBLIC_ORIGIN must be an absolute HTTP(S) origin")
	}
	level := strings.ToLower(strings.TrimSpace(os.Getenv("LOG_LEVEL")))
	if level == "" {
		level = "info"
	}
	if level != "debug" && level != "info" && level != "warn" && level != "error" {
		return Config{}, errors.New("LOG_LEVEL must be debug, info, warn, or error")
	}
	return Config{Addr: fmt.Sprintf(":%d", n), ShutdownTimeout: timeout, DBPath: dbPath, PublicOrigin: origin, AdminPassword: os.Getenv("ADMIN_PASSWORD"), LogLevel: level}, nil
}
