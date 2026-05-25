---
slug: title
section: primitives
title: "Title"
description: "Один heading по prop Order (h1–h6)."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Текст"
    href: /docs/primitives/text/
  - label: "Блок"
    href: /docs/primitives/block/
  - label: "Стек"
    href: /docs/primitives/stack/
api:
  - name: "Order"
    type: "int"
    description: "Уровень заголовка 1–6"
  - name: "Class"
    type: "string"
    description: "Дополнительные utility-классы"
  - name: "value"
    type: "string"
    description: "Текст заголовка (второй аргумент компонента)"
---

Title рендерит один заголовок по prop Order. TitleTag сопоставляет Order с h1–h6.

## H1

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Title(ui.TitleProps{Order: 1}, "Заголовок страницы")
}
```

## H2

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Title(ui.TitleProps{Order: 2}, "Заголовок секции")
}
```

## H3

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Title(ui.TitleProps{Order: 3, Class: "text-lg"}, "Заголовок подсекции")
}
```
