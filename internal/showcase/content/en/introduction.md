---
slug: introduction
section: getting-started
title: "Introduction"
description: "FastyGo UI is a wireframe-first component gallery on Go and templ. Pages show live previews, copy-pasteable templ snippets, and prop tables—similar to shadcn/ui, without shipping a monolithic UI package."
source: github.com/fastygo/ui
package: internal/site
related:
  - label: "Installation"
    href: /docs/installation/
  - label: "Components index"
    href: /docs/
---

FastyGo UI is a wireframe-first component gallery on Go and templ. Pages show live previews, copy-pasteable templ snippets, and prop tables—similar to shadcn/ui, without shipping a monolithic UI package.

## Wireframe scope

Structure, semantics, and accessibility come first; visual brand polish is a later phase.

{{demo id="introduction.overview"}}

```templ
@ui.Stack(ui.StackProps{}) {
  @ui.Title(ui.TitleProps{Order: 1}, "Component gallery")
  @ui.Text(ui.TextProps{}, "Wireframe previews + API tables.")
}
```
