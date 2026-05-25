package specvalidate

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"gopkg.in/yaml.v3"
)

// Doc describes a synced showcase spec markdown file.
type Doc struct {
	Path string
	FM   map[string]any
	Body string
}

// Error is one validation failure.
type Error struct {
	File    string
	Field   string
	Message string
}

func (e Error) String() string {
	if e.Field != "" {
		return fmt.Sprintf("%s: %s: %s", e.File, e.Field, e.Message)
	}
	return fmt.Sprintf("%s: %s", e.File, e.Message)
}

var (
	reExampleHeading = regexp.MustCompile(`(?m)^## Example ([^\n]+)\s*$`)
	reFence          = regexp.MustCompile("(?s)```(templ|go)\\s*\\n(.*?)```")
)

// Load reads a spec-format markdown file.
func Load(path string) (*Doc, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	text := strings.ReplaceAll(string(raw), "\r\n", "\n")
	if !strings.HasPrefix(text, "---\n") {
		return nil, fmt.Errorf("missing YAML front matter")
	}
	end := strings.Index(text[4:], "\n---")
	if end < 0 {
		return nil, fmt.Errorf("unclosed YAML front matter")
	}
	fmText := text[4 : 4+end]
	body := strings.TrimLeft(text[4+end+4:], "\n")
	fm := map[string]any{}
	if err := yaml.Unmarshal([]byte(fmText), &fm); err != nil {
		return nil, fmt.Errorf("front matter YAML: %w", err)
	}
	return &Doc{Path: path, FM: fm, Body: body}, nil
}

// Discover finds spec docs under showcase content roots.
func Discover(contentRoot string) ([]string, error) {
	var paths []string
	for _, section := range []string{"primitives", "utils"} {
		root := filepath.Join(contentRoot, section)
		err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if d.IsDir() || !strings.HasSuffix(path, ".md") {
				return nil
			}
			paths = append(paths, path)
			return nil
		})
		if err != nil && !os.IsNotExist(err) {
			return nil, err
		}
	}
	compRoot := filepath.Join(contentRoot, "components")
	compMatches, _ := filepath.Glob(filepath.Join(compRoot, "*.md"))
	for _, path := range compMatches {
		doc, err := Load(path)
		if err != nil {
			continue
		}
		if stringField(doc.FM, "layer") != "" {
			paths = append(paths, path)
		}
	}
	return paths, nil
}

// Validate checks one synced spec document.
func Validate(doc *Doc, contentRoot string) []Error {
	rel, _ := filepath.Rel(contentRoot, doc.Path)
	rel = filepath.ToSlash(rel)
	var errs []Error

	layer := stringField(doc.FM, "layer")
	if layer == "" {
		return []Error{{File: rel, Field: "layer", Message: "missing layer (not a spec document)"}}
	}
	wantSection, ok := sectionForLayer(layer)
	if !ok {
		errs = append(errs, Error{rel, "layer", fmt.Sprintf("unknown layer %q", layer)})
	}
	parts := strings.Split(rel, "/")
	if len(parts) < 2 {
		errs = append(errs, Error{rel, "", "expected {section}/{slug}.md path"})
	} else if ok && parts[0] != wantSection {
		errs = append(errs, Error{rel, "layer", fmt.Sprintf("layer %q belongs in %q/, not %q/", layer, wantSection, parts[0])})
	}
	slug := strings.TrimSuffix(parts[len(parts)-1], ".md")
	if slug == "" {
		errs = append(errs, Error{rel, "slug", "missing slug in filename"})
	}

	required := []string{"id", "layer", "kind", "package", "facade", "semantics"}
	for _, key := range required {
		if _, ok := doc.FM[key]; !ok {
			errs = append(errs, Error{rel, key, "required front matter key missing"})
		}
	}

	switch layer {
	case "helper":
		if _, ok := doc.FM["exports"]; !ok {
			errs = append(errs, Error{rel, "exports", "helper spec requires exports"})
		}
		errs = append(errs, validateHelperExamples(doc, rel)...)
	default:
		for _, key := range []string{"api", "showcase"} {
			if _, ok := doc.FM[key]; !ok {
				errs = append(errs, Error{rel, key, "required for brick specs"})
			}
		}
		errs = append(errs, validateShowcaseExamples(doc, rel)...)
	}
	return errs
}

func sectionForLayer(layer string) (string, bool) {
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

func validateHelperExamples(doc *Doc, rel string) []Error {
	var ids []string
	for _, item := range getSlice(doc.FM, "examples") {
		m, _ := item.(map[string]any)
		if m == nil {
			continue
		}
		if id := stringField(m, "id"); id != "" {
			ids = append(ids, id)
		}
	}
	if len(ids) == 0 {
		return nil
	}
	return validateExampleSections(doc, rel, ids, "go")
}

func validateShowcaseExamples(doc *Doc, rel string) []Error {
	showcase := getSlice(doc.FM, "showcase")
	if showcase == nil {
		return nil
	}
	var ids []string
	for _, item := range showcase {
		m, _ := item.(map[string]any)
		if m == nil {
			continue
		}
		id := stringField(m, "id")
		if id == "" {
			return []Error{{rel, "showcase", "entry missing id"}}
		}
		ids = append(ids, id)
	}
	return validateExampleSections(doc, rel, ids, "templ")
}

func validateExampleSections(doc *Doc, rel string, showcaseIDs []string, fenceLang string) []Error {
	var errs []Error
	showcaseSet := map[string]bool{}
	for _, id := range showcaseIDs {
		showcaseSet[id] = true
	}
	exampleIDs := map[string]bool{}
	for _, m := range reExampleHeading.FindAllStringSubmatch(doc.Body, -1) {
		id := strings.TrimSpace(m[1])
		exampleIDs[id] = true
		if !showcaseSet[id] {
			errs = append(errs, Error{rel, "showcase", fmt.Sprintf("## Example %s has no matching showcase entry", id)})
		}
	}
	for _, id := range showcaseIDs {
		if !exampleIDs[id] {
			errs = append(errs, Error{rel, "showcase", fmt.Sprintf("missing ## Example %s section", id)})
			continue
		}
		section := extractExampleSection(doc.Body, id)
		fences := reFence.FindAllStringSubmatch(section, -1)
		if len(fences) != 1 {
			errs = append(errs, Error{rel, id, fmt.Sprintf("expected exactly one fenced code block, found %d", len(fences))})
			continue
		}
		if fences[0][1] != fenceLang {
			errs = append(errs, Error{rel, id, fmt.Sprintf("expected ```%s fence, found ```%s", fenceLang, fences[0][1])})
		}
	}
	return errs
}

func extractExampleSection(body, id string) string {
	startMarker := "## Example " + id
	start := strings.Index(body, startMarker)
	if start < 0 {
		return ""
	}
	rest := body[start+len(startMarker):]
	next := strings.Index(rest, "\n## ")
	if next >= 0 {
		return rest[:next]
	}
	return rest
}

func stringField(fm map[string]any, key string) string {
	v, _ := fm[key].(string)
	return v
}

func getSlice(fm map[string]any, key string) []any {
	v, ok := fm[key]
	if !ok {
		return nil
	}
	s, _ := v.([]any)
	return s
}
