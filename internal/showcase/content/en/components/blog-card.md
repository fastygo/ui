---
slug: blog-card
section: components
title: "Blog Card"
description: "Reusable blog list cards with media placeholder — vertical and horizontal layouts."
source: github.com/fastygo/ui/internal/ui/components/blogcard
package: github.com/fastygo/ui/internal/ui/components/blogcard
related:
  - label: "Card"
    href: /docs/components/card/
  - label: "Stack"
    href: /docs/components/stack/
api:
  - name: "Title"
    type: "string"
    description: "Article title"
  - name: "Excerpt"
    type: "string"
    description: "Short summary"
  - name: "Href"
    type: "string"
    description: "Read-more link target"
  - name: "MediaURL"
    type: "string"
    description: "Reserved for future media primitive; wireframe uses muted placeholder"
  - name: "MediaAlt"
    type: "string"
    description: "Accessible label for media placeholder"
  - name: "DateLabel"
    type: "string"
    description: "Published date (pre-formatted)"
---

Reusable blog list cards with media placeholder — vertical and horizontal layouts.

## Vertical

Stacked media and copy for grids and feeds.

{{demo id="blog-card.vertical"}}

```templ
@blogcard.VerticalBlogCard(blogcard.DefaultVertical())
```

## Horizontal

Side-by-side media and copy for dense lists.

{{demo id="blog-card.horizontal"}}

```templ
@blogcard.HorizontalBlogCard(blogcard.DefaultHorizontal())
```
