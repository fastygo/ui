---
slug: menubar
section: components
title: "Панель меню"
description: "Горизонтальная панель меню."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Меню навигации"
    href: /docs/components/navigation-menu/
  - label: "Выпадающее меню"
    href: /docs/components/dropdown-menu/
api:
  - name: "Group"
    type: "Group"
    description: "Горизонтальный ряд кнопок"
---

Горизонтальная панель меню.

## По умолчанию

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Group(ui.GroupProps{Class: "flex gap-1 rounded-md border border-border bg-card p-1 text-sm"}) {
		@ui.Button(ui.ButtonProps{Variant: "ghost", Size: "sm"}) {
			Файл
		}
		@ui.Button(ui.ButtonProps{Variant: "ghost", Size: "sm"}) {
			Правка
		}
		@ui.Button(ui.ButtonProps{Variant: "ghost", Size: "sm"}) {
			Вид
		}
	}
}
```
