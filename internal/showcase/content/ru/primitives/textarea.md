---
slug: textarea
section: primitives
title: "Textarea"
description: "Нативное многострочное текстовое поле."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Поле ввода"
    href: /docs/primitives/input/
  - label: "Метка"
    href: /docs/primitives/label/
  - label: "Форма"
    href: /docs/components/form/
api:
  - name: "Rows"
    type: "int"
    description: "Число видимых строк (по умолчанию 4)"
  - name: "Name"
    type: "string"
    description: "Имя поля формы"
  - name: "Placeholder"
    type: "string"
    description: "Подсказка в пустом поле"
  - name: "Disabled"
    type: "bool"
    description: "Неактивное состояние"
---

Textarea рендерит нативное многострочное текстовое поле. Rows по умолчанию равен четырём, если не задан.

## По умолчанию

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Textarea(ui.TextareaProps{
		Name:        "message",
		Placeholder: "Ваше сообщение",
	})
}
```

## Неактивное

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Textarea(ui.TextareaProps{
		Disabled: true,
		Name:     "locked",
		Value:    "Фиксированный текст",
	})
}
```

## Пользовательские строки

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Textarea(ui.TextareaProps{
		Rows:        6,
		Name:        "bio",
		Placeholder: "О себе",
	})
}
```
