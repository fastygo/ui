---
slug: stack
section: components
title: "Stack"
description: "Vertical flex column for stacking children with gap utilities."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Group"
    href: /docs/components/group/
  - label: "Box"
    href: /docs/components/box/
api:
  - name: "Class"
    type: "string"
    description: "Tailwind utilities including gap"
  - name: "Tag"
    type: "string"
    description: "div | nav | section | ul | …"
  - name: "Attrs"
    type: "templ.Attributes"
    description: "Extra HTML attributes"
---

Vertical flex column for stacking children with gap utilities.

## Default

{{demo id="stack.default"}}

```templ
@ui.Stack(ui.StackProps{Class: "gap-2"}) { … }
```

## Row

{{demo id="stack.horizontal"}}

```templ
@ui.Stack(ui.StackProps{Class: "flex-row gap-2"}) { … }
```

## Nav tag

{{demo id="stack.nav"}}

```templ
@ui.Stack(ui.StackProps{Tag: "nav"}) { … }
```
