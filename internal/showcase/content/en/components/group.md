---
slug: group
section: components
title: "Group"
description: "Horizontal flex row for grouping controls."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Stack"
    href: /docs/components/stack/
  - label: "Box"
    href: /docs/components/box/
api:
  - name: "Class"
    type: "string"
    description: "Tailwind utilities"
  - name: "Tag"
    type: "string"
    description: "div | span"
  - name: "Attrs"
    type: "templ.Attributes"
    description: "Extra attributes"
---

Horizontal flex row for grouping controls.

## Default

{{demo id="group.default"}}

```templ
@ui.Group(ui.GroupProps{Class: "flex items-center gap-2"}) { … }
```

## Wrap

{{demo id="group.wrap"}}

```templ
@ui.Group(ui.GroupProps{Class: "flex flex-wrap gap-2"}) { … }
```
