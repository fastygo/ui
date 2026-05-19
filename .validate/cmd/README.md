# snapshot-render

Go entrypoint: **`go run ./.validate/cmd/snapshot-render`** (from repo root).

- **`-list`** — supported `-route` keys.
- **`-route`** — comma-separated: `home`, `sample`.
- **`-lang`** — fixture locale (`en`, `ru`, …), default `en`.
- **`-out`** — output directory, default `.validate/html-snapshots/generated`.

Draft HTML is written under **`html-snapshots/generated/`** (run **`go run ./.validate/cmd/snapshot-render -route=home`** — **`-list` only prints keys and creates nothing**). Copy conforming files to **`html-snapshots/nu/`** for `bun run validate:html`.

npm/bun: **`bun run snapshot:html -- -route=home`**
