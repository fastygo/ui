---
slug: alert
section: components
title: "Alert"
description: "Callout для важных сообщений."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Карточка"
    href: /docs/components/card/
  - label: "Badge"
    href: /docs/components/badge/
api:
  - name: "Variant"
    type: "string"
    description: "default | destructive"
  - name: "Class"
    type: "string"
    description: "Дополнительные утилиты"
---

Callout для важных сообщений.

## По умолчанию

```templ
import cmp "github.com/fastygo/templ/components"
import "github.com/fastygo/templ/ui"

templ Example() {
	@cmp.Alert(cmp.AlertProps{}) {
		@ui.Title(ui.TitleProps{Order: 4, Class: "text-sm font-semibold"}, "Обратите внимание")
		@ui.Text(ui.TextProps{Class: "text-sm"}, "Компоненты можно добавлять из галереи.")
	}
}
```

## Деструктивный

```templ
import cmp "github.com/fastygo/templ/components"
import "github.com/fastygo/templ/ui"

templ Example() {
	@cmp.Alert(cmp.AlertProps{Variant: "destructive"}) {
		@ui.Title(ui.TitleProps{Order: 4, Class: "text-sm font-semibold"}, "Ошибка")
		@ui.Text(ui.TextProps{Class: "text-sm"}, "Что-то пошло не так.")
	}
}
```
