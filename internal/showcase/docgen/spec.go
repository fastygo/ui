package docgen

import (
	"fmt"
	"path/filepath"
	"sort"
	"strings"

	"gopkg.in/yaml.v3"
)

func isSpecFrontMatter(fm map[string]any) bool {
	layer := strings.TrimSpace(stringFieldFM(fm, "layer"))
	return layer != ""
}

func pageMetaFromSpec(fm map[string]any, body, sourceFile string) (PageMeta, error) {
	slug := slugFromSpecSource(sourceFile)
	if slug == "" {
		return PageMeta{}, fmt.Errorf("cannot derive slug from %q", sourceFile)
	}
	section, err := sectionFromSpecLayer(stringFieldFM(fm, "layer"))
	if err != nil {
		return PageMeta{}, err
	}
	title := strings.TrimSpace(stringFieldFM(fm, "templ"))
	if title == "" {
		title = humanTitle(slug)
	}
	desc := summaryFirstSentence(body)
	if desc == "" {
		desc = title
	}
	source := strings.TrimSpace(stringFieldFM(fm, "facade"))
	if source == "" {
		source = "github.com/fastygo/templ/ui"
	}
	pkg := strings.TrimSpace(stringFieldFM(fm, "package"))
	return PageMeta{
		Slug:        slug,
		Section:     section,
		Title:       title,
		Description: desc,
		Source:      source,
		Package:     pkg,
		API:         flattenSpecAPI(fm),
	}, nil
}

func slugFromSpecSource(sourceFile string) string {
	base := strings.TrimSuffix(filepath.Base(sourceFile), ".md")
	base = strings.TrimSuffix(base, ".spec")
	if base != "" && base != "index" {
		return strings.ToLower(base)
	}
	parts := strings.Split(filepath.ToSlash(sourceFile), "/")
	if len(parts) >= 2 {
		return strings.ToLower(parts[len(parts)-2])
	}
	return ""
}

func sectionFromSpecLayer(layer string) (string, error) {
	switch strings.ToLower(strings.TrimSpace(layer)) {
	case "primitive":
		return "primitives", nil
	case "helper":
		return "utils", nil
	case "composite":
		return "components", nil
	default:
		return "", fmt.Errorf("unknown spec layer %q", layer)
	}
}

func humanTitle(slug string) string {
	parts := strings.Split(strings.ReplaceAll(slug, "-", " "), " ")
	for i, p := range parts {
		if p == "" {
			continue
		}
		parts[i] = strings.ToUpper(p[:1]) + p[1:]
	}
	return strings.Join(parts, " ")
}

func summaryFirstSentence(body string) string {
	body = strings.TrimSpace(body)
	idx := strings.Index(body, "## Summary")
	if idx < 0 {
		return ""
	}
	rest := strings.TrimSpace(body[idx+len("## Summary"):])
	if rest == "" {
		return ""
	}
	if nl := strings.Index(rest, "\n## "); nl >= 0 {
		rest = rest[:nl]
	}
	rest = strings.TrimSpace(rest)
	if rest == "" {
		return ""
	}
	end := strings.IndexAny(rest, ".\n")
	if end < 0 {
		return rest
	}
	return strings.TrimSpace(rest[:end+1])
}

func flattenSpecAPI(fm map[string]any) []APIField {
	apiRaw, ok := fm["api"].(map[string]any)
	if !ok || len(apiRaw) == 0 {
		return flattenSpecExports(fm)
	}
	names := make([]string, 0, len(apiRaw))
	for name := range apiRaw {
		names = append(names, name)
	}
	sortStrings(names)
	var out []APIField
	for _, name := range names {
		field, _ := apiRaw[name].(map[string]any)
		if field == nil {
			continue
		}
		out = append(out, APIField{
			Name:        name,
			Type:        stringFieldFM(field, "type"),
			Description: describeSpecAPIField(field),
		})
	}
	return out
}

func flattenSpecExports(fm map[string]any) []APIField {
	exports, ok := fm["exports"].(map[string]any)
	if !ok {
		return nil
	}
	var out []APIField
	for group, raw := range exports {
		items, _ := raw.([]any)
		for _, item := range items {
			m, _ := item.(map[string]any)
			if m == nil {
				continue
			}
			name := stringFieldFM(m, "name")
			if name == "" {
				continue
			}
			desc := stringFieldFM(m, "role")
			if sig := stringFieldFM(m, "signature"); sig != "" {
				if desc != "" {
					desc += "; "
				}
				desc += sig
			}
			if typ := stringFieldFM(m, "type"); typ != "" {
				if desc != "" {
					desc += "; "
				}
				desc += "type: " + typ
			}
			if desc == "" {
				desc = group
			}
			out = append(out, APIField{Name: name, Type: group, Description: desc})
		}
	}
	return out
}

func describeSpecAPIField(field map[string]any) string {
	var parts []string
	if role := stringFieldFM(field, "role"); role != "" {
		parts = append(parts, "role: "+role)
	}
	if enum := stringSliceFM(field["enum"]); len(enum) > 0 {
		parts = append(parts, "enum: "+strings.Join(enum, "|"))
	}
	if def := stringFieldFM(field, "default"); def != "" {
		parts = append(parts, "default: "+def)
	}
	if src := stringFieldFM(field, "allow-list-source"); src != "" {
		parts = append(parts, "source: "+src)
	}
	return strings.Join(parts, "; ")
}

func stringFieldFM(fm map[string]any, key string) string {
	v, _ := fm[key].(string)
	return v
}

func stringSliceFM(v any) []string {
	switch t := v.(type) {
	case []any:
		var out []string
		for _, item := range t {
			if s, ok := item.(string); ok {
				out = append(out, s)
			}
		}
		return out
	case []string:
		return t
	default:
		return nil
	}
}

func sortStrings(values []string) {
	sort.Slice(values, func(i, j int) bool { return values[i] < values[j] })
}

func splitFrontMatterRaw(raw []byte) (map[string]any, PageMeta, string, bool, error) {
	s := string(raw)
	if !strings.HasPrefix(s, "---\n") && !strings.HasPrefix(s, "---\r\n") {
		return nil, PageMeta{}, "", false, fmt.Errorf("missing YAML front matter")
	}
	s = strings.TrimPrefix(s, "---")
	s = strings.TrimPrefix(s, "\r\n")
	s = strings.TrimPrefix(s, "\n")
	end := strings.Index(s, "\n---")
	if end < 0 {
		return nil, PageMeta{}, "", false, fmt.Errorf("unterminated front matter")
	}
	fmText := s[:end]
	body := strings.TrimPrefix(s[end+4:], "\n")
	body = strings.TrimPrefix(body, "\r\n")

	var fm map[string]any
	if err := yaml.Unmarshal([]byte(fmText), &fm); err != nil {
		return nil, PageMeta{}, "", false, fmt.Errorf("front matter: %w", err)
	}
	if isSpecFrontMatter(fm) {
		return fm, PageMeta{}, body, true, nil
	}
	var meta PageMeta
	if err := yaml.Unmarshal([]byte(fmText), &meta); err != nil {
		return nil, PageMeta{}, "", false, fmt.Errorf("front matter: %w", err)
	}
	return nil, meta, body, false, nil
}
