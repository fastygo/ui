# Docs showcase (static compiler)

Documentation is built at compile time from localized Markdown plus a Go preview registry.

## Layout

```
internal/showcase/
  content/
    en/**/*.md          # English source (generated via migrate:docs or hand-edited)
    ru/**/*.md          # Russian overrides (fallback to en when missing)
  docgen/               # Loader, parser, validator, HTML builder
  previews/             # Demo ID → templ.Component registry (+ bridge from catalog)
  catalog/              # Legacy runtime registry (still used for preview registration)
  register.go           # blank-import catalog packages for init registration
cmd/docgen/             # Static HTML generator CLI
cmd/migrate-docs/       # One-shot MD export from registry.Page
```

## Commands

- `bun run migrate:docs` — regenerate English Markdown from `catalog/*/doc.go`
- `bun run docs:build` — write `web/static/docs/` (deploy artifact)
- `bun run docs:snapshot` — write Nu validation snapshots under `.validate/html-snapshots/nu/docs/`

## Markdown contract

- YAML front matter: `slug`, `section`, `title`, `description`, optional `source`, `package`, `related`, `api`
- Body: prose + `{{demo id="slug.variant"}}` + fenced code block (templ/go)
- No Tailwind classes or raw HTML in prose (code fences exempt)

## Serving

When `web/static/docs/` exists, `internal/site` serves `/docs/` and `/ru/docs/` as static HTML. Runtime templ docs handlers are skipped.

## Real components

Implement UI in `internal/ui/components/` (e.g. `blogcard/`). Document in Markdown; register previews via `previews.Register` or catalog bridge (`previews.RegisterFromRegistry`).
