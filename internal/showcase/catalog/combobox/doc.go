package combobox

import (
	"github.com/fastygo/ui/internal/registry"
	"github.com/fastygo/ui/internal/showcase/showcaseutil"
)

func init() {
	registry.Register(registry.Page{
		Slug:        "combobox",
		Title:       "Combobox",
		Section:     "components",
		Description: "Searchable select wireframe (data-ui8kit combobox).",
		Source:      "github.com/fastygo/templ/ui",
		Package:     "github.com/fastygo/templ/ui",
		Variants: []registry.Variant{
			showcaseutil.Variant("default", "Default", "", `@ui.Input + @ui.List { data-ui8kit=combobox }`, previewDefault),
		},
		API: []registry.APIField{
			{Name: "Input", Type: "Input", Description: "Filter field"},
			{Name: "List", Type: "List", Description: "Options listbox"},
		},
		Related: []registry.RelatedLink{
			{Label: "Select", Href: "/docs/components/select"},
			{Label: "Command", Href: "/docs/components/command"},
		},
	})
}
