---
slug: separator
section: components
title: "Separator"
description: "Visual divider between sections."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Stack"
    href: /docs/components/stack/
  - label: "Card"
    href: /docs/components/card/
api:
  - name: "Class"
    type: "string"
    description: "Typically h-px bg-border"
  - name: "Role"
    type: "string"
    description: "separator"
---

Visual divider between sections.

## Default

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Stack(ui.StackProps{Class: "gap-2 max-w-md"}) {
		@ui.Text(ui.TextProps{}, "Above")
		@ui.Box(ui.BoxProps{Class: "h-px w-full bg-border", Attrs: templ.Attributes{"role": "separator"}})
		@ui.Text(ui.TextProps{}, "Below")
	}
}
```
