package views

import (
	"github.com/fastygo/ui/internal/ui/components/toggles"
	"github.com/fastygo/ui/internal/ui/layout"
)

// AssetPaths are URLs for CSS and JS bundles.
type AssetPaths struct {
	CSS     string
	ThemeJS string
	AppJS   string
}

// LayoutData drives the default app shell (sidebar + header).
type LayoutData struct {
	PageTitle      string // visible page name in the header bar
	Lang           string
	Brand          string // appended to PageTitle in the document <title> only
	Active         string
	NavItems       []layout.NavItem
	Assets         AssetPaths
	Theme          layout.ThemeToggleProps
	LanguageSwitch toggles.LanguageSwitchProps
}

// DocumentTitle returns the SEO document title for <title>.
func (d LayoutData) DocumentTitle() string {
	return FormatDocumentTitle(d.PageTitle, d.Brand)
}

// FormatDocumentTitle builds "Page · Brand" for the document head.
func FormatDocumentTitle(pageTitle, brand string) string {
	if brand == "" {
		brand = "FastyGo UI"
	}
	return pageTitle + " · " + brand
}

// DashboardDocLink is a CTA on the home page pointing at static docs.
type DashboardDocLink struct {
	Label string
	Href  string
}

// DashboardData is the home page body inside the shell.
type DashboardData struct {
	Title       string
	Description string
	Body        string
	DocLinks    []DashboardDocLink
}

// SampleData is a second stub route for onboarding new pages.
type SampleData struct {
	Title       string
	Description string
	Body        string
}
