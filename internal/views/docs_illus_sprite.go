package views

import (
	"github.com/fastygo/ui/internal/views/docsstatic"
)

func docsIllusSpriteTitle(section string) string {
	switch section {
	case docsstatic.IllusSectionPrimitives:
		return "Primitive illustration sprites"
	case docsstatic.IllusSectionComponents:
		return "Component illustration sprites"
	default:
		return "Illustration sprites"
	}
}

func docsIllusSpriteLead(section string) string {
	switch section {
	case docsstatic.IllusSectionPrimitives:
		return "Light-only sprite sheet for all published primitive docs cards. Each cell uses production illustration size on a white background."
	case docsstatic.IllusSectionComponents:
		return "Light-only sprite sheet for all published component docs cards. Each cell uses production illustration size on a white background."
	default:
		return "Light-only sprite sheet for docs index illustrations."
	}
}

func docsIllusSpriteNavLinks() string {
	return "Sprites: /lab/docs-index-illus/primitives · /lab/docs-index-illus/components · Export: /lab/docs-index-illus/export/primitives · /lab/docs-index-illus/export/components · Lab: /lab/docs-index-illus"
}

func docsIllusSpriteEntries(section string) []docsstatic.IndexIllustrationEntry {
	return docsstatic.IndexIllustrationEntriesForSection(section)
}

func docsIllusLabCuratedSamples(section string) []docsstatic.IndexIllustrationEntry {
	suffixes := docsIllusLabCuratedSuffixes(section)
	byHref := make(map[string]docsstatic.IndexIllustrationEntry, len(suffixes))
	for _, entry := range docsstatic.IndexIllustrationEntriesForSection(section) {
		byHref[entry.Href] = entry
	}
	out := make([]docsstatic.IndexIllustrationEntry, 0, len(suffixes))
	for _, suffix := range suffixes {
		href := "/docs" + suffix
		entry, ok := byHref[href]
		if !ok {
			continue
		}
		out = append(out, entry)
	}
	return out
}

func docsIllusLabCuratedSuffixes(section string) []string {
	switch section {
	case docsstatic.IllusSectionPrimitives:
		return []string{
			"/primitives/button/",
			"/primitives/input/",
			"/primitives/checkbox/",
			"/primitives/select/",
			"/primitives/grid/",
			"/primitives/image/",
			"/primitives/switch/",
			"/primitives/textarea/",
		}
	case docsstatic.IllusSectionComponents:
		return []string{
			"/components/accordion/",
			"/components/alert/",
			"/components/dialog/",
			"/components/form/",
			"/components/pagination/",
			"/components/tabs/",
			"/components/table/",
			"/components/avatar/",
		}
	default:
		return nil
	}
}
