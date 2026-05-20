package breadcrumb

import (
	"github.com/fastygo/ui/internal/registry"
	"github.com/fastygo/ui/internal/showcase/showcaseutil"
)

func init() {
	registry.Register(registry.Page{
		Slug:        "breadcrumb",
		Title:       "Breadcrumb",
		Section:     "components",
		Description: "Hierarchy navigation trail.",
		Source:      "github.com/fastygo/templ/ui",
		Package:     "github.com/fastygo/templ/ui",
		Variants: []registry.Variant{
			showcaseutil.Variant("default", "Default", "", `@cmp.Breadcrumb(cmp.BreadcrumbProps{Items: items})`, previewDefault),
			showcaseutil.Variant("current", "Current page", "", `Last item with Current: true`, previewCurrent),
		},
		API: []registry.APIField{
			{Name: "Items", Type: "[]BreadcrumbItem", Description: "Label, Href, Current, Disabled"},
			{Name: "Class", Type: "string", Description: "Nav wrapper utilities"},
		},
		Related: []registry.RelatedLink{
			{Label: "Navigation Menu", Href: "/docs/components/navigation-menu"},
			{Label: "Tabs", Href: "/docs/components/tabs"},
		},
	})
}
