---
slug: input
section: primitives
title: "Поле ввода"
description: "Однострочное нативное текстовое поле."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Метка"
    href: /docs/primitives/label/
  - label: "Многострочное поле"
    href: /docs/primitives/textarea/
  - label: "Форма"
    href: /docs/components/form/
api:
  - name: "Type"
    type: "string"
    description: "text | email | password | number | search | tel | url"
  - name: "Name"
    type: "string"
    description: "Имя поля формы"
  - name: "Placeholder"
    type: "string"
    description: "Подсказка в пустом поле"
  - name: "Disabled"
    type: "bool"
    description: "Неактивное состояние"
  - name: "Required"
    type: "bool"
    description: "Обязательное поле"
---

Input рендерит однострочное нативное текстовое поле. Используйте для email, поиска и коротких значений.

## По умолчанию

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Input(ui.InputProps{
		Name:        "email",
		Placeholder: "name@example.com",
	})
}
```

## Email

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Input(ui.InputProps{
		ID:   "demo-email",
		Type: "email",
		Name: "email",
	})
}
```

## Малый размер

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Input(ui.InputProps{
		Size:        "sm",
		Name:        "q",
		Placeholder: "Поиск",
	})
}
```

## Неактивное

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Input(ui.InputProps{
		Disabled: true,
		Name:     "locked",
		Value:    "Только чтение",
	})
}
```
