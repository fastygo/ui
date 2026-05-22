---
slug: marketing-hero
section: blocks
title: "Marketing Hero"
description: "Block scaffold — section wireframe with placeholder copy for future github.com/fastygo/blocks extraction."
source: internal/ui/blocks
package: internal/ui/blocks
related:
  - label: "Card"
    href: /docs/components/card/
  - label: "Stack"
    href: /docs/components/stack/
api:
  - name: "Title"
    type: "string"
    description: "Section heading"
  - name: "Body"
    type: "string"
    description: "Supporting copy"
---

Block scaffold — section wireframe with placeholder copy for future github.com/fastygo/blocks extraction.

## Wireframe

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Stack(ui.StackProps{Class: "gap-4 max-w-2xl"}) {
		@ui.Title(ui.TitleProps{Order: 2}, "Marketing Hero")
		@ui.Text(ui.TextProps{Class: "text-sm text-muted-foreground leading-relaxed"}, "Wireframe hero headline and call-to-action for marketing blocks.")
		@ui.Button(ui.ButtonProps{Variant: "outline", Size: "sm"}) {
			Action
		}
	}
}
```

## Compact

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Stack(ui.StackProps{Class: "gap-2 max-w-2xl"}) {
		@ui.Title(ui.TitleProps{Order: 2}, "Marketing Hero")
		@ui.Text(ui.TextProps{Class: "text-sm text-muted-foreground leading-relaxed"}, "Wireframe hero headline and call-to-action for marketing blocks.")
		@ui.Button(ui.ButtonProps{Variant: "outline", Size: "sm"}) {
			Action
		}
	}
}
```
