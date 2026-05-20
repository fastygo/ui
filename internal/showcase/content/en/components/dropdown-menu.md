---
slug: dropdown-menu
section: components
title: "Dropdown Menu"
description: "Menu triggered by a button."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Context Menu"
    href: /docs/components/context-menu/
  - label: "Menubar"
    href: /docs/components/menubar/
api:
  - name: "List"
    type: "List"
    description: "menu tag for items"
---

Menu triggered by a button.

## Default

Wireframe composition from ui primitives.

{{demo id="dropdown-menu.default"}}

```templ
@ui.Group { trigger + menu list }
```
