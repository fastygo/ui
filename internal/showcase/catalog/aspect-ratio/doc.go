package aspectratio

import (
	"github.com/fastygo/ui/internal/registry"
	"github.com/fastygo/ui/internal/showcase/showcaseutil"
)

func init() {
	registry.Register(registry.Page{
		Slug:        "aspect-ratio",
		Title:       "Aspect Ratio",
		Section:     "components",
		Description: "Fixed aspect container.",
		Source:      "github.com/fastygo/templ/ui",
		Package:     "github.com/fastygo/templ/ui",
		Variants: []registry.Variant{
			showcaseutil.Variant("default", "Default", "Wireframe composition from ui primitives.", `@ui.Box(ui.BoxProps{Class: "aspect-video …"})`, previewDefault),
		},
		API: []registry.APIField{
			{Name: "Class", Type: "string", Description: "aspect-video | aspect-square"},
		},
		Related: []registry.RelatedLink{
			{Label: "Carousel", Href: "/docs/components/carousel"},
			{Label: "Card", Href: "/docs/components/card"},
		},
	})
}
