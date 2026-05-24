---
slug: box
section: primitives
title: "Box"
description: "Внутренняя обёртка для layout внутри Block; не для landmarks верхнего уровня."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Блок"
    href: /docs/primitives/block/
  - label: "Стек"
    href: /docs/primitives/stack/
  - label: "Группа"
    href: /docs/primitives/group/
api:
  - name: "Tag"
    type: "string"
    description: "div | section | article | aside | header | footer | main | nav | figure"
  - name: "Class"
    type: "string"
    description: "Дополнительные utility-классы"
  - name: "Attrs"
    type: "templ.Attributes"
    description: "Дополнительные HTML-атрибуты"
---

Box оборачивает внутренние секции layout внутри Block. Не используйте для landmarks верхнего уровня страницы.

## По умолчанию

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Box(ui.BoxProps{Class: "mx-auto max-w-3xl px-6"}) {
		Внутреннее содержимое
	}
}
```

## Section

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Box(ui.BoxProps{Tag: "section", Class: "p-4"}) {
		Тело секции
	}
}
```
