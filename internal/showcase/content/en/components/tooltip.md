---
slug: tooltip
section: components
title: "Tooltip"
description: "Hint on hover/focus (wireframe)."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Popover"
    href: /docs/components/popover/
  - label: "Button"
    href: /docs/components/button/
api:
  - name: "Role"
    type: "string"
    description: "tooltip on hint box"
---

Hint on hover/focus (wireframe).

## Default

Wireframe composition from ui primitives.

{{demo id="tooltip.default"}}

```templ
@ui.Button + @ui.Box[role=tooltip]
```
