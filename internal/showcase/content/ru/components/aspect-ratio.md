---
slug: aspect-ratio
section: components
title: "Соотношение сторон"
description: "Контейнер с фиксированным aspect ratio."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Карусель"
    href: /docs/components/carousel/
  - label: "Карточка"
    href: /docs/components/card/
api:
  - name: "Class"
    type: "string"
    description: "aspect-video | aspect-square"
---

Контейнер с фиксированным aspect ratio.

## По умолчанию

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
