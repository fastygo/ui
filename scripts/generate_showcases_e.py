# Part 5 — remaining wireframe, grid, icon, language-toggle, blocks, main

page("grid", "Grid", "CSS grid layout with columns.",
    [("default", "Two columns", "", '@ui.Grid(ui.GridProps{Class: "grid-cols-2 gap-4"}) { … }', "previewDefault"),
     ("three", "Three columns", "", '@ui.Grid(ui.GridProps{Class: "grid-cols-3 gap-2"}) { … }', "previewThree")],
    [("Class", "string", "grid-cols-* and gap utilities"), ("GridCol", "component", "Column cell wrapper")],
    [("Stack", "/docs/components/stack"), ("Container", "/docs/components/container")],
    '''func previewDefault(ctx context.Context, w io.Writer) error { return gridDemo(ctx, w, "grid-cols-2 gap-4", 2) }
func previewThree(ctx context.Context, w io.Writer) error { return gridDemo(ctx, w, "grid-cols-3 gap-2", 3) }
func gridDemo(ctx context.Context, w io.Writer, class string, n int) error {
\treturn showcaseutil.RenderWithChildren(ctx, w, ui.Grid(ui.GridProps{Class: class + " max-w-md"}), func(ctx context.Context, w io.Writer) error {
\t\tfor i := 1; i <= n; i++ {
\t\t\tlabel := fmt.Sprint(i)
\t\t\tif err := ui.GridCol(ui.GridColProps{}).Render(templ.WithChildren(ctx, showcaseutil.Child(func(ctx context.Context, w io.Writer) error {
\t\t\t\treturn ui.Box(ui.BoxProps{Class: "rounded border border-border p-3 text-center text-sm"}).Render(templ.WithChildren(ctx, showcaseutil.Child(func(ctx context.Context, w io.Writer) error {
\t\t\t\t\treturn ui.Text(ui.TextProps{}, label).Render(ctx, w)
\t\t\t\t})), w)
\t\t\t})), w); err != nil { return err }
\t\t}
\t\treturn nil
\t})
}''')

simple_wf("navigation-menu", "Navigation Menu", "Site navigation with sections.",
    '@ui.List[Tag=menu] { links }',
    '''func previewDefault(ctx context.Context, w io.Writer) error {
\treturn showcaseutil.RenderWithChildren(ctx, w, ui.List(ui.ListProps{Tag: "menu", Class: "flex gap-4 text-sm"}), func(ctx context.Context, w io.Writer) error {
\t\tfor _, label := range []string{"Home", "Docs", "Blog"} {
\t\t\tif err := ui.ListItem(ui.ListItemProps{}).Render(templ.WithChildren(ctx, showcaseutil.Child(func(ctx context.Context, w io.Writer) error {
\t\t\t\treturn ui.Button(ui.ButtonProps{Variant: "link"}, label).Render(ctx, w)
\t\t\t})), w); err != nil { return err }
\t\t}
\t\treturn nil
\t})
}''',
    [("Tag", "string", "menu for menubar-style lists")],
    [("Menubar", "/docs/components/menubar"), ("Breadcrumb", "/docs/components/breadcrumb")])

simple_wf("menubar", "Menubar", "Horizontal menu bar.",
    '@ui.Group + menu triggers',
    '''func previewDefault(ctx context.Context, w io.Writer) error {
\treturn showcaseutil.RenderWithChildren(ctx, w, ui.Group(ui.GroupProps{Class: "flex gap-1 rounded-md border border-border bg-card p-1 text-sm"}), func(ctx context.Context, w io.Writer) error {
\t\tfor _, label := range []string{"File", "Edit", "View"} {
\t\t\tif err := ui.Button(ui.ButtonProps{Variant: "ghost", Size: "sm"}, label).Render(ctx, w); err != nil { return err }
\t\t}
\t\treturn nil
\t})
}''',
    [("Group", "Group", "Horizontal button row")],
    [("Navigation Menu", "/docs/components/navigation-menu"), ("Dropdown Menu", "/docs/components/dropdown-menu")])

page("combobox", "Combobox", "Searchable select wireframe (data-ui8kit combobox).",
    [("default", "Default", "", '@ui.Input + @ui.List { data-ui8kit=combobox }', "previewDefault")],
    [("Input", "Input", "Filter field"), ("List", "List", "Options listbox")],
    [("Select", "/docs/components/select"), ("Command", "/docs/components/command")],
    WF_STACK + '''func previewDefault(ctx context.Context, w io.Writer) error {
\treturn wfStack(ctx, w, ''' + (A % "combobox") + ''', func(ctx context.Context, w io.Writer) error {
\t\tif err := ui.Input(ui.InputProps{Placeholder: "Search framework…", Attrs: templ.Attributes{"role": "combobox", "aria-expanded": "true"}}).Render(ctx, w); err != nil { return err }
\t\treturn showcaseutil.RenderWithChildren(ctx, w, ui.List(ui.ListProps{Class: "rounded-md border border-border bg-card p-1 text-sm", Attrs: templ.Attributes{"role": "listbox"}}), func(ctx context.Context, w io.Writer) error {
\t\t\tfor _, opt := range []string{"Go", "TypeScript", "Rust"} {
\t\t\t\tif err := ui.ListItem(ui.ListItemProps{}).Render(templ.WithChildren(ctx, showcaseutil.Child(func(ctx context.Context, w io.Writer) error {
\t\t\t\t\treturn ui.Button(ui.ButtonProps{Variant: "ghost", Class: "h-8 w-full justify-start"}, opt).Render(ctx, w)
\t\t\t\t})), w); err != nil { return err }
\t\t\t}
\t\t\treturn nil
\t\t})
\t})
}''')

simple_wf("data-table", "Data Table", "Table with toolbar (wireframe).",
    '@ui.Table inside @ui.Stack',
    '''func previewDefault(ctx context.Context, w io.Writer) error {
\treturn showcaseutil.RenderWithChildren(ctx, w, ui.Stack(ui.StackProps{Class: "gap-3 max-w-lg"}), func(ctx context.Context, w io.Writer) error {
\t\tif err := ui.Group(ui.GroupProps{Class: "flex items-center justify-between"}).Render(templ.WithChildren(ctx, showcaseutil.Child(func(ctx context.Context, w io.Writer) error {
\t\t\tif err := ui.Input(ui.InputProps{Placeholder: "Filter rows…", Class: "max-w-xs"}).Render(ctx, w); err != nil { return err }
\t\t\treturn ui.Button(ui.ButtonProps{Size: "sm", Variant: "outline"}, "Columns").Render(ctx, w)
\t\t})), w); err != nil { return err }
\t\treturn showcaseutil.RenderWithChildren(ctx, w, ui.Table(ui.TableProps{Class: "w-full text-sm border border-border"}), func(ctx context.Context, w io.Writer) error {
\t\t\treturn showcaseutil.RenderWithChildren(ctx, w, ui.TableBody(ui.TableSectionProps{}), func(ctx context.Context, w io.Writer) error {
\t\t\t\treturn showcaseutil.RenderWithChildren(ctx, w, ui.TableRow(ui.TableRowProps{}), func(ctx context.Context, w io.Writer) error {
\t\t\t\t\tif err := ui.TableCell(ui.TableCellProps{}, "Row A").Render(ctx, w); err != nil { return err }
\t\t\t\t\treturn ui.TableCell(ui.TableCellProps{}, "Active").Render(ctx, w)
\t\t\t\t})
\t\t\t})
\t\t})
\t})
}''',
    [("Table", "Table", "Semantic table"), ("Input", "Input", "Filter field")],
    [("Table", "/docs/components/table"), ("Pagination", "/docs/components/pagination")])

simple_wf("alert-dialog", "Alert Dialog", "Modal that interrupts flow.",
    '@ui.Box { alert copy + confirm }',
    '''func previewDefault(ctx context.Context, w io.Writer) error {
\treturn showcaseutil.RenderWithChildren(ctx, w, ui.Box(ui.BoxProps{Class: "max-w-sm rounded-lg border border-border bg-card p-4"}), func(ctx context.Context, w io.Writer) error {
\t\tif err := ui.Title(ui.TitleProps{Order: 3, Class: "text-base font-semibold"}, "Are you sure?").Render(ctx, w); err != nil { return err }
\t\tif err := ui.Text(ui.TextProps{Class: "text-sm text-muted-foreground"}, "This action cannot be undone.").Render(ctx, w); err != nil { return err }
\t\treturn ui.Group(ui.GroupProps{Class: "mt-4 flex justify-end gap-2"}).Render(templ.WithChildren(ctx, showcaseutil.Child(func(ctx context.Context, w io.Writer) error {
\t\t\tif err := ui.Button(ui.ButtonProps{Variant: "outline"}, "Cancel").Render(ctx, w); err != nil { return err }
\t\t\treturn ui.Button(ui.ButtonProps{Variant: "destructive"}, "Delete").Render(ctx, w)
\t\t})), w)
\t})
}''',
    [("Title", "Title", "Alert heading")],
    [("Dialog", "/docs/components/dialog"), ("Alert", "/docs/components/alert")])

simple_wf("drawer", "Drawer", "Bottom sheet drawer wireframe.",
    '@ui.Box anchored to bottom',
    '''func previewDefault(ctx context.Context, w io.Writer) error {
\treturn ui.Box(ui.BoxProps{Class: "w-full max-w-md rounded-t-xl border border-border bg-card p-4"}).Render(templ.WithChildren(ctx, showcaseutil.Child(func(ctx context.Context, w io.Writer) error {
\t\tif err := ui.Box(ui.BoxProps{Class: "mx-auto mb-3 h-1 w-10 rounded-full bg-muted"}).Render(ctx, w); err != nil { return err }
\t\treturn ui.Text(ui.TextProps{}, "Drawer content.").Render(ctx, w)
\t})), w)
}''',
    [("Class", "string", "Bottom sheet surface")],
    [("Sheet", "/docs/components/sheet"), ("Dialog", "/docs/components/dialog")])

simple_wf("hover-card", "Hover Card", "Rich preview on hover.",
    '@ui.Group { trigger + card }',
    '''func previewDefault(ctx context.Context, w io.Writer) error {
\treturn showcaseutil.RenderWithChildren(ctx, w, ui.Group(ui.GroupProps{Class: "flex items-start gap-3"}), func(ctx context.Context, w io.Writer) error {
\t\tif err := ui.Button(ui.ButtonProps{Variant: "link"}, "@user").Render(ctx, w); err != nil { return err }
\t\treturn ui.Box(ui.BoxProps{Class: "w-56 rounded-lg border border-border bg-card p-3 text-sm"}).Render(templ.WithChildren(ctx, showcaseutil.Child(func(ctx context.Context, w io.Writer) error {
\t\t\treturn ui.Text(ui.TextProps{}, "Profile preview card.").Render(ctx, w)
\t\t})), w)
\t})
}''',
    [("Box", "Box", "Preview card surface")],
    [("Popover", "/docs/components/popover"), ("Tooltip", "/docs/components/tooltip")])

page("collapsible", "Collapsible", "Single expand/collapse region.",
    [("default", "Default", "", '@ui.Button + @ui.Box { data-ui8kit=disclosure }', "previewDefault")],
    [("Trigger", "Button", "Expands panel"), ("Panel", "Box", "Collapsible content")],
    [("Accordion", "/docs/components/accordion"), ("Sheet", "/docs/components/sheet")],
    WF_STACK + '''func previewDefault(ctx context.Context, w io.Writer) error {
\treturn wfStack(ctx, w, ''' + (A % "disclosure") + ''', func(ctx context.Context, w io.Writer) error {
\t\tif err := ui.Button(ui.ButtonProps{Variant: "ghost", Attrs: templ.Attributes{"data-ui8kit-trigger": "", "aria-expanded": "false"}}, "Show more").Render(ctx, w); err != nil { return err }
\t\treturn ui.Box(ui.BoxProps{Class: "rounded border border-border p-3 text-sm", Attrs: templ.Attributes{"data-ui8kit-panel": "", "hidden": true}}).Render(templ.WithChildren(ctx, showcaseutil.Child(func(ctx context.Context, w io.Writer) error {
\t\t\treturn ui.Text(ui.TextProps{}, "Hidden details.").Render(ctx, w)
\t\t})), w)
\t})
}''')

simple_wf("aspect-ratio", "Aspect Ratio", "Fixed aspect container.",
    '@ui.Box(ui.BoxProps{Class: "aspect-video …"})',
    '''func previewDefault(ctx context.Context, w io.Writer) error {
\treturn ui.Box(ui.BoxProps{Class: "aspect-video w-full max-w-xs overflow-hidden rounded-lg border border-border bg-muted"}).Render(templ.WithChildren(ctx, showcaseutil.Child(func(ctx context.Context, w io.Writer) error {
\t\treturn ui.Box(ui.BoxProps{Class: "flex h-full items-center justify-center text-sm text-muted-foreground"}).Render(templ.WithChildren(ctx, showcaseutil.Child(func(ctx context.Context, w io.Writer) error {
\t\t\treturn ui.Text(ui.TextProps{}, "16:9").Render(ctx, w)
\t\t})), w)
\t})), w)
}''',
    [("Class", "string", "aspect-video | aspect-square")],
    [("Carousel", "/docs/components/carousel"), ("Card", "/docs/components/card")])

page("carousel", "Carousel", "Horizontal scrolling list.",
    [("default", "Default", "", '@ui.Group { slides }', "previewDefault")],
    [("Group", "Group", "Horizontal flex row")],
    [("Aspect Ratio", "/docs/components/aspect-ratio"), ("Card", "/docs/components/card")],
    '''func previewDefault(ctx context.Context, w io.Writer) error {
\treturn showcaseutil.RenderWithChildren(ctx, w, ui.Group(ui.GroupProps{Class: "flex gap-2 overflow-x-auto max-w-md"}), func(ctx context.Context, w io.Writer) error {
\t\tfor i := 1; i <= 3; i++ {
\t\t\tlabel := fmt.Sprintf("Slide %d", i)
\t\t\tif err := ui.Box(ui.BoxProps{Class: "min-w-[8rem] shrink-0 rounded-lg border border-border bg-card p-4 text-sm"}).Render(templ.WithChildren(ctx, showcaseutil.Child(func(ctx context.Context, w io.Writer) error {
\t\t\t\treturn ui.Text(ui.TextProps{}, label).Render(ctx, w)
\t\t\t})), w); err != nil { return err }
\t\t}
\t\treturn nil
\t})
}''')

simple_wf("calendar", "Calendar", "Date grid placeholder.",
    '@ui.Box { month label + day grid }',
    '''func previewDefault(ctx context.Context, w io.Writer) error {
\treturn showcaseutil.RenderWithChildren(ctx, w, ui.Box(ui.BoxProps{Class: "w-64 rounded-lg border border-border p-3"}), func(ctx context.Context, w io.Writer) error {
\t\tif err := ui.Text(ui.TextProps{Class: "text-sm font-medium"}, "May 2026 — placeholder").Render(ctx, w); err != nil { return err }
\t\treturn ui.Grid(ui.GridProps{Class: "mt-2 grid-cols-7 gap-1 text-center text-xs text-muted-foreground"}).Render(templ.WithChildren(ctx, showcaseutil.Child(func(ctx context.Context, w io.Writer) error {
\t\t\tfor d := 1; d <= 7; d++ {
\t\t\t\tif err := ui.Text(ui.TextProps{}, fmt.Sprint(d)).Render(ctx, w); err != nil { return err }
\t\t\t}
\t\t\treturn nil
\t\t})), w)
\t})
}''',
    [("Class", "string", "Calendar frame utilities")],
    [("Command", "/docs/components/command"), ("Popover", "/docs/components/popover")])

page("command", "Command", "Command palette: search + list.",
    [("default", "Default", "", '@ui.Input + @ui.List', "previewDefault")],
    [("Input", "Input", "Search field"), ("List", "List", "Commands")],
    [("Combobox", "/docs/components/combobox"), ("Dialog", "/docs/components/dialog")],
    '''func previewDefault(ctx context.Context, w io.Writer) error {
\treturn showcaseutil.RenderWithChildren(ctx, w, ui.Stack(ui.StackProps{Class: "gap-2 max-w-sm rounded-lg border border-border bg-card p-2"}), func(ctx context.Context, w io.Writer) error {
\t\tif err := ui.Input(ui.InputProps{Placeholder: "Type a command…"}).Render(ctx, w); err != nil { return err }
\t\treturn showcaseutil.RenderWithChildren(ctx, w, ui.List(ui.ListProps{Class: "text-sm"}), func(ctx context.Context, w io.Writer) error {
\t\t\tfor _, cmd := range []string{"Open docs", "Toggle theme", "Go home"} {
\t\t\t\tif err := ui.ListItem(ui.ListItemProps{}).Render(templ.WithChildren(ctx, showcaseutil.Child(func(ctx context.Context, w io.Writer) error {
\t\t\t\t\treturn ui.Button(ui.ButtonProps{Variant: "ghost", Class: "h-8 w-full justify-start"}, cmd).Render(ctx, w)
\t\t\t\t})), w); err != nil { return err }
\t\t\t}
\t\t\treturn nil
\t\t})
\t})
}''')

simple_wf("context-menu", "Context Menu", "Stub menu list on right-click target.",
    '@ui.Box + @ui.List[menu]',
    '''func previewDefault(ctx context.Context, w io.Writer) error {
\treturn showcaseutil.RenderWithChildren(ctx, w, ui.Stack(ui.StackProps{Class: "gap-2"}), func(ctx context.Context, w io.Writer) error {
\t\tif err := ui.Box(ui.BoxProps{Class: "rounded-lg border border-dashed border-border p-8 text-center text-sm text-muted-foreground"}).Render(templ.WithChildren(ctx, showcaseutil.Child(func(ctx context.Context, w io.Writer) error {
\t\t\treturn ui.Text(ui.TextProps{}, "Right-click area (stub)").Render(ctx, w)
\t\t})), w); err != nil { return err }
\t\treturn showcaseutil.RenderWithChildren(ctx, w, ui.List(ui.ListProps{Tag: "menu", Class: "w-40 rounded-md border border-border bg-card p-1 text-sm"}), func(ctx context.Context, w io.Writer) error {
\t\t\tfor _, label := range []string{"Copy", "Paste"} {
\t\t\t\tif err := ui.ListItem(ui.ListItemProps{}).Render(templ.WithChildren(ctx, showcaseutil.Child(func(ctx context.Context, w io.Writer) error {
\t\t\t\t\treturn ui.Text(ui.TextProps{}, label).Render(ctx, w)
\t\t\t\t})), w); err != nil { return err }
\t\t\t}
\t\t\treturn nil
\t\t})
\t})
}''',
    [("List", "List", "menu tag")],
    [("Dropdown Menu", "/docs/components/dropdown-menu"), ("Menubar", "/docs/components/menubar")])

# App registry
page("icon", "Icon", "Latty icon mask (app registry).",
    [("default", "Default", "", '@icon.Icon(icon.IconProps{Name: "home"})', "previewDefault"),
     ("sizes", "Sizes", "", 'Size: xs | sm | md | lg', "previewSizes")],
    [("Name", "string", "Latty icon name"), ("Size", "string", "xs | sm | md | lg"), ("Class", "string", "Extra utilities")],
    [("Button", "/docs/components/button"), ("Badge", "/docs/components/badge")],
    '''func previewDefault(ctx context.Context, w io.Writer) error {
\treturn icon.Icon(icon.IconProps{Name: "home"}).Render(ctx, w)
}
func previewSizes(ctx context.Context, w io.Writer) error {
\tfor _, sz := range []string{"xs", "sm", "md", "lg"} {
\t\tif err := icon.Icon(icon.IconProps{Name: "settings", Size: sz, Class: "mr-2"}).Render(ctx, w); err != nil { return err }
\t}
\treturn nil
}''',
    imports=ICON_IMP)

page("language-toggle", "Language Toggle", "Locale switcher (app toggles package).",
    [("default", "Default", "", '@toggles.LanguageToggle(data)', "previewDefault")],
    [("CurrentLabel", "string", "Visible label"), ("NextHref", "string", "Link to next locale"), ("CurrentLocale", "string", "Active locale code")],
    [("Button", "/docs/components/button"), ("Icon", "/docs/components/icon")],
    '''func previewDefault(ctx context.Context, w io.Writer) error {
\treturn toggles.LanguageToggle(view.LanguageToggleData{
\t\tCurrentLabel: "EN", CurrentLocale: "en", NextLocale: "ru", NextHref: "/?lang=ru",
\t\tDefaultLocale: "en", AvailableLocales: []string{"en", "ru"}, Label: "Switch language",
\t}).Render(ctx, w)
}''',
    imports=ICON_IMP)

# Blocks
def block_page(slug, title, copy):
    page(slug, title, "Block scaffold — section wireframe with placeholder copy for future github.com/fastygo/blocks extraction.",
        [("default", "Wireframe", "", f'@ui.Stack {{ @ui.Title … "{title}" }}', "previewDefault"),
         ("compact", "Compact", "", "Denser spacing variant", "previewCompact")],
        [("Title", "string", "Section heading"), ("Body", "string", "Supporting copy")],
        [("Card", "/docs/components/card"), ("Stack", "/docs/components/stack")],
        f'''func previewDefault(ctx context.Context, w io.Writer) error {{ return blockBody(ctx, w, "{copy}", false) }}
func previewCompact(ctx context.Context, w io.Writer) error {{ return blockBody(ctx, w, "{copy}", true) }}
func blockBody(ctx context.Context, w io.Writer, body string, compact bool) error {{
\tgap := "gap-4"
\tif compact {{ gap = "gap-2" }}
\treturn showcaseutil.RenderWithChildren(ctx, w, ui.Stack(ui.StackProps{{Class: gap + " max-w-2xl"}}), func(ctx context.Context, w io.Writer) error {{
\t\tif err := ui.Title(ui.TitleProps{{Order: 2}}, "{title}").Render(ctx, w); err != nil {{ return err }}
\t\tif err := ui.Text(ui.TextProps{{Class: "text-sm text-muted-foreground leading-relaxed"}}, body).Render(ctx, w); err != nil {{ return err }}
\t\treturn ui.Button(ui.ButtonProps{{Variant: "outline", Size: "sm"}}, "Action").Render(ctx, w)
\t}})
}}''',
        section="blocks", source="internal/ui/blocks", pkg="internal/ui/blocks", imports=BLOCK_IMP)

block_page("dashboard-overview", "Dashboard Overview", "Wireframe metrics and summary row for dashboard blocks.")
block_page("marketing-hero", "Marketing Hero", "Wireframe hero headline and call-to-action for marketing blocks.")
block_page("docs-article", "Docs Article", "Wireframe article body for documentation blocks.")


def main():
    for slug, spec in PAGES.items():
        title, desc, section, source, pkg, variants, api, related, imports, body = spec
        write_go(slug, doc_go(slug, title, desc, section, source, pkg, variants, api, related),
                 show_go(go_pkg(slug), imports, body))

    slugs = sorted(PAGES.keys())
    lines = [f'\t_ "github.com/fastygo/ui/internal/showcase/catalog/{s}"' for s in slugs]
    reg = OUT / "register.go"
    reg.write_text(
        "// Package components registers all showcase doc pages via init.\npackage components\n\nimport (\n"
        + "\n".join(lines)
        + "\n)\n",
        encoding="utf-8",
    )
    print(f"Generated {len(slugs)} showcase packages -> {OUT}")

if __name__ == "__main__":
    main()
