# APG checklists (manual)

Use this when adding or changing **interactive** patterns (dialog, sheet, tabs, combobox, menu, switch, …).

## Per-pattern record

Create a short markdown file under `apg/` (e.g. `apg/dialog-sheet.md`) with:

- **APG reference** — link to the relevant WAI-ARIA Authoring Practices Guide pattern.
- **Roles & properties** — required roles, `aria-*`, labels, `aria-controls` / `aria-expanded` pairs.
- **Keyboard** — tab order, Escape, arrow keys, Enter/Space as specified.
- **Focus** — initial focus, return focus on close, focus trap expectations.
- **Our markup contract** — `data-ui8kit` values, `@ui8kit/aria` pattern name, manifest entry in `web/static/js/manifest.json`.

**Automation complements this:** run `bun run test:e2e` (axe) and `bun run validate:aria`; they do **not** replace keyboard review.
