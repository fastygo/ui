---
slug: progress
section: components
title: "Progress"
description: "Progress indicator wireframe."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Skeleton"
    href: /docs/components/skeleton/
  - label: "Slider"
    href: /docs/components/slider/
api:
  - name: "Role"
    type: "string"
    description: "progressbar"
  - name: "AriaValuenow"
    type: "string"
    description: "Current value"
---

Progress indicator wireframe.

## Default

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Box(ui.BoxProps{
		Class: "h-2 w-full max-w-xs overflow-hidden rounded-full bg-muted",
		Attrs: templ.Attributes{
			"role":          "progressbar",
			"aria-valuenow": "60",
			"aria-valuemin": "0",
			"aria-valuemax": "100",
		},
	}) {
		@ui.Box(ui.BoxProps{Class: "h-full w-3/5 bg-primary"})
	}
}
```
