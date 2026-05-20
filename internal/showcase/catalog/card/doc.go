package card

import (
	"github.com/fastygo/ui/internal/registry"
	"github.com/fastygo/ui/internal/showcase/showcaseutil"
)

func init() {
	registry.Register(registry.Page{
		Slug:        "card",
		Title:       "Card",
		Section:     "components",
		Description: "Grouped content with header, body, and footer.",
		Source:      "github.com/fastygo/templ/ui",
		Package:     "github.com/fastygo/templ/ui",
		Variants: []registry.Variant{
			showcaseutil.Variant("default", "Default", "", `@cmp.Card(cmp.CardProps{}) { @cmp.CardHeader … }`, previewDefault),
			showcaseutil.Variant("footer", "With footer", "", `@cmp.CardFooter …`, previewFooter),
		},
		API: []registry.APIField{
			{Name: "Variant", Type: "string", Description: "Surface variant"},
			{Name: "CardHeader", Type: "component", Description: "Title area"},
			{Name: "CardContent", Type: "component", Description: "Main body"},
		},
		Related: []registry.RelatedLink{
			{Label: "Alert", Href: "/docs/components/alert"},
			{Label: "Dialog", Href: "/docs/components/dialog"},
		},
	})
}
