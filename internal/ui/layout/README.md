# App shell layout (`registry:layout`)

Application-owned structural shell (ported from UI8Kit `layout` for the templ registry migration).

- **Shell** — document, mobile sheet nav, desktop sidebar, sticky header, main slot
- **Sidebar** — nav links with latty icons (`internal/ui/components/icon`)
- **Header** — menu trigger, title, theme toggle, trailing slot (language)

Uses `github.com/fastygo/templ/ui` for primitives and preserves `data-ui8kit-*` hooks for `theme.js` / `ui8kit.js` / `@ui8kit/aria`.
