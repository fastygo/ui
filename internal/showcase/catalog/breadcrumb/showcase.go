package breadcrumb

import (
	"context"
	"io"

	cmp "github.com/fastygo/templ/components"
)

var crumbItems = []cmp.BreadcrumbItem{{Label: "Home", Href: "/"}, {Label: "Docs", Href: "/docs"}, {Label: "Button", Current: true}}

func previewDefault(ctx context.Context, w io.Writer) error {
	return cmp.Breadcrumb(cmp.BreadcrumbProps{Items: crumbItems}).Render(ctx, w)
}
func previewCurrent(ctx context.Context, w io.Writer) error {
	items := []cmp.BreadcrumbItem{{Label: "Home", Href: "/"}, {Label: "Settings", Current: true}}
	return cmp.Breadcrumb(cmp.BreadcrumbProps{Items: items}).Render(ctx, w)
}
