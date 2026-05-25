---
slug: select
section: primitives
title: "Select"
description: "Нативный dropdown из слайса Options."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Метка"
    href: /docs/primitives/label/
  - label: "Поле ввода"
    href: /docs/primitives/input/
  - label: "Combobox"
    href: /docs/components/combobox/
api:
  - name: "Name"
    type: "string"
    description: "Имя поля формы"
  - name: "Value"
    type: "string"
    description: "Выбранное значение"
  - name: "Options"
    type: "[]Option"
    description: "Список вариантов Value и Label"
  - name: "Disabled"
    type: "bool"
    description: "Неактивное состояние"
---

Select рендерит нативный выпадающий список из слайса Options. Выбранный option отмечается, когда Value совпадает с Option Value.

## По умолчанию

```templ
import "github.com/fastygo/templ/ui"
import "github.com/fastygo/templ/ui/select"

templ Example() {
	@ui.Select(ui.SelectProps{
		Name:  "country",
		Value: "us",
		Options: []selectfield.Option{
			{Value: "us", Label: "США"},
			{Value: "ca", Label: "Канада"},
		},
	})
}
```

## Неактивный

```templ
import "github.com/fastygo/templ/ui"
import "github.com/fastygo/templ/ui/select"

templ Example() {
	@ui.Select(ui.SelectProps{
		Disabled: true,
		Name:     "locked",
		Options: []selectfield.Option{
			{Value: "a", Label: "Вариант A"},
		},
	})
}
```
