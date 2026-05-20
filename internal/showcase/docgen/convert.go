package docgen

import "github.com/fastygo/ui/internal/views/docsstatic"

// ToPageData converts a parsed doc page into a templ view model.
func ToPageData(page DocPage) docsstatic.PageData {
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
	return docsstatic.PageData{
		Title:       page.Meta.Title,
		Description: page.Meta.Description,
		Source:      page.Meta.Source,
		Blocks:      blocks,
		API:         api,
		Related:     related,
	}
}

func convertBlock(b Block) docsstatic.Block {
	switch x := b.(type) {
	case ParagraphBlock:
		return docsstatic.ParagraphBlock{Text: x.Text}
	case HeadingBlock:
		return docsstatic.HeadingBlock{Level: x.Level, Text: x.Text}
	case ListBlock:
		return docsstatic.ListBlock{Items: x.Items}
	case DemoBlock:
		return docsstatic.DemoBlock{ID: x.ID, CodeSource: x.Code.Source}
	case CodeBlock:
		return docsstatic.CodeBlock{Source: x.Source}
	default:
		return nil
	}
}

// BuildIndexSections groups pages into index sections for one locale.
func BuildIndexSections(pages []DocPage, locale string) []docsstatic.IndexSection {
	bySection := map[string][]DocPage{}
	for _, p := range pages {
		if p.Locale != locale {
			continue
		}
		bySection[p.Meta.Section] = append(bySection[p.Meta.Section], p)
	}
	order := []string{"getting-started", "components", "blocks"}
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
		out = append(out, docsstatic.IndexSection{Label: docsstatic.SectionLabel(sec), Links: links})
	}
	return out
}
