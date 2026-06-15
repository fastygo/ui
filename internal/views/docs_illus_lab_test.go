package views

import (
	"bytes"
	"context"
	"strings"
	"testing"

	"github.com/fastygo/ui/internal/ui/layout"
	"github.com/fastygo/ui/internal/views/docsstatic"
)

func TestDocsIllusLabPage_rendersCuratedSamples(t *testing.T) {
	d := layoutDataFixture()

	var buf bytes.Buffer
	if err := DocsIllusLabPage(d).Render(context.Background(), &buf); err != nil {
		t.Fatal(err)
	}
	html := buf.String()
	for _, want := range []string{
		"bg-muted-foreground/32",
		"Sample primitives (8)",
		"Sample components (8)",
		"Input",
		"Avatar",
		"Index card preview",
	} {
		if !strings.Contains(html, want) {
			t.Fatalf("html missing %q", want)
		}
	}
	for _, absent := range []string{
		"Production recipes (all index cards)",
	} {
		if strings.Contains(html, absent) {
			t.Fatalf("html should not contain %q", absent)
		}
	}
}

func TestDocsIllusSpritePage_rendersAllSectionEntries(t *testing.T) {
	d := layoutDataFixture()

	t.Run("primitives", func(t *testing.T) {
		var buf bytes.Buffer
		if err := DocsIllusSpritePage(d, docsstatic.IllusSectionPrimitives).Render(context.Background(), &buf); err != nil {
			t.Fatal(err)
		}
		html := buf.String()
		if !strings.Contains(html, "docs-illus-sprite-cell") {
			t.Fatal("sprite page should render cells")
		}
		if !strings.Contains(html, "Input") {
			t.Fatal("sprite page should include Input label")
		}
		if got := strings.Count(html, "docs-illus-sprite-cell"); got != 26 {
			t.Fatalf("expected 26 primitive sprite cells, got %d", got)
		}
	})

	t.Run("components", func(t *testing.T) {
		var buf bytes.Buffer
		if err := DocsIllusSpritePage(d, docsstatic.IllusSectionComponents).Render(context.Background(), &buf); err != nil {
			t.Fatal(err)
		}
		html := buf.String()
		if got := strings.Count(html, "docs-illus-sprite-cell"); got != 26 {
			t.Fatalf("expected 26 component sprite cells, got %d", got)
		}
		if !strings.Contains(html, "bg-background") {
			t.Fatal("sprite page should use light background cells")
		}
	})
}

func TestDocsIllusExportPage_rendersTransparentGrid(t *testing.T) {
	d := layoutDataFixture()

	for _, section := range []string{docsstatic.IllusSectionPrimitives, docsstatic.IllusSectionComponents} {
		t.Run(section, func(t *testing.T) {
			var buf bytes.Buffer
			if err := DocsIllusExportPage(d, section).Render(context.Background(), &buf); err != nil {
				t.Fatal(err)
			}
			html := buf.String()
			for _, want := range []string{
				"docs-illus-export-grid",
				"docs-illus-export-cell",
				"docs-illus-export-page",
				"data-illus-slug=",
				`data-illus-export-section="` + section + `"`,
			} {
				if !strings.Contains(html, want) {
					t.Fatalf("html missing %q", want)
				}
			}
			if got := strings.Count(html, "docs-illus-export-cell"); got != 26 {
				t.Fatalf("expected 26 export cells, got %d", got)
			}
			for _, absent := range []string{
				"docs-illus-sprite-label",
				"border-border",
				"docs-illus-sprite-canvas",
			} {
				if strings.Contains(html, absent) {
					t.Fatalf("export page should not contain %q", absent)
				}
			}
		})
	}
}

func layoutDataFixture() LayoutData {
	return LayoutData{
		Lang:  "en",
		Brand: "FastyGo UI",
		Assets: AssetPaths{
			CSS:     "/static/css/app.css",
			ThemeJS: "/static/js/theme.js",
			AppJS:   "/static/js/ui8kit.js",
		},
		Theme: layout.ThemeToggleProps{},
	}
}
