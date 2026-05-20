package sheet

import (
	"github.com/fastygo/ui/internal/registry"
	"github.com/fastygo/ui/internal/showcase/showcaseutil"
)

func init() {
	registry.Register(registry.Page{
		Slug:        "sheet",
		Title:       "Sheet",
		Section:     "components",
		Description: "Slide-over panel wireframe.",
		Source:      "github.com/fastygo/templ/ui",
		Package:     "github.com/fastygo/templ/ui",
		Variants: []registry.Variant{
			showcaseutil.Variant("default", "Default", "Wireframe composition from ui primitives.", `@ui.Box { header + body }`, previewDefault),
		},
		API: []registry.APIField{
			{Name: "Class", Type: "string", Description: "Panel surface utilities"},
		},
		Related: []registry.RelatedLink{
			{Label: "Dialog", Href: "/docs/components/dialog"},
			{Label: "Drawer", Href: "/docs/components/drawer"},
		},
	})
}
