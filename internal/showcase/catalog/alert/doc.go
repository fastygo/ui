package alert

import (
	"github.com/fastygo/ui/internal/registry"
	"github.com/fastygo/ui/internal/showcase/showcaseutil"
)

func init() {
	registry.Register(registry.Page{
		Slug:        "alert",
		Title:       "Alert",
		Section:     "components",
		Description: "Callout for important messages.",
		Source:      "github.com/fastygo/templ/ui",
		Package:     "github.com/fastygo/templ/ui",
		Variants: []registry.Variant{
			showcaseutil.Variant("default", "Default", "", `@cmp.Alert(cmp.AlertProps{}) { … }`, previewDefault),
			showcaseutil.Variant("destructive", "Destructive", "", `@cmp.Alert(cmp.AlertProps{Variant: "destructive"}) { … }`, previewDestructive),
		},
		API: []registry.APIField{
			{Name: "Variant", Type: "string", Description: "default | destructive"},
			{Name: "Class", Type: "string", Description: "Extra utilities"},
		},
		Related: []registry.RelatedLink{
			{Label: "Card", Href: "/docs/components/card"},
			{Label: "Badge", Href: "/docs/components/badge"},
		},
	})
}
