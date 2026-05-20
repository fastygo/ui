package box

import (
	"github.com/fastygo/ui/internal/registry"
	"github.com/fastygo/ui/internal/showcase/showcaseutil"
)

func init() {
	registry.Register(registry.Page{
		Slug:        "box",
		Title:       "Box",
		Section:     "components",
		Description: "Generic block wrapper without landmark semantics.",
		Source:      "github.com/fastygo/templ/ui",
		Package:     "github.com/fastygo/templ/ui",
		Variants: []registry.Variant{
			showcaseutil.Variant("default", "Default", "", `@ui.Box(ui.BoxProps{Class: "rounded-lg border border-border p-4"}) { … }`, previewDefault),
			showcaseutil.Variant("pre", "Pre tag", "", `@ui.Box(ui.BoxProps{Tag: "pre"}) { … }`, previewPre),
		},
		API: []registry.APIField{
			{Name: "Class", Type: "string", Description: "Tailwind utilities"},
			{Name: "Tag", Type: "string", Description: "motion.div | pre | span"},
			{Name: "Attrs", Type: "templ.Attributes", Description: "Extra attributes"},
		},
		Related: []registry.RelatedLink{
			{Label: "Block", Href: "/docs/components/block"},
			{Label: "Stack", Href: "/docs/components/stack"},
		},
	})
}
