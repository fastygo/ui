---
slug: text
section: primitives
title: "Text"
description: "Обычный текст в семантическом теге; копия передаётся вторым аргументом."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Заголовок"
    href: /docs/primitives/title/
  - label: "Стек"
    href: /docs/primitives/stack/
  - label: "Бейдж"
    href: /docs/primitives/badge/
api:
  - name: "Tag"
    type: "string"
    description: "p | span | small | code | blockquote | em | strong | …"
  - name: "Class"
    type: "string"
    description: "Дополнительные utility-классы"
  - name: "value"
    type: "string"
    description: "Текст копии (второй аргумент компонента)"
---

Text рендерит обычный текст в семантическом теге. Копию передают вторым аргументом value.

## Абзац

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Text(ui.TextProps{}, "Текст абзаца body.")
}
```

## Подпись small

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Text(ui.TextProps{Tag: "small", Class: "text-muted-foreground"}, "Вспомогательная строка.")
}
```

## Code

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Text(ui.TextProps{Tag: "code", Class: "text-sm"}, "npm install")
}
```
