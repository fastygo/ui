# `registry:components`

**Presentation-only** building blocks for the showcase app. Props in, templ out — **no** HTTP, domain logic, or caching.

## Current

| Package | Role |
|---------|------|
| `icon/` | Latty mask icons (`latty latty-*`) |
| `toggles/` | Language toggle (framework `view.LanguageToggleData`) |

## Heuristic

If it only formats props and renders markup → **`components`**. If it fetches or orchestrates side effects → **`widgets`**.
