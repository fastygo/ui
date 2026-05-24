---
slug: checkbox
section: primitives
title: "Чекбокс"
description: "Нативный булев контрол формы."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Метка"
    href: /docs/primitives/label/
  - label: "Радиокнопка"
    href: /docs/primitives/radio/
  - label: "Форма"
    href: /docs/components/form/
api:
  - name: "Name"
    type: "string"
    description: "Имя поля формы"
  - name: "Checked"
    type: "bool"
    description: "Отмеченное состояние"
  - name: "Disabled"
    type: "bool"
    description: "Неактивное состояние"
  - name: "AriaLabel"
    type: "string"
    description: "Доступное имя без видимой метки"
---

Чекбокс рендерит нативный булев контрол формы. Используйте для согласий и переключателей настроек.

## Не отмечен

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Checkbox(ui.CheckboxProps{Name: "terms", ID: "terms"})
}
```

## Отмечен

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Checkbox(ui.CheckboxProps{Name: "terms", ID: "terms", Checked: true})
}
```

## Неактивный

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Checkbox(ui.CheckboxProps{Name: "locked", Disabled: true, Checked: true})
}
```
