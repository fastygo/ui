package radio

import (
	"github.com/fastygo/ui/internal/registry"
	"github.com/fastygo/ui/internal/showcase/showcaseutil"
)

func init() {
	registry.Register(registry.Page{
		Slug:        "radio",
		Title:       "Radio",
		Section:     "components",
		Description: "Single choice within a group.",
		Source:      "github.com/fastygo/templ/ui",
		Package:     "github.com/fastygo/templ/ui",
		Variants: []registry.Variant{
			showcaseutil.Variant("default", "Default", "", `@ui.Radio(ui.RadioProps{Name: "plan", Value: "free"})`, previewDefault),
			showcaseutil.Variant("checked", "Selected", "", `@ui.Radio(ui.RadioProps{Checked: true})`, previewChecked),
		},
		API: []registry.APIField{
			{Name: "Name", Type: "string", Description: "Group name"},
			{Name: "Value", Type: "string", Description: "Option value"},
			{Name: "Checked", Type: "bool", Description: "Selected state"},
		},
		Related: []registry.RelatedLink{
			{Label: "Checkbox", Href: "/docs/components/checkbox"},
			{Label: "Select", Href: "/docs/components/select"},
		},
	})
}
