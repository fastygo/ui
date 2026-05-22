package docgen

import (
	"io/fs"
	"strings"
	"testing"

	showcasecontent "github.com/fastygo/ui/internal/showcase/content"
)

func TestLoadAll_enHasNoLegacyDemoDirectives(t *testing.T) {
	pages, err := LoadAll(LoadOptions{Locales: []string{"en"}})
	if err != nil {
		t.Fatal(err)
	}
	for _, page := range pages {
		if len(page.DemoIDs) != 0 {
			t.Fatalf("%s: expected no legacy demo ids, got %v", page.SourceFile, page.DemoIDs)
		}
	}
	err = fs.WalkDir(showcasecontent.FS, "en", func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() || !strings.HasSuffix(path, ".md") {
			return err
		}
		raw, err := showcasecontent.FS.ReadFile(path)
		if err != nil {
			return err
		}
		if strings.Contains(string(raw), "{{demo") {
			t.Errorf("%s: still contains legacy {{demo}} directive", path)
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}
