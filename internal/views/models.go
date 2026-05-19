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

// DashboardData is the home page body inside the shell.
type DashboardData struct {
	Title       string
	Description string
	Body        string
}

// SampleData is a second stub route for onboarding new pages.
type SampleData struct {
	Title       string
	Description string
	Body        string
}
