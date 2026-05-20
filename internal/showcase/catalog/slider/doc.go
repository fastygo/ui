package slider

import (
	"github.com/fastygo/ui/internal/registry"
	"github.com/fastygo/ui/internal/showcase/showcaseutil"
)

func init() {
	registry.Register(registry.Page{
		Slug:        "slider",
		Title:       "Slider",
		Section:     "components",
		Description: "Native range input styled via ui.Input.",
		Source:      "github.com/fastygo/templ/ui",
		Package:     "github.com/fastygo/templ/ui",
		Variants: []registry.Variant{
			showcaseutil.Variant("default", "Default", "", `@ui.Input(ui.InputProps{Type: "range", Min: "0", Max: "100"})`, previewDefault),
			showcaseutil.Variant("value", "With value", "", `@ui.Input(ui.InputProps{Type: "range", Value: "50"})`, previewValue),
		},
		API: []registry.APIField{
			{Name: "Min", Type: "string", Description: "Minimum value"},
			{Name: "Max", Type: "string", Description: "Maximum value"},
			{Name: "Value", Type: "string", Description: "Current value"},
		},
		Related: []registry.RelatedLink{
			{Label: "Switch", Href: "/docs/components/switch"},
			{Label: "Input", Href: "/docs/components/input"},
		},
	})
}
