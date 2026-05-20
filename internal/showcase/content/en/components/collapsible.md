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

{{demo id="collapsible.default"}}

```templ
@ui.Button + @ui.Box { data-ui8kit=disclosure }
```
