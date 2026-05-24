---
slug: select
section: components
title: "Select"
description: "Нативный выпадающий список (ui.Select / selectfield)."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Combobox"
    href: /docs/components/combobox/
  - label: "Radio"
    href: /docs/primitives/radio/
api:
  - name: "Options"
    type: "[]ui.Option"
    description: "Пары value/label"
  - name: "Name"
    type: "string"
    description: "Имя поля формы"
  - name: "Value"
    type: "string"
    description: "Выбранное значение"
---

Нативный выпадающий список (ui.Select / selectfield).

## По умолчанию

```templ
import (
	"github.com/fastygo/templ/ui"
	"github.com/fastygo/templ/ui/selectfield"
)

templ Example() {
	@ui.Select(ui.SelectProps{
		Name: "role",
		Options: []selectfield.Option{
			{Value: "viewer", Label: "Наблюдатель"},
			{Value: "editor", Label: "Редактор"},
		},
		Value: "viewer",
	})
}
```

## Отключено

```templ
import (
	"github.com/fastygo/templ/ui"
	"github.com/fastygo/templ/ui/selectfield"
)

templ Example() {
	@ui.Select(ui.SelectProps{
		Name: "role",
		Options: []selectfield.Option{
			{Value: "viewer", Label: "Наблюдатель"},
			{Value: "editor", Label: "Редактор"},
		},
		Disabled: true,
	})
}
```
