---
slug: button
section: primitives
title: "Кнопка"
description: "Действие по клику или навигация при заданном Href."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Форма"
    href: /docs/components/form/
  - label: "Группа"
    href: /docs/primitives/group/
  - label: "Поле ввода"
    href: /docs/primitives/input/
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
    description: "URL для ссылки вместо button"
  - name: "Disabled"
    type: "bool"
    description: "Неактивное состояние"
---

Кнопка запускает действие по клику. При заданном Href рендерится как ссылка.

## Основная

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Button(ui.ButtonProps{Variant: "default"}) {
		Основная
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

## Ссылка

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Button(ui.ButtonProps{Variant: "link", Href: "/docs"}) {
		Подробнее
	}
}
```

## Отправка формы

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Button(ui.ButtonProps{Type: "submit"}) {
		Отправить
	}
}
```

## Неактивная

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Button(ui.ButtonProps{Disabled: true}) {
		Недоступно
	}
}
```
