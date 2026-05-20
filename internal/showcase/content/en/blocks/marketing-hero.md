---
slug: marketing-hero
section: blocks
title: "Marketing Hero"
description: "Block scaffold — section wireframe with placeholder copy for future github.com/fastygo/blocks extraction."
source: internal/ui/blocks
package: internal/ui/blocks
related:
  - label: "Card"
    href: /docs/components/card/
  - label: "Stack"
    href: /docs/components/stack/
api:
  - name: "Title"
    type: "string"
    description: "Section heading"
  - name: "Body"
    type: "string"
    description: "Supporting copy"
---

Block scaffold — section wireframe with placeholder copy for future github.com/fastygo/blocks extraction.

## Wireframe

{{demo id="marketing-hero.default"}}

```templ
@ui.Stack { @ui.Title … "Marketing Hero" }
```

## Compact

{{demo id="marketing-hero.compact"}}

```go
Denser spacing variant
```
