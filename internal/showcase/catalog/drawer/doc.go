package drawer

import (
	"github.com/fastygo/ui/internal/registry"
	"github.com/fastygo/ui/internal/showcase/showcaseutil"
)

func init() {
	registry.Register(registry.Page{
		Slug:        "drawer",
		Title:       "Drawer",
		Section:     "components",
		Description: "Bottom sheet drawer wireframe.",
		Source:      "github.com/fastygo/templ/ui",
		Package:     "github.com/fastygo/templ/ui",
		Variants: []registry.Variant{
			showcaseutil.Variant("default", "Default", "Wireframe composition from ui primitives.", `@ui.Box anchored to bottom`, previewDefault),
		},
		API: []registry.APIField{
			{Name: "Class", Type: "string", Description: "Bottom sheet surface"},
		},
		Related: []registry.RelatedLink{
			{Label: "Sheet", Href: "/docs/components/sheet"},
			{Label: "Dialog", Href: "/docs/components/dialog"},
		},
	})
}
