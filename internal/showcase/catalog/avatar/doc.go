package avatar

import (
	"github.com/fastygo/ui/internal/registry"
	"github.com/fastygo/ui/internal/showcase/showcaseutil"
)

func init() {
	registry.Register(registry.Page{
		Slug:        "avatar",
		Title:       "Avatar",
		Section:     "components",
		Description: "User avatar placeholder.",
		Source:      "github.com/fastygo/templ/ui",
		Package:     "github.com/fastygo/templ/ui",
		Variants: []registry.Variant{
			showcaseutil.Variant("default", "Default", "Wireframe composition from ui primitives.", `@ui.Box(ui.BoxProps{Class: "rounded-full …"})`, previewDefault),
		},
		API: []registry.APIField{
			{Name: "Class", Type: "string", Description: "rounded-full size utilities"},
		},
		Related: []registry.RelatedLink{
			{Label: "Badge", Href: "/docs/components/badge"},
			{Label: "Card", Href: "/docs/components/card"},
		},
	})
}
