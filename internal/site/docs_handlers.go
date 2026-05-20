package site

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/fastygo/framework/pkg/web"
	"github.com/fastygo/ui/internal/registry"
	"github.com/fastygo/ui/internal/views"
	viewdocs "github.com/fastygo/ui/internal/views/docs"
)

func (f *Feature) layoutDataDocs(ctx context.Context, r *http.Request, title, active string) views.LayoutData {
	d := f.layoutData(ctx, r, title, active)
	d.NavItems = docsNavItems(active)
	return d
}

func (f *Feature) registerDocsRoutes(mux *http.ServeMux) {
	f.registerStaticDocsRoutes(mux)
	if _, err := os.Stat(staticDocsRoot); err == nil {
		return
	}
	mux.HandleFunc("GET /docs/components/{slug}", f.getDocsComponentSlug)
	mux.HandleFunc("GET /docs/blocks/{slug}", f.getDocsBlockSlug)
	mux.HandleFunc("GET /docs/{slug}", f.getDocsSlug)
	mux.HandleFunc("GET /docs", f.getDocsIndex)
}

func (f *Feature) getDocsIndex(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fix := f.fixtureLocale(ctx)
	layout := f.layoutDataDocs(ctx, r, fix.Docs.IndexTitle, "/docs")
	body := viewdocs.Index(viewdocs.IndexData{
		Title:       fix.Docs.IndexTitle,
		Description: fix.Docs.IndexDescription,
		Sections:    buildDocsIndexSections(),
	})
	_ = web.Render(ctx, w, views.SiteShell(layout, body))
}

func (f *Feature) getDocsSlug(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	if page, ok := registry.PageBySlug(slug); ok && page.Section == "getting-started" {
		f.renderDocsPage(w, r, page)
		return
	}
	http.NotFound(w, r)
}

func (f *Feature) getDocsComponentSlug(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	if page, ok := registry.PageBySlug(slug); ok {
		f.renderDocsPage(w, r, page)
		return
	}
	http.NotFound(w, r)
}

func (f *Feature) getDocsBlockSlug(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	if page, ok := registry.PageBySlug(slug); ok {
		f.renderDocsPage(w, r, page)
		return
	}
	http.NotFound(w, r)
}

func (f *Feature) renderDocsPage(w http.ResponseWriter, r *http.Request, page registry.Page) {
	ctx := r.Context()
	layout := f.layoutDataDocs(ctx, r, page.Title, page.Path)
	_ = web.Render(ctx, w, views.SiteShell(layout, viewdocs.ComponentDoc(page)))
}

func buildDocsIndexSections() []viewdocs.IndexSection {
	var out []viewdocs.IndexSection
	for _, sec := range registry.Sections() {
		var links []viewdocs.IndexLink
		for _, p := range registry.PagesInSection(sec.ID) {
			links = append(links, viewdocs.IndexLink{
				Title:       p.Title,
				Description: truncate(p.Description, 120),
				Href:        p.Path,
			})
		}
		out = append(out, viewdocs.IndexSection{Label: sec.Label, Links: links})
	}
	return out
}

func truncate(s string, max int) string {
	s = strings.TrimSpace(s)
	if len(s) <= max {
		return s
	}
	return s[:max] + "…"
}
