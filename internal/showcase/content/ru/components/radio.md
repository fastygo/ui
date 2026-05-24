---
slug: radio
section: components
title: "Radio"
description: "Radio-кнопка для выбора одного варианта."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Чекбокс"
    href: /docs/primitives/checkbox/
  - label: "Switch"
    href: /docs/primitives/switch/
api:
  - name: "Name"
    type: "string"
    description: "Имя поля формы"
  - name: "Value"
    type: "string"
    description: "Значение варианта"
  - name: "Checked"
    type: "bool"
    description: "Начальное состояние checked"
---

Radio-кнопка для выбора одного варианта.

## По умолчанию

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Radio(ui.RadioProps{Name: "plan", Value: "free", AriaLabel: "Бесплатный план"})
}
```

## Отмечен

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Radio(ui.RadioProps{Name: "plan", Value: "pro", Checked: true, AriaLabel: "Pro-план"})
}
```
