package togglegroup

import (
	"github.com/fastygo/ui/internal/registry"
	"github.com/fastygo/ui/internal/showcase/showcaseutil"
)

func init() {
	registry.Register(registry.Page{
		Slug:        "toggle-group",
		Title:       "Toggle Group",
		Section:     "components",
		Description: "Grouped toggle buttons.",
		Source:      "github.com/fastygo/templ/ui",
		Package:     "github.com/fastygo/templ/ui",
		Variants: []registry.Variant{
			showcaseutil.Variant("default", "Default", "", `@ui.Group { @ui.Button[aria-pressed] … }`, previewDefault),
		},
		API: []registry.APIField{
			{Name: "Class", Type: "string", Description: "Group layout utilities"},
		},
		Related: []registry.RelatedLink{
			{Label: "Toggle", Href: "/docs/components/toggle"},
			{Label: "Tabs", Href: "/docs/components/tabs"},
		},
	})
}
