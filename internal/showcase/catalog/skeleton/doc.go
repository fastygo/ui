package skeleton

import (
	"github.com/fastygo/ui/internal/registry"
	"github.com/fastygo/ui/internal/showcase/showcaseutil"
)

func init() {
	registry.Register(registry.Page{
		Slug:        "skeleton",
		Title:       "Skeleton",
		Section:     "components",
		Description: "Loading placeholder blocks.",
		Source:      "github.com/fastygo/templ/ui",
		Package:     "github.com/fastygo/templ/ui",
		Variants: []registry.Variant{
			showcaseutil.Variant("default", "Default", "Wireframe composition from ui primitives.", `@ui.Box(ui.BoxProps{Class: "animate-pulse bg-muted"})`, previewDefault),
		},
		API: []registry.APIField{
			{Name: "Class", Type: "string", Description: "animate-pulse bg-muted shapes"},
		},
		Related: []registry.RelatedLink{
			{Label: "Progress", Href: "/docs/components/progress"},
			{Label: "Card", Href: "/docs/components/card"},
		},
	})
}
