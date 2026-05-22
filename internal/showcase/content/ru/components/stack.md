---
slug: stack
section: components
title: "Стек"
description: "Вертикальная flex-колонка для дочерних элементов с gap-утилитами."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Группа"
    href: /docs/components/group/
  - label: "Box"
    href: /docs/components/box/
api:
  - name: "Class"
    type: "string"
    description: "Tailwind-утилиты, включая gap"
  - name: "Tag"
    type: "string"
    description: "div | nav | section | ul | …"
  - name: "Attrs"
    type: "templ.Attributes"
    description: "Дополнительные HTML-атрибуты"
---

Вертикальная flex-колонка для дочерних элементов с gap-утилитами.

## По умолчанию

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Stack(ui.StackProps{Class: "gap-2"}) {
		@ui.Title(ui.TitleProps{Order: 3}, "Стек")
		@ui.Text(ui.TextProps{Class: "text-sm text-muted-foreground"}, "Дочерние элементы выстраиваются вертикально.")
	}
}
```

## Строка

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Stack(ui.StackProps{Class: "flex-row items-center gap-2"}) {
		@ui.Title(ui.TitleProps{Order: 3}, "Стек")
		@ui.Text(ui.TextProps{Class: "text-sm text-muted-foreground"}, "Дочерние элементы в одной строке.")
	}
}
```

## Тег nav

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Stack(ui.StackProps{Tag: "nav", Class: "gap-1"}) {
		@ui.Title(ui.TitleProps{Order: 3}, "Стек")
		@ui.Text(ui.TextProps{Class: "text-sm text-muted-foreground"}, "Landmark nav с вертикальным списком.")
	}
}
```
