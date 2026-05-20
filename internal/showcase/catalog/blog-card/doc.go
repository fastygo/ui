package blogcardcatalog

import (
	"context"
	"io"

	"github.com/fastygo/ui/internal/registry"
	"github.com/fastygo/ui/internal/showcase/showcaseutil"
	"github.com/fastygo/ui/internal/ui/components/blogcard"
)

func init() {
	registry.Register(registry.Page{
		Slug:        "blog-card",
		Title:       "Blog Card",
		Section:     "components",
		Description: "Reusable blog list cards with media placeholder — vertical and horizontal layouts.",
		Source:      "github.com/fastygo/ui/internal/ui/components/blogcard",
		Package:     "github.com/fastygo/ui/internal/ui/components/blogcard",
		Variants: []registry.Variant{
			showcaseutil.Variant(
				"vertical",
				"Vertical",
				"Stacked media and copy for grids and feeds.",
				`@blogcard.VerticalBlogCard(blogcard.DefaultVertical())`,
				previewVertical,
			),
			showcaseutil.Variant(
				"horizontal",
				"Horizontal",
				"Side-by-side media and copy for dense lists.",
				`@blogcard.HorizontalBlogCard(blogcard.DefaultHorizontal())`,
				previewHorizontal,
			),
		},
		API: []registry.APIField{
			{Name: "Title", Type: "string", Description: "Article title"},
			{Name: "Excerpt", Type: "string", Description: "Short summary"},
			{Name: "Href", Type: "string", Description: "Read-more link target"},
			{Name: "MediaURL", Type: "string", Description: "Reserved for future media primitive; wireframe uses muted placeholder"},
			{Name: "MediaAlt", Type: "string", Description: "Accessible label for media placeholder"},
			{Name: "DateLabel", Type: "string", Description: "Published date (pre-formatted)"},
		},
		Related: []registry.RelatedLink{
			{Label: "Card", Href: "/docs/components/card"},
			{Label: "Stack", Href: "/docs/components/stack"},
		},
	})
}

func previewVertical(ctx context.Context, w io.Writer) error {
	return blogcard.VerticalBlogCard(blogcard.DefaultVertical()).Render(ctx, w)
}

func previewHorizontal(ctx context.Context, w io.Writer) error {
	return blogcard.HorizontalBlogCard(blogcard.DefaultHorizontal()).Render(ctx, w)
}
