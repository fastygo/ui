// Command migrate-docs generates Markdown sources from the runtime docs registry (one-time migration aid).
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/fastygo/ui/internal/registry"
	"github.com/fastygo/ui/internal/showcase/docgen"

	_ "github.com/fastygo/ui/internal/showcase"
)

func main() {
	out := flag.String("out", "internal/showcase/content/en", "output root for generated English markdown")
	flag.Parse()

	if err := os.MkdirAll(*out, 0o755); err != nil {
		log.Fatal(err)
	}

	for _, page := range registry.AllPages() {
		if err := writePage(*out, page); err != nil {
			log.Fatalf("%s: %v", page.Slug, err)
		}
	}
	fmt.Printf("migrate-docs: wrote %d pages to %s\n", len(registry.AllPages()), *out)
}

func writePage(outRoot string, page registry.Page) error {
	relDir, err := sectionDir(page.Section)
	if err != nil {
		return err
	}
	dir := filepath.Join(outRoot, relDir)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	path := filepath.Join(dir, page.Slug+".md")
	body := renderMarkdown(page)
	return os.WriteFile(path, []byte(body), 0o644)
}

func sectionDir(section string) (string, error) {
	switch strings.ToLower(section) {
	case "getting-started":
		return ".", nil
	case "components":
		return "components", nil
	case "blocks":
		return "blocks", nil
	default:
		return "", fmt.Errorf("unknown section %q", section)
	}
}

func renderMarkdown(page registry.Page) string {
	var b strings.Builder
	b.WriteString("---\n")
	b.WriteString(fmt.Sprintf("slug: %s\n", page.Slug))
	b.WriteString(fmt.Sprintf("section: %s\n", page.Section))
	b.WriteString(fmt.Sprintf("title: %s\n", yamlQuote(page.Title)))
	b.WriteString(fmt.Sprintf("description: %s\n", yamlQuote(page.Description)))
	if page.Source != "" {
		b.WriteString(fmt.Sprintf("source: %s\n", page.Source))
	}
	if page.Package != "" {
		b.WriteString(fmt.Sprintf("package: %s\n", page.Package))
	}
	if len(page.Related) > 0 {
		b.WriteString("related:\n")
		for _, r := range page.Related {
			href := docgen.NormalizeHref(r.Href)
			b.WriteString(fmt.Sprintf("  - label: %s\n    href: %s\n", yamlQuote(r.Label), href))
		}
	}
	if len(page.API) > 0 {
		b.WriteString("api:\n")
		for _, f := range page.API {
			b.WriteString(fmt.Sprintf("  - name: %s\n    type: %s\n    description: %s\n", yamlQuote(f.Name), yamlQuote(f.Type), yamlQuote(f.Description)))
		}
	}
	b.WriteString("---\n\n")
	b.WriteString(page.Description)
	b.WriteString("\n\n")
	for _, v := range page.Variants {
		b.WriteString(fmt.Sprintf("## %s\n\n", v.Title))
		if strings.TrimSpace(v.Description) != "" {
			b.WriteString(v.Description)
			b.WriteString("\n\n")
		}
		demoID := page.Slug + "." + v.ID
		b.WriteString(fmt.Sprintf("{{demo id=\"%s\"}}\n\n", demoID))
		lang := "templ"
		code := strings.TrimSpace(v.Code)
		if code != "" && !strings.Contains(code, "@") && !strings.Contains(code, "Render") {
			lang = "go"
		}
		b.WriteString("```" + lang + "\n")
		b.WriteString(code)
		b.WriteString("\n```\n\n")
	}
	return b.String()
}

func yamlQuote(s string) string {
	return fmt.Sprintf("%q", s)
}
