# FastyGo UI

Wireframe **structure lab** and reference site for the [FastyGo](https://github.com/fastygo) stack: [FastyGo Framework](https://github.com/fastygo/framework), [`github.com/fastygo/templ`](https://github.com/fastygo/templ) (atoms + kit composites), and [UI8Kit](https://github.com/fastygo/ui8kit) for client ARIA behavior.

The app ships a full **docs / component gallery** (`/docs/`) with embedded **en** / **ru** locale JSON, sidebar shell, theme toggle, and accessibility checks. There is no login or control plane — use it as a starting point for documentation sites, registry showcases, or marketing pages that share the same chrome.

**License:** [MIT](LICENSE)

## What this repo is

| Layer | Location | Role |
|-------|----------|------|
| Atoms & kit | `github.com/fastygo/templ` (`ui/`, `components/`) | Buttons, stacks, forms, cards — vendored in Go |
| App registry | [`internal/ui/`](internal/ui/) | Shell, blocks, widgets, showcase components (staging before extraction) |
| Docs content | [`internal/showcase/content/`](internal/showcase/content/) | Markdown sources per locale |
| Docs compiler | [`cmd/docgen`](cmd/docgen) + [`internal/showcase/docgen/`](internal/showcase/docgen/) | Static HTML → `web/static/docs/` |
| Pages | [`internal/views/`](internal/views/) | `templ` layouts and docs rendering |

Design policy (ui8px, composition, fixtures, validation) lives in [`.cursor/rules/`](.cursor/rules/) and [`.ui8px/policy/`](.ui8px/policy/).

## Prerequisites

- Go 1.25+
- [Bun](https://bun.sh) (CSS build, ui8px, Playwright)

## Quick start

```bash
bun install
go mod download
bun run go
```

`bun run go` runs `templ generate`, `build:css`, then [`scripts/run-server.mjs`](scripts/run-server.mjs) (`go run ./cmd/server` from the repo root so `web/static` resolves correctly). **Ctrl+C** forwards to the Go process and frees the port.

Open:

- [http://127.0.0.1:8080/](http://127.0.0.1:8080/) — home
- [http://127.0.0.1:8080/docs/](http://127.0.0.1:8080/docs/) — docs index and component gallery
- [http://127.0.0.1:8080/sample](http://127.0.0.1:8080/sample) — second stub route

Closing a browser tab does **not** stop the server — stop the terminal job or see [Troubleshooting](#troubleshooting).

### Docs static HTML

`/docs/` is served from prebuilt files under [`web/static/docs/`](web/static/docs/). After changing docs markdown, `templ` views, or index card layout, regenerate and commit:

```bash
bun run docs:build
```

Source markdown: `internal/showcase/content/{en,ru}/`. Output layout follows locale routing (`/docs/…` → `en/`, `/ru/docs/…` → `ru/`).

## Environment

| Variable | Default | Purpose |
|----------|---------|---------|
| `APP_BIND` | `127.0.0.1:8080` | HTTP listen address |
| `APP_STATIC_DIR` | `web/static` when unset | Static files under `/static/`. Set an absolute path if the server cwd is not the repo root. |
| `SESSION_KEY` | dev-only fallback (logged) | Reserved for future session features |
| `APP_DEFAULT_LOCALE` | `en` | Default locale (unprefixed `/docs/` URLs) |
| `APP_AVAILABLE_LOCALES` | `en,ru` | Locales for the header switcher |

Health: `GET /healthz`, `GET /readyz` ([`cmd/server/main.go`](cmd/server/main.go)).

## Deploy on Vercel (static docs)

Connect the Git repository — **no Go runtime on Vercel**.

[`vercel.json`](vercel.json) build:

```bash
bun install && bun run build:css && node scripts/vercel-static-export.mjs
```

**Docs HTML is not built on Vercel.** Commit `web/static/docs/` from local `bun run docs:build` when content or templates change.

| Step | Where |
|------|--------|
| `docs:build` (Go + templ) | Local / CI before push |
| `build:css` (Tailwind v4) | Vercel build |
| `vercel-static-export` | Vercel build → `public/` |

| URL | Source in repo |
|-----|----------------|
| `/docs/…` | `web/static/docs/en/…` |
| `/ru/docs/…` | `web/static/docs/ru/…` |
| `/static/…` | `web/static/{css,js,img,fonts}` |

Redirects: `/` → `/docs/`; `/en/docs/…` → `/docs/…` when `en` is default.

Local preview: `bun run vercel:build && npx serve public`

For a **Go binary** on your own host, use `make deploy`.

## Project layout

| Path | Role |
|------|------|
| [`cmd/server/`](cmd/server/) | HTTP entrypoint |
| [`cmd/docgen/`](cmd/docgen/) | Static docs generator CLI |
| [`internal/serverapp/`](internal/serverapp/) | App wiring (config, locales, features) |
| [`internal/site/`](internal/site/) | Routes: `/`, `/sample`, `/docs/…`, illustration lab under `/lab/` |
| [`internal/fixtures/`](internal/fixtures/) | Embedded locale JSON and typed `Locale` model |
| [`internal/views/`](internal/views/) | `templ` pages, shell, docs static renderer |
| [`internal/ui/`](internal/ui/) | In-app UI registry — see [`internal/ui/README.md`](internal/ui/README.md) |
| [`internal/showcase/`](internal/showcase/) | Markdown sources + docgen pipeline |
| [`web/static/`](web/static/) | Built `app.css`, UI8Kit JS, theme script, fonts, images, committed docs HTML |
| [`.validate/`](.validate/) | Nu HTML snapshots, APG notes, spec validation |
| [`scripts/`](scripts/) | Dev server, Vercel export, HTML validation |

Static assets (committed): Tailwind output, [`web/static/css/tweakcn.css`](web/static/css/tweakcn.css) tokens, [`web/static/css/fonts.css`](web/static/css/fonts.css), Latty icon masks, `@ui8kit/aria` bundle + [`web/static/js/manifest.json`](web/static/js/manifest.json).

## Verification

```bash
bun run verify
```

Pipeline: `templ generate` → `build:css` → spec sync/validate → `docs:build` → docs snapshot copy → `ui8px lint` → `ui8px validate aria` → Nu HTML on `.validate/html-snapshots/nu/**/*.html` (network; empty dir exits 0) → `go test ./...` → Playwright + axe (home route; `color-contrast` disabled until brand phase).

First-time e2e browsers: `bun run test:e2e:install`.

Refresh HTML snapshots: `go run ./.validate/cmd/snapshot-render`, then copy conforming files into `.validate/html-snapshots/nu/`.

Details: [`.cursor/rules/fastygo-ui-validation-testing.mdc`](.cursor/rules/fastygo-ui-validation-testing.mdc).

## Troubleshooting

### Port already in use

Another process (often a previous `go run`) holds the port:

```bash
export APP_BIND=127.0.0.1:8081
bun run go
```

Windows: `netstat -ano | findstr :8080` then `taskkill /PID <pid> /F`.

### Static files 404

Run from the repo root (or `bun run go`), run `bun run build:css`, and check `APP_STATIC_DIR`.

### Docs look stale after template changes

Run `bun run docs:build` and refresh the browser. `/docs/` serves files from `web/static/docs/`, not live Go rendering.

## Using as a template

1. Copy the repo (or the app subtree: `cmd/`, `internal/`, `web/static/`, `package.json`, `.ui8px/`, `.validate/`, `tests/`, `playwright.config.ts`, `scripts/`).
2. Change the module path in `go.mod` and imports.
3. Add routes in [`internal/site/feature.go`](internal/site/feature.go); extend nav copy in [`internal/fixtures/locale/`](internal/fixtures/locale/).
4. Add `templ` under [`internal/views/`](internal/views/) and registry UI under [`internal/ui/`](internal/ui/); keep Tailwind `@source` in [`web/static/css/input.css`](web/static/css/input.css) in sync.

## Roadmap

- **Now:** wireframe IA, semantics, a11y, registry fill — explicit Tailwind utilities, no brand polish.
- **Later:** visual identity, token contrast lock, extract stable blocks/widgets to `github.com/fastygo/blocks` / `github.com/fastygo/widgets`.

The [`.fastygo/`](.fastygo/) directory in some workspaces is reference-only and is **not** imported at runtime.
