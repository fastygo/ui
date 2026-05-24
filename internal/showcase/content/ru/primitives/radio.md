---
slug: radio
section: primitives
title: "Радиокнопка"
description: "Один вариант в именованной группе radio; общий Name."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Чекбокс"
    href: /docs/primitives/checkbox/
  - label: "Группа"
    href: /docs/primitives/group/
  - label: "Метка"
    href: /docs/primitives/label/
api:
  - name: "Name"
    type: "string"
    description: "Общее имя группы"
  - name: "Value"
    type: "string"
    description: "Значение варианта"
  - name: "Checked"
    type: "bool"
    description: "Выбранный вариант в группе"
  - name: "Disabled"
    type: "bool"
    description: "Неактивное состояние"
---

Radio рендерит один вариант в именованной группе. Все радиокнопки группы используют одно значение Name.

## Группа

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Stack(ui.StackProps{Class: "gap-2"}) {
		@ui.Radio(ui.RadioProps{Name: "plan", Value: "free", ID: "plan-free", Checked: true})
		@ui.Radio(ui.RadioProps{Name: "plan", Value: "pro", ID: "plan-pro"})
	}
}
```

## Неактивная

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Radio(ui.RadioProps{Name: "plan", Value: "pro", ID: "plan-pro", Disabled: true})
}
```
