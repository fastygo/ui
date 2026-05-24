---
slug: navigation-menu
section: components
title: "Меню навигации"
description: "Навигация по сайту с разделами."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Панель меню"
    href: /docs/components/menubar/
  - label: "Хлебные крошки"
    href: /docs/components/breadcrumb/
api:
  - name: "Tag"
    type: "string"
    description: "menu для списков в стиле menubar"
---

Навигация по сайту с разделами.

## По умолчанию

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.List(ui.ListProps{Tag: "menu", Class: "flex gap-4 text-sm"}) {
		@ui.ListItem(ui.ListItemProps{}) {
			@ui.Button(ui.ButtonProps{Variant: "link"}) {
				Главная
			}
		}
		@ui.ListItem(ui.ListItemProps{}) {
			@ui.Button(ui.ButtonProps{Variant: "link"}) {
				Документация
			}
		}
		@ui.ListItem(ui.ListItemProps{}) {
			@ui.Button(ui.ButtonProps{Variant: "link"}) {
				Блог
			}
		}
	}
}
```
