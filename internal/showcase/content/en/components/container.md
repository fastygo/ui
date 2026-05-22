---
slug: container
section: components
title: "Container"
description: "Centers content with a max-width constraint."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Stack"
    href: /docs/components/stack/
  - label: "Box"
    href: /docs/components/box/
api:
  - name: "Class"
    type: "string"
    description: "Tailwind utilities"
  - name: "Tag"
    type: "string"
    description: "motion.div | section | main"
  - name: "Attrs"
    type: "templ.Attributes"
    description: "Extra attributes"
---

Centers content with a max-width constraint.

## Default

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Container(ui.ContainerProps{Class: "mx-auto max-w-3xl border border-border rounded-lg p-4"}) {
		@ui.Text(ui.TextProps{}, "Container content.")
	}
}
```

## Section

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Container(ui.ContainerProps{Tag: "section", Class: "mx-auto max-w-2xl border border-border rounded-lg p-4"}) {
		@ui.Text(ui.TextProps{}, "Section landmark.")
	}
}
```
