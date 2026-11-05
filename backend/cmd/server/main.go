package main

import (
	"context"
	"errors"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/amdkna/www-freelanceops/backend/internal/config"
	"github.com/amdkna/www-freelanceops/backend/internal/httpapi"
)

func main() {
	if err := run(); err != nil {
		slog.Error("server failed", "error", err)
		os.Exit(1)
	}
}

func run() error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}
	listener, err := net.Listen("tcp", cfg.Addr)
	if err != nil {
		return err
	}
	defer listener.Close()

	health := httpapi.NewHealth()
	server := &http.Server{
		Handler:           httpapi.Router(health),
		ReadHeaderTimeout: 5_000_000_000,
		IdleTimeout:       60_000_000_000,
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	serverErrors := make(chan error, 1)
	go func() {
		serverErrors <- server.Serve(listener)
	}()
	slog.Info("server listening", "address", listener.Addr().String())

	select {
	case err := <-serverErrors:
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	case <-ctx.Done():
		health.SetReady(false)
		slog.Info("server shutting down")
		shutdownCtx, cancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeout)
		defer cancel()
		if err := server.Shutdown(shutdownCtx); err != nil {
			_ = server.Close()
			return err
		}
		if err := <-serverErrors; err != nil && !errors.Is(err, http.ErrServerClosed) {
			return err
		}
		slog.Info("server stopped")
		return nil
	}
}
