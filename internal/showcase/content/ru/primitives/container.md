---
slug: container
section: primitives
title: "Контейнер"
description: "Оболочка ширины страницы с div, main или section."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Блок"
    href: /docs/primitives/block/
  - label: "Box"
    href: /docs/primitives/box/
  - label: "Стек"
    href: /docs/primitives/stack/
api:
  - name: "Tag"
    type: "string"
    description: "div | main | section"
  - name: "Class"
    type: "string"
    description: "Дополнительные utility-классы"
  - name: "Attrs"
    type: "templ.Attributes"
    description: "Дополнительные HTML-атрибуты"
---

Container оборачивает содержимое страницы в оболочку ширины. Поддерживает корневые теги div, main и section.

## По умолчанию

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Container(ui.ContainerProps{Class: "mx-auto max-w-5xl px-6"}) {
		Контент по ширине страницы
	}
}
```

## Main

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Container(ui.ContainerProps{Tag: "main", Class: "mx-auto max-w-5xl px-6"}) {
		Содержимое основной области
	}
}
```
