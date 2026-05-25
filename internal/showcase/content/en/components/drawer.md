---
slug: drawer
section: components
title: "Drawer"
description: "Bottom sheet drawer wireframe."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Alert Dialog"
    href: /docs/components/alert-dialog/
  - label: "Dialog"
    href: /docs/components/dialog/
api:
  - name: "Class"
    type: "string"
    description: "Bottom sheet surface"
---

Bottom sheet drawer wireframe.

## Default

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Box(ui.BoxProps{Class: "w-full max-w-md rounded-t-xl border border-border bg-card p-4"}) {
		@ui.Box(ui.BoxProps{Class: "mx-auto mb-3 h-1 w-10 rounded-full bg-muted"})
		@ui.Text(ui.TextProps{}, "Drawer content.")
	}
}
```
