---
slug: popover
section: components
title: "Popover"
description: "Плавающая панель с контентом."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Dropdown Menu"
    href: /docs/components/dropdown-menu/
  - label: "Tooltip"
    href: /docs/components/tooltip/
api:
  - name: "Class"
    type: "string"
    description: "Утилиты плавающей панели"
---

Плавающая панель с контентом.

## По умолчанию

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Stack(ui.StackProps{Class: "gap-2"}) {
		@ui.Button(ui.ButtonProps{Variant: "outline"}) {
			Открыть popover
		}
		@ui.Box(ui.BoxProps{Class: "w-56 rounded-lg border border-border bg-card p-3 text-sm shadow"}) {
			@ui.Text(ui.TextProps{}, "Текст внутри popover.")
		}
	}
}
```
