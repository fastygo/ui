---
slug: language-toggle
section: components
title: "Language Toggle"
description: "Locale switcher (app toggles package)."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Button"
    href: /docs/components/button/
  - label: "Icon"
    href: /docs/components/icon/
api:
  - name: "CurrentLabel"
    type: "string"
    description: "Visible label"
  - name: "NextHref"
    type: "string"
    description: "Link to next locale"
  - name: "CurrentLocale"
    type: "string"
    description: "Active locale code"
---

Locale switcher (app toggles package).

## Default

```templ
import "github.com/fastygo/framework/pkg/web/view"
import toggles "github.com/fastygo/ui/internal/ui/components/toggles"

templ Example() {
	@toggles.LanguageToggle(view.LanguageToggleData{
		CurrentLabel:     "EN",
		CurrentLocale:    "en",
		NextLocale:       "ru",
		NextHref:         "/?lang=ru",
		DefaultLocale:    "en",
		AvailableLocales: []string{"en", "ru"},
		Label:            "Switch language",
	})
}
```
