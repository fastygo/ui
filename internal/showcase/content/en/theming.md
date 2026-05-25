---
slug: theming
section: getting-started
title: "Theming"
description: "Semantic tokens, typography, spacing, and elevation from tweakcn-style CSS variables. Live previews match the FastyGo design-system reference; the header theme toggle switches light/dark via theme.js."
source: web/static/css/tweakcn.css
package: web/static/js/theme.js
related:
  - label: "Introduction"
    href: /docs/introduction/
  - label: "Button"
    href: /docs/primitives/button/
---

Semantic tokens, typography, spacing, and elevation come from tweakcn-style CSS variables in `web/static/css/tweakcn.css`. The header theme toggle applies `.dark` on `<html>` via `theme.js`; every preview below respects the active theme.

## Semantic tokens

Use token utilities such as `bg-background` and `text-foreground` in templ class strings. See **Colors** below for the full core and semantic palette.

```go
<body class="bg-background text-foreground">
  <!-- theme.js toggles .dark on html -->
</body>
```

## Hero

Reference board header: logo, title, subtitle, and call-to-action buttons.

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Stack(ui.StackProps{Class: "gap-4 max-w-2xl"}) {
		@ui.Group(ui.GroupProps{Class: "items-center gap-2"}) {
			@ui.Box(ui.BoxProps{Class: "flex h-10 w-10 items-center justify-center rounded-md bg-primary"}) {
				@ui.Text(ui.TextProps{Class: "font-semibold text-primary-foreground"}, "F")
			}
			@ui.Text(ui.TextProps{Class: "font-semibold text-xl"}, "FastyGo")
		}
		@ui.Title(ui.TitleProps{Order: 1}, "UI Design System")
		@ui.Text(ui.TextProps{Class: "text-sm text-muted-foreground"}, "A library of ready-made UI components built with Templ Go")
		@ui.Group(ui.GroupProps{Class: "gap-2"}) {
			@ui.Button(ui.ButtonProps{}) {
				Get Started
			}
			@ui.Button(ui.ButtonProps{Variant: "outline"}) {
				View Documentation
			}
		}
	}
}
```

## Colors

Core and semantic swatches map to CSS variables. Reference hex values (light board) are listed in prose; use semantic utilities in templ.

- **Primary** — `--primary` (#E4572A)
- **Secondary** — `--secondary` (#2E477D)
- **Accent** — `--accent` (#D6E4F3)
- **Background** — `--background` (#F8FAFC)
- **Foreground** — `--foreground` (#333333)
- **Success** — `--success` (#22C55E)
- **Warning** — `--warning` (#EAB308)
- **Destructive** — `--destructive` (#EF4444)
- **Muted** — `--muted` (#F7F9FA)
- **Border** — `--border` (#CCCCCC)

### Core palette

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Group(ui.GroupProps{Class: "flex-wrap gap-3"}) {
		@ui.Stack(ui.StackProps{Class: "gap-1 items-center"}) {
			@ui.Box(ui.BoxProps{Class: "h-10 w-16 rounded-md bg-primary border border-border"})
			@ui.Text(ui.TextProps{Class: "text-xs text-muted-foreground"}, "Primary")
		}
		@ui.Stack(ui.StackProps{Class: "gap-1 items-center"}) {
			@ui.Box(ui.BoxProps{Class: "h-10 w-16 rounded-md bg-secondary border border-border"})
			@ui.Text(ui.TextProps{Class: "text-xs text-muted-foreground"}, "Secondary")
		}
		@ui.Stack(ui.StackProps{Class: "gap-1 items-center"}) {
			@ui.Box(ui.BoxProps{Class: "h-10 w-16 rounded-md bg-accent border border-border"})
			@ui.Text(ui.TextProps{Class: "text-xs text-muted-foreground"}, "Accent")
		}
		@ui.Stack(ui.StackProps{Class: "gap-1 items-center"}) {
			@ui.Box(ui.BoxProps{Class: "h-10 w-16 rounded-md bg-background border border-border"})
			@ui.Text(ui.TextProps{Class: "text-xs text-muted-foreground"}, "Background")
		}
		@ui.Stack(ui.StackProps{Class: "gap-1 items-center"}) {
			@ui.Box(ui.BoxProps{Class: "h-10 w-16 rounded-md bg-foreground border border-border"})
			@ui.Text(ui.TextProps{Class: "text-xs text-muted-foreground"}, "Foreground")
		}
	}
}
```

### Semantic palette

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Group(ui.GroupProps{Class: "flex-wrap gap-3"}) {
		@ui.Stack(ui.StackProps{Class: "gap-1 items-center"}) {
			@ui.Box(ui.BoxProps{Class: "h-10 w-16 rounded-md bg-success border border-border"})
			@ui.Text(ui.TextProps{Class: "text-xs text-muted-foreground"}, "Success")
		}
		@ui.Stack(ui.StackProps{Class: "gap-1 items-center"}) {
			@ui.Box(ui.BoxProps{Class: "h-10 w-16 rounded-md bg-warning border border-border"})
			@ui.Text(ui.TextProps{Class: "text-xs text-muted-foreground"}, "Warning")
		}
		@ui.Stack(ui.StackProps{Class: "gap-1 items-center"}) {
			@ui.Box(ui.BoxProps{Class: "h-10 w-16 rounded-md bg-destructive border border-border"})
			@ui.Text(ui.TextProps{Class: "text-xs text-muted-foreground"}, "Destructive")
		}
		@ui.Stack(ui.StackProps{Class: "gap-1 items-center"}) {
			@ui.Box(ui.BoxProps{Class: "h-10 w-16 rounded-md bg-muted border border-border"})
			@ui.Text(ui.TextProps{Class: "text-xs text-muted-foreground"}, "Muted")
		}
		@ui.Stack(ui.StackProps{Class: "gap-1 items-center"}) {
			@ui.Box(ui.BoxProps{Class: "h-10 w-16 rounded-md bg-background border-2 border-border"})
			@ui.Text(ui.TextProps{Class: "text-xs text-muted-foreground"}, "Border")
		}
	}
}
```

## Typography

Google Sans (`--font-sans`) for UI copy; Source Code Pro (`--font-mono`) for code. Fonts load from `web/static/css/fonts.css`.

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Stack(ui.StackProps{Class: "gap-3"}) {
		@ui.Title(ui.TitleProps{Order: 1, Class: "text-3xl font-semibold"}, "Heading 1 — 32/40 Semibold")
		@ui.Title(ui.TitleProps{Order: 2, Class: "text-2xl font-semibold"}, "Heading 2 — 24/32 Semibold")
		@ui.Title(ui.TitleProps{Order: 3, Class: "text-xl font-medium"}, "Heading 3 — 20/28 Medium")
		@ui.Title(ui.TitleProps{Order: 4, Class: "text-lg font-medium"}, "Heading 4 — 18/24 Medium")
		@ui.Text(ui.TextProps{Class: "text-base"}, "Body Large — 16/24 Regular")
		@ui.Text(ui.TextProps{Class: "text-sm"}, "Body — 14/20 Regular")
		@ui.Text(ui.TextProps{Tag: "small", Class: "text-xs"}, "Small — 12/16 Regular")
		@ui.Text(ui.TextProps{Tag: "code", Class: "font-mono text-sm"}, `fmt.Println("FastyGo UI Design System")`)
	}
}
```

## Spacing and radius

Spacing scale: 4, 8, 12, 16, 24, 32, 40, 48, 64 px (Tailwind base `--spacing: 0.25rem` → 4 px per unit). Border radius: 4, 6, 8, 12, 24 px and full (pill).

### Spacing scale

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Group(ui.GroupProps{Class: "items-end gap-2 flex-wrap"}) {
		@ui.Stack(ui.StackProps{Class: "gap-1 items-center"}) {
			@ui.Box(ui.BoxProps{Class: "h-10 w-4 rounded-sm bg-primary"})
			@ui.Text(ui.TextProps{Class: "text-xs text-muted-foreground"}, "4")
		}
		@ui.Stack(ui.StackProps{Class: "gap-1 items-center"}) {
			@ui.Box(ui.BoxProps{Class: "h-10 w-8 rounded-sm bg-primary"})
			@ui.Text(ui.TextProps{Class: "text-xs text-muted-foreground"}, "8")
		}
		@ui.Stack(ui.StackProps{Class: "gap-1 items-center"}) {
			@ui.Box(ui.BoxProps{Class: "h-10 w-12 rounded-sm bg-primary"})
			@ui.Text(ui.TextProps{Class: "text-xs text-muted-foreground"}, "12")
		}
		@ui.Stack(ui.StackProps{Class: "gap-1 items-center"}) {
			@ui.Box(ui.BoxProps{Class: "h-10 w-16 rounded-sm bg-primary"})
			@ui.Text(ui.TextProps{Class: "text-xs text-muted-foreground"}, "16")
		}
		@ui.Stack(ui.StackProps{Class: "gap-1 items-center"}) {
			@ui.Box(ui.BoxProps{Class: "h-10 w-24 rounded-sm bg-primary"})
			@ui.Text(ui.TextProps{Class: "text-xs text-muted-foreground"}, "24")
		}
		@ui.Stack(ui.StackProps{Class: "gap-1 items-center"}) {
			@ui.Box(ui.BoxProps{Class: "h-10 w-32 rounded-sm bg-primary"})
			@ui.Text(ui.TextProps{Class: "text-xs text-muted-foreground"}, "32")
		}
		@ui.Stack(ui.StackProps{Class: "gap-1 items-center"}) {
			@ui.Box(ui.BoxProps{Class: "h-10 w-40 rounded-sm bg-primary"})
			@ui.Text(ui.TextProps{Class: "text-xs text-muted-foreground"}, "40")
		}
		@ui.Stack(ui.StackProps{Class: "gap-1 items-center"}) {
			@ui.Box(ui.BoxProps{Class: "h-10 w-48 rounded-sm bg-primary"})
			@ui.Text(ui.TextProps{Class: "text-xs text-muted-foreground"}, "48")
		}
		@ui.Stack(ui.StackProps{Class: "gap-1 items-center"}) {
			@ui.Box(ui.BoxProps{Class: "h-10 w-64 rounded-sm bg-primary"})
			@ui.Text(ui.TextProps{Class: "text-xs text-muted-foreground"}, "64")
		}
	}
}
```

### Border radius

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Group(ui.GroupProps{Class: "gap-3 flex-wrap"}) {
		@ui.Stack(ui.StackProps{Class: "gap-1 items-center"}) {
			@ui.Box(ui.BoxProps{Class: "h-10 w-10 rounded-sm bg-muted border border-border"})
			@ui.Text(ui.TextProps{Class: "text-xs text-muted-foreground"}, "4")
		}
		@ui.Stack(ui.StackProps{Class: "gap-1 items-center"}) {
			@ui.Box(ui.BoxProps{Class: "h-10 w-10 rounded-md bg-muted border border-border"})
			@ui.Text(ui.TextProps{Class: "text-xs text-muted-foreground"}, "6")
		}
		@ui.Stack(ui.StackProps{Class: "gap-1 items-center"}) {
			@ui.Box(ui.BoxProps{Class: "h-10 w-10 rounded-lg bg-muted border border-border"})
			@ui.Text(ui.TextProps{Class: "text-xs text-muted-foreground"}, "8")
		}
		@ui.Stack(ui.StackProps{Class: "gap-1 items-center"}) {
			@ui.Box(ui.BoxProps{Class: "h-10 w-10 rounded-xl bg-muted border border-border"})
			@ui.Text(ui.TextProps{Class: "text-xs text-muted-foreground"}, "12")
		}
		@ui.Stack(ui.StackProps{Class: "gap-1 items-center"}) {
			@ui.Box(ui.BoxProps{Class: "h-10 w-10 rounded-full bg-muted border border-border"})
			@ui.Text(ui.TextProps{Class: "text-xs text-muted-foreground"}, "Full")
		}
	}
}
```

## Shadows

Elevation tokens `--shadow-sm` through `--shadow-xl` in tweakcn map to Tailwind `shadow-*` utilities.

```templ
import "github.com/fastygo/templ/ui"

templ Example() {
	@ui.Group(ui.GroupProps{Class: "gap-4 flex-wrap"}) {
		@ui.Stack(ui.StackProps{Class: "gap-1 items-center"}) {
			@ui.Box(ui.BoxProps{Class: "h-16 w-24 rounded-md bg-card border border-border shadow-sm"})
			@ui.Text(ui.TextProps{Class: "text-xs text-muted-foreground"}, "sm")
		}
		@ui.Stack(ui.StackProps{Class: "gap-1 items-center"}) {
			@ui.Box(ui.BoxProps{Class: "h-16 w-24 rounded-md bg-card border border-border shadow-md"})
			@ui.Text(ui.TextProps{Class: "text-xs text-muted-foreground"}, "md")
		}
		@ui.Stack(ui.StackProps{Class: "gap-1 items-center"}) {
			@ui.Box(ui.BoxProps{Class: "h-16 w-24 rounded-md bg-card border border-border shadow-lg"})
			@ui.Text(ui.TextProps{Class: "text-xs text-muted-foreground"}, "lg")
		}
		@ui.Stack(ui.StackProps{Class: "gap-1 items-center"}) {
			@ui.Box(ui.BoxProps{Class: "h-16 w-24 rounded-md bg-card border border-border shadow-xl"})
			@ui.Text(ui.TextProps{Class: "text-xs text-muted-foreground"}, "xl")
		}
	}
}
```
