---
slug: box
section: components
title: "Box"
description: "Generic block wrapper without landmark semantics."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Block"
    href: /docs/components/block/
  - label: "Stack"
    href: /docs/components/stack/
api:
  - name: "Class"
    type: "string"
    description: "Tailwind utilities"
  - name: "Tag"
    type: "string"
    description: "motion.div | pre | span"
  - name: "Attrs"
    type: "templ.Attributes"
    description: "Extra attributes"
---

Generic block wrapper without landmark semantics.

## Default

{{demo id="box.default"}}

```templ
@ui.Box(ui.BoxProps{Class: "rounded-lg border border-border p-4"}) { … }
```

## Pre tag

{{demo id="box.pre"}}

```templ
@ui.Box(ui.BoxProps{Tag: "pre"}) { … }
```
