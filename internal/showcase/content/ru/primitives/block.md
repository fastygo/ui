---
slug: block
section: primitives
title: "Block"
description: "Landmark верхнего уровня страницы; не вкладывайте Block в Block."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Box"
    href: /docs/primitives/box/
  - label: "Контейнер"
    href: /docs/primitives/container/
  - label: "Стек"
    href: /docs/primitives/stack/
api:
  - name: "Tag"
    type: "string"
    description: "main | section | header | footer | nav | article | aside | figure | div"
  - name: "Class"
    type: "string"
    description: "Дополнительные utility-классы"
  - name: "ID"
    type: "string"
    description: "Идентификатор элемента"
---

Block рендерит landmark верхнего уровня страницы. Не вкладывайте Block внутрь другого Block.

## Main

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Block(ui.BlockProps{Tag: "main", Class: "min-h-screen"}) {
		Содержимое страницы
	}
}
```

## Section

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Block(ui.BlockProps{Tag: "section", Class: "py-8"}) {
		Содержимое секции
	}
}
```

## Header

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Block(ui.BlockProps{Tag: "header", Class: "border-b border-border"}) {
		Шапка сайта
	}
}
```

## Nav

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Block(ui.BlockProps{Tag: "nav", Class: "py-2"}) {
		Основная навигация
	}
}
```
