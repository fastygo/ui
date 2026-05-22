---
slug: alert-dialog
section: components
title: "Alert Dialog"
description: "Modal that interrupts flow."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Dialog"
    href: /docs/components/dialog/
  - label: "Alert"
    href: /docs/components/alert/
api:
  - name: "Title"
    type: "Title"
    description: "Alert heading"
---

Modal that interrupts flow.

## Default

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Box(ui.BoxProps{Class: "max-w-sm rounded-lg border border-border bg-card p-4"}) {
		@ui.Title(ui.TitleProps{Order: 3, Class: "text-base font-semibold"}, "Are you sure?")
		@ui.Text(ui.TextProps{Class: "text-sm text-muted-foreground"}, "This action cannot be undone.")
		@ui.Group(ui.GroupProps{Class: "mt-4 flex justify-end gap-2"}) {
			@ui.Button(ui.ButtonProps{Variant: "outline"}) {
				Cancel
			}
			@ui.Button(ui.ButtonProps{Variant: "destructive"}) {
				Delete
			}
		}
	}
}
```
