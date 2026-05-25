---
slug: dropdown-menu
section: components
title: "Dropdown Menu"
description: "Меню, открываемое по кнопке."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Меню навигации"
    href: /docs/components/navigation-menu/
  - label: "Панель меню"
    href: /docs/components/menubar/
api:
  - name: "List"
    type: "List"
    description: "Тег menu для пунктов"
---

Меню, открываемое по кнопке.

## По умолчанию

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Stack(ui.StackProps{Class: "gap-2"}) {
		@ui.Button(ui.ButtonProps{Variant: "outline"}) {
			Открыть меню
		}
		@ui.List(ui.ListProps{Tag: "menu", Class: "w-40 rounded-md border border-border bg-card p-1 text-sm"}) {
			@ui.ListItem(ui.ListItemProps{Class: "rounded px-2 py-1 hover:bg-accent"}) {
				@ui.Button(ui.ButtonProps{Variant: "ghost", Class: "h-8 w-full justify-start px-2"}) {
					Профиль
				}
			}
			@ui.ListItem(ui.ListItemProps{Class: "rounded px-2 py-1 hover:bg-accent"}) {
				@ui.Button(ui.ButtonProps{Variant: "ghost", Class: "h-8 w-full justify-start px-2"}) {
					Настройки
				}
			}
			@ui.ListItem(ui.ListItemProps{Class: "rounded px-2 py-1 hover:bg-accent"}) {
				@ui.Button(ui.ButtonProps{Variant: "ghost", Class: "h-8 w-full justify-start px-2"}) {
					Выйти
				}
			}
		}
	}
}
```
