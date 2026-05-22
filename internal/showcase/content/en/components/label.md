---
slug: label
section: components
title: "Label"
description: "Accessible label for form controls."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Input"
    href: /docs/components/input/
  - label: "Checkbox"
    href: /docs/components/checkbox/
api:
  - name: "HTMLFor"
    type: "string"
    description: "id of associated control"
  - name: "Class"
    type: "string"
    description: "Tailwind utilities"
---

Accessible label for form controls.

## Default

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Stack(ui.StackProps{}) {
		@ui.Label(ui.LabelProps{HTMLFor: "showcase-email"}) {
			Email
		}
		@ui.Input(ui.InputProps{ID: "showcase-email", Placeholder: "you@example.com"})
	}
}
```

## Required hint

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Stack(ui.StackProps{}) {
		@ui.Label(ui.LabelProps{HTMLFor: "showcase-name"}) {
			Name
		}
		@ui.Input(ui.InputProps{ID: "showcase-name", Required: true, Placeholder: "Required"})
	}
}
```
