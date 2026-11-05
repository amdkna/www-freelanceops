package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	Addr            string
	ShutdownTimeout time.Duration
}

func Load() (Config, error) {
	port := strings.TrimSpace(os.Getenv("PORT"))
	if port == "" {
		return Config{}, errors.New("PORT is required")
	}
	n, err := strconv.Atoi(port)
	if err != nil || n < 1 || n > 65535 {
		return Config{}, fmt.Errorf("PORT must be an integer between 1 and 65535")
	}
	timeout := 10 * time.Second
	if raw := strings.TrimSpace(os.Getenv("SHUTDOWN_TIMEOUT")); raw != "" {
		timeout, err = time.ParseDuration(raw)
		if err != nil || timeout <= 0 {
			return Config{}, fmt.Errorf("SHUTDOWN_TIMEOUT must be a positive Go duration")
		}
	}
	return Config{Addr: fmt.Sprintf(":%d", n), ShutdownTimeout: timeout}, nil
}
