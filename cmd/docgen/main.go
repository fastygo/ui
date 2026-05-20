// Command docgen builds static documentation HTML from localized Markdown and preview registry.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/fastygo/ui/internal/showcase/docgen"
	"github.com/fastygo/ui/internal/showcase/previews"

	_ "github.com/fastygo/ui/internal/showcase"
)

func main() {
	out := flag.String("out", "web/static/docs", "output directory for static docs")
	locales := flag.String("locales", "en,ru", "comma-separated locale codes")
	strict := flag.Bool("strict-locale", false, "fail when a non-default locale page is missing")
	flag.Parse()

	if err := previews.RegisterFromRegistry(); err != nil {
		log.Fatalf("previews: %v", err)
	}

	localeList := splitLocales(*locales)
	pages, err := docgen.LoadAll(docgen.LoadOptions{
		Locales:       localeList,
		DefaultLocale: "en",
		StrictLocale:  *strict,
	})
	if err != nil {
		log.Fatalf("load: %v", err)
	}
	if err := docgen.ResolveDemos(pages); err != nil {
		log.Fatalf("resolve: %v", err)
	}
	if err := docgen.Build(context.Background(), pages, docgen.BuildConfig{
		OutputDir: *out,
		Locales:   localeList,
	}); err != nil {
		log.Fatalf("build: %v", err)
	}
	fmt.Printf("docgen: wrote %d page(s) across %v -> %s\n", len(pages), localeList, *out)
}

func splitLocales(s string) []string {
	var out []string
	for _, part := range strings.Split(s, ",") {
		part = strings.TrimSpace(part)
		if part != "" {
			out = append(out, part)
		}
	}
	if len(out) == 0 {
		return []string{"en"}
	}
	return out
}
