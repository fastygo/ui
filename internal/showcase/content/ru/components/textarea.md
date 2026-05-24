---
slug: textarea
section: components
title: "Многострочное поле"
description: "Многострочный текстовый ввод."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Поле ввода"
    href: /docs/primitives/input/
  - label: "Форма"
    href: /docs/components/form/
api:
  - name: "Placeholder"
    type: "string"
    description: "Текст placeholder"
  - name: "Rows"
    type: "int"
    description: "Число видимых строк"
  - name: "Disabled"
    type: "bool"
    description: "Отключает элемент"
---

Многострочный текстовый ввод.

## По умолчанию

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Textarea(ui.TextareaProps{Placeholder: "Ваше сообщение"})
}
```

## Отключено

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Textarea(ui.TextareaProps{Placeholder: "Отключено", Disabled: true})
}
```
