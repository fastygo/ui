---
slug: tooltip
section: components
title: "Tooltip"
description: "Подсказка при hover/focus (wireframe)."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Popover"
    href: /docs/components/popover/
  - label: "Кнопка"
    href: /docs/primitives/button/
api:
  - name: "Role"
    type: "string"
    description: "tooltip на блоке подсказки"
---

Подсказка при hover/focus (wireframe).

## По умолчанию

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Stack(ui.StackProps{Class: "gap-2 items-start"}) {
		@ui.Button(ui.ButtonProps{Attrs: templ.Attributes{"aria-describedby": "tip-demo"}}) {
			Наведите курсор
		}
		@ui.Box(ui.BoxProps{Class: "rounded border border-border bg-popover px-2 py-1 text-xs", Attrs: templ.Attributes{"id": "tip-demo", "role": "tooltip"}}) {
			@ui.Text(ui.TextProps{}, "Текст подсказки")
		}
	}
}
```
