---
slug: title
section: components
title: "Title"
description: "Semantic heading with order 1–6."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Text"
    href: /docs/components/text/
  - label: "Stack"
    href: /docs/components/stack/
api:
  - name: "Order"
    type: "int"
    description: "1–6 maps to h1–h6"
  - name: "Class"
    type: "string"
    description: "Tailwind utilities"
---

Semantic heading with order 1–6.

## Heading 1

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Title(ui.TitleProps{Order: 1}, "Page title")
}
```

## Heading 2

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Title(ui.TitleProps{Order: 2}, "Section")
}
```

## Heading 3

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Title(ui.TitleProps{Order: 3}, "Subsection")
}
```
