---
slug: separator
section: components
title: "Разделитель"
description: "Визуальный разделитель между секциями."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Stack"
    href: /docs/primitives/stack/
  - label: "Карточка"
    href: /docs/components/card/
api:
  - name: "Class"
    type: "string"
    description: "Обычно h-px bg-border"
  - name: "Role"
    type: "string"
    description: "separator"
---

Визуальный разделитель между секциями.

## По умолчанию

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Stack(ui.StackProps{Class: "gap-2 max-w-md"}) {
		@ui.Text(ui.TextProps{}, "Сверху")
		@ui.Box(ui.BoxProps{Class: "h-px w-full bg-border", Attrs: templ.Attributes{"role": "separator"}})
		@ui.Text(ui.TextProps{}, "Снизу")
	}
}
```
