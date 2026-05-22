---
slug: installation
section: getting-started
title: "Installation"
description: "Use this repository as the reference app, or copy individual components from the gallery into your project (future fastygo add workflow)."
source: github.com/fastygo/ui
package: cmd/server
related:
  - label: "Theming"
    href: /docs/theming/
  - label: "Button"
    href: /docs/components/button/
---

Use this repository as the reference app, or copy individual components from the gallery into your project (future fastygo add workflow).

## From git

Clone the app, install Bun for CSS, run templ generate, then start the server.

```go
bun install
go mod download
bun run build:css
go tool templ generate ./...
bun run go
```
