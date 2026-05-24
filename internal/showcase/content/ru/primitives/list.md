---
slug: list
section: primitives
title: "Список"
description: "Контейнер ul, ol, dl или menu с элементами ListItem."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Стек"
    href: /docs/primitives/stack/
  - label: "Текст"
    href: /docs/primitives/text/
  - label: "Навигационное меню"
    href: /docs/components/navigation-menu/
api:
  - name: "Tag"
    type: "string"
    description: "ul | ol | dl | menu"
  - name: "Class"
    type: "string"
    description: "Дополнительные utility-классы"
  - name: "ListItem"
    type: "component"
    description: "Один элемент li внутри List"
---

List рендерит контейнеры ul, ol, dl или menu. ListItem рендерит одну строку li внутри List.

## Маркированный

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.List(ui.ListProps{Class: "gap-2"}) {
		@ui.ListItem(ui.ListItemProps{}) { Первый пункт }
		@ui.ListItem(ui.ListItemProps{}) { Второй пункт }
	}
}
```

## Нумерованный

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.List(ui.ListProps{Tag: "ol", Class: "list-decimal pl-4 gap-2"}) {
		@ui.ListItem(ui.ListItemProps{Value: 1}) { Шаг один }
		@ui.ListItem(ui.ListItemProps{Value: 2}) { Шаг два }
	}
}
```
