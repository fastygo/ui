package docgen

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/fastygo/ui/internal/fixtures"
	"github.com/fastygo/ui/internal/ui/layout"
	"github.com/fastygo/ui/internal/views"
	"github.com/fastygo/ui/internal/views/docsstatic"
)

// BuildConfig controls static HTML generation.
type BuildConfig struct {
	OutputDir   string
	Locales     []string
	ModuleRoot  string
	Incremental bool // skip unchanged pages and artifacts
	Force       bool // rebuild everything regardless of stamps
}

// BuildStats reports incremental build activity.
type BuildStats struct {
	PagesWritten, PagesSkipped int
	ArtifactsWritten           int
}

// Build writes static HTML and machine-readable artifacts.
func Build(ctx context.Context, pages []DocPage, cfg BuildConfig) (BuildStats, error) {
	var stats BuildStats
	if err := os.MkdirAll(cfg.OutputDir, 0o755); err != nil {
		return stats, err
	}
	if err := ValidateLinks(pages); err != nil {
		return stats, err
	}

	root, err := resolveModuleRoot(cfg.ModuleRoot)
	if err != nil {
		return stats, err
	}
	globalHash, err := globalBuildInput(root)
	if err != nil {
		return stats, fmt.Errorf("build input hash: %w", err)
	}

	incremental := cfg.Incremental && !cfg.Force

	for _, locale := range cfg.Locales {
		locPages := filterLocale(pages, locale)
		fix, err := fixtures.LoadLocale(locale)
		if err != nil {
			fix, _ = fixtures.LoadLocale("en")
		}

		idxPath := indexOutputPath(cfg.OutputDir, locale)
		idxInput := indexBuildInput(locale, locPages, globalHash)
		if incremental && pageUpToDate(root, idxPath, idxInput) {
			stats.PagesSkipped++
		} else {
			if err := writeIndex(ctx, cfg.OutputDir, locale, locPages, fix); err != nil {
				return stats, err
			}
			if err := writePageStamp(root, idxPath, idxInput); err != nil {
				return stats, err
			}
			stats.PagesWritten++
		}

		for _, page := range locPages {
			outFull := filepath.Join(cfg.OutputDir, filepath.FromSlash(page.OutputPath))
			input := pageBuildInput(page, globalHash)
			if incremental && pageUpToDate(root, outFull, input) {
				stats.PagesSkipped++
				continue
			}
			if err := writePage(ctx, cfg.OutputDir, locale, page, locPages, fix); err != nil {
				return stats, err
			}
			if err := writePageStamp(root, outFull, input); err != nil {
				return stats, err
			}
			stats.PagesWritten++
		}
	}

	searchPath := filepath.Join(cfg.OutputDir, "search-index.json")
	searchInput := searchIndexBuildInput(pages, globalHash)
	if !incremental || !pageUpToDate(root, searchPath, searchInput) {
		if err := writeSearchIndex(cfg.OutputDir, pages); err != nil {
			return stats, err
		}
		if err := writePageStamp(root, searchPath, searchInput); err != nil {
			return stats, err
		}
		stats.ArtifactsWritten++
	}

	manifestPath := filepath.Join(cfg.OutputDir, "registry-manifest.json")
	manifestInput := registryManifestBuildInput(pages, globalHash)
	if !incremental || !pageUpToDate(root, manifestPath, manifestInput) {
		if err := writeRegistryManifest(cfg.OutputDir, pages); err != nil {
			return stats, err
		}
		if err := writePageStamp(root, manifestPath, manifestInput); err != nil {
			return stats, err
		}
		stats.ArtifactsWritten++
	}

	sitemapPath := filepath.Join(cfg.OutputDir, "sitemap.xml")
	sitemapInput := sitemapBuildInput(pages, globalHash)
	if !incremental || !pageUpToDate(root, sitemapPath, sitemapInput) {
		if err := writeSitemap(cfg.OutputDir, pages); err != nil {
			return stats, err
		}
		if err := writePageStamp(root, sitemapPath, sitemapInput); err != nil {
			return stats, err
		}
		stats.ArtifactsWritten++
	}

	return stats, nil
}

func searchIndexBuildInput(pages []DocPage, globalHash string) string {
	h := sha256.New()
	h.Write([]byte("search-index"))
	h.Write([]byte(globalHash))
	sorted := append([]DocPage(nil), pages...)
	sort.Slice(sorted, func(i, j int) bool {
		if sorted[i].Locale != sorted[j].Locale {
			return sorted[i].Locale < sorted[j].Locale
		}
		return sorted[i].OutputPath < sorted[j].OutputPath
	})
	for _, p := range sorted {
		h.Write([]byte(pageBuildInput(p, globalHash)))
	}
	return hex.EncodeToString(h.Sum(nil))
}

func registryManifestBuildInput(pages []DocPage, globalHash string) string {
	h := sha256.New()
	h.Write([]byte("registry-manifest"))
	h.Write([]byte(globalHash))
	for _, p := range pages {
		if p.Locale != "en" {
			continue
		}
		h.Write([]byte(pageBuildInput(p, globalHash)))
	}
	return hex.EncodeToString(h.Sum(nil))
}

func sitemapBuildInput(pages []DocPage, globalHash string) string {
	h := sha256.New()
	h.Write([]byte("sitemap"))
	h.Write([]byte(globalHash))
	var paths []string
	seen := map[string]struct{}{}
	for _, p := range pages {
		if _, ok := seen[p.PublicPath]; ok {
			continue
		}
		seen[p.PublicPath] = struct{}{}
		paths = append(paths, p.PublicPath)
	}
	sort.Strings(paths)
	for _, path := range paths {
		h.Write([]byte(path))
	}
	return hex.EncodeToString(h.Sum(nil))
}

func filterLocale(pages []DocPage, locale string) []DocPage {
	var out []DocPage
	for _, p := range pages {
		if p.Locale == locale {
			out = append(out, p)
		}
	}
	return out
}

func writePage(ctx context.Context, outRoot, locale string, page DocPage, all []DocPage, fix fixtures.Locale) error {
	body := docsstatic.Page(ToPageData(page))
	css, themeJS, appJS := docsstatic.StaticAssetPaths()
	layout := views.LayoutData{
		Title:    docsstatic.FormatPageTitle(page.Meta.Title, fix.Brand),
		Lang:     localeLang(locale),
		Brand:    fix.Brand,
		Active:   page.PublicPath,
		NavItems: BuildNavItems(all, locale, page.PublicPath),
		Assets: views.AssetPaths{
			CSS:     css,
			ThemeJS: themeJS,
			AppJS:   appJS,
		},
		Theme: layout.ThemeToggleProps{
			Label:              fix.Theme.Label,
			SwitchToDarkLabel:  fix.Theme.SwitchToDarkLabel,
			SwitchToLightLabel: fix.Theme.SwitchToLight,
		},
	}
	shell := views.SiteShell(layout, body)
	full := filepath.Join(outRoot, filepath.FromSlash(page.OutputPath))
	if err := os.MkdirAll(filepath.Dir(full), 0o755); err != nil {
		return err
	}
	f, err := os.Create(full)
	if err != nil {
		return err
	}
	defer f.Close()
	if err := shell.Render(ctx, f); err != nil {
		return fmt.Errorf("render %s: %w", page.OutputPath, err)
	}
	return nil
}

func writeIndex(ctx context.Context, outRoot, locale string, pages []DocPage, fix fixtures.Locale) error {
	sections := BuildIndexSections(pages, locale)
	body := docsstatic.Index(fix.Docs.IndexTitle, fix.Docs.IndexDescription, sections)
	css, themeJS, appJS := docsstatic.StaticAssetPaths()
	active := docsHomePath(locale)
	layout := views.LayoutData{
		Title:    docsstatic.FormatPageTitle(fix.Docs.IndexTitle, fix.Brand),
		Lang:     localeLang(locale),
		Brand:    fix.Brand,
		Active:   active,
		NavItems: BuildNavItems(pages, locale, active),
		Assets: views.AssetPaths{
			CSS:     css,
			ThemeJS: themeJS,
			AppJS:   appJS,
		},
		Theme: layout.ThemeToggleProps{
			Label:              fix.Theme.Label,
			SwitchToDarkLabel:  fix.Theme.SwitchToDarkLabel,
			SwitchToLightLabel: fix.Theme.SwitchToLight,
		},
	}
	shell := views.SiteShell(layout, body)
	full := indexOutputPath(outRoot, locale)
	if err := os.MkdirAll(filepath.Dir(full), 0o755); err != nil {
		return err
	}
	f, err := os.Create(full)
	if err != nil {
		return err
	}
	defer f.Close()
	return shell.Render(ctx, f)
}

func localeLang(locale string) string {
	if locale == "" {
		return "en"
	}
	return locale
}

func writeSearchIndex(outRoot string, pages []DocPage) error {
	var entries []SearchEntry
	for _, p := range pages {
		var headings []string
		for _, h := range p.Headings {
			headings = append(headings, h.Text)
		}
		entries = append(entries, SearchEntry{
			Locale:      p.Locale,
			Section:     p.Meta.Section,
			Slug:        p.Meta.Slug,
			Title:       p.Meta.Title,
			Description: p.Meta.Description,
			Href:        p.PublicPath,
			Headings:    headings,
		})
	}
	sort.Slice(entries, func(i, j int) bool {
		if entries[i].Locale != entries[j].Locale {
			return entries[i].Locale < entries[j].Locale
		}
		return entries[i].Title < entries[j].Title
	})
	return writeJSON(filepath.Join(outRoot, "search-index.json"), entries)
}

func writeRegistryManifest(outRoot string, pages []DocPage) error {
	seen := map[string]struct{}{}
	var items []RegistryItem
	for _, p := range pages {
		if p.Locale != "en" {
			continue
		}
		key := p.Meta.Section + "/" + p.Meta.Slug
		if _, ok := seen[key]; ok {
			continue
		}
		seen[key] = struct{}{}
		items = append(items, RegistryItem{
			Slug:    p.Meta.Slug,
			Section: p.Meta.Section,
			Title:   p.Meta.Title,
			Source:  p.Meta.Source,
			Package: p.Meta.Package,
		})
	}
	sort.Slice(items, func(i, j int) bool {
		if items[i].Section != items[j].Section {
			return items[i].Section < items[j].Section
		}
		return items[i].Title < items[j].Title
	})
	return writeJSON(filepath.Join(outRoot, "registry-manifest.json"), items)
}

func writeSitemap(outRoot string, pages []DocPage) error {
	var urls []string
	seen := map[string]struct{}{}
	for _, p := range pages {
		if _, ok := seen[p.PublicPath]; ok {
			continue
		}
		seen[p.PublicPath] = struct{}{}
		urls = append(urls, p.PublicPath)
	}
	sort.Strings(urls)
	var b strings.Builder
	b.WriteString(`<?xml version="1.0" encoding="UTF-8"?>` + "\n")
	b.WriteString(`<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">` + "\n")
	for _, u := range urls {
		b.WriteString("  <url><loc>")
		b.WriteString(xmlEscape(u))
		b.WriteString("</loc></url>\n")
	}
	b.WriteString("</urlset>\n")
	return os.WriteFile(filepath.Join(outRoot, "sitemap.xml"), []byte(b.String()), 0o644)
}

func xmlEscape(s string) string {
	s = strings.ReplaceAll(s, "&", "&amp;")
	s = strings.ReplaceAll(s, "<", "&lt;")
	s = strings.ReplaceAll(s, ">", "&gt;")
	return s
}

func writeJSON(path string, v any) error {
	raw, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, raw, 0o644)
}
