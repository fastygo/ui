package separator

import (
	"github.com/fastygo/ui/internal/registry"
	"github.com/fastygo/ui/internal/showcase/showcaseutil"
)

func init() {
	registry.Register(registry.Page{
		Slug:        "separator",
		Title:       "Separator",
		Section:     "components",
		Description: "Visual divider between sections.",
		Source:      "github.com/fastygo/templ/ui",
		Package:     "github.com/fastygo/templ/ui",
		Variants: []registry.Variant{
			showcaseutil.Variant("default", "Default", "Wireframe composition from ui primitives.", `@ui.Box(ui.BoxProps{Class: "h-px bg-border"})`, previewDefault),
		},
		API: []registry.APIField{
			{Name: "Class", Type: "string", Description: "Typically h-px bg-border"},
			{Name: "Role", Type: "string", Description: "separator"},
		},
		Related: []registry.RelatedLink{
			{Label: "Stack", Href: "/docs/components/stack"},
			{Label: "Card", Href: "/docs/components/card"},
		},
	})
}
