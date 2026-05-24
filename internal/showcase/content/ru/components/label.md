---
slug: label
section: components
title: "Метка"
description: "Доступная метка для элементов формы."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Поле ввода"
    href: /docs/primitives/input/
  - label: "Чекбокс"
    href: /docs/primitives/checkbox/
api:
  - name: "HTMLFor"
    type: "string"
    description: "id связанного элемента управления"
  - name: "Class"
    type: "string"
    description: "Tailwind-утилиты"
---

Доступная метка для элементов формы.

## По умолчанию

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Stack(ui.StackProps{}) {
		@ui.Label(ui.LabelProps{HTMLFor: "showcase-email"}) {
			Email
		}
		@ui.Input(ui.InputProps{ID: "showcase-email", Placeholder: "you@example.com"})
	}
}
```

## Обязательное поле

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Stack(ui.StackProps{}) {
		@ui.Label(ui.LabelProps{HTMLFor: "showcase-name"}) {
			Имя
		}
		@ui.Input(ui.InputProps{ID: "showcase-name", Required: true, Placeholder: "Обязательно"})
	}
}
```
