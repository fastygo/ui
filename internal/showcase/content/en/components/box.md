---
slug: box
section: components
title: "Box"
description: "Generic block wrapper without landmark semantics."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Block"
    href: /docs/components/block/
  - label: "Stack"
    href: /docs/components/stack/
api:
  - name: "Class"
    type: "string"
    description: "Tailwind utilities"
  - name: "Tag"
    type: "string"
    description: "motion.div | pre | span"
  - name: "Attrs"
    type: "templ.Attributes"
    description: "Extra attributes"
---

Generic block wrapper without landmark semantics.

## Default

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Box(ui.BoxProps{Class: "rounded-lg border border-border p-4"}) {
		@ui.Text(ui.TextProps{}, "Box content.")
	}
}
```

## Pre tag

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Box(ui.BoxProps{Tag: "pre", Class: "rounded-md border border-border bg-muted/30 p-3 text-xs font-mono"}) {
		@ui.Text(ui.TextProps{Tag: "code"}, "code snippet")
	}
}
```
