# Docs showcase (static compiler)

Documentation is built at compile time from localized Markdown with `templ Example()` preview fences.

## Layout

```
internal/showcase/
  content/
    en/**/*.md          # English source
    ru/**/*.md          # Russian overrides (fallback to en when missing)
  docgen/               # Loader, parser, preview compiler, HTML builder
cmd/docgen/             # Static HTML generator CLI
```

## Commands

- `bun run docs:build` — write `web/static/docs/` (deploy artifact)
- `bun run docs:snapshot` — write Nu validation snapshots under `.validate/html-snapshots/nu/docs/`

## Markdown contract

- YAML front matter: `slug`, `section`, `title`, `description`, optional `source`, `package`, `related`, `api`
- Body: prose + ` ```templ ` fences with `templ Example() { … }` for live previews
- Other fence languages (`go`, `bash`, …) render as plain code blocks
- No Tailwind classes or raw HTML in prose (code fences exempt)

## Serving

When `web/static/docs/` exists, `internal/site` serves `/docs/` and `/ru/docs/` as static HTML.

## Real components

Implement UI in `internal/ui/components/` (e.g. `blogcard/`). Document in Markdown with self-contained `templ Example()` snippets.
