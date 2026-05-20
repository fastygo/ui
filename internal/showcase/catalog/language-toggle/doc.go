package languagetoggle

import (
	"github.com/fastygo/ui/internal/registry"
	"github.com/fastygo/ui/internal/showcase/showcaseutil"
)

func init() {
	registry.Register(registry.Page{
		Slug:        "language-toggle",
		Title:       "Language Toggle",
		Section:     "components",
		Description: "Locale switcher (app toggles package).",
		Source:      "github.com/fastygo/templ/ui",
		Package:     "github.com/fastygo/templ/ui",
		Variants: []registry.Variant{
			showcaseutil.Variant("default", "Default", "", `@toggles.LanguageToggle(data)`, previewDefault),
		},
		API: []registry.APIField{
			{Name: "CurrentLabel", Type: "string", Description: "Visible label"},
			{Name: "NextHref", Type: "string", Description: "Link to next locale"},
			{Name: "CurrentLocale", Type: "string", Description: "Active locale code"},
		},
		Related: []registry.RelatedLink{
			{Label: "Button", Href: "/docs/components/button"},
			{Label: "Icon", Href: "/docs/components/icon"},
		},
	})
}
