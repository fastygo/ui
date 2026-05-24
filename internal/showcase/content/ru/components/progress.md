---
slug: progress
section: components
title: "Прогресс"
description: "Wireframe индикатора прогресса."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Слайдер"
    href: /docs/components/slider/
  - label: "Соотношение сторон"
    href: /docs/components/aspect-ratio/
api:
  - name: "Role"
    type: "string"
    description: "progressbar"
  - name: "AriaValuenow"
    type: "string"
    description: "Текущее значение"
---

Wireframe индикатора прогресса.

## По умолчанию

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
