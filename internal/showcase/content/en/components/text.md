---
slug: text
section: components
title: "Text"
description: "Inline or block text with configurable tag."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Title"
    href: /docs/components/title/
  - label: "Label"
    href: /docs/components/label/
api:
  - name: "Tag"
    type: "string"
    description: "p | span | code | …"
  - name: "Class"
    type: "string"
    description: "Tailwind utilities"
---

Inline or block text with configurable tag.

## Paragraph

{{demo id="text.default"}}

```templ
@ui.Text(ui.TextProps{}, "Body copy.")
```

## Muted

{{demo id="text.muted"}}

```templ
@ui.Text(ui.TextProps{Class: "text-sm text-muted-foreground"}, "…")
```

## Code

{{demo id="text.code"}}

```templ
@ui.Text(ui.TextProps{Tag: "code"}, "npm install")
```
