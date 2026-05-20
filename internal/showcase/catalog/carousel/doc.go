package carousel

import (
	"github.com/fastygo/ui/internal/registry"
	"github.com/fastygo/ui/internal/showcase/showcaseutil"
)

func init() {
	registry.Register(registry.Page{
		Slug:        "carousel",
		Title:       "Carousel",
		Section:     "components",
		Description: "Horizontal scrolling list.",
		Source:      "github.com/fastygo/templ/ui",
		Package:     "github.com/fastygo/templ/ui",
		Variants: []registry.Variant{
			showcaseutil.Variant("default", "Default", "", `@ui.Group { slides }`, previewDefault),
		},
		API: []registry.APIField{
			{Name: "Group", Type: "Group", Description: "Horizontal flex row"},
		},
		Related: []registry.RelatedLink{
			{Label: "Aspect Ratio", Href: "/docs/components/aspect-ratio"},
			{Label: "Card", Href: "/docs/components/card"},
		},
	})
}
