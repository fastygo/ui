# Component: `<name>`

## Purpose

One sentence: what user problem this solves.

## Props / data (Go or templ)

| Field / prop | Type | Required | Notes |
|----------------|------|----------|--------|
| | | | |

## Wireframe example

Where the live prototype lives (`internal/ui/…`, `internal/views/…`, or showcase route).

## Markup & ARIA

- **Roles:**
- **Required `aria-*` / labels:** (static via `github.com/fastygo/base/utils/aria.go` where applicable)
- **`data-ui8kit` / behavior:** pattern name; listed in `web/static/js/manifest.json` if JS behavior applies.

## Keyboard (APG)

Bullet list of expected keys and focus behavior.

## Tests

- [ ] `go test` render / snapshot (if applicable)
- [ ] `bun run test:e2e` (axe) — **wireframe:** serious/critical except `color-contrast` (disabled until brand); **brand:** remove `color-contrast` disable in `tests/e2e/*.spec.ts`
- [ ] `bun run validate:aria` after manifest or markup changes
- [ ] Nu HTML — copy snapshot into `.validate/html-snapshots/nu/` and run `bun run validate:html`

## References

- Link to APG pattern
- Link to `@ui8kit/aria` pattern doc if any
