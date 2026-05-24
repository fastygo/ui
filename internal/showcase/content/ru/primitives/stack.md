---
slug: stack
section: primitives
title: "Стек"
description: "Вертикальный flex-столбец; базовый примитив отступов."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Группа"
    href: /docs/primitives/group/
  - label: "Box"
    href: /docs/primitives/box/
  - label: "Текст"
    href: /docs/primitives/text/
api:
  - name: "Tag"
    type: "string"
    description: "div | nav | section | ul | ol | …"
  - name: "Class"
    type: "string"
    description: "Дополнительные utility-классы"
  - name: "Attrs"
    type: "templ.Attributes"
    description: "Дополнительные HTML-атрибуты"
---

Stack располагает дочерние элементы в вертикальном flex-столбце. Это базовый примитив отступов между секциями.

## По умолчанию

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Stack(ui.StackProps{Class: "gap-4"}) {
		@ui.Text(ui.TextProps{}, "Первая строка")
		@ui.Text(ui.TextProps{}, "Вторая строка")
	}
}
```

## Список ul

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Stack(ui.StackProps{Tag: "ul", Class: "gap-2"}) {
		@ui.Text(ui.TextProps{Tag: "span"}, "Пункт один")
	}
}
```
