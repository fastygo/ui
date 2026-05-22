package serverapp

import (
	"errors"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/fastygo/framework/pkg/app"
	"github.com/fastygo/framework/pkg/web/locale"
	"github.com/fastygo/framework/pkg/web/security"
	"github.com/fastygo/ui/internal/platform"
	"github.com/fastygo/ui/internal/site"
)

var (
	once    sync.Once
	cached  *app.App
	initErr error
)

// New builds (or returns) the assembled HTTP application.
func New() (*app.App, error) {
	once.Do(func() {
		cfg, err := platform.Load()
		if err != nil {
			initErr = err
			return
		}

		logger := newLogger(cfg.LogLevel, cfg.LogFormat)
		slog.SetDefault(logger)

		feat := site.NewFeature(cfg.AvailableLocales, cfg.DefaultLocale, filepath.Join(cfg.StaticDir, "docs"))

		cached = app.New(cfg.Config).
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
			WithFeature(feat).
			Build()
	})
	if initErr != nil {
		return nil, initErr
	}
	if cached == nil {
		return nil, errors.New("serverapp: application not initialized")
	}
	return cached, nil
}

// Handler returns the root HTTP handler for serverless or embedded use.
func Handler() http.Handler {
	application, err := New()
	if err != nil {
		slog.Error("serverapp init", "error", err)
		return http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			http.Error(w, "service unavailable", http.StatusServiceUnavailable)
		})
	}
	return application.Handler()
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
