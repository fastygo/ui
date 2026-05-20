#!/usr/bin/env python3
"""Generate internal/showcase/catalog/<slug>/{doc.go,showcase.go}."""
from __future__ import annotations

import re
from pathlib import Path

ROOT = Path(__file__).resolve().parents[1]
OUT = ROOT / "internal" / "showcase" / "catalog"
REG = ROOT / "internal" / "registry" / "component_imports.go"


def go_pkg(slug: str) -> str:
    if slug == "select":
        return "selectfield"
    if slug == "switch":
        return "formswitch"
    return re.sub(r"[^a-z0-9]", "", slug.lower())


def write_go(slug: str, doc: str, showcase: str) -> None:
    d = OUT / slug
    d.mkdir(parents=True, exist_ok=True)
    (d / "doc.go").write_text(doc, encoding="utf-8")
    (d / "showcase.go").write_text(showcase, encoding="utf-8")


def doc_go(slug: str, title: str, desc: str, section: str, source: str, pkg: str,
           variants: list, api: list, related: list) -> str:
    p = go_pkg(slug)
    vs = "\n".join(
        f'\t\t\tshowcaseutil.Variant("{a}", "{b}", "{c}", `{d}`, {e}),'
        for a, b, c, d, e in variants
    )
    ap = "\n".join(f'\t\t\t{{Name: "{a}", Type: "{b}", Description: "{c}"}},' for a, b, c in api)
    rl = "\n".join(f'\t\t\t{{Label: "{a}", Href: "{b}"}},' for a, b in related)
    return f"""package {p}

import (
\t"github.com/fastygo/ui/internal/registry"
\t"github.com/fastygo/ui/internal/showcase/showcaseutil"
)

func init() {{
\tregistry.Register(registry.Page{{
\t\tSlug:        "{slug}",
\t\tTitle:       "{title}",
\t\tSection:     "{section}",
\t\tDescription: "{desc}",
\t\tSource:      "{source}",
\t\tPackage:     "{pkg}",
\t\tVariants: []registry.Variant{{
{vs}
\t\t}},
\t\tAPI: []registry.APIField{{
{ap}
\t\t}},
\t\tRelated: []registry.RelatedLink{{
{rl}
\t\t}},
\t}})
}}
"""


def show_go(pkg: str, imports: str, body: str) -> str:
    return f"package {pkg}\n\nimport (\n{imports}\n)\n\n{body}"


UI = (
    '\t"context"\n\t"fmt"\n\t"io"\n\n\t"github.com/a-h/templ"\n'
    '\t"github.com/fastygo/templ/ui"\n'
    '\t"github.com/fastygo/ui/internal/showcase/showcaseutil"'
)
CMP = (
    '\t"context"\n\t"fmt"\n\t"io"\n\n\t"github.com/a-h/templ"\n'
    '\tcmp "github.com/fastygo/templ/components"\n'
    '\t"github.com/fastygo/templ/ui"\n'
    '\t"github.com/fastygo/ui/internal/showcase/showcaseutil"'
)
ICON_IMP = (
    '\t"context"\n\t"io"\n\n\t"github.com/fastygo/framework/pkg/web/view"\n'
    '\t"github.com/fastygo/ui/internal/ui/components/icon"\n'
    '\t"github.com/fastygo/ui/internal/ui/components/toggles"'
)
BLOCK_IMP = (
    '\t"context"\n\t"io"\n\n\t"github.com/fastygo/templ/ui"\n'
    '\t"github.com/fastygo/ui/internal/showcase/showcaseutil"'
)

A = 'templ.Attributes{"data-ui8kit": "%s"}'

# slug -> (title, desc, section, source, pkg, variants, api, related, imports, body)
PAGES: dict[str, tuple] = {}

def page(slug, title, desc, variants, api, related, body, *,
         section="components", source="github.com/fastygo/templ/ui",
         pkg=None, imports=UI):
    PAGES[slug] = (title, desc, section, source, pkg or source, variants, api, related, imports, body)


# --- generated body helpers referenced in PAGES ---
# Layout
page("container", "Container", "Centers content with a max-width constraint.",
    [("default", "Default", "", '@ui.Container(ui.ContainerProps{Class: "mx-auto max-w-3xl px-4"}) { … }', "previewDefault"),
     ("section", "Section", "", '@ui.Container(ui.ContainerProps{Tag: "section"}) { … }', "previewSection")],
    [("Class", "string", "Tailwind utilities"), ("Tag", "string", "motion.div | section | main"), ("Attrs", "templ.Attributes", "Extra attributes")],
    [("Stack", "/docs/components/stack"), ("Box", "/docs/components/box")],
    '''func previewDefault(ctx context.Context, w io.Writer) error {
\treturn showcaseutil.RenderWithChildren(ctx, w, ui.Container(ui.ContainerProps{Class: "mx-auto max-w-3xl border border-border rounded-lg p-4"}), func(ctx context.Context, w io.Writer) error {
\t\treturn ui.Text(ui.TextProps{}, "Container content.").Render(ctx, w)
\t})
}
func previewSection(ctx context.Context, w io.Writer) error {
\treturn showcaseutil.RenderWithChildren(ctx, w, ui.Container(ui.ContainerProps{Tag: "section", Class: "mx-auto max-w-2xl border border-border rounded-lg p-4"}), func(ctx context.Context, w io.Writer) error {
\t\treturn ui.Text(ui.TextProps{}, "Section landmark.").Render(ctx, w)
\t})
}''')

page("group", "Group", "Horizontal flex row for grouping controls.",
    [("default", "Default", "", '@ui.Group(ui.GroupProps{Class: "flex items-center gap-2"}) { … }', "previewDefault"),
     ("wrap", "Wrap", "", '@ui.Group(ui.GroupProps{Class: "flex flex-wrap gap-2"}) { … }', "previewWrap")],
    [("Class", "string", "Tailwind utilities"), ("Tag", "string", "div | span"), ("Attrs", "templ.Attributes", "Extra attributes")],
    [("Stack", "/docs/components/stack"), ("Box", "/docs/components/box")],
    '''func previewDefault(ctx context.Context, w io.Writer) error {
\treturn showcaseutil.RenderWithChildren(ctx, w, ui.Group(ui.GroupProps{Class: "flex items-center gap-2"}), func(ctx context.Context, w io.Writer) error {
\t\tif err := ui.Button(ui.ButtonProps{Size: "sm"}, "One").Render(ctx, w); err != nil { return err }
\t\treturn ui.Button(ui.ButtonProps{Size: "sm", Variant: "outline"}, "Two").Render(ctx, w)
\t})
}
func previewWrap(ctx context.Context, w io.Writer) error {
\treturn showcaseutil.RenderWithChildren(ctx, w, ui.Group(ui.GroupProps{Class: "flex flex-wrap gap-2 max-w-xs"}), func(ctx context.Context, w io.Writer) error {
\t\tfor i := 1; i <= 4; i++ {
\t\t\tif err := ui.Badge(ui.BadgeProps{}).Render(templ.WithChildren(ctx, showcaseutil.Child(func(ctx context.Context, w io.Writer) error {
\t\t\t\treturn ui.Text(ui.TextProps{}, fmt.Sprint(i)).Render(ctx, w)
\t\t\t})), w); err != nil { return err }
\t\t}
\t\treturn nil
\t})
}''')

for _part in ("b", "c", "d", "e"):
    exec(open(ROOT / f"scripts/generate_showcases_{_part}.py").read(), globals())

if __name__ == "__main__":
    main()
