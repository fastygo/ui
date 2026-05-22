---
slug: checkbox
section: components
title: "Чекбокс"
description: "Логический checkbox."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Radio"
    href: /docs/components/radio/
  - label: "Switch"
    href: /docs/components/switch/
api:
  - name: "Name"
    type: "string"
    description: "Имя поля формы"
  - name: "Checked"
    type: "bool"
    description: "Начальное состояние checked"
  - name: "Disabled"
    type: "bool"
    description: "Отключает элемент"
---

Логический checkbox.

## По умолчанию

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Checkbox(ui.CheckboxProps{Name: "terms", AriaLabel: "Принять условия"})
}
```

## Отмечен

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Checkbox(ui.CheckboxProps{Name: "terms", Checked: true, AriaLabel: "Условия приняты"})
}
```
