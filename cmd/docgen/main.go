// Command docgen builds static documentation HTML from localized Markdown.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fastygo/ui/internal/showcase/docgen"
)

func main() {
	out := flag.String("out", "web/static/docs", "output directory for static docs")
	locales := flag.String("locales", envOr("APP_AVAILABLE_LOCALES", "en,ru"), "comma-separated locale codes")
	defaultLocale := flag.String("default-locale", envOr("APP_DEFAULT_LOCALE", "en"), "default locale (unprefixed /docs/ URLs)")
	strict := flag.Bool("strict-locale", false, "fail when a non-default locale page is missing")
	force := flag.Bool("force", false, "rebuild all pages and previews, ignoring incremental stamps")
	cleanPreviews := flag.Bool("clean-previews", false, "clear persistent preview store before build")
	flag.Parse()

	localeList := splitLocales(*locales)
	pages, err := docgen.LoadAll(docgen.LoadOptions{
		Locales:       localeList,
		DefaultLocale: *defaultLocale,
		StrictLocale:  *strict,
	})
	if err != nil {
		log.Fatalf("load: %v", err)
	}
	if err := docgen.HighlightCodeBlocks(pages); err != nil {
		log.Fatalf("highlight: %v", err)
	}

	previewStats, err := docgen.CompilePreviews(pages, docgen.PreviewCacheConfig{
		Force:      *force,
		CleanStore: *cleanPreviews,
	})
	if err != nil {
		log.Fatalf("preview compile: %v", err)
	}

	buildStats, err := docgen.Build(context.Background(), pages, docgen.BuildConfig{
		OutputDir:     *out,
		Locales:       localeList,
		DefaultLocale: *defaultLocale,
		Incremental:   !*force,
		Force:         *force,
	})
	if err != nil {
		log.Fatalf("build: %v", err)
	}

	fmt.Printf(
		"docgen: %d page(s) across %v (default=%s) -> %s (previews: %d cached, %d compiled, %d unique; pages: %d written, %d skipped; artifacts: %d written)\n",
		len(pages),
		localeList,
		*defaultLocale,
		*out,
		previewStats.Cached,
		previewStats.Compiled,
		previewStats.Unique,
		buildStats.PagesWritten,
		buildStats.PagesSkipped,
		buildStats.ArtifactsWritten,
	)
}

func envOr(key, fallback string) string {
	if v := strings.TrimSpace(os.Getenv(key)); v != "" {
		return v
	}
	return fallback
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
