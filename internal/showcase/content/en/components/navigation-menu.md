---
slug: navigation-menu
section: components
title: "Navigation Menu"
description: "Site navigation with sections."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Menubar"
    href: /docs/components/menubar/
  - label: "Breadcrumb"
    href: /docs/components/breadcrumb/
api:
  - name: "Tag"
    type: "string"
    description: "menu for menubar-style lists"
---

Site navigation with sections.

## Default

Wireframe composition from ui primitives.

{{demo id="navigation-menu.default"}}

```templ
@ui.List[Tag=menu] { links }
```
