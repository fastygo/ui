---
slug: tooltip
section: components
title: "Tooltip"
description: "Hint on hover/focus (wireframe)."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Popover"
    href: /docs/components/popover/
  - label: "Button"
    href: /docs/primitives/button/
api:
  - name: "Role"
    type: "string"
    description: "tooltip on hint box"
---

Hint on hover/focus (wireframe).

## Default

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Stack(ui.StackProps{Class: "gap-2 items-start"}) {
		@ui.Button(ui.ButtonProps{Attrs: templ.Attributes{"aria-describedby": "tip-demo"}}) {
			Hover me
		}
		@ui.Box(ui.BoxProps{Class: "rounded border border-border bg-popover px-2 py-1 text-xs", Attrs: templ.Attributes{"id": "tip-demo", "role": "tooltip"}}) {
			@ui.Text(ui.TextProps{}, "Tooltip text")
		}
	}
}
```
