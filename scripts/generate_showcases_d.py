# Part 4 — more wireframe patterns, icon, language-toggle, blocks, main

def simple_wf(slug, title, desc, code, preview_body, api, related):
    page(slug, title, desc,
        [("default", "Default", "Wireframe composition from ui primitives.", code, "previewDefault")],
        api, related, WF_STACK + preview_body)

simple_wf("sheet", "Sheet", "Slide-over panel wireframe.",
    '@ui.Box { header + body }',
    '''func previewDefault(ctx context.Context, w io.Writer) error {
\treturn wfStack(ctx, w, nil, func(ctx context.Context, w io.Writer) error {
\t\tif err := ui.Box(ui.BoxProps{Class: "w-64 rounded-l-lg border border-border bg-card p-4"}).Render(templ.WithChildren(ctx, showcaseutil.Child(func(ctx context.Context, w io.Writer) error {
\t\t\tif err := ui.Title(ui.TitleProps{Order: 3, Class: "text-sm font-semibold"}, "Sheet").Render(ctx, w); err != nil { return err }
\t\t\treturn ui.Text(ui.TextProps{Class: "text-sm text-muted-foreground"}, "Side panel content.").Render(ctx, w)
\t\t})), w); err != nil { return err }
\t\treturn nil
\t})
}''',
    [("Class", "string", "Panel surface utilities")],
    [("Dialog", "/docs/components/dialog"), ("Drawer", "/docs/components/drawer")])

simple_wf("separator", "Separator", "Visual divider between sections.",
    '@ui.Box(ui.BoxProps{Class: "h-px bg-border"})',
    '''func previewDefault(ctx context.Context, w io.Writer) error {
\treturn wfStack(ctx, w, nil, func(ctx context.Context, w io.Writer) error {
\t\tif err := ui.Text(ui.TextProps{}, "Above").Render(ctx, w); err != nil { return err }
\t\tif err := ui.Box(ui.BoxProps{Class: "h-px w-full bg-border", Attrs: templ.Attributes{"role": "separator"}}).Render(ctx, w); err != nil { return err }
\t\treturn ui.Text(ui.TextProps{}, "Below").Render(ctx, w)
\t})
}''',
    [("Class", "string", "Typically h-px bg-border"), ("Role", "string", "separator")],
    [("Stack", "/docs/components/stack"), ("Card", "/docs/components/card")])

simple_wf("skeleton", "Skeleton", "Loading placeholder blocks.",
    '@ui.Box(ui.BoxProps{Class: "animate-pulse bg-muted"})',
    '''func previewDefault(ctx context.Context, w io.Writer) error {
\treturn wfStack(ctx, w, nil, func(ctx context.Context, w io.Writer) error {
\t\tif err := ui.Box(ui.BoxProps{Class: "h-4 w-3/4 max-w-xs animate-pulse rounded bg-muted"}).Render(ctx, w); err != nil { return err }
\t\treturn ui.Box(ui.BoxProps{Class: "h-4 w-1/2 max-w-xs animate-pulse rounded bg-muted"}).Render(ctx, w)
\t})
}''',
    [("Class", "string", "animate-pulse bg-muted shapes")],
    [("Progress", "/docs/components/progress"), ("Card", "/docs/components/card")])

simple_wf("progress", "Progress", "Progress indicator wireframe.",
    '@ui.Box { track + fill }',
    '''func previewDefault(ctx context.Context, w io.Writer) error {
\treturn ui.Box(ui.BoxProps{Class: "h-2 w-full max-w-xs overflow-hidden rounded-full bg-muted", Attrs: templ.Attributes{"role": "progressbar", "aria-valuenow": "60", "aria-valuemin": "0", "aria-valuemax": "100"}}).Render(templ.WithChildren(ctx, showcaseutil.Child(func(ctx context.Context, w io.Writer) error {
\t\treturn ui.Box(ui.BoxProps{Class: "h-full w-3/5 bg-primary"}).Render(ctx, w)
\t})), w)
}''',
    [("Role", "string", "progressbar"), ("AriaValuenow", "string", "Current value")],
    [("Skeleton", "/docs/components/skeleton"), ("Slider", "/docs/components/slider")])

simple_wf("avatar", "Avatar", "User avatar placeholder.",
    '@ui.Box(ui.BoxProps{Class: "rounded-full …"})',
    '''func previewDefault(ctx context.Context, w io.Writer) error {
\treturn ui.Group(ui.GroupProps{Class: "flex items-center gap-3"}).Render(templ.WithChildren(ctx, showcaseutil.Child(func(ctx context.Context, w io.Writer) error {
\t\tif err := ui.Box(ui.BoxProps{Class: "flex h-10 w-10 items-center justify-center rounded-full bg-muted text-sm font-medium"}).Render(templ.WithChildren(ctx, showcaseutil.Child(func(ctx context.Context, w io.Writer) error {
\t\t\treturn ui.Text(ui.TextProps{}, "AB").Render(ctx, w)
\t\t})), w); err != nil { return err }
\t\treturn ui.Text(ui.TextProps{Class: "text-sm font-medium"}, "Ada Lovelace").Render(ctx, w)
\t})), w)
}''',
    [("Class", "string", "rounded-full size utilities")],
    [("Badge", "/docs/components/badge"), ("Card", "/docs/components/card")])

simple_wf("toast", "Toast", "Transient notification wireframe.",
    '@ui.Box { message + action }',
    '''func previewDefault(ctx context.Context, w io.Writer) error {
\treturn ui.Box(ui.BoxProps{Class: "flex max-w-sm items-center justify-between gap-4 rounded-lg border border-border bg-card p-4 shadow"}).Render(templ.WithChildren(ctx, showcaseutil.Child(func(ctx context.Context, w io.Writer) error {
\t\tif err := ui.Text(ui.TextProps{Class: "text-sm"}, "Saved successfully.").Render(ctx, w); err != nil { return err }
\t\treturn ui.Button(ui.ButtonProps{Variant: "outline", Size: "sm"}, "Undo").Render(ctx, w)
\t})), w)
}''',
    [("Class", "string", "Card-like surface")],
    [("Alert", "/docs/components/alert"), ("Dialog", "/docs/components/dialog")])

page("toggle", "Toggle", "Pressable toggle button (wireframe).",
    [("default", "Default", "", '@ui.Button(ui.ButtonProps{Variant: "outline", Attrs: templ.Attributes{"aria-pressed": "false"}})', "previewDefault"),
     ("pressed", "Pressed", "", 'aria-pressed="true"', "previewPressed")],
    [("AriaPressed", "string", "true | false"), ("Variant", "string", "Button variant")],
    [("Toggle Group", "/docs/components/toggle-group"), ("Switch", "/docs/components/switch")],
  '''func previewDefault(ctx context.Context, w io.Writer) error {
\treturn ui.Button(ui.ButtonProps{Variant: "outline", Attrs: templ.Attributes{"aria-pressed": "false"}}, "Bold").Render(ctx, w)
}
func previewPressed(ctx context.Context, w io.Writer) error {
\treturn ui.Button(ui.ButtonProps{Variant: "secondary", Attrs: templ.Attributes{"aria-pressed": "true"}}, "Bold").Render(ctx, w)
}''')

page("toggle-group", "Toggle Group", "Grouped toggle buttons.",
    [("default", "Default", "", '@ui.Group { @ui.Button[aria-pressed] … }', "previewDefault")],
    [("Class", "string", "Group layout utilities")],
    [("Toggle", "/docs/components/toggle"), ("Tabs", "/docs/components/tabs")],
  '''func previewDefault(ctx context.Context, w io.Writer) error {
\treturn showcaseutil.RenderWithChildren(ctx, w, ui.Group(ui.GroupProps{Class: "inline-flex rounded-md border border-border"}), func(ctx context.Context, w io.Writer) error {
\t\tif err := ui.Button(ui.ButtonProps{Variant: "secondary", Size: "sm", Attrs: templ.Attributes{"aria-pressed": "true"}}, "Left").Render(ctx, w); err != nil { return err }
\t\treturn ui.Button(ui.ButtonProps{Variant: "ghost", Size: "sm", Attrs: templ.Attributes{"aria-pressed": "false"}}, "Right").Render(ctx, w)
\t})
}''')

simple_wf("pagination", "Pagination", "Page navigation controls.",
    '@ui.Group { prev / numbers / next }',
    '''func previewDefault(ctx context.Context, w io.Writer) error {
\treturn showcaseutil.RenderWithChildren(ctx, w, ui.Group(ui.GroupProps{Class: "flex items-center gap-1"}), func(ctx context.Context, w io.Writer) error {
\t\tif err := ui.Button(ui.ButtonProps{Variant: "outline", Size: "sm"}, "Prev").Render(ctx, w); err != nil { return err }
\t\tif err := ui.Button(ui.ButtonProps{Variant: "secondary", Size: "sm"}, "1").Render(ctx, w); err != nil { return err }
\t\treturn ui.Button(ui.ButtonProps{Variant: "outline", Size: "sm"}, "Next").Render(ctx, w)
\t})
}''',
    [("Buttons", "Button", "Prev / page / Next")],
    [("Data Table", "/docs/components/data-table"), ("Breadcrumb", "/docs/components/breadcrumb")])

simple_wf("tooltip", "Tooltip", "Hint on hover/focus (wireframe).",
    '@ui.Button + @ui.Box[role=tooltip]',
    '''func previewDefault(ctx context.Context, w io.Writer) error {
\treturn showcaseutil.RenderWithChildren(ctx, w, ui.Stack(ui.StackProps{Class: "gap-2 items-start"}), func(ctx context.Context, w io.Writer) error {
\t\tif err := ui.Button(ui.ButtonProps{Attrs: templ.Attributes{"aria-describedby": "tip-demo"}}, "Hover me").Render(ctx, w); err != nil { return err }
\t\treturn ui.Box(ui.BoxProps{Class: "rounded border border-border bg-popover px-2 py-1 text-xs", Attrs: templ.Attributes{"id": "tip-demo", "role": "tooltip"}}, ).Render(templ.WithChildren(ctx, showcaseutil.Child(func(ctx context.Context, w io.Writer) error {
\t\t\treturn ui.Text(ui.TextProps{}, "Tooltip text").Render(ctx, w)
\t\t})), w)
\t})
}''',
    [("Role", "string", "tooltip on hint box")],
    [("Popover", "/docs/components/popover"), ("Button", "/docs/components/button")])

simple_wf("popover", "Popover", "Floating content panel.",
    '@ui.Button + floating @ui.Box',
    '''func previewDefault(ctx context.Context, w io.Writer) error {
\treturn showcaseutil.RenderWithChildren(ctx, w, ui.Stack(ui.StackProps{Class: "gap-2"}), func(ctx context.Context, w io.Writer) error {
\t\tif err := ui.Button(ui.ButtonProps{Variant: "outline"}, "Open popover").Render(ctx, w); err != nil { return err }
\t\treturn ui.Box(ui.BoxProps{Class: "w-56 rounded-lg border border-border bg-card p-3 text-sm shadow"}).Render(templ.WithChildren(ctx, showcaseutil.Child(func(ctx context.Context, w io.Writer) error {
\t\t\treturn ui.Text(ui.TextProps{}, "Popover body copy.").Render(ctx, w)
\t\t})), w)
\t})
}''',
    [("Class", "string", "Floating panel utilities")],
    [("Dropdown Menu", "/docs/components/dropdown-menu"), ("Tooltip", "/docs/components/tooltip")])

simple_wf("dropdown-menu", "Dropdown Menu", "Menu triggered by a button.",
    '@ui.Group { trigger + menu list }',
    '''func previewDefault(ctx context.Context, w io.Writer) error {
\treturn showcaseutil.RenderWithChildren(ctx, w, ui.Stack(ui.StackProps{Class: "gap-2"}), func(ctx context.Context, w io.Writer) error {
\t\tif err := ui.Button(ui.ButtonProps{Variant: "outline"}, "Open menu").Render(ctx, w); err != nil { return err }
\t\treturn showcaseutil.RenderWithChildren(ctx, w, ui.List(ui.ListProps{Tag: "menu", Class: "w-40 rounded-md border border-border bg-card p-1 text-sm"}), func(ctx context.Context, w io.Writer) error {
\t\t\tfor _, label := range []string{"Profile", "Settings", "Sign out"} {
\t\t\t\tif err := ui.ListItem(ui.ListItemProps{Class: "rounded px-2 py-1 hover:bg-accent"}).Render(templ.WithChildren(ctx, showcaseutil.Child(func(ctx context.Context, w io.Writer) error {
\t\t\t\t\treturn ui.Button(ui.ButtonProps{Variant: "ghost", Class: "h-8 w-full justify-start px-2"}, label).Render(ctx, w)
\t\t\t\t})), w); err != nil { return err }
\t\t\t}
\t\t\treturn nil
\t\t})
\t})
}''',
    [("List", "List", "menu tag for items")],
    [("Context Menu", "/docs/components/context-menu"), ("Menubar", "/docs/components/menubar")])

exec(open(ROOT / "scripts" / "generate_showcases_e.py").read(), globals())
