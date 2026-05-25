---
slug: switch
section: primitives
title: "Switch"
description: "Toggle с ARIA role switch и aria-checked."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Чекбокс"
    href: /docs/primitives/checkbox/
  - label: "Метка"
    href: /docs/primitives/label/
  - label: "Форма"
    href: /docs/components/form/
api:
  - name: "Name"
    type: "string"
    description: "Имя поля формы"
  - name: "Checked"
    type: "bool"
    description: "Включённое состояние"
  - name: "Disabled"
    type: "bool"
    description: "Неактивное состояние"
  - name: "AriaLabel"
    type: "string"
    description: "Доступное имя переключателя"
---

Switch рендерит toggle с ARIA-семантикой switch. aria-checked отражает prop Checked при рендере.

## Выключен

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Switch(ui.SwitchProps{
		Name:      "notifications",
		ID:        "notifications",
		AriaLabel: "Уведомления",
	})
}
```

## Включён

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Switch(ui.SwitchProps{
		Name:      "notifications",
		ID:        "notifications",
		Checked:   true,
		AriaLabel: "Уведомления",
	})
}
```

## Неактивный

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Switch(ui.SwitchProps{
		Name:      "locked",
		Disabled:  true,
		Checked:   true,
		AriaLabel: "Заблокированная настройка",
	})
}
```
