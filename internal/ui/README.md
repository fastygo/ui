# In-app UI registry (`internal/ui`)

Staging area for the FastyGoUI showcase (shadcn-style structure lab). **Single source of truth during development** — do not depend on `github.com/fastygo/blocks` or `github.com/fastygo/widgets` until the set is frozen and extracted.

## Layout

```
internal/ui/
  layout/       # App shell — stays in github.com/fastygo/ui after extraction
  components/   # Small showcase-only UI (icon, toggles, …)
  blocks/       # Section organisms → later github.com/fastygo/blocks
  widgets/      # UI + behavior/API → later github.com/fastygo/widgets
  variants/     # Optional wireframe utility maps
  utils/        # Thin helpers on github.com/fastygo/templ/utils
```

## Atoms & molecules (external)

- Primitives: `import "github.com/fastygo/templ/ui"` → `@ui.*`
- Composites: `import cmp "github.com/fastygo/templ/components"` → `@cmp.*`
- Helpers: `github.com/fastygo/templ/utils` (tags, Cn, ARIA, CVA)

## Extraction (later)

| Path | Future module |
|------|----------------|
| `blocks/*` | `github.com/fastygo/blocks` |
| `widgets/*` | `github.com/fastygo/widgets` |
| `layout/` | Stays in this app |

See `.cursor/rules/fastygo-ui-design-system-registry.mdc`.
