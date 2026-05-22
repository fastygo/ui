.PHONY: help tidy css generate docs build verify run clean deploy install vercel-build

APP_NAME := blank
BIN_DIR := bin
BINARY := $(BIN_DIR)/$(APP_NAME)

# Default target shows available commands
help:
	@echo "FastyGo UI - Makefile targets:"
	@echo ""
	@echo "  tidy         Update go.mod/go.sum and refresh vendor/"
	@echo "  css          Build minified Tailwind CSS (bun)"
	@echo "  docs         Generate static docs HTML (en + ru)"
	@echo "  generate     Run templ code generation"
	@echo "  build        Production build: generate + css + go build -mod=vendor"
	@echo "  verify       Full check: templ, css, ui8px lint, go test"
	@echo "  run          Start dev server via bun script (with signal handling)"
	@echo "  install      Install bun dependencies"
	@echo "  clean        Remove build artifacts (bin/)"
	@echo "  deploy       Full production pipeline (tidy + generate + css + build)"
	@echo "  vercel-build Vercel static export (→ public/)"
	@echo ""
	@echo "Server deploy example:"
	@echo "  git pull && make deploy"
	@echo "  sudo systemctl restart $(APP_NAME)"
	@echo ""
	@echo "Vercel: connect repo — zero config (see vercel.json + bun run vercel:build)"

# Ensure Go dependencies and vendor directory are up to date.
# Run this after git pull on the server or when adding new imports.
tidy:
	go mod tidy
	go mod vendor

# Build frontend assets.
css:
	bun run build:css

# Generate static documentation under web/static/docs/.
docs:
	bun run docs:build

# Generate Go code from .templ files.
# Must run before building if templates changed.
generate:
	go tool templ generate ./...

# Production binary build.
# Uses vendor/ for reproducible offline builds.
# Strips debug info with -ldflags="-s -w".
build: generate css
	@mkdir -p $(BIN_DIR)
	go build -mod=vendor -ldflags="-s -w" -o $(BINARY) ./cmd/server

# Vercel static export: templ + css + docs → public/ (see vercel.json).
vercel-build:
	bun run vercel:build

# Run the full verification pipeline (matches package.json "verify").
verify:
	bun run verify

# Start development server (uses scripts/run-server.mjs for proper signals).
run:
	bun run go

# Install bun dependencies (run once after clone or package.json changes).
install:
	bun install

# Remove compiled binary and other build outputs.
clean:
	rm -rf $(BIN_DIR)/*

# Production deploy pipeline.
# Intended to be run on the server right after `git pull`.
deploy: tidy generate css build
	@echo ""
	@echo "✅ Deploy build finished."
	@echo "   Binary ready: $(BINARY)"
	@echo ""
	@echo "Next steps on server:"
	@echo "   sudo systemctl restart $(APP_NAME)   # if using systemd"
	@echo "   # or: ./$(BINARY)                   # for manual start"
