---
slug: introduction
section: getting-started
title: "Introduction"
description: "FastyGo UI is a wireframe-first component gallery on Go and templ. Pages show live previews, copy-pasteable templ snippets, and prop tables—similar to shadcn/ui, without shipping a monolithic UI package."
source: github.com/fastygo/ui
package: internal/site
related:
  - label: "Installation"
    href: /docs/installation/
  - label: "Components index"
    href: /docs/
---

FastyGo UI is a wireframe-first component gallery on Go and templ. Pages show live previews, copy-pasteable templ snippets, and prop tables—similar to shadcn/ui, without shipping a monolithic UI package.

## Wireframe scope

Structure, semantics, and accessibility come first; visual brand polish is a later phase.

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Stack(ui.StackProps{Class: "gap-2"}) {
		@ui.Title(ui.TitleProps{Order: 2}, "Component gallery")
		@ui.Text(ui.TextProps{Class: "text-sm text-muted-foreground"}, "Wireframe previews and API tables.")
	}
}
```
