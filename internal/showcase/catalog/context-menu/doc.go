package contextmenu

import (
	"github.com/fastygo/ui/internal/registry"
	"github.com/fastygo/ui/internal/showcase/showcaseutil"
)

func init() {
	registry.Register(registry.Page{
		Slug:        "context-menu",
		Title:       "Context Menu",
		Section:     "components",
		Description: "Stub menu list on right-click target.",
		Source:      "github.com/fastygo/templ/ui",
		Package:     "github.com/fastygo/templ/ui",
		Variants: []registry.Variant{
			showcaseutil.Variant("default", "Default", "Wireframe composition from ui primitives.", `@ui.Box + @ui.List[menu]`, previewDefault),
		},
		API: []registry.APIField{
			{Name: "List", Type: "List", Description: "menu tag"},
		},
		Related: []registry.RelatedLink{
			{Label: "Dropdown Menu", Href: "/docs/components/dropdown-menu"},
			{Label: "Menubar", Href: "/docs/components/menubar"},
		},
	})
}
