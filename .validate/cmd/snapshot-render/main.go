// Command snapshot-render writes full HTML5 documents under .validate/html-snapshots/generated/
// for Nu HTML validation. Use -route for one or more logical screens (comma-separated).
//
// From repository root:
//
//	go run ./.validate/cmd/snapshot-render -route=home
//	go run ./.validate/cmd/snapshot-render -route=home,sample -lang=en
//	go run ./.validate/cmd/snapshot-render -list
package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"

	"github.com/a-h/templ"
	"github.com/fastygo/framework/pkg/web/view"
	"github.com/fastygo/ui/internal/fixtures"
	"github.com/fastygo/ui/internal/ui/layout"
	"github.com/fastygo/ui/internal/views"
)

const defaultOutDir = ".validate/html-snapshots/generated"

var routeKeys = map[string]string{
	"home":   "GET / — default shell + home body (fixture copy)",
	"sample": "GET /sample — default shell + sample stub",
}

func main() {
	list := flag.Bool("list", false, "print supported route keys and exit")
	routeFlag := flag.String("route", "", "comma-separated route keys (see -list)")
	lang := flag.String("lang", "en", "fixture locale code (internal/fixtures/locale/{lang}.json)")
	outDir := flag.String("out", defaultOutDir, "output directory for .html files")
	flag.Parse()

	if *list {
		fmt.Println("Supported -route keys:")
		for k, desc := range routeKeys {
			fmt.Printf("  %-20s %s\n", k, desc)
		}
		fmt.Println()
		fmt.Println("-list does not write any files. To generate HTML snapshots, run for example:")
		fmt.Printf("  go run ./.validate/cmd/snapshot-render -route=home\n")
		fmt.Printf("  go run ./.validate/cmd/snapshot-render -route=home,sample -lang=en\n")
		fmt.Println()
		fmt.Printf("Default output directory: %s\n", defaultOutDir)
		fmt.Println("Nu CI (bun run validate:html) only reads .validate/html-snapshots/nu/ — copy conforming files there when ready.")
		return
	}
	if strings.TrimSpace(*routeFlag) == "" {
		fmt.Fprintln(os.Stderr, "snapshot-render: -route is required (comma-separated), or use -list")
		os.Exit(2)
	}

	routes := splitRoutes(*routeFlag)
	if err := os.MkdirAll(*outDir, 0o755); err != nil {
		fmt.Fprintln(os.Stderr, "snapshot-render:", err)
		os.Exit(1)
	}

	fix, err := fixtures.LoadLocale(strings.ToLower(strings.TrimSpace(*lang)))
	if err != nil {
		fmt.Fprintln(os.Stderr, "snapshot-render: fixtures:", err)
		os.Exit(1)
	}
	langCode := strings.ToLower(strings.TrimSpace(*lang))

	for _, key := range routes {
		if _, ok := routeKeys[key]; !ok {
			fmt.Fprintf(os.Stderr, "snapshot-render: unknown route %q (use -list)\n", key)
			os.Exit(2)
		}
		var comp templ.Component
		switch key {
		case "home":
			comp, err = renderHome(fix, langCode)
		case "sample":
			comp, err = renderSample(fix, langCode)
		default:
			err = fmt.Errorf("unhandled route %q", key)
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "snapshot-render %s: %v\n", key, err)
			os.Exit(1)
		}
		outPath := filepath.Join(*outDir, fmt.Sprintf("%s-%s.html", key, langCode))
		if err := writeComponent(outPath, comp); err != nil {
			fmt.Fprintf(os.Stderr, "snapshot-render %s: %v\n", key, err)
			os.Exit(1)
		}
		fmt.Println("wrote", outPath)
	}
}

func splitRoutes(s string) []string {
	var out []string
	for _, p := range strings.Split(s, ",") {
		p = strings.TrimSpace(p)
		if p != "" {
			out = append(out, p)
		}
	}
	return out
}

func languageToggle(fix fixtures.Locale, lang, path string) view.LanguageToggleData {
	req := httptest.NewRequest("GET", "http://127.0.0.1"+path+"?lang="+lang, nil)
	return view.BuildLanguageToggle(view.LanguageToggleConfig{
		CurrentLocale: lang,
		DefaultLocale: "en",
		Available:     []string{"en", "ru"},
		Label:         fix.LanguageToggleLabel,
		LocaleLabels:  map[string]string{"en": "EN", "ru": "RU"},
		Request:       req,
		EnhanceWithJS: true,
		SPATarget:     "main",
	})
}

func assetPaths() views.AssetPaths {
	return views.AssetPaths{
		CSS:     "/static/css/app.css",
		ThemeJS: "/static/js/theme.js",
		AppJS:   "/static/js/ui8kit.js",
	}
}

func siteNav(fix fixtures.Locale) []layout.NavItem {
	return []layout.NavItem{
		{Label: fix.Nav.Home, Path: "/", Icon: "home"},
		{Label: fix.Nav.Sample, Path: "/sample", Icon: "layout-dashboard"},
	}
}

func layoutData(fix fixtures.Locale, lang, title, active, langTogglePath string) views.LayoutData {
	return views.LayoutData{
		Title:      title + " · " + fix.Brand,
		Lang:       lang,
		Brand:      fix.Brand,
		Active:     active,
		NavItems:   siteNav(fix),
		Assets:     assetPaths(),
		Theme:      themeProps(fix),
		LanguageToggle: languageToggle(fix, lang, langTogglePath),
	}
}

func themeProps(fix fixtures.Locale) layout.ThemeToggleProps {
	return layout.ThemeToggleProps{
		Label:              fix.Theme.Label,
		SwitchToDarkLabel:  fix.Theme.SwitchToDarkLabel,
		SwitchToLightLabel: fix.Theme.SwitchToLight,
	}
}

func renderHome(fix fixtures.Locale, lang string) (templ.Component, error) {
	layout := layoutData(fix, lang, fix.Dashboard.Title, "/", "/")
	body := views.DashboardPage(views.DashboardData{
		Title:       fix.Dashboard.Title,
		Description: fix.Dashboard.Description,
		Body:        fix.Dashboard.Body,
	})
	return views.SiteShell(layout, body), nil
}

func renderSample(fix fixtures.Locale, lang string) (templ.Component, error) {
	layout := layoutData(fix, lang, fix.SampleStub.Title, "/sample", "/sample")
	body := views.SamplePage(views.SampleData{
		Title:       fix.SampleStub.Title,
		Description: fix.SampleStub.Description,
		Body:        fix.SampleStub.Body,
	})
	return views.SiteShell(layout, body), nil
}

func writeComponent(path string, c templ.Component) error {
	var buf bytes.Buffer
	if err := c.Render(context.Background(), &buf); err != nil {
		return err
	}
	return os.WriteFile(path, buf.Bytes(), 0o644)
}
