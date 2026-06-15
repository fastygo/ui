package docsstatic_test

import (
	"strings"
	"testing"

	"github.com/fastygo/ui/internal/showcase/docgen"
	"github.com/fastygo/ui/internal/views/docsstatic"
)

func TestIllustrationRegistry_matchesPublishedDocs(t *testing.T) {
	pages, err := docgen.LoadAll(docgen.LoadOptions{Locales: []string{"en"}})
	if err != nil {
		t.Fatal(err)
	}

	registry := make(map[string]struct{}, len(docsstatic.IndexIllustrationEntries()))
	for _, entry := range docsstatic.IndexIllustrationEntries() {
		registry[entry.Href] = struct{}{}
	}

	var missing []string
	for _, page := range pages {
		if page.Locale != "en" {
			continue
		}
		switch page.Meta.Section {
		case docsstatic.IllusSectionPrimitives, docsstatic.IllusSectionComponents:
		default:
			continue
		}
		if _, ok := registry[page.PublicPath]; !ok {
			missing = append(missing, page.PublicPath)
		}
	}

	if len(missing) > 0 {
		t.Fatalf("registry missing %d published pages: %s", len(missing), strings.Join(missing, ", "))
	}

	if got := len(docsstatic.IndexIllustrationEntries()); got != 52 {
		t.Fatalf("expected 52 registry entries, got %d", got)
	}
	if got := len(docsstatic.IndexIllustrationEntriesForSection(docsstatic.IllusSectionPrimitives)); got != 26 {
		t.Fatalf("expected 26 primitive entries, got %d", got)
	}
	if got := len(docsstatic.IndexIllustrationEntriesForSection(docsstatic.IllusSectionComponents)); got != 26 {
		t.Fatalf("expected 26 component entries, got %d", got)
	}
}
