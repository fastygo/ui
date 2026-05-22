---
slug: input
section: components
title: "Поле ввода"
description: "Текстовое поле ввода."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Многострочное поле"
    href: /docs/components/textarea/
  - label: "Форма"
    href: /docs/components/form/
api:
  - name: "Type"
    type: "string"
    description: "HTML-тип input"
  - name: "Placeholder"
    type: "string"
    description: "Текст placeholder"
  - name: "Disabled"
    type: "bool"
    description: "Отключает элемент"
---

Текстовое поле ввода.

## По умолчанию

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Input(ui.InputProps{Placeholder: "Email"})
}
```

## Отключено

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Input(ui.InputProps{Placeholder: "Отключено", Disabled: true})
}
```

## Файл

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Input(ui.InputProps{Type: "file"})
}
```
