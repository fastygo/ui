package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

// Copies Templ *.spec.md into FastyGoUI showcase content as *.md (no YAML transform).
func main() {
	templRoot := flag.String("templ", envOr("TEMPL_ROOT", defaultTemplRoot()), "path to github.com/fastygo/templ repo")
	outRoot := flag.String("out", "internal/showcase/content/en", "showcase content locale root")
	dryRun := flag.Bool("dry-run", false, "print actions without writing")
	flag.Parse()

	type job struct {
		pattern  string
		skipDirs map[string]bool
	}
	jobs := []job{
		{pattern: "ui/*/*.spec.md", skipDirs: map[string]bool{"formswitch": true, "selectfield": true}},
		{pattern: "components/*/*.spec.md", skipDirs: nil},
		{pattern: "utils/*.spec.md", skipDirs: nil},
	}

	var copied int
	for _, job := range jobs {
		pattern := filepath.Join(*templRoot, filepath.FromSlash(job.pattern))
		matches, err := filepath.Glob(pattern)
		if err != nil {
			fmt.Fprintf(os.Stderr, "sync-specs: glob %q: %v\n", pattern, err)
			os.Exit(1)
		}
		for _, src := range matches {
			dir := filepath.Base(filepath.Dir(src))
			if job.skipDirs[dir] {
				continue
			}
			layer, err := readSpecLayer(src)
			if err != nil {
				fmt.Fprintf(os.Stderr, "sync-specs: %s: %v\n", src, err)
				os.Exit(1)
			}
			destDir, ok := destSection(layer)
			if !ok {
				fmt.Fprintf(os.Stderr, "sync-specs: %s: unknown layer %q\n", src, layer)
				os.Exit(1)
			}
			slug := strings.TrimSuffix(filepath.Base(src), ".spec.md")
			dest := filepath.Join(*outRoot, destDir, slug+".md")
			if err := copyFile(src, dest, *dryRun); err != nil {
				fmt.Fprintf(os.Stderr, "sync-specs: %v\n", err)
				os.Exit(1)
			}
			copied++
		}
	}
	fmt.Printf("sync-specs: %d file(s) %s -> %s\n", copied, *templRoot, *outRoot)
}

func readSpecLayer(path string) (string, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	text := strings.ReplaceAll(string(raw), "\r\n", "\n")
	if !strings.HasPrefix(text, "---\n") {
		return "", fmt.Errorf("missing front matter")
	}
	end := strings.Index(text[4:], "\n---")
	if end < 0 {
		return "", fmt.Errorf("unclosed front matter")
	}
	fm := map[string]any{}
	if err := yaml.Unmarshal([]byte(text[4:4+end]), &fm); err != nil {
		return "", err
	}
	layer, _ := fm["layer"].(string)
	if strings.TrimSpace(layer) == "" {
		return "", fmt.Errorf("missing layer")
	}
	return layer, nil
}

func destSection(layer string) (string, bool) {
	switch strings.ToLower(strings.TrimSpace(layer)) {
	case "primitive":
		return "primitives", true
	case "helper":
		return "utils", true
	case "composite":
		return "components", true
	default:
		return "", false
	}
}

func copyFile(src, dest string, dryRun bool) error {
	if dryRun {
		fmt.Printf("would copy %s -> %s\n", src, dest)
		return nil
	}
	if err := os.MkdirAll(filepath.Dir(dest), 0o755); err != nil {
		return err
	}
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()
	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()
	if _, err := io.Copy(out, in); err != nil {
		return err
	}
	return out.Close()
}

func envOr(key, fallback string) string {
	if v := strings.TrimSpace(os.Getenv(key)); v != "" {
		return v
	}
	return fallback
}

func defaultTemplRoot() string {
	candidates := []string{
		filepath.Join("..", "@Templ"),
		filepath.Join("..", "..", "@Templ"),
	}
	for _, c := range candidates {
		if st, err := os.Stat(filepath.Join(c, "go.mod")); err == nil && !st.IsDir() {
			if _, err := os.Stat(filepath.Join(c, "ui")); err == nil {
				return c
			}
		}
	}
	return filepath.Join("..", "@Templ")
}
