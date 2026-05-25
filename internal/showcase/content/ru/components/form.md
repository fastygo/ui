---
slug: form
section: components
title: "Form"
description: "Landmark формы с вспомогательными элементами FormItem."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Поле ввода"
    href: /docs/primitives/input/
  - label: "Кнопка"
    href: /docs/primitives/button/
api:
  - name: "Action"
    type: "string"
    description: "URL action формы"
  - name: "Method"
    type: "string"
    description: "GET | POST"
  - name: "FormItem"
    type: "component"
    description: "Группа label + control"
---

Landmark формы с вспомогательными элементами FormItem.

## Вход

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Form(ui.FormProps{Class: "max-w-sm"}) {
		@ui.FormItem(ui.FormItemProps{}) {
			@ui.Label(ui.LabelProps{HTMLFor: "login-email"}) {
				Email
			}
			@ui.Input(ui.InputProps{ID: "login-email", Type: "email", Placeholder: "you@example.com"})
		}
		@ui.Button(ui.ButtonProps{Type: "submit"}) {
			Войти
		}
	}
}
```

## В строку

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Form(ui.FormProps{Class: "max-w-md"}) {
		@ui.Group(ui.GroupProps{Class: "flex items-end gap-2"}) {
			@ui.Input(ui.InputProps{Placeholder: "Поиск"})
			@ui.Button(ui.ButtonProps{}) {
				Найти
			}
		}
	}
}
```
