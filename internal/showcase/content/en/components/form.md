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

{{demo id="form.default"}}

```templ
@ui.Form(ui.FormProps{}) { @ui.FormItem … }
```

## Inline

{{demo id="form.inline"}}

```go
Compact horizontal FormItem layout
```
