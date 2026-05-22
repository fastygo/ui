package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/fastygo/framework/pkg/app"
	"github.com/fastygo/framework/pkg/web/locale"
	"github.com/fastygo/framework/pkg/web/security"
	"github.com/fastygo/ui/internal/platform"
	"github.com/fastygo/ui/internal/site"
)

func main() {
	cfg, err := platform.Load()
	if err != nil {
		slog.Error("config", "error", err)
		os.Exit(1)
	}

	logger := newLogger(cfg.LogLevel, cfg.LogFormat)
	slog.SetDefault(logger)

	feat := site.NewFeature(cfg.AvailableLocales, cfg.DefaultLocale, filepath.Join(cfg.StaticDir, "docs"))

	builder := app.New(cfg.Config).
		WithLogger(logger).
		WithSecurity(security.LoadConfig()).
		WithLocales(app.LocalesConfig{
			Default:   cfg.DefaultLocale,
			Available: cfg.AvailableLocales,
			Strategy:  nil,
			Cookie: locale.CookieOptions{
				Enabled: true,
				Name:    "lang",
			},
			SPA: true,
		}).
		WithHealthEndpoints("/healthz", "/readyz").
		WithFeature(feat)

	application := builder.Build()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	if err := application.Run(ctx); err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Error("server", "error", err)
		os.Exit(1)
	}
}

func newLogger(level, format string) *slog.Logger {
	var lvl slog.Level
	switch strings.ToLower(strings.TrimSpace(level)) {
	case "debug":
		lvl = slog.LevelDebug
	case "warn", "warning":
		lvl = slog.LevelWarn
	case "error":
		lvl = slog.LevelError
	default:
		lvl = slog.LevelInfo
	}
	opts := &slog.HandlerOptions{Level: lvl}
	var h slog.Handler
	if strings.ToLower(strings.TrimSpace(format)) == "json" {
		h = slog.NewJSONHandler(os.Stdout, opts)
	} else {
		h = slog.NewTextHandler(os.Stdout, opts)
	}
	return slog.New(h)
}
