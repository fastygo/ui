---
slug: label
section: primitives
title: "Label"
description: "Связывает видимый текст с одним контролом формы."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Поле ввода"
    href: /docs/primitives/input/
  - label: "Чекбокс"
    href: /docs/primitives/checkbox/
  - label: "Форма"
    href: /docs/components/form/
api:
  - name: "HTMLFor"
    type: "string"
    description: "ID связанного контрола (атрибут for)"
  - name: "Class"
    type: "string"
    description: "Дополнительные utility-классы"
  - name: "Attrs"
    type: "templ.Attributes"
    description: "Дополнительные HTML-атрибуты"
---

Label связывает видимый текст с одним контролом формы. Атрибут for задаётся через HTMLFor.

## С контролом

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Label(ui.LabelProps{HTMLFor: "demo-email"}) {
		Email
	}
}
```

## С полем ввода

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Stack(ui.StackProps{Class: "gap-2"}) {
		@ui.Label(ui.LabelProps{HTMLFor: "demo-email"}) {
			Email
		}
		@ui.Input(ui.InputProps{ID: "demo-email", Type: "email", Name: "email"})
	}
}
```
