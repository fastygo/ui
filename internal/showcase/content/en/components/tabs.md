---
slug: tabs
section: components
title: "Tabs"
description: "Tabbed interface (wireframe; data-ui8kit tabs)."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Accordion"
    href: /docs/components/accordion/
  - label: "Navigation Menu"
    href: /docs/components/navigation-menu/
api:
  - name: "Attrs"
    type: "templ.Attributes"
    description: "data-ui8kit=tabs on root"
  - name: "Tab"
    type: "Button"
    description: "data-ui8kit-tab"
  - name: "Panel"
    type: "Box"
    description: "data-ui8kit-panel"
---

Tabbed interface (wireframe; data-ui8kit tabs).

## Default

Wireframe composition from ui primitives.

{{demo id="tabs.default"}}

```templ
@ui.Stack { tablist + panels }
```
