---
slug: docs-article
section: blocks
title: "Docs Article"
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

{{demo id="docs-article.default"}}

```templ
@ui.Stack { @ui.Title … "Docs Article" }
```

## Compact

{{demo id="docs-article.compact"}}

```go
Denser spacing variant
```
