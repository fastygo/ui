---
slug: list
section: components
title: "List"
description: "Semantic ul/ol/dl list containers."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Table"
    href: /docs/components/table/
  - label: "Breadcrumb"
    href: /docs/components/breadcrumb/
api:
  - name: "Tag"
    type: "string"
    description: "ul | ol | dl | menu"
  - name: "Class"
    type: "string"
    description: "Tailwind utilities"
  - name: "Attrs"
    type: "templ.Attributes"
    description: "Extra attributes"
---

Semantic ul/ol/dl list containers.

## Unordered

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.List(ui.ListProps{Class: "list-disc pl-5 text-sm"}) {
		@ui.ListItem(ui.ListItemProps{}) {
			@ui.Text(ui.TextProps{}, "First item")
		}
		@ui.ListItem(ui.ListItemProps{}) {
			@ui.Text(ui.TextProps{}, "Second item")
		}
	}
}
```

## Ordered

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.List(ui.ListProps{Tag: "ol", Class: "list-decimal pl-5 text-sm"}) {
		@ui.ListItem(ui.ListItemProps{}) {
			@ui.Text(ui.TextProps{}, "First item")
		}
		@ui.ListItem(ui.ListItemProps{}) {
			@ui.Text(ui.TextProps{}, "Second item")
		}
	}
}
```
