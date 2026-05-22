package doclocale

import (
	"strings"

	"github.com/fastygo/ui/internal/ui/components/toggles"
)

// BuildLanguageSwitch builds the header En / Ru control for a docs page.
func (r Routing) BuildLanguageSwitch(currentLocale, activePath, ariaLabel string) toggles.LanguageSwitchProps {
	r = r.Normalize()
	current := strings.ToLower(strings.TrimSpace(currentLocale))
	if current == "" {
		current = r.Default
	}
	var items []toggles.LanguageSwitchItem
	for _, loc := range r.Locales {
		items = append(items, toggles.LanguageSwitchItem{
			Locale: loc,
			Label:  r.Label(loc),
			Href:   r.AlternatePublicPath(activePath, loc),
			Active: loc == current,
		})
	}
	return toggles.LanguageSwitchProps{
		AriaLabel: ariaLabel,
		Items:     items,
	}
}
