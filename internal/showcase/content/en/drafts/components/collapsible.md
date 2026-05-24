---
slug: collapsible
section: components
title: "Collapsible"
description: "Single expand/collapse region."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Accordion"
    href: /docs/components/accordion/
  - label: "Sheet"
    href: /docs/components/sheet/
api:
  - name: "Trigger"
    type: "Button"
    description: "Expands panel"
  - name: "Panel"
    type: "Box"
    description: "Collapsible content"
---

Single expand/collapse region.

## Default

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Stack(ui.StackProps{Class: "gap-2 max-w-md", Attrs: templ.Attributes{"data-ui8kit": "disclosure"}}) {
		@ui.Button(ui.ButtonProps{Variant: "ghost", Attrs: templ.Attributes{"data-ui8kit-trigger": "", "aria-expanded": "false"}}) {
			Show more
		}
		@ui.Box(ui.BoxProps{Class: "rounded border border-border p-3 text-sm", Attrs: templ.Attributes{"data-ui8kit-panel": "", "hidden": true}}) {
			@ui.Text(ui.TextProps{}, "Hidden details.")
		}
	}
}
```
