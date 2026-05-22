---
slug: block
section: components
title: "Block"
description: "Top-level landmark sections (do not nest Block in Block)."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Box"
    href: /docs/components/box/
  - label: "Container"
    href: /docs/components/container/
api:
  - name: "Tag"
    type: "string"
    description: "main | section | aside | nav | …"
  - name: "Class"
    type: "string"
    description: "Tailwind utilities"
  - name: "Attrs"
    type: "templ.Attributes"
    description: "Extra attributes"
---

Top-level landmark sections (do not nest Block in Block).

## Main

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Block(ui.BlockProps{Tag: "main", Class: "rounded-lg border border-border p-4"}) {
		@ui.Text(ui.TextProps{}, "Main landmark block.")
	}
}
```

## Aside

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Block(ui.BlockProps{Tag: "aside", Class: "rounded-lg border border-border p-4 w-48"}) {
		@ui.Text(ui.TextProps{}, "Aside block.")
	}
}
```
