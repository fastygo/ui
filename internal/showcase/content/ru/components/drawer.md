---
slug: drawer
section: components
title: "Drawer"
description: "Wireframe нижней выезжающей панели (bottom sheet)."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Alert Dialog"
    href: /docs/components/alert-dialog/
  - label: "Диалог"
    href: /docs/components/dialog/
api:
  - name: "Class"
    type: "string"
    description: "Поверхность bottom sheet"
---

Wireframe нижней выезжающей панели (bottom sheet).

## По умолчанию

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Box(ui.BoxProps{Class: "w-full max-w-md rounded-t-xl border border-border bg-card p-4"}) {
		@ui.Box(ui.BoxProps{Class: "mx-auto mb-3 h-1 w-10 rounded-full bg-muted"})
		@ui.Text(ui.TextProps{}, "Содержимое drawer.")
	}
}
```
