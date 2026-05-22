---
slug: carousel
section: components
title: "Carousel"
description: "Horizontal scrolling list."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Aspect Ratio"
    href: /docs/components/aspect-ratio/
  - label: "Card"
    href: /docs/components/card/
api:
  - name: "Group"
    type: "Group"
    description: "Horizontal flex row"
---

Horizontal scrolling list.

## Default

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Group(ui.GroupProps{Class: "flex gap-2 overflow-x-auto max-w-md"}) {
		@ui.Box(ui.BoxProps{Class: "min-w-[8rem] shrink-0 rounded-lg border border-border bg-card p-4 text-sm"}) {
			@ui.Text(ui.TextProps{}, "Slide 1")
		}
		@ui.Box(ui.BoxProps{Class: "min-w-[8rem] shrink-0 rounded-lg border border-border bg-card p-4 text-sm"}) {
			@ui.Text(ui.TextProps{}, "Slide 2")
		}
		@ui.Box(ui.BoxProps{Class: "min-w-[8rem] shrink-0 rounded-lg border border-border bg-card p-4 text-sm"}) {
			@ui.Text(ui.TextProps{}, "Slide 3")
		}
	}
}
```
