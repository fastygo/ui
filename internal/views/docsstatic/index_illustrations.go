package docsstatic

import (
	"context"
	"io"
	"strings"

	"github.com/a-h/templ"
	"github.com/fastygo/templ/ui"
)

type indexIllustration struct {
	Class   string
	Surface string
	Shapes  []indexIllustrationShape
}

type indexIllustrationShape struct {
	Kind string
	Tone string
	X    string
	Y    string
	W    string
	H    string
}

var indexIllustrationBySuffix = map[string]indexIllustration{
	"/components/accordion/": {
		Class: "docs-index-illus-accordion",
		Shapes: []indexIllustrationShape{
			pill("ink", "0rem", "0.25rem", "6.5rem", "1rem"),
			pill("ink", "0rem", "2rem", "6.5rem", "1rem"),
		},
	},
	"/components/alert/": {
		Class: "docs-index-illus-alert",
		Shapes: []indexIllustrationShape{
			dot("accent", "0.15rem", "1.15rem", "0.9rem"),
			pill("ink", "1.8rem", "0.8rem", "4.75rem", "0.75rem"),
			pill("soft", "1.8rem", "2.15rem", "6.25rem", "0.55rem"),
		},
	},
	"/components/alert-dialog/": dialogIllustration(),
	"/components/card/": {
		Class: "docs-index-illus-card",
		Shapes: []indexIllustrationShape{
			rect("ink", "1rem", "0rem", "6rem", "1.75rem"),
			pill("soft", "1rem", "2.65rem", "4.8rem", "0.55rem"),
			pill("soft", "1rem", "3.65rem", "3.1rem", "0.45rem"),
		},
	},
	"/components/dialog/": dialogIllustration(),
	"/components/form/": {
		Class: "docs-index-illus-form",
		Shapes: []indexIllustrationShape{
			pill("ink", "0.8rem", "0rem", "6.5rem", "0.75rem"),
			pill("ink", "0.8rem", "1.55rem", "6.5rem", "0.75rem"),
			pill("ink", "0.8rem", "3.1rem", "4rem", "0.75rem"),
			pill("accent", "5.25rem", "3.1rem", "2rem", "0.75rem"),
		},
	},
	"/components/pagination/": {
		Class: "docs-index-illus-pagination",
		Shapes: []indexIllustrationShape{
			rect("soft", "0rem", "1.1rem", "1.25rem", "1.25rem"),
			rect("accent", "1.75rem", "1.1rem", "1.25rem", "1.25rem"),
			rect("soft", "3.5rem", "1.1rem", "1.25rem", "1.25rem"),
			rect("soft", "5.25rem", "1.1rem", "1.25rem", "1.25rem"),
		},
	},
	"/components/table/": {
		Class:   "docs-index-illus-table",
		Surface: "outline",
		Shapes: []indexIllustrationShape{
			pill("border", "0rem", "1.4rem", "7rem", "0.12rem"),
			pill("border", "0rem", "2.75rem", "7rem", "0.12rem"),
			pill("border", "2.25rem", "0rem", "0.12rem", "4.2rem"),
			pill("border", "4.7rem", "0rem", "0.12rem", "4.2rem"),
		},
	},
	"/components/tabs/": {
		Class: "docs-index-illus-tabs",
		Shapes: []indexIllustrationShape{
			pill("accent", "0rem", "0.65rem", "2rem", "1rem"),
			pill("soft", "2.35rem", "0.65rem", "2rem", "1rem"),
			pill("soft", "4.7rem", "0.65rem", "2rem", "1rem"),
			pill("border", "0rem", "2.25rem", "6.7rem", "0.12rem"),
		},
	},
	"/primitives/badge/": {
		Class: "docs-index-illus-badge",
		Shapes: []indexIllustrationShape{
			pill("ink", "0rem", "0.3rem", "5.15rem", "1.5rem"),
			dot("accent", "5.5rem", "0.7rem", "0.7rem"),
		},
	},
	"/primitives/box/": {
		Class: "docs-index-illus-box",
		Shapes: []indexIllustrationShape{
			rect("ink", "1rem", "0rem", "5.75rem", "2.35rem"),
			pill("soft", "2rem", "3rem", "3.75rem", "0.45rem"),
		},
	},
	"/primitives/button/": {
		Class: "docs-index-illus-button",
		Shapes: []indexIllustrationShape{
			pill("accent", "1.15rem", "1.35rem", "5.75rem", "1.35rem"),
		},
	},
	"/primitives/radio/": {
		Class: "docs-index-illus-radio",
		Shapes: []indexIllustrationShape{
			dot("accent", "0.2rem", "1rem", "1.3rem"),
			dot("cutout", "0.59rem", "1.39rem", "0.52rem"),
			pill("ink", "2rem", "1.15rem", "4.2rem", "0.9rem"),
		},
	},
}

func dialogIllustration() indexIllustration {
	return indexIllustration{
		Class: "docs-index-illus-dialog",
		Shapes: []indexIllustrationShape{
			dot("ink", "6.58rem", "0.13rem", "0.64rem"),
			pill("ink", "0rem", "0.25rem", "3.75rem", "0.65rem"),
			rect("soft", "0rem", "1.55rem", "6.9rem", "1.65rem"),
			pill("soft", "0rem", "3.85rem", "4.25rem", "0.45rem"),
			pill("accent", "4.8rem", "3.65rem", "2.25rem", "0.85rem"),
		},
	}
}

func rect(tone, x, y, w, h string) indexIllustrationShape {
	return indexIllustrationShape{Kind: "rect", Tone: tone, X: x, Y: y, W: w, H: h}
}

func pill(tone, x, y, w, h string) indexIllustrationShape {
	return indexIllustrationShape{Kind: "pill", Tone: tone, X: x, Y: y, W: w, H: h}
}

func dot(tone, x, y, size string) indexIllustrationShape {
	return indexIllustrationShape{Kind: "dot", Tone: tone, X: x, Y: y, W: size, H: size}
}

func indexCardIllustration(href string) (indexIllustration, bool) {
	for suffix, illustration := range indexIllustrationBySuffix {
		if strings.HasSuffix(href, suffix) {
			if illustration.Surface == "" {
				illustration.Surface = "filled"
			}
			return illustration, true
		}
	}
	return indexIllustration{}, false
}

func renderIndexCardIllustration(ctx context.Context, w io.Writer, illustration indexIllustration) error {
	return ui.Box(ui.BoxProps{
		Class: "docs-index-card-illustration",
		Attrs: templ.Attributes{"aria-hidden": "true"},
	}).Render(templ.WithChildren(ctx, templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		if err := ui.Box(ui.BoxProps{Class: "docs-index-illus-surface docs-index-illus-surface-" + illustration.Surface}).Render(ctx, w); err != nil {
			return err
		}
		for _, shape := range illustration.Shapes {
			if err := ui.Box(ui.BoxProps{
				Class: "docs-index-illus-shape docs-index-illus-" + shape.Kind + " docs-index-illus-tone-" + shape.Tone,
				Attrs: templ.Attributes{"style": shapeStyle(shape)},
			}).Render(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})), w)
}

func shapeStyle(shape indexIllustrationShape) string {
	return "--x:" + shape.X + ";--y:" + shape.Y + ";--w:" + shape.W + ";--h:" + shape.H + ";"
}
