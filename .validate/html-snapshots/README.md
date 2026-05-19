# HTML snapshots

## `generated/` (draft)

- **`go run ./.validate/cmd/snapshot-render`** writes here (see `-route`, `-list`, `-lang`, `-out`).
- Draft files are **not** run through `bun run validate:html` by default.

## `nu/` (Nu gate)

- Put only HTML you want **enforced in CI**: copy from `generated/` when Nu-clean, or hand-curate.
- **`bun run validate:html`** walks **only** `.validate/html-snapshots/nu/**/*.html` (W3C Nu, network). Empty tree → skip.

## Document requirements

- Full documents: `<!doctype html>`, `<html>`, `<head>`, `<body>`.
