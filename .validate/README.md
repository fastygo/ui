# Validation layout (FastyGoUI)

This directory holds **inputs and docs** for standards checks—not application runtime code.

| Path | Purpose |
|------|---------|
| [`cmd/`](./cmd/) | `snapshot-render` — Go CLI for HTML drafts (see [`cmd/README.md`](./cmd/README.md)). |
| [`html-snapshots/`](./html-snapshots/) | **`generated/`** — `go run ./.validate/cmd/snapshot-render`; **`nu/`** — only these files are checked by `bun run validate:html` (W3C Nu). |
| [`apg/`](./apg/) | **Manual** APG-oriented checklists per widget (roles, keyboard, states)—run when adding or changing interactive patterns. |
| [`docs/`](./docs/) | Templates for **props / examples / ARIA** notes per component. |

**W3C CSS Validator** is intentionally **out of scope** until the **brand** phase (Tailwind output and token layers are noisy for a strict CSS validator gate).

See `.cursor/rules/fastygo-ui-validation-testing.mdc` for the full policy.
