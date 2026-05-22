---
slug: form
section: components
title: "Form"
description: "Form landmark with item helpers."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Input"
    href: /docs/components/input/
  - label: "Button"
    href: /docs/components/button/
api:
  - name: "Action"
    type: "string"
    description: "Form action URL"
  - name: "Method"
    type: "string"
    description: "GET | POST"
  - name: "FormItem"
    type: "component"
    description: "Label + control group"
---

Form landmark with item helpers.

## Login

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Form(ui.FormProps{Class: "max-w-sm"}) {
		@ui.FormItem(ui.FormItemProps{}) {
			@ui.Label(ui.LabelProps{HTMLFor: "login-email"}) {
				Email
			}
			@ui.Input(ui.InputProps{ID: "login-email", Type: "email", Placeholder: "you@example.com"})
		}
		@ui.Button(ui.ButtonProps{Type: "submit"}) {
			Sign in
		}
	}
}
```

## Inline

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Form(ui.FormProps{Class: "max-w-md"}) {
		@ui.Group(ui.GroupProps{Class: "flex items-end gap-2"}) {
			@ui.Input(ui.InputProps{Placeholder: "Search"})
			@ui.Button(ui.ButtonProps{}) {
				Go
			}
		}
	}
}
```
