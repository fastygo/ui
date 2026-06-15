package docsstatic

import (
	"context"
	"io"

	"github.com/a-h/templ"
	"github.com/fastygo/templ/ui"
	templutils "github.com/fastygo/templ/utils"
)

const indexCardIllustratedClass = "docs-index-card-illustrated"

// IllustrationLayout controls canvas positioning on index cards vs lab standalone cells.
type IllustrationLayout string

const (
	IllustrationEmbedded   IllustrationLayout = "embedded"
	IllustrationStandalone IllustrationLayout = "standalone"
)

type indexIllustration struct {
	render func(context.Context, io.Writer) error
}

// IndexIllustrationEntry is one illustrated docs index card.
type IndexIllustrationEntry struct {
	Label   string
	Section string
	Href    string
}

func indexCardIllustration(href string) (indexIllustration, bool) {
	for suffix, illustration := range indexIllustrationBySuffix {
		if hasIndexIllustrationHref(href, suffix) {
			return illustration, true
		}
	}
	return indexIllustration{}, false
}

func hasIndexIllustrationHref(href, suffix string) bool {
	return len(href) >= len(suffix) && href[len(href)-len(suffix):] == suffix
}

// IndexIllustrationEntries returns the illustrated index cards in display order.
func IndexIllustrationEntries() []IndexIllustrationEntry {
	out := make([]IndexIllustrationEntry, len(indexIllustrationSpecs))
	for i, spec := range indexIllustrationSpecs {
		out[i] = IndexIllustrationEntry{
			Label:   spec.label,
			Section: spec.section,
			Href:    "/docs" + spec.suffix,
		}
	}
	return out
}

// IndexIllustrationEntriesForSection returns entries for primitives or components.
func IndexIllustrationEntriesForSection(section string) []IndexIllustrationEntry {
	all := IndexIllustrationEntries()
	out := make([]IndexIllustrationEntry, 0, len(all)/2)
	for _, entry := range all {
		if entry.Section == section {
			out = append(out, entry)
		}
	}
	return out
}

// IndexIllustrationSpriteComponent renders the PNG sprite for a docs index card.
func IndexIllustrationSpriteComponent(href string, layout IllustrationLayout) templ.Component {
	if layout == "" {
		layout = IllustrationEmbedded
	}
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		sprite, ok := indexIllustrationSpriteMeta(href)
		if !ok {
			return nil
		}
		return renderIndexCardSprite(ctx, w, sprite, layout)
	})
}

// IndexIllustrationComponent renders the HTML illustration recipe (lab / export).
func IndexIllustrationComponent(href string, layout IllustrationLayout) templ.Component {
	if layout == "" {
		layout = IllustrationEmbedded
	}
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		illustration, ok := indexCardIllustration(href)
		if !ok {
			return nil
		}
		return renderIndexCardIllustration(ctx, w, illustration, layout)
	})
}

// IndexCardIllustrationClass returns the illustrated card modifier class for href.
func IndexCardIllustrationClass(href string) string {
	if _, ok := indexIllustrationSpriteMeta(href); !ok {
		return ""
	}
	return indexCardIllustratedClass
}

func renderIndexCardSprite(ctx context.Context, w io.Writer, sprite indexIllustrationSprite, layout IllustrationLayout) error {
	if layout == "" {
		layout = IllustrationEmbedded
	}
	classes := []string{
		illustrationRootClass(layout),
		"docs-index-card-illus",
		sprite.sheetClass(),
	}
	if layout == IllustrationEmbedded {
		classes = append(classes, "docs-index-card-illus-fade")
	}
	return ui.Box(ui.BoxProps{
		Class: templutils.Cn(classes...),
		Attrs: templ.Attributes{
			"aria-hidden": "true",
			"style":       sprite.backgroundPositionStyle(),
		},
	}).Render(ctx, w)
}

func renderIndexCardIllustration(ctx context.Context, w io.Writer, illustration indexIllustration, layout IllustrationLayout) error {
	if layout == "" {
		layout = IllustrationEmbedded
	}
	return ui.Box(ui.BoxProps{
		Class: illustrationRootClass(layout),
		Attrs: templ.Attributes{"aria-hidden": "true"},
	}).Render(templ.WithChildren(ctx, templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return illustration.render(ctx, w)
	})), w)
}

func illustrationRootClass(layout IllustrationLayout) string {
	base := "pointer-events-none z-0 h-20 w-36 shrink-0"
	if layout == IllustrationEmbedded {
		return templutils.Cn(base, "absolute right-4 top-4")
	}
	return templutils.Cn(base, "relative")
}
