---
slug: toast
section: components
title: "Toast"
description: "Transient notification wireframe."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Alert"
    href: /docs/components/alert/
  - label: "Dialog"
    href: /docs/components/dialog/
api:
  - name: "Class"
    type: "string"
    description: "Card-like surface"
---

Transient notification wireframe.

## Default

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Box(ui.BoxProps{Class: "flex max-w-sm items-center justify-between gap-4 rounded-lg border border-border bg-card p-4 shadow"}) {
		@ui.Text(ui.TextProps{Class: "text-sm"}, "Saved successfully.")
		@ui.Button(ui.ButtonProps{Variant: "outline", Size: "sm"}) {
			Undo
		}
	}
}
```
