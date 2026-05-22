---
slug: calendar
section: components
title: "Calendar"
description: "Date grid placeholder."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Command"
    href: /docs/components/command/
  - label: "Popover"
    href: /docs/components/popover/
api:
  - name: "Class"
    type: "string"
    description: "Calendar frame utilities"
---

Date grid placeholder.

## Default

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Box(ui.BoxProps{Class: "w-64 rounded-lg border border-border p-3"}) {
		@ui.Text(ui.TextProps{Class: "text-sm font-medium"}, "May 2026 — placeholder")
		@ui.Grid(ui.GridProps{Class: "mt-2 grid-cols-7 gap-1 text-center text-xs text-muted-foreground"}) {
			@ui.Text(ui.TextProps{}, "1")
			@ui.Text(ui.TextProps{}, "2")
			@ui.Text(ui.TextProps{}, "3")
			@ui.Text(ui.TextProps{}, "4")
			@ui.Text(ui.TextProps{}, "5")
			@ui.Text(ui.TextProps{}, "6")
			@ui.Text(ui.TextProps{}, "7")
		}
	}
}
```
