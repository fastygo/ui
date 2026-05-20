package views

import (
	"github.com/fastygo/framework/pkg/web/view"
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
	Title          string
	Lang           string
	Brand          string
	Active         string
	NavItems       []layout.NavItem
	Assets         AssetPaths
	Theme          layout.ThemeToggleProps
	LanguageToggle view.LanguageToggleData
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
