---
slug: switch
section: components
title: "Switch"
description: "Переключатель toggle."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Чекбокс"
    href: /docs/components/checkbox/
  - label: "Toggle"
    href: /docs/components/toggle/
api:
  - name: "Name"
    type: "string"
    description: "Имя поля формы"
  - name: "Checked"
    type: "bool"
    description: "Начальное состояние «включено»"
  - name: "Disabled"
    type: "bool"
    description: "Отключает элемент"
---

Переключатель toggle.

## По умолчанию

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Switch(ui.SwitchProps{Name: "airplane", AriaLabel: "Режим полёта"})
}
```

## Включён

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Switch(ui.SwitchProps{Name: "airplane", Checked: true, AriaLabel: "Режим полёта включён"})
}
```
