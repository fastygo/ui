package hovercard

import (
	"github.com/fastygo/ui/internal/registry"
	"github.com/fastygo/ui/internal/showcase/showcaseutil"
)

func init() {
	registry.Register(registry.Page{
		Slug:        "hover-card",
		Title:       "Hover Card",
		Section:     "components",
		Description: "Rich preview on hover.",
		Source:      "github.com/fastygo/templ/ui",
		Package:     "github.com/fastygo/templ/ui",
		Variants: []registry.Variant{
			showcaseutil.Variant("default", "Default", "Wireframe composition from ui primitives.", `@ui.Group { trigger + card }`, previewDefault),
		},
		API: []registry.APIField{
			{Name: "Box", Type: "Box", Description: "Preview card surface"},
		},
		Related: []registry.RelatedLink{
			{Label: "Popover", Href: "/docs/components/popover"},
			{Label: "Tooltip", Href: "/docs/components/tooltip"},
		},
	})
}
