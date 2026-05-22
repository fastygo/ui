package docgen

import (
	"fmt"
	"io/fs"
	"path"
	"sort"
	"strings"

	showcasecontent "github.com/fastygo/ui/internal/showcase/content"
)

// LoadOptions configures documentation loading.
type LoadOptions struct {
	Locales        []string
	DefaultLocale  string
	StrictLocale   bool // fail if a page is missing for a non-default locale
}

// LoadAll discovers and parses all markdown pages for the given locales.
func LoadAll(opts LoadOptions) ([]DocPage, error) {
	if opts.DefaultLocale == "" {
		opts.DefaultLocale = "en"
	}
	if len(opts.Locales) == 0 {
		opts.Locales = []string{opts.DefaultLocale}
	}

	byKey := map[string]map[string]DocPage{} // locale -> content key -> page
	for _, locale := range opts.Locales {
		pages, err := loadLocale(locale)
		if err != nil {
			return nil, err
		}
		byKey[locale] = map[string]DocPage{}
		for _, p := range pages {
			key := pageKey(p.Meta)
			byKey[locale][key] = p
		}
	}

	var out []DocPage
	enPages := byKey[opts.DefaultLocale]
	if enPages == nil {
		return nil, fmt.Errorf("default locale %q has no pages", opts.DefaultLocale)
	}

	keys := sortedKeys(enPages)
	for _, locale := range opts.Locales {
		for _, key := range keys {
			p, ok := byKey[locale][key]
			if ok {
				out = append(out, p)
				continue
			}
			if locale == opts.DefaultLocale {
				continue
			}
			base, ok := enPages[key]
			if !ok {
				continue
			}
			if opts.StrictLocale {
				return nil, fmt.Errorf("locale %q missing page %s", locale, key)
			}
			fallback := base
			fallback.Locale = locale
			fallback.OutputPath = OutputRelPath(locale, fallback.Meta)
			fallback.PublicPath = PublicPath(locale, fallback.Meta)
			fallback.FallbackEN = true
			fallback.SourceFile = fallback.SourceFile + " (fallback:en)"
			fallback.ContentHash = base.ContentHash
			out = append(out, fallback)
		}
	}
	return out, nil
}

func loadLocale(locale string) ([]DocPage, error) {
	root := locale
	var pages []DocPage
	err := fs.WalkDir(showcasecontent.FS, root, func(filePath string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || !strings.HasSuffix(filePath, ".md") {
			return nil
		}
		raw, err := showcasecontent.FS.ReadFile(filePath)
		if err != nil {
			return err
		}
		page, err := ParseFile(locale, filePath, raw)
		if err != nil {
			return err
		}
		pages = append(pages, page)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return pages, nil
}

func pageKey(meta PageMeta) string {
	return path.Join(strings.ToLower(meta.Section), strings.ToLower(meta.Slug))
}

func sortedKeys(m map[string]DocPage) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
