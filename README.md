# FastyGo UI

Reference app for [FastyGo Framework](https://github.com/fastygo/framework) and [UI8Kit](https://github.com/fastygo/ui8kit) `v0.4.0`: **default app shell** (sidebar, mobile sheet, header, language and theme toggles) with **embedded locale JSON** (`en` / `ru`). No login, no control-plane panel — suitable as the base for a docs / component-gallery / marketing site while keeping the same chrome.

## Prerequisites

- Go 1.25+
- [Bun](https://bun.sh) (for CSS build and `ui8px`)

## Quick start

```bash
bun install
go mod download
bun run build:css
go tool templ generate ./...
bun run go
```

UI8Kit static CSS/JS, theme scripts, and **Google Sans** (`gfonts.css` + `fonts/google-sans/`) are committed under [`web/static/`](web/static/).

`bun run go` runs [`scripts/run-server.mjs`](scripts/run-server.mjs): the server always starts with the **repository root as cwd** (correct `web/static`), and **Ctrl+C** is forwarded to the Go process so the port is released.

Closing a **browser tab does not** stop an HTTP server. Stop the job in the terminal (**Ctrl+C**) or close the terminal panel; if the port stays busy, an old `go` process is still running (see troubleshooting below).

Open [http://127.0.0.1:8080/](http://127.0.0.1:8080/) for the home page, [http://127.0.0.1:8080/docs](http://127.0.0.1:8080/docs) for the component gallery, or [http://127.0.0.1:8080/sample](http://127.0.0.1:8080/sample) for the second stub route.

## Environment

| Variable | Default | Purpose |
|----------|---------|---------|
| `APP_BIND` | `127.0.0.1:8080` | HTTP listen address |
| `APP_STATIC_DIR` | `web/static` when env omitted | Static files under `/static/`. Framework’s built-in default points at a CMS-style folder; this app **forces** `web/static` whenever `APP_STATIC_DIR` is not set in the environment. Use an absolute path if you do not start the server from the repo root. |
| `SESSION_KEY` | dev-only fallback (logged) | Reserved for future session-backed features (framework config) |
| `APP_DEFAULT_LOCALE` | `en` | Default locale |
| `APP_AVAILABLE_LOCALES` | `en,ru` | Locales for the header switcher (query + cookie) |

Probes: `GET /healthz` and `GET /readyz` are registered in [`cmd/server/main.go`](cmd/server/main.go).

## Deploy on Vercel (static docs)

Connect the Git repository in Vercel — **no project settings, no Go on the server**.

[`vercel.json`](vercel.json) runs only:

```bash
bun install && bun run build:css && node scripts/vercel-static-export.mjs
```

**Docs HTML is not generated on Vercel** — it must already live in the repo under `web/static/docs/` (from local `bun run docs:build`, which needs Go). Commit those files when you change markdown or templates.

| Step | Where |
|------|--------|
| `docs:build` (Go + templ) | **Local** / CI before push |
| `build:css` (Tailwind) | Vercel build |
| `vercel-static-export` | Vercel build → `public/` |

Export layout:

| URL | Source in repo |
|-----|----------------|
| `/docs/…` | `web/static/docs/en/…` |
| `/ru/docs/…` | `web/static/docs/ru/…` |
| `/static/…` | `web/static/{css,js,img,fonts}` |

Redirects: `/` → `/docs/`; `/en/docs/…` → `/docs/…`.

Local preview: `bun run vercel:build && npx serve public`

For a **Go binary** on your own server, use `make deploy` instead.

## Troubleshooting

### `listen tcp ... bind: Only one usage of each socket address`

Another process (often a previous `go run`) is still bound to that port. Stop it or use another port:

```bash
export APP_BIND=127.0.0.1:8081
bun run go
```

On Windows, find and end the listener, for example: `netstat -ano | findstr :8080` then `taskkill /PID <pid> /F`.

### Static files 404

Run from the repo root (or use `bun run go`), run `bun run build:css`, and ensure `web/static` exists. See `APP_STATIC_DIR` in the table above.

## Project layout

| Path | Role |
|------|------|
| [`cmd/server/main.go`](cmd/server/main.go) | Composition root: config, locales, health, site feature |
| [`internal/serverapp/`](internal/serverapp/) | Shared app assembly for local/server deploy |
| [`scripts/vercel-static-export.mjs`](scripts/vercel-static-export.mjs) | Maps docgen output → `public/` for Vercel |
| [`internal/site/`](internal/site/) | HTTP routes: `GET /`, `GET /sample`, `GET /docs/...` (component gallery) |
| [`internal/fixtures/locale/`](internal/fixtures/locale/) | Embedded JSON copy per locale |
| [`internal/views/`](internal/views/) | `templ` pages, [`layout.templ`](internal/views/layout.templ) (`SiteShell` + UI8Kit `Shell`), [`partials/header_trailing.templ`](internal/views/partials/header_trailing.templ) (language switch) |
| [`internal/ui/components/`](internal/ui/components/) | App-level UI pieces (e.g. `components/toggles` language control) |
| [`web/static/`](web/static/) | `app.css` (Tailwind build), `css/latty-icons.css` (mask SVG hooks for `Icon`), `css/tweakcn.css`, `css/gfonts.css`, `fonts/google-sans/*`, `js/theme.js`, `js/ui8kit.js` |

## Verification

```bash
bun run verify
```

`bun run verify` runs `playwright install chromium` before e2e tests (no-op when browsers are cached). To pre-download only: `bun run test:e2e:install`.

`verify` runs: `templ generate` → Tailwind `build:css` → `ui8px lint` → **`ui8px validate aria`** (see `web/static/js/manifest.json`) → **Nu HTML** on `.validate/html-snapshots/nu/**/*.html` (network; skips if none) → **`go test ./...`** → **Playwright + axe** on `/` (axe runs without `color-contrast` until brand phase—see rule file). Use **`go run ./.validate/cmd/snapshot-render`** to refresh `html-snapshots/generated/`; copy conforming files into **`html-snapshots/nu/`** for the Nu gate.

**W3C CSS Validator** is intentionally not part of `verify` until the brand phase (see `.cursor/rules/fastygo-ui-validation-testing.mdc`).

## Cloning this template

1. Copy the repository (or subtree: `cmd/server`, `internal/site`, `internal/ui`, `internal/views`, `internal/fixtures`, `web/static`, `package.json`, `.ui8px`, `.validate`, `tests`, `playwright.config.ts`, `scripts`).
2. Change the Go module path in `go.mod` and imports.
3. Add routes in [`internal/site/feature.go`](internal/site/feature.go) and nav labels in [`internal/fixtures/locale/`](internal/fixtures/locale/).
4. Add or extend `templ` under [`internal/views/`](internal/views/) (and Tailwind `@source` in [`web/static/css/input.css`](web/static/css/input.css) for new paths).

## Roadmap

- Landing and marketing shell when you are ready (current default remains sidebar chrome).
- Docs and component examples on top of the same shell.

The [`.fastygo/`](.fastygo/) directory in some workspaces is reference-only and is **not** imported at runtime; this module builds standalone.
