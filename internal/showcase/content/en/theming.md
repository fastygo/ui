---
slug: theming
section: getting-started
title: "Theming"
description: "Semantic tokens (background, foreground, primary, …) come from tweakcn-style CSS variables. The header theme toggle switches light/dark via theme.js."
source: web/static/css/tweakcn.css
package: web/static/js/theme.js
related:
  - label: "Introduction"
    href: /docs/introduction/
---

Semantic tokens (background, foreground, primary, …) come from tweakcn-style CSS variables. The header theme toggle switches light/dark via theme.js.

## Semantic tokens

Use token utilities such as bg-background and text-foreground in templ class strings.

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Group(ui.GroupProps{Class: "gap-2"}) {
		@ui.Box(ui.BoxProps{Class: "inline-block h-10 w-16 rounded-md border border-border bg-background"})
		@ui.Box(ui.BoxProps{Class: "inline-block h-10 w-16 rounded-md bg-primary"})
		@ui.Box(ui.BoxProps{Class: "inline-block h-10 w-16 rounded-md bg-muted"})
	}
}
```

```go
<body class="bg-background text-foreground">
  <!-- theme.js toggles .dark on html -->
</body>
```
