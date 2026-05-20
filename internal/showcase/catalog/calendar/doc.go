package calendar

import (
	"github.com/fastygo/ui/internal/registry"
	"github.com/fastygo/ui/internal/showcase/showcaseutil"
)

func init() {
	registry.Register(registry.Page{
		Slug:        "calendar",
		Title:       "Calendar",
		Section:     "components",
		Description: "Date grid placeholder.",
		Source:      "github.com/fastygo/templ/ui",
		Package:     "github.com/fastygo/templ/ui",
		Variants: []registry.Variant{
			showcaseutil.Variant("default", "Default", "Wireframe composition from ui primitives.", `@ui.Box { month label + day grid }`, previewDefault),
		},
		API: []registry.APIField{
			{Name: "Class", Type: "string", Description: "Calendar frame utilities"},
		},
		Related: []registry.RelatedLink{
			{Label: "Command", Href: "/docs/components/command"},
			{Label: "Popover", Href: "/docs/components/popover"},
		},
	})
}
