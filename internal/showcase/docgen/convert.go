package docgen

import (
	"github.com/fastygo/ui/internal/fixtures"
	"github.com/fastygo/ui/internal/views/docsstatic"
)

// ToPageData converts a parsed doc page into a templ view model.
func ToPageData(page DocPage, fix fixtures.Locale) docsstatic.PageData {
	apiHeading := fix.Docs.APIHeading
	if apiHeading == "" {
		apiHeading = "API"
	}
	relatedHeading := fix.Docs.RelatedHeading
	if relatedHeading == "" {
		relatedHeading = "Related"
	}
	var blocks []docsstatic.Block
	for _, b := range page.Blocks {
		blocks = append(blocks, convertBlock(b))
	}
	var api []docsstatic.APIField
	for _, f := range page.Meta.API {
		api = append(api, docsstatic.APIField{
			Name:        f.Name,
			Type:        f.Type,
			Description: f.Description,
		})
	}
	var related []docsstatic.RelatedLink
	for _, r := range page.Meta.Related {
		related = append(related, docsstatic.RelatedLink{
			Label: r.Label,
			Href:  NormalizeHref(r.Href),
		})
	}
	toc := make([]docsstatic.TOCHeading, 0, len(page.Headings)+2)
	for _, h := range page.Headings {
		toc = append(toc, docsstatic.TOCHeading{Level: h.Level, Text: h.Text, ID: h.ID})
	}
	if len(api) > 0 {
		toc = append(toc, docsstatic.TOCHeading{Level: 2, Text: apiHeading, ID: "api"})
	}
	if len(related) > 0 {
		toc = append(toc, docsstatic.TOCHeading{Level: 2, Text: relatedHeading, ID: "related"})
	}
	return docsstatic.PageData{
		Title:               page.Meta.Title,
		Description:         page.Meta.Description,
		Source:              page.Meta.Source,
		Blocks:              blocks,
		API:                 api,
		Related:             related,
		TOC:                 toc,
		APISectionTitle:     apiHeading,
		RelatedSectionTitle: relatedHeading,
	}
}

func convertBlock(b Block) docsstatic.Block {
	switch x := b.(type) {
	case ParagraphBlock:
		return docsstatic.ParagraphBlock{Text: x.Text}
	case HeadingBlock:
		return docsstatic.HeadingBlock{Level: x.Level, Text: x.Text, ID: x.ID}
	case ListBlock:
		return docsstatic.ListBlock{Items: x.Items}
	case PreviewCodeBlock:
		return docsstatic.PreviewCodeBlock{
			ID:              x.ID,
			Source:          x.Source,
			HTML:            x.HTML,
			HighlightedHTML: x.HighlightedHTML,
			SourceFile:      x.SourceFile,
			FenceIndex:      x.FenceIndex,
		}
	case CodeBlock:
		return docsstatic.CodeBlock{
			Language:        x.Language,
			Source:          x.Source,
			HighlightedHTML: x.HighlightedHTML,
		}
	default:
		return nil
	}
}

// BuildIndexSections groups pages into index sections for one locale.
func BuildIndexSections(pages []DocPage, locale string, fix fixtures.Locale) []docsstatic.IndexSection {
	bySection := map[string][]DocPage{}
	for _, p := range pages {
		if p.Locale != locale {
			continue
		}
		bySection[p.Meta.Section] = append(bySection[p.Meta.Section], p)
	}
	order := []string{"getting-started", "primitives", "utils", "components", "blocks"}
	var out []docsstatic.IndexSection
	for _, sec := range order {
		items := bySection[sec]
		if len(items) == 0 {
			continue
		}
		var links []docsstatic.IndexLink
		for _, p := range items {
			links = append(links, docsstatic.IndexLink{
				Title:       p.Meta.Title,
				Description: docsstatic.Truncate(p.Meta.Description, 120),
				Href:        p.PublicPath,
			})
		}
		out = append(out, docsstatic.IndexSection{Label: sectionLabel(fix, sec), Links: links})
	}
	return out
}
