package table

import (
	"github.com/fastygo/ui/internal/registry"
	"github.com/fastygo/ui/internal/showcase/showcaseutil"
)

func init() {
	registry.Register(registry.Page{
		Slug:        "table",
		Title:       "Table",
		Section:     "components",
		Description: "Semantic data table structure.",
		Source:      "github.com/fastygo/templ/ui",
		Package:     "github.com/fastygo/templ/ui",
		Variants: []registry.Variant{
			showcaseutil.Variant("default", "Default", "", `@ui.Table(ui.TableProps{}) { @ui.TableHead … }`, previewDefault),
			showcaseutil.Variant("striped", "Compact", "", `Dense row styling via TableCell Class`, previewCompact),
		},
		API: []registry.APIField{
			{Name: "Class", Type: "string", Description: "Table wrapper utilities"},
			{Name: "TableHead", Type: "component", Description: "thead section"},
			{Name: "TableBody", Type: "component", Description: "tbody section"},
		},
		Related: []registry.RelatedLink{
			{Label: "Data Table", Href: "/docs/components/data-table"},
			{Label: "List", Href: "/docs/components/list"},
		},
	})
}
