---
slug: skeleton
section: components
title: "Skeleton"
description: "Loading placeholder blocks."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Progress"
    href: /docs/components/progress/
  - label: "Card"
    href: /docs/components/card/
api:
  - name: "Class"
    type: "string"
    description: "animate-pulse bg-muted shapes"
---

Loading placeholder blocks.

## Default

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Stack(ui.StackProps{Class: "gap-2 max-w-md"}) {
		@ui.Box(ui.BoxProps{Class: "h-4 w-3/4 max-w-xs animate-pulse rounded bg-muted"})
		@ui.Box(ui.BoxProps{Class: "h-4 w-1/2 max-w-xs animate-pulse rounded bg-muted"})
	}
}
```
