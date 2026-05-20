package selectfield

import (
	"github.com/fastygo/ui/internal/registry"
	"github.com/fastygo/ui/internal/showcase/showcaseutil"
)

func init() {
	registry.Register(registry.Page{
		Slug:        "select",
		Title:       "Select",
		Section:     "components",
		Description: "Native select dropdown (ui.Select / selectfield).",
		Source:      "github.com/fastygo/templ/ui",
		Package:     "github.com/fastygo/templ/ui",
		Variants: []registry.Variant{
			showcaseutil.Variant("default", "Default", "", `@ui.Select(ui.SelectProps{Name: "role", Options: opts})`, previewDefault),
			showcaseutil.Variant("disabled", "Disabled", "", `@ui.Select(ui.SelectProps{Disabled: true})`, previewDisabled),
		},
		API: []registry.APIField{
			{Name: "Options", Type: "[]ui.Option", Description: "Value/label pairs"},
			{Name: "Name", Type: "string", Description: "Form field name"},
			{Name: "Value", Type: "string", Description: "Selected value"},
		},
		Related: []registry.RelatedLink{
			{Label: "Combobox", Href: "/docs/components/combobox"},
			{Label: "Radio", Href: "/docs/components/radio"},
		},
	})
}
