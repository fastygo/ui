---
slug: aspect-ratio
section: components
title: "Aspect Ratio"
description: "Fixed aspect container."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Carousel"
    href: /docs/components/carousel/
  - label: "Card"
    href: /docs/components/card/
api:
  - name: "Class"
    type: "string"
    description: "aspect-video | aspect-square"
---

Fixed aspect container.

## Default

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Box(ui.BoxProps{Class: "aspect-video w-full max-w-xs overflow-hidden rounded-lg border border-border bg-muted"}) {
		@ui.Box(ui.BoxProps{Class: "flex h-full items-center justify-center text-sm text-muted-foreground"}) {
			@ui.Text(ui.TextProps{}, "16:9")
		}
	}
}
```
