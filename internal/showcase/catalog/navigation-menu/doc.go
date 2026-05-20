package navigationmenu

import (
	"github.com/fastygo/ui/internal/registry"
	"github.com/fastygo/ui/internal/showcase/showcaseutil"
)

func init() {
	registry.Register(registry.Page{
		Slug:        "navigation-menu",
		Title:       "Navigation Menu",
		Section:     "components",
		Description: "Site navigation with sections.",
		Source:      "github.com/fastygo/templ/ui",
		Package:     "github.com/fastygo/templ/ui",
		Variants: []registry.Variant{
			showcaseutil.Variant("default", "Default", "Wireframe composition from ui primitives.", `@ui.List[Tag=menu] { links }`, previewDefault),
		},
		API: []registry.APIField{
			{Name: "Tag", Type: "string", Description: "menu for menubar-style lists"},
		},
		Related: []registry.RelatedLink{
			{Label: "Menubar", Href: "/docs/components/menubar"},
			{Label: "Breadcrumb", Href: "/docs/components/breadcrumb"},
		},
	})
}
