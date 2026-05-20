package block

import (
	"github.com/fastygo/ui/internal/registry"
	"github.com/fastygo/ui/internal/showcase/showcaseutil"
)

func init() {
	registry.Register(registry.Page{
		Slug:        "block",
		Title:       "Block",
		Section:     "components",
		Description: "Top-level landmark sections (do not nest Block in Block).",
		Source:      "github.com/fastygo/templ/ui",
		Package:     "github.com/fastygo/templ/ui",
		Variants: []registry.Variant{
			showcaseutil.Variant("main", "Main", "", `@ui.Block(ui.BlockProps{Tag: "main"}) { … }`, previewMain),
			showcaseutil.Variant("aside", "Aside", "", `@ui.Block(ui.BlockProps{Tag: "aside"}) { … }`, previewAside),
		},
		API: []registry.APIField{
			{Name: "Tag", Type: "string", Description: "main | section | aside | nav | …"},
			{Name: "Class", Type: "string", Description: "Tailwind utilities"},
			{Name: "Attrs", Type: "templ.Attributes", Description: "Extra attributes"},
		},
		Related: []registry.RelatedLink{
			{Label: "Box", Href: "/docs/components/box"},
			{Label: "Container", Href: "/docs/components/container"},
		},
	})
}
