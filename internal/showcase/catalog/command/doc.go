package command

import (
	"github.com/fastygo/ui/internal/registry"
	"github.com/fastygo/ui/internal/showcase/showcaseutil"
)

func init() {
	registry.Register(registry.Page{
		Slug:        "command",
		Title:       "Command",
		Section:     "components",
		Description: "Command palette: search + list.",
		Source:      "github.com/fastygo/templ/ui",
		Package:     "github.com/fastygo/templ/ui",
		Variants: []registry.Variant{
			showcaseutil.Variant("default", "Default", "", `@ui.Input + @ui.List`, previewDefault),
		},
		API: []registry.APIField{
			{Name: "Input", Type: "Input", Description: "Search field"},
			{Name: "List", Type: "List", Description: "Commands"},
		},
		Related: []registry.RelatedLink{
			{Label: "Combobox", Href: "/docs/components/combobox"},
			{Label: "Dialog", Href: "/docs/components/dialog"},
		},
	})
}
