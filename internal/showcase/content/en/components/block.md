---
slug: block
section: components
title: "Block"
description: "Top-level landmark sections (do not nest Block in Block)."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Box"
    href: /docs/components/box/
  - label: "Container"
    href: /docs/components/container/
api:
  - name: "Tag"
    type: "string"
    description: "main | section | aside | nav | …"
  - name: "Class"
    type: "string"
    description: "Tailwind utilities"
  - name: "Attrs"
    type: "templ.Attributes"
    description: "Extra attributes"
---

Top-level landmark sections (do not nest Block in Block).

## Main

{{demo id="block.main"}}

```templ
@ui.Block(ui.BlockProps{Tag: "main"}) { … }
```

## Aside

{{demo id="block.aside"}}

```templ
@ui.Block(ui.BlockProps{Tag: "aside"}) { … }
```
