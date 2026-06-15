package docsstatic

import (
	"context"
	"io"

	"github.com/a-h/templ"
	"github.com/fastygo/templ/ui"
	templutils "github.com/fastygo/templ/utils"
)

const (
	illusBgSoft   = "bg-muted-foreground/16"
	illusBgBorder = "bg-muted-foreground/24"
	illusBgInk    = "bg-muted-foreground/32"
	illusBgAccent = "bg-muted-foreground/48"
)

func renderStackBars(ctx context.Context, w io.Writer, count, accentIndex int) error {
	return ui.Stack(ui.StackProps{Class: "flex h-full flex-col justify-center gap-1 p-2"}).Render(
		templ.WithChildren(ctx, templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
			for i := 0; i < count; i++ {
				tone := illusBgInk
				if i == accentIndex {
					tone = illusBgAccent
				}
				if err := ui.Box(ui.BoxProps{
					Class: templutils.Cn("h-2 w-full rounded-full", tone),
				}).Render(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})), w)
}

func renderSingleAccent(ctx context.Context, w io.Writer) error {
	return ui.Box(ui.BoxProps{Class: "flex h-full items-center justify-center p-2"}).Render(
		templ.WithChildren(ctx, templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
			return ui.Box(ui.BoxProps{
				Class: templutils.Cn("h-5 w-16 rounded-full", illusBgAccent),
			}).Render(ctx, w)
		})), w)
}

func renderBlockCaption(ctx context.Context, w io.Writer) error {
	return ui.Stack(ui.StackProps{Class: "flex h-full flex-col justify-center gap-1 p-2"}).Render(
		templ.WithChildren(ctx, templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
			if err := ui.Box(ui.BoxProps{
				Class: templutils.Cn("h-8 w-full rounded-md", illusBgInk),
			}).Render(ctx, w); err != nil {
				return err
			}
			if err := ui.Box(ui.BoxProps{
				Class: templutils.Cn("h-2 w-32 rounded-full", illusBgSoft),
			}).Render(ctx, w); err != nil {
				return err
			}
			return ui.Box(ui.BoxProps{
				Class: templutils.Cn("h-2 w-24 rounded-full", illusBgSoft),
			}).Render(ctx, w)
		})), w)
}

func renderTabButtons(ctx context.Context, w io.Writer) error {
	return ui.Box(ui.BoxProps{Class: "flex h-full items-center justify-center gap-1 p-2"}).Render(
		templ.WithChildren(ctx, templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
			tones := []string{illusBgSoft, illusBgAccent, illusBgSoft}
			for _, tone := range tones {
				if err := ui.Box(ui.BoxProps{
					Class: templutils.Cn("h-6 flex-1 rounded-md", tone),
				}).Render(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})), w)
}

func renderPaginationButtons(ctx context.Context, w io.Writer) error {
	return ui.Box(ui.BoxProps{Class: "flex h-full items-center justify-center gap-1 p-2"}).Render(
		templ.WithChildren(ctx, templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
			for i := 0; i < 4; i++ {
				tone := illusBgSoft
				if i == 1 {
					tone = illusBgAccent
				}
				if err := ui.Box(ui.BoxProps{
					Class: templutils.Cn("h-8 w-8 shrink-0 rounded-sm", tone),
				}).Render(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})), w)
}

func renderMarkLine(ctx context.Context, w io.Writer) error {
	return ui.Box(ui.BoxProps{Class: "flex h-full items-center gap-2 p-2"}).Render(
		templ.WithChildren(ctx, templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
			if err := ui.Box(ui.BoxProps{
				Class: templutils.Cn("h-4 w-4 shrink-0 rounded-full", illusBgAccent),
			}).Render(ctx, w); err != nil {
				return err
			}
			return ui.Stack(ui.StackProps{Class: "flex min-w-0 flex-1 flex-col justify-center gap-1"}).Render(
				templ.WithChildren(ctx, templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
					if err := ui.Box(ui.BoxProps{
						Class: templutils.Cn("h-2 w-full rounded-full", illusBgInk),
					}).Render(ctx, w); err != nil {
						return err
					}
					return ui.Box(ui.BoxProps{
						Class: templutils.Cn("h-2 w-32 rounded-full", illusBgSoft),
					}).Render(ctx, w)
				})), w)
		})), w)
}

func renderCellGrid(ctx context.Context, w io.Writer, cols, rows int) error {
	gridClass := "grid gap-1 p-2"
	switch cols {
	case 3:
		gridClass += " grid-cols-3"
	case 4:
		gridClass += " grid-cols-4"
	}
	return ui.Grid(ui.GridProps{Class: gridClass}).Render(
		templ.WithChildren(ctx, templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
			for i := 0; i < cols*rows; i++ {
				tone := illusBgSoft
				if i == cols*rows/2 {
					tone = illusBgInk
				}
				if err := ui.Box(ui.BoxProps{
					Class: templutils.Cn("h-3 rounded-sm", tone),
				}).Render(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})), w)
}

func renderPanel(ctx context.Context, w io.Writer) error {
	return ui.Stack(ui.StackProps{Class: "flex h-full flex-col justify-center gap-1 p-2"}).Render(
		templ.WithChildren(ctx, templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
			if err := ui.Box(ui.BoxProps{
				Class: templutils.Cn("h-2 w-24 rounded-full", illusBgInk),
			}).Render(ctx, w); err != nil {
				return err
			}
			if err := ui.Box(ui.BoxProps{
				Class: templutils.Cn("h-8 w-full rounded-md", illusBgSoft),
			}).Render(ctx, w); err != nil {
				return err
			}
			return ui.Box(ui.BoxProps{Class: "grid grid-cols-4 gap-1"}).Render(
				templ.WithChildren(ctx, templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
					if err := ui.Box(ui.BoxProps{
						Class: templutils.Cn("col-span-3 h-2 rounded-full", illusBgSoft),
					}).Render(ctx, w); err != nil {
						return err
					}
					return ui.Box(ui.BoxProps{
						Class: templutils.Cn("h-2 rounded-full", illusBgAccent),
					}).Render(ctx, w)
				})), w)
		})), w)
}

func renderFieldLines(ctx context.Context, w io.Writer) error {
	return ui.Stack(ui.StackProps{Class: "flex h-full flex-col justify-center gap-1 p-2"}).Render(
		templ.WithChildren(ctx, templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
			if err := ui.Box(ui.BoxProps{
				Class: templutils.Cn("h-2 w-16 rounded-full", illusBgInk),
			}).Render(ctx, w); err != nil {
				return err
			}
			return ui.Box(ui.BoxProps{
				Class: templutils.Cn("h-6 w-full rounded-md", illusBgSoft),
			}).Render(ctx, w)
		})), w)
}

func renderFieldStack(ctx context.Context, w io.Writer, fields int) error {
	return ui.Stack(ui.StackProps{Class: "flex h-full flex-col justify-center gap-1 p-2"}).Render(
		templ.WithChildren(ctx, templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
			for i := 0; i < fields; i++ {
				if err := ui.Box(ui.BoxProps{
					Class: templutils.Cn("h-2 w-full rounded-md", illusBgSoft),
				}).Render(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})), w)
}

func renderToggleControl(ctx context.Context, w io.Writer) error {
	return ui.Box(ui.BoxProps{Class: "flex h-full items-center gap-2 p-2"}).Render(
		templ.WithChildren(ctx, templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
			if err := ui.Box(ui.BoxProps{
				Class: templutils.Cn("h-4 w-4 shrink-0 rounded-sm", illusBgAccent),
			}).Render(ctx, w); err != nil {
				return err
			}
			return ui.Box(ui.BoxProps{
				Class: templutils.Cn("h-2 flex-1 rounded-full", illusBgInk),
			}).Render(ctx, w)
		})), w)
}

func renderSwitchControl(ctx context.Context, w io.Writer) error {
	return ui.Box(ui.BoxProps{Class: "flex h-full items-center p-2"}).Render(
		templ.WithChildren(ctx, templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
			return ui.Box(ui.BoxProps{
				Class: templutils.Cn("flex h-5 w-10 items-center justify-end rounded-full p-1", illusBgSoft),
			}).Render(templ.WithChildren(ctx, templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
				return ui.Box(ui.BoxProps{
					Class: templutils.Cn("h-4 w-4 rounded-full", illusBgAccent),
				}).Render(ctx, w)
			})), w)
		})), w)
}

func renderMenuList(ctx context.Context, w io.Writer) error {
	widths := []string{"w-full", "w-32", "w-24"}
	return ui.Stack(ui.StackProps{Class: "flex h-full flex-col justify-center gap-1 p-2"}).Render(
		templ.WithChildren(ctx, templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
			for i, width := range widths {
				tone := illusBgSoft
				if i == 1 {
					tone = illusBgAccent
				}
				if err := ui.Box(ui.BoxProps{
					Class: templutils.Cn("h-2 rounded-full", width, tone),
				}).Render(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})), w)
}

func renderProgressLine(ctx context.Context, w io.Writer) error {
	return ui.Box(ui.BoxProps{Class: "flex h-full items-center p-2"}).Render(
		templ.WithChildren(ctx, templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
			return ui.Box(ui.BoxProps{
				Class: templutils.Cn("relative h-2 w-full rounded-full", illusBgSoft),
			}).Render(templ.WithChildren(ctx, templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
				return ui.Box(ui.BoxProps{
					Class: templutils.Cn("h-2 w-24 rounded-full", illusBgAccent),
				}).Render(ctx, w)
			})), w)
		})), w)
}

func renderSeparatorLine(ctx context.Context, w io.Writer) error {
	return ui.Box(ui.BoxProps{Class: "flex h-full items-center p-2"}).Render(
		templ.WithChildren(ctx, templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
			return ui.Box(ui.BoxProps{
				Class: templutils.Cn("h-2 w-full rounded-full", illusBgBorder),
			}).Render(ctx, w)
		})), w)
}

func renderIconMark(ctx context.Context, w io.Writer) error {
	return ui.Box(ui.BoxProps{Class: "flex h-full items-center justify-center p-2"}).Render(
		templ.WithChildren(ctx, templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
			return ui.Box(ui.BoxProps{
				Class: templutils.Cn("h-8 w-8 rounded-md", illusBgAccent),
			}).Render(ctx, w)
		})), w)
}

func renderLinkPill(ctx context.Context, w io.Writer) error {
	return ui.Box(ui.BoxProps{Class: "flex h-full items-center p-2"}).Render(
		templ.WithChildren(ctx, templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
			return ui.Box(ui.BoxProps{
				Class: templutils.Cn("h-2 w-24 rounded-full", illusBgInk),
			}).Render(ctx, w)
		})), w)
}

// IllusOpacityToneClass returns the utility class for a lab opacity swatch tone.
func IllusOpacityToneClass(tone string) string {
	switch tone {
	case "soft":
		return illusBgSoft
	case "border":
		return illusBgBorder
	case "ink":
		return illusBgInk
	case "accent":
		return illusBgAccent
	default:
		return illusBgInk
	}
}
