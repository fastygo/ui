package site

import (
	"context"
	"net/http"
	"strings"

	"github.com/fastygo/framework/pkg/app"
	"github.com/fastygo/framework/pkg/web"
	"github.com/fastygo/framework/pkg/web/locale"
	"github.com/fastygo/ui/internal/fixtures"
	"github.com/fastygo/ui/internal/ui/components/toggles"
	"github.com/fastygo/ui/internal/ui/layout"
	"github.com/fastygo/ui/internal/views"
)

// Feature wires public HTTP routes for the default app shell (sidebar, i18n, theme).
type Feature struct {
	available      []string
	defaultLocale  string
	staticDocsRoot string
}

// NewFeature constructs the site feature.
func NewFeature(available []string, defaultLocale, staticDocsRoot string) *Feature {
	return &Feature{
		available:      available,
		defaultLocale:  defaultLocale,
		staticDocsRoot: resolveDocsRoot(staticDocsRoot),
	}
}

// SetNavItems implements app.NavProvider.
func (f *Feature) SetNavItems(_ []app.NavItem) {}

// ID implements app.Feature.
func (f *Feature) ID() string {
	return "site"
}

// NavItems implements app.Feature.
func (f *Feature) NavItems() []app.NavItem {
	return nil
}

func (f *Feature) fixtureLocale(ctx context.Context) fixtures.Locale {
	code := locale.From(ctx)
	if code == "" {
		code = f.defaultLocale
	}
	loc, err := fixtures.LoadLocale(code)
	if err != nil {
		loc, _ = fixtures.LoadLocale(f.defaultLocale)
	}
	return loc
}


func (f *Feature) assetPaths() views.AssetPaths {
	return views.AssetPaths{
		CSS:     "/static/css/app.css",
		ThemeJS: "/static/js/theme.js",
		AppJS:   "/static/js/ui8kit.js",
	}
}

func (f *Feature) layoutData(ctx context.Context, r *http.Request, title, active string) views.LayoutData {
	fix := f.fixtureLocale(ctx)
	return views.LayoutData{
		PageTitle:      title,
		Lang:           locale.From(ctx),
		Brand:          fix.Brand,
		Active:         active,
		NavItems:       f.siteNav(ctx, fix),
		Assets:         f.assetPaths(),
		Theme: layout.ThemeToggleProps{
			Label:              fix.Theme.Label,
			SwitchToDarkLabel:  fix.Theme.SwitchToDarkLabel,
			SwitchToLightLabel: fix.Theme.SwitchToLight,
		},
		LanguageSwitch: f.languageSwitch(ctx, r, fix),
	}
}

func (f *Feature) languageSwitch(ctx context.Context, r *http.Request, fix fixtures.Locale) toggles.LanguageSwitchProps {
	routing := f.docsRouting()
	current := strings.ToLower(strings.TrimSpace(locale.From(ctx)))
	if current == "" {
		current = routing.Default
	}
	var items []toggles.LanguageSwitchItem
	for _, loc := range routing.Locales {
		items = append(items, toggles.LanguageSwitchItem{
			Locale: loc,
			Label:  routing.Label(loc),
			Href:   locale.BuildLangHref(r, loc, f.defaultLocale),
			Active: loc == current,
		})
	}
	return toggles.LanguageSwitchProps{
		AriaLabel: fix.LanguageToggleLabel,
		Items:     items,
	}
}

// Routes implements app.Feature.
func (f *Feature) Routes(mux *http.ServeMux) {
	mux.HandleFunc("GET /{$}", f.getHome)
	mux.HandleFunc("GET /sample", f.getSample)
	f.registerDocsRoutes(mux)
}

func (f *Feature) getHome(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fix := f.fixtureLocale(ctx)
	layout := f.layoutData(ctx, r, fix.Dashboard.Title, "/")
	_ = web.Render(ctx, w, views.SiteShell(layout, views.DashboardPage(views.DashboardData{
		Title:       fix.Dashboard.Title,
		Description: fix.Dashboard.Description,
		Body:        fix.Dashboard.Body,
		DocLinks:    dashboardDocLinks(fix),
	})))
}

func dashboardDocLinks(fix fixtures.Locale) []views.DashboardDocLink {
	out := make([]views.DashboardDocLink, 0, len(fix.DocNav))
	for _, link := range fix.DocNav {
		out = append(out, views.DashboardDocLink{Label: link.Label, Href: link.Path})
	}
	return out
}

func (f *Feature) getSample(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fix := f.fixtureLocale(ctx)
	layout := f.layoutData(ctx, r, fix.SampleStub.Title, "/sample")
	_ = web.Render(ctx, w, views.SiteShell(layout, views.SamplePage(views.SampleData{
		Title:       fix.SampleStub.Title,
		Description: fix.SampleStub.Description,
		Body:        fix.SampleStub.Body,
	})))
}
