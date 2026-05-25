---
slug: group
section: primitives
title: "Group"
description: "Горизонтальный flex-ряд; поддерживает fieldset для связанных контролов."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Стек"
    href: /docs/primitives/stack/
  - label: "Кнопка"
    href: /docs/primitives/button/
  - label: "Радиокнопка"
    href: /docs/primitives/radio/
api:
  - name: "Tag"
    type: "string"
    description: "div | fieldset | section | nav | …"
  - name: "Class"
    type: "string"
    description: "Дополнительные utility-классы"
  - name: "Attrs"
    type: "templ.Attributes"
    description: "Дополнительные HTML-атрибуты"
---

Group располагает дочерние элементы в горизонтальном flex-ряду. Поддерживает fieldset для связанных контролов формы.

## Кнопки в ряд

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Group(ui.GroupProps{Class: "gap-2"}) {
		@ui.Button(ui.ButtonProps{Variant: "default"}) { Сохранить }
		@ui.Button(ui.ButtonProps{Variant: "outline"}) { Отмена }
	}
}
```

## Fieldset

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Group(ui.GroupProps{Tag: "fieldset", Class: "gap-4"}) {
		@ui.Radio(ui.RadioProps{Name: "plan", Value: "free", ID: "plan-free"})
		@ui.Radio(ui.RadioProps{Name: "plan", Value: "pro", ID: "plan-pro"})
	}
}
```
