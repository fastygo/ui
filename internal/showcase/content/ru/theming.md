---
slug: theming
section: getting-started
title: "Theming"
description: "Семантические токены (background, foreground, primary, …) задаются CSS-переменными в стиле tweakcn. Переключатель темы в шапке меняет светлую/тёмную тему через theme.js."
source: web/static/css/tweakcn.css
package: web/static/js/theme.js
related:
  - label: "Введение"
    href: /docs/introduction/
---

Семантические токены (background, foreground, primary, …) задаются CSS-переменными в стиле tweakcn. Переключатель темы в шапке меняет светлую/тёмную тему через theme.js.

## Семантические токены

Используйте утилиты токенов — например bg-background и text-foreground — в class-строках templ.

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Group(ui.GroupProps{Class: "gap-2"}) {
		@ui.Box(ui.BoxProps{Class: "inline-block h-10 w-16 rounded-md border border-border bg-background"})
		@ui.Box(ui.BoxProps{Class: "inline-block h-10 w-16 rounded-md bg-primary"})
		@ui.Box(ui.BoxProps{Class: "inline-block h-10 w-16 rounded-md bg-muted"})
	}
}
```

```go
<body class="bg-background text-foreground">
  <!-- theme.js переключает .dark на html -->
</body>
```
