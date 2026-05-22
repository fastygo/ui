// Command docgen builds static documentation HTML from localized Markdown.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/fastygo/ui/internal/showcase/docgen"
)

func main() {
	out := flag.String("out", "web/static/docs", "output directory for static docs")
	locales := flag.String("locales", "en,ru", "comma-separated locale codes")
	strict := flag.Bool("strict-locale", false, "fail when a non-default locale page is missing")
	force := flag.Bool("force", false, "rebuild all pages and previews, ignoring incremental stamps")
	cleanPreviews := flag.Bool("clean-previews", false, "clear persistent preview store before build")
	flag.Parse()

	localeList := splitLocales(*locales)
	pages, err := docgen.LoadAll(docgen.LoadOptions{
		Locales:       localeList,
		DefaultLocale: "en",
		StrictLocale:  *strict,
	})
	if err != nil {
		log.Fatalf("load: %v", err)
	}

	previewStats, err := docgen.CompilePreviews(pages, docgen.PreviewCacheConfig{
		Force:      *force,
		CleanStore: *cleanPreviews,
	})
	if err != nil {
		log.Fatalf("preview compile: %v", err)
	}

	buildStats, err := docgen.Build(context.Background(), pages, docgen.BuildConfig{
		OutputDir:   *out,
		Locales:     localeList,
		Incremental: !*force,
		Force:       *force,
	})
	if err != nil {
		log.Fatalf("build: %v", err)
	}

	fmt.Printf(
		"docgen: %d page(s) across %v -> %s (previews: %d cached, %d compiled; pages: %d written, %d skipped; artifacts: %d written)\n",
		len(pages),
		localeList,
		*out,
		previewStats.Cached,
		previewStats.Compiled,
		buildStats.PagesWritten,
		buildStats.PagesSkipped,
		buildStats.ArtifactsWritten,
	)
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
