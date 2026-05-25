---
slug: language-toggle
section: components
title: "Language Toggle"
description: "Переключатель локали (пакет toggles приложения)."
source: github.com/fastygo/templ/ui
package: github.com/fastygo/templ/ui
related:
  - label: "Кнопка"
    href: /docs/primitives/button/
  - label: "Иконка"
    href: /docs/components/icon/
api:
  - name: "CurrentLabel"
    type: "string"
    description: "Видимая подпись"
  - name: "NextHref"
    type: "string"
    description: "Ссылка на следующую локаль"
  - name: "CurrentLocale"
    type: "string"
    description: "Код активной локали"
---

Переключатель локали (пакет toggles приложения).

## По умолчанию

```templ
import "github.com/fastygo/framework/pkg/web/view"
import toggles "github.com/fastygo/ui/internal/ui/components/toggles"

templ Example() {
	@toggles.LanguageToggle(view.LanguageToggleData{
		CurrentLabel:     "RU",
		CurrentLocale:    "ru",
		NextLocale:       "en",
		NextHref:         "/?lang=en",
		DefaultLocale:    "en",
		AvailableLocales: []string{"en", "ru"},
		Label:            "Переключить язык",
	})
}
```
