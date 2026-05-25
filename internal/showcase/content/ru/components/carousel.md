---
slug: carousel
section: components
title: "Carousel"
description: "Горизонтальный прокручиваемый список."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Соотношение сторон"
    href: /docs/components/aspect-ratio/
  - label: "Карточка"
    href: /docs/components/card/
api:
  - name: "Group"
    type: "Group"
    description: "Горизонтальный flex-ряд"
---

Горизонтальный прокручиваемый список.

## По умолчанию

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Group(ui.GroupProps{Class: "flex gap-2 overflow-x-auto max-w-md"}) {
		@ui.Box(ui.BoxProps{Class: "min-w-[8rem] shrink-0 rounded-lg border border-border bg-card p-4 text-sm"}) {
			@ui.Text(ui.TextProps{}, "Слайд 1")
		}
		@ui.Box(ui.BoxProps{Class: "min-w-[8rem] shrink-0 rounded-lg border border-border bg-card p-4 text-sm"}) {
			@ui.Text(ui.TextProps{}, "Слайд 2")
		}
		@ui.Box(ui.BoxProps{Class: "min-w-[8rem] shrink-0 rounded-lg border border-border bg-card p-4 text-sm"}) {
			@ui.Text(ui.TextProps{}, "Слайд 3")
		}
	}
}
```
