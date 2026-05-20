---
slug: accordion
section: components
title: "Accordion"
description: "Vertically stacked expandable sections (wireframe; data-ui8kit accordion)."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Collapsible"
    href: /docs/components/collapsible/
  - label: "Tabs"
    href: /docs/components/tabs/
api:
  - name: "Attrs"
    type: "templ.Attributes"
    description: "data-ui8kit on root"
  - name: "Trigger"
    type: "Button"
    description: "data-ui8kit-trigger"
  - name: "Panel"
    type: "Box"
    description: "data-ui8kit-panel"
---

Vertically stacked expandable sections (wireframe; data-ui8kit accordion).

## Default

Wireframe composition from ui primitives.

{{demo id="accordion.default"}}

```templ
@ui.Stack(templ.Attributes{"data-ui8kit": "accordion"}) { @ui.Button[data-ui8kit-trigger] … }
```
