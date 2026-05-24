---
slug: pagination
section: components
title: "Пагинация"
description: "Элементы навигации по страницам."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Таблица"
    href: /docs/components/table/
  - label: "Хлебные крошки"
    href: /docs/components/breadcrumb/
api:
  - name: "Buttons"
    type: "Button"
    description: "Назад / страница / Вперёд"
---

Элементы навигации по страницам.

## По умолчанию

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Group(ui.GroupProps{Class: "flex items-center gap-1"}) {
		@ui.Button(ui.ButtonProps{Variant: "outline", Size: "sm"}) {
			Назад
		}
		@ui.Button(ui.ButtonProps{Variant: "secondary", Size: "sm"}) {
			1
		}
		@ui.Button(ui.ButtonProps{Variant: "outline", Size: "sm"}) {
			Вперёд
		}
	}
}
```
