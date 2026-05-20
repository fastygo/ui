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

{{demo id="label.default"}}

```templ
@showcaseutil.RenderLabel(ctx, w, ui.LabelProps{HTMLFor: "email"}, "Email")
```

## Required hint

{{demo id="label.required"}}

```go
Pair with aria-required on control
```
