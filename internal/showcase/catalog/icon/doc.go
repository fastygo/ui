package iconcatalog

import (
	"github.com/fastygo/ui/internal/registry"
	"github.com/fastygo/ui/internal/showcase/showcaseutil"
)

func init() {
	registry.Register(registry.Page{
		Slug:        "icon",
		Title:       "Icon",
		Section:     "components",
		Description: "Latty icon mask (app registry).",
		Source:      "github.com/fastygo/ui/internal/ui/components/icon",
		Package:     "github.com/fastygo/ui/internal/ui/components/icon",
		Variants: []registry.Variant{
			showcaseutil.Variant("default", "Default", "", `@icon.Icon(icon.IconProps{Name: "home"})`, previewDefault),
			showcaseutil.Variant("sizes", "Sizes", "", `Size: xs | sm | md | lg`, previewSizes),
		},
		API: []registry.APIField{
			{Name: "Name", Type: "string", Description: "Latty icon name"},
			{Name: "Size", Type: "string", Description: "xs | sm | md | lg"},
			{Name: "Class", Type: "string", Description: "Extra utilities"},
		},
		Related: []registry.RelatedLink{
			{Label: "Button", Href: "/docs/components/button"},
			{Label: "Badge", Href: "/docs/components/badge"},
		},
	})
}
