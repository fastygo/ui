---
slug: button
section: components
title: "Кнопка"
description: "Запускает действие или переход, если отрендерена как ссылка. Основана на github.com/fastygo/templ/ui."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Переключатель"
    href: /docs/components/toggle/
  - label: "Форма"
    href: /docs/components/form/
api:
  - name: "Variant"
    type: "string"
    description: "default | secondary | destructive | outline | ghost | link | unstyled"
  - name: "Size"
    type: "string"
    description: "default | sm | lg | icon"
  - name: "Type"
    type: "string"
    description: "button | submit | reset"
  - name: "Href"
    type: "string"
    description: "Если задано — рендерится якорь вместо button"
  - name: "Disabled"
    type: "bool"
    description: "Отключает взаимодействие"
  - name: "Class"
    type: "string"
    description: "Дополнительные Tailwind-утилиты"
  - name: "AriaLabel"
    type: "string"
    description: "Доступное имя, когда видимого текста недостаточно"
---

Запускает действие или переход, если отрендерена как ссылка. Основана на github.com/fastygo/templ/ui.

## По умолчанию

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Button(ui.ButtonProps{}) {
		Кнопка
	}
}
```

## Вторичная

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Button(ui.ButtonProps{Variant: "secondary"}) {
		Вторичная
	}
}
```

## Контурная

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Button(ui.ButtonProps{Variant: "outline"}) {
		Контурная
	}
}
```

## Деструктивная

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Button(ui.ButtonProps{Variant: "destructive"}) {
		Удалить
	}
}
```

## Ghost

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Button(ui.ButtonProps{Variant: "ghost"}) {
		Ghost
	}
}
```

## Ссылка

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Button(ui.ButtonProps{Variant: "link"}) {
		Ссылка
	}
}
```

## Размеры

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Group(ui.GroupProps{Class: "gap-2"}) {
		@ui.Button(ui.ButtonProps{Size: "sm"}) {
			Малая
		}
		@ui.Button(ui.ButtonProps{}) {
			Обычная
		}
		@ui.Button(ui.ButtonProps{Size: "lg"}) {
			Крупная
		}
	}
}
```
