---
slug: hover-card
section: components
title: "Hover Card"
description: "Rich preview on hover."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Popover"
    href: /docs/components/popover/
  - label: "Tooltip"
    href: /docs/components/tooltip/
api:
  - name: "Box"
    type: "Box"
    description: "Preview card surface"
---

Rich preview on hover.

## Default

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Group(ui.GroupProps{Class: "flex items-start gap-3"}) {
		@ui.Button(ui.ButtonProps{Variant: "link"}) {
			{ "@user" }
		}
		@ui.Box(ui.BoxProps{Class: "w-56 rounded-lg border border-border bg-card p-3 text-sm"}) {
			@ui.Text(ui.TextProps{}, "Profile preview card.")
		}
	}
}
```
