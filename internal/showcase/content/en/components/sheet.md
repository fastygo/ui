---
slug: sheet
section: components
title: "Sheet"
description: "Slide-over panel wireframe."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Dialog"
    href: /docs/components/dialog/
  - label: "Drawer"
    href: /docs/components/drawer/
api:
  - name: "Class"
    type: "string"
    description: "Panel surface utilities"
---

Slide-over panel wireframe.

## Default

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Stack(ui.StackProps{Class: "gap-2 max-w-md"}) {
		@ui.Box(ui.BoxProps{Class: "w-64 rounded-l-lg border border-border bg-card p-4"}) {
			@ui.Title(ui.TitleProps{Order: 3, Class: "text-sm font-semibold"}, "Sheet")
			@ui.Text(ui.TextProps{Class: "text-sm text-muted-foreground"}, "Side panel content.")
		}
	}
}
```
