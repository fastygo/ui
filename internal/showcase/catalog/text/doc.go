package text

import (
	"github.com/fastygo/ui/internal/registry"
	"github.com/fastygo/ui/internal/showcase/showcaseutil"
)

func init() {
	registry.Register(registry.Page{
		Slug:        "text",
		Title:       "Text",
		Section:     "components",
		Description: "Inline or block text with configurable tag.",
		Source:      "github.com/fastygo/templ/ui",
		Package:     "github.com/fastygo/templ/ui",
		Variants: []registry.Variant{
			showcaseutil.Variant("default", "Paragraph", "", `@ui.Text(ui.TextProps{}, "Body copy.")`, previewDefault),
			showcaseutil.Variant("muted", "Muted", "", `@ui.Text(ui.TextProps{Class: "text-sm text-muted-foreground"}, "…")`, previewMuted),
			showcaseutil.Variant("code", "Code", "", `@ui.Text(ui.TextProps{Tag: "code"}, "npm install")`, previewCode),
		},
		API: []registry.APIField{
			{Name: "Tag", Type: "string", Description: "p | span | code | …"},
			{Name: "Class", Type: "string", Description: "Tailwind utilities"},
		},
		Related: []registry.RelatedLink{
			{Label: "Title", Href: "/docs/components/title"},
			{Label: "Label", Href: "/docs/components/label"},
		},
	})
}
