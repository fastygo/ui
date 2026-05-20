package menubar

import (
	"github.com/fastygo/ui/internal/registry"
	"github.com/fastygo/ui/internal/showcase/showcaseutil"
)

func init() {
	registry.Register(registry.Page{
		Slug:        "menubar",
		Title:       "Menubar",
		Section:     "components",
		Description: "Horizontal menu bar.",
		Source:      "github.com/fastygo/templ/ui",
		Package:     "github.com/fastygo/templ/ui",
		Variants: []registry.Variant{
			showcaseutil.Variant("default", "Default", "Wireframe composition from ui primitives.", `@ui.Group + menu triggers`, previewDefault),
		},
		API: []registry.APIField{
			{Name: "Group", Type: "Group", Description: "Horizontal button row"},
		},
		Related: []registry.RelatedLink{
			{Label: "Navigation Menu", Href: "/docs/components/navigation-menu"},
			{Label: "Dropdown Menu", Href: "/docs/components/dropdown-menu"},
		},
	})
}
