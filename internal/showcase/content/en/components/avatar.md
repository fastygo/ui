---
slug: avatar
section: components
title: "Avatar"
description: "User avatar placeholder."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Badge"
    href: /docs/primitives/badge/
  - label: "Card"
    href: /docs/components/card/
api:
  - name: "Class"
    type: "string"
    description: "rounded-full size utilities"
---

User avatar placeholder.

## Default

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Group(ui.GroupProps{Class: "flex items-center gap-3"}) {
		@ui.Box(ui.BoxProps{Class: "flex h-10 w-10 items-center justify-center rounded-full bg-muted text-sm font-medium"}) {
			@ui.Text(ui.TextProps{}, "AB")
		}
		@ui.Text(ui.TextProps{Class: "text-sm font-medium"}, "Ada Lovelace")
	}
}
```
