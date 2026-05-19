package views

import (
	"bytes"
	"context"
	"strings"
	"testing"

	"github.com/fastygo/framework/pkg/web/view"
	"github.com/fastygo/ui/internal/ui/layout"
)

func TestSiteShell_dashboardRenders(t *testing.T) {
	d := LayoutData{
		Title:    "Home · FastyGo UI",
		Lang:     "en",
		Brand:    "FastyGo UI",
		Active:   "/",
		NavItems: []layout.NavItem{{Label: "Home", Path: "/", Icon: "home"}},
		Assets:   AssetPaths{CSS: "/static/css/app.css", ThemeJS: "/static/js/theme.js", AppJS: "/static/js/ui8kit.js"},
		Theme:    layout.ThemeToggleProps{},
		LanguageToggle: view.LanguageToggleData{
			CurrentLocale: "en",
			CurrentLabel:  "EN",
		},
	}
	body := DashboardPage(DashboardData{Title: "Home", Description: "d", Body: "b"})
	var buf bytes.Buffer
	if err := SiteShell(d, body).Render(context.Background(), &buf); err != nil {
		t.Fatal(err)
	}
	html := buf.String()
	if !strings.Contains(strings.ToLower(html), "<!doctype html>") {
		t.Fatal("expected full document with doctype")
	}
	if !strings.Contains(html, `data-ui8kit="sheet"`) {
		t.Fatal("expected shell mobile sheet markup")
	}
	if !strings.Contains(html, `id="ui8kit-theme-toggle"`) {
		t.Fatal("expected theme toggle control")
	}
	if !strings.Contains(html, `id="language-toggle"`) {
		t.Fatal("expected language toggle control")
	}
}
