package dropdownmenu

import (
	"github.com/fastygo/ui/internal/registry"
	"github.com/fastygo/ui/internal/showcase/showcaseutil"
)

func init() {
	registry.Register(registry.Page{
		Slug:        "dropdown-menu",
		Title:       "Dropdown Menu",
		Section:     "components",
		Description: "Menu triggered by a button.",
		Source:      "github.com/fastygo/templ/ui",
		Package:     "github.com/fastygo/templ/ui",
		Variants: []registry.Variant{
			showcaseutil.Variant("default", "Default", "Wireframe composition from ui primitives.", `@ui.Group { trigger + menu list }`, previewDefault),
		},
		API: []registry.APIField{
			{Name: "List", Type: "List", Description: "menu tag for items"},
		},
		Related: []registry.RelatedLink{
			{Label: "Context Menu", Href: "/docs/components/context-menu"},
			{Label: "Menubar", Href: "/docs/components/menubar"},
		},
	})
}
