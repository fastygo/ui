package docgen

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	demoLineRe     = regexp.MustCompile(`^\{\{demo\s+id="([^"]+)"\s*\}\}\s*$`)
	rawHTMLRe      = regexp.MustCompile(`(?i)<\s*/?\s*[a-z]`)
	tailwindHintRe = regexp.MustCompile(`(?:bg-[a-z]|text-[a-z]|gap-\d|p-\d|m-\d|rounded-[a-z]|border-border|md:[a-z]|lg:[a-z]|w-full|h-\d|grid-cols-|grid-rows-|flex-row|flex-col|items-center|justify-)`)
	classAttrRe    = regexp.MustCompile(`(?i)\bclass\s*=`)
)

// ValidatePage runs structural checks on a parsed page.
func ValidatePage(page DocPage) error {
	if err := ValidateMeta(page.Meta); err != nil {
		return fmt.Errorf("%s: %w", page.SourceFile, err)
	}
	h1 := 0
	for _, b := range page.Blocks {
		if h, ok := b.(HeadingBlock); ok && h.Level == 1 {
			h1++
		}
	}
	if h1 > 1 {
		return fmt.Errorf("%s: expected at most one h1, found %d", page.SourceFile, h1)
	}
	seenDemos := map[string]struct{}{}
	for _, id := range page.DemoIDs {
		if _, dup := seenDemos[id]; dup {
			return fmt.Errorf("%s: duplicate demo id %q", page.SourceFile, id)
		}
		seenDemos[id] = struct{}{}
	}
	for _, link := range page.Meta.Related {
		if !link.External && strings.HasPrefix(link.Href, "/docs") {
			page.Meta.Related[len(page.Meta.Related)-1].Href = NormalizeHref(link.Href)
		}
	}
	return nil
}

// ValidateBodyRules rejects raw HTML and Tailwind-like utility strings in markdown prose (not code fences).
func ValidateBodyRules(sourceFile, body string) error {
	inFence := false
	for _, line := range strings.Split(body, "\n") {
		trim := strings.TrimSpace(line)
		if strings.HasPrefix(trim, "```") {
			inFence = !inFence
			continue
		}
		if inFence || trim == "" || strings.HasPrefix(trim, "---") || demoLineRe.MatchString(trim) {
			continue
		}
		if rawHTMLRe.MatchString(trim) {
			return fmt.Errorf("%s: raw HTML is not allowed in markdown", sourceFile)
		}
		if classAttrRe.MatchString(trim) {
			return fmt.Errorf("%s: class attributes are not allowed in markdown", sourceFile)
		}
		if tailwindHintRe.MatchString(trim) {
			return fmt.Errorf("%s: line looks like Tailwind utilities (forbidden in markdown): %q", sourceFile, trim)
		}
	}
	return nil
}

// ValidateLinks checks internal related/demo hrefs against the set of generated public paths.
func ValidateLinks(pages []DocPage) error {
	paths := make(map[string]struct{}, len(pages)*2)
	for _, p := range pages {
		paths[p.PublicPath] = struct{}{}
		paths[strings.TrimSuffix(p.PublicPath, "/")] = struct{}{}
	}
	for _, p := range pages {
		for _, link := range p.Meta.Related {
			if link.External {
				continue
			}
			href := NormalizeHref(link.Href)
			if strings.HasPrefix(href, "/docs") {
				if _, ok := paths[href]; !ok {
					if _, ok2 := paths[strings.TrimSuffix(href, "/")]; !ok2 {
						return fmt.Errorf("%s: broken internal link %q", p.SourceFile, href)
					}
				}
			}
		}
	}
	return nil
}

var allowedFenceLangs = map[string]struct{}{
	"templ": {}, "go": {}, "bash": {}, "json": {}, "markdown": {}, "md": {}, "": {},
}

func validateFenceLang(sourceFile, lang string) error {
	lang = strings.ToLower(strings.TrimSpace(lang))
	if _, ok := allowedFenceLangs[lang]; !ok {
		return fmt.Errorf("%s: unsupported code fence language %q", sourceFile, lang)
	}
	return nil
}
