# Part 2 of generate_showcases.py — page definitions and main()

page("box", "Box", "Generic block wrapper without landmark semantics.",
    [("default", "Default", "", '@ui.Box(ui.BoxProps{Class: "rounded-lg border border-border p-4"}) { … }', "previewDefault"),
     ("pre", "Pre tag", "", '@ui.Box(ui.BoxProps{Tag: "pre"}) { … }', "previewPre")],
    [("Class", "string", "Tailwind utilities"), ("Tag", "string", "motion.div | pre | span"), ("Attrs", "templ.Attributes", "Extra attributes")],
    [("Block", "/docs/components/block"), ("Stack", "/docs/components/stack")],
    '''func previewDefault(ctx context.Context, w io.Writer) error {
\treturn showcaseutil.RenderWithChildren(ctx, w, ui.Box(ui.BoxProps{Class: "rounded-lg border border-border p-4"}), func(ctx context.Context, w io.Writer) error {
\t\treturn ui.Text(ui.TextProps{}, "Box content.").Render(ctx, w)
\t})
}
func previewPre(ctx context.Context, w io.Writer) error {
\treturn showcaseutil.RenderWithChildren(ctx, w, ui.Box(ui.BoxProps{Tag: "pre", Class: "rounded-md border border-border bg-muted/30 p-3 text-xs font-mono"}), func(ctx context.Context, w io.Writer) error {
\t\treturn ui.Text(ui.TextProps{Tag: "code"}, "code snippet").Render(ctx, w)
\t})
}''')

page("block", "Block", "Top-level landmark sections (do not nest Block in Block).",
    [("main", "Main", "", '@ui.Block(ui.BlockProps{Tag: "main"}) { … }', "previewMain"),
     ("aside", "Aside", "", '@ui.Block(ui.BlockProps{Tag: "aside"}) { … }', "previewAside")],
    [("Tag", "string", "main | section | aside | nav | …"), ("Class", "string", "Tailwind utilities"), ("Attrs", "templ.Attributes", "Extra attributes")],
    [("Box", "/docs/components/box"), ("Container", "/docs/components/container")],
    '''func previewMain(ctx context.Context, w io.Writer) error {
\treturn showcaseutil.RenderWithChildren(ctx, w, ui.Block(ui.BlockProps{Tag: "main", Class: "rounded-lg border border-border p-4"}), func(ctx context.Context, w io.Writer) error {
\t\treturn ui.Text(ui.TextProps{}, "Main landmark block.").Render(ctx, w)
\t})
}
func previewAside(ctx context.Context, w io.Writer) error {
\treturn showcaseutil.RenderWithChildren(ctx, w, ui.Block(ui.BlockProps{Tag: "aside", Class: "rounded-lg border border-border p-4 w-48"}), func(ctx context.Context, w io.Writer) error {
\t\treturn ui.Text(ui.TextProps{}, "Aside block.").Render(ctx, w)
\t})
}''')

page("title", "Title", "Semantic heading with order 1–6.",
    [("h1", "Heading 1", "", '@ui.Title(ui.TitleProps{Order: 1}, "Page title")', "previewH1"),
     ("h2", "Heading 2", "", '@ui.Title(ui.TitleProps{Order: 2}, "Section")', "previewH2"),
     ("h3", "Heading 3", "", '@ui.Title(ui.TitleProps{Order: 3}, "Subsection")', "previewH3")],
    [("Order", "int", "1–6 maps to h1–h6"), ("Class", "string", "Tailwind utilities")],
    [("Text", "/docs/components/text"), ("Stack", "/docs/components/stack")],
    '''func previewH1(ctx context.Context, w io.Writer) error { return ui.Title(ui.TitleProps{Order: 1}, "Page title").Render(ctx, w) }
func previewH2(ctx context.Context, w io.Writer) error { return ui.Title(ui.TitleProps{Order: 2}, "Section").Render(ctx, w) }
func previewH3(ctx context.Context, w io.Writer) error { return ui.Title(ui.TitleProps{Order: 3}, "Subsection").Render(ctx, w) }''')

page("text", "Text", "Inline or block text with configurable tag.",
    [("default", "Paragraph", "", '@ui.Text(ui.TextProps{}, "Body copy.")', "previewDefault"),
     ("muted", "Muted", "", '@ui.Text(ui.TextProps{Class: "text-sm text-muted-foreground"}, "…")', "previewMuted"),
     ("code", "Code", "", '@ui.Text(ui.TextProps{Tag: "code"}, "npm install")', "previewCode")],
    [("Tag", "string", "p | span | code | …"), ("Class", "string", "Tailwind utilities")],
    [("Title", "/docs/components/title"), ("Label", "/docs/components/label")],
    '''func previewDefault(ctx context.Context, w io.Writer) error { return ui.Text(ui.TextProps{}, "Body copy.").Render(ctx, w) }
func previewMuted(ctx context.Context, w io.Writer) error { return ui.Text(ui.TextProps{Class: "text-sm text-muted-foreground"}, "Muted supporting text.").Render(ctx, w) }
func previewCode(ctx context.Context, w io.Writer) error { return ui.Text(ui.TextProps{Tag: "code", Class: "font-mono text-xs"}, "npm install").Render(ctx, w) }''')

page("list", "List", "Semantic ul/ol/dl list containers.",
    [("unordered", "Unordered", "", '@ui.List(ui.ListProps{Class: "list-disc pl-5"}) { … }', "previewUnordered"),
     ("ordered", "Ordered", "", '@ui.List(ui.ListProps{Tag: "ol", Class: "list-decimal pl-5"}) { … }', "previewOrdered")],
    [("Tag", "string", "ul | ol | dl | menu"), ("Class", "string", "Tailwind utilities"), ("Attrs", "templ.Attributes", "Extra attributes")],
    [("Table", "/docs/components/table"), ("Breadcrumb", "/docs/components/breadcrumb")],
    '''func previewUnordered(ctx context.Context, w io.Writer) error { return listPreview(ctx, w, ui.ListProps{Class: "list-disc pl-5 text-sm"}) }
func previewOrdered(ctx context.Context, w io.Writer) error { return listPreview(ctx, w, ui.ListProps{Tag: "ol", Class: "list-decimal pl-5 text-sm"}) }
func listPreview(ctx context.Context, w io.Writer, props ui.ListProps) error {
\treturn showcaseutil.RenderWithChildren(ctx, w, ui.List(props), func(ctx context.Context, w io.Writer) error {
\t\tfor _, item := range []string{"First item", "Second item"} {
\t\t\tif err := ui.ListItem(ui.ListItemProps{}).Render(templ.WithChildren(ctx, showcaseutil.Child(func(ctx context.Context, w io.Writer) error {
\t\t\t\treturn ui.Text(ui.TextProps{}, item).Render(ctx, w)
\t\t\t})), w); err != nil { return err }
\t\t}
\t\treturn nil
\t})
}''')

page("badge", "Badge", "Small status label chip.",
    [("default", "Default", "", '@ui.Badge(ui.BadgeProps{}, "Badge")', "previewDefault"),
     ("secondary", "Secondary", "", '@ui.Badge(ui.BadgeProps{Variant: "secondary"}, "Secondary")', "previewSecondary"),
     ("outline", "Outline", "", '@ui.Badge(ui.BadgeProps{Variant: "outline"}, "Outline")', "previewOutline"),
     ("destructive", "Destructive", "", '@ui.Badge(ui.BadgeProps{Variant: "destructive"}, "Alert")', "previewDestructive")],
    [("Variant", "string", "default | secondary | destructive | outline"), ("Size", "string", "default | sm | lg"), ("Class", "string", "Extra utilities")],
    [("Button", "/docs/components/button"), ("Alert", "/docs/components/alert")],
    '''func badge(ctx context.Context, w io.Writer, p ui.BadgeProps, label string) error {
\treturn showcaseutil.RenderWithChildren(ctx, w, ui.Badge(p), func(ctx context.Context, w io.Writer) error {
\t\treturn ui.Text(ui.TextProps{}, label).Render(ctx, w)
\t})
}
func previewDefault(ctx context.Context, w io.Writer) error { return badge(ctx, w, ui.BadgeProps{}, "Badge") }
func previewSecondary(ctx context.Context, w io.Writer) error { return badge(ctx, w, ui.BadgeProps{Variant: "secondary"}, "Secondary") }
func previewOutline(ctx context.Context, w io.Writer) error { return badge(ctx, w, ui.BadgeProps{Variant: "outline"}, "Outline") }
func previewDestructive(ctx context.Context, w io.Writer) error { return badge(ctx, w, ui.BadgeProps{Variant: "destructive"}, "Alert") }''')

page("table", "Table", "Semantic data table structure.",
    [("default", "Default", "", '@ui.Table(ui.TableProps{}) { @ui.TableHead … }', "previewDefault"),
     ("striped", "Compact", "", "Dense row styling via TableCell Class", "previewCompact")],
    [("Class", "string", "Table wrapper utilities"), ("TableHead", "component", "thead section"), ("TableBody", "component", "tbody section")],
    [("Data Table", "/docs/components/data-table"), ("List", "/docs/components/list")],
    '''func previewDefault(ctx context.Context, w io.Writer) error { return renderSampleTable(ctx, w, "") }
func previewCompact(ctx context.Context, w io.Writer) error { return renderSampleTable(ctx, w, "text-xs") }
func renderSampleTable(ctx context.Context, w io.Writer, cellClass string) error {
\treturn showcaseutil.RenderWithChildren(ctx, w, ui.Table(ui.TableProps{Class: "w-full text-sm border border-border"}), func(ctx context.Context, w io.Writer) error {
\t\thead := func(ctx context.Context, w io.Writer) error {
\t\t\treturn showcaseutil.RenderWithChildren(ctx, w, ui.TableHead(ui.TableSectionProps{}), func(ctx context.Context, w io.Writer) error {
\t\t\t\treturn showcaseutil.RenderWithChildren(ctx, w, ui.TableRow(ui.TableRowProps{}), func(ctx context.Context, w io.Writer) error {
\t\t\t\t\tif err := ui.TableHeadCell(ui.TableCellProps{Class: cellClass}, "Name").Render(ctx, w); err != nil { return err }
\t\t\t\t\treturn ui.TableHeadCell(ui.TableCellProps{Class: cellClass}, "Role").Render(ctx, w)
\t\t\t\t})
\t\t\t})
\t\t}
\t\tbody := func(ctx context.Context, w io.Writer) error {
\t\t\treturn showcaseutil.RenderWithChildren(ctx, w, ui.TableBody(ui.TableSectionProps{}), func(ctx context.Context, w io.Writer) error {
\t\t\t\trows := [][2]string{{"Ada", "Admin"}, {"Lin", "Editor"}}
\t\t\t\tfor _, row := range rows {
\t\t\t\t\tif err := showcaseutil.RenderWithChildren(ctx, w, ui.TableRow(ui.TableRowProps{}), func(ctx context.Context, w io.Writer) error {
\t\t\t\t\t\tif err := ui.TableCell(ui.TableCellProps{Class: cellClass}, row[0]).Render(ctx, w); err != nil { return err }
\t\t\t\t\t\treturn ui.TableCell(ui.TableCellProps{Class: cellClass}, row[1]).Render(ctx, w)
\t\t\t\t\t}); err != nil { return err }
\t\t\t\t}
\t\t\t\treturn nil
\t\t\t})
\t\t}
\t\tif err := head(ctx, w); err != nil { return err }
\t\treturn body(ctx, w)
\t})
}''')

# Forms
page("input", "Input", "Single-line text input control.",
    [("default", "Default", "", '@ui.Input(ui.InputProps{Placeholder: "Email"})', "previewDefault"),
     ("disabled", "Disabled", "", '@ui.Input(ui.InputProps{Disabled: true})', "previewDisabled"),
     ("file", "File", "", '@ui.Input(ui.InputProps{Type: "file"})', "previewFile")],
    [("Type", "string", "text | email | password | file | range | …"), ("Placeholder", "string", "Placeholder text"), ("Disabled", "bool", "Disables input"), ("Class", "string", "Tailwind utilities")],
    [("Textarea", "/docs/components/textarea"), ("Form", "/docs/components/form")],
    '''func previewDefault(ctx context.Context, w io.Writer) error { return ui.Input(ui.InputProps{Placeholder: "Email"}).Render(ctx, w) }
func previewDisabled(ctx context.Context, w io.Writer) error { return ui.Input(ui.InputProps{Placeholder: "Disabled", Disabled: true}).Render(ctx, w) }
func previewFile(ctx context.Context, w io.Writer) error { return ui.Input(ui.InputProps{Type: "file"}).Render(ctx, w) }''')

page("textarea", "Textarea", "Multi-line text input.",
    [("default", "Default", "", '@ui.Textarea(ui.TextareaProps{Placeholder: "Message"})', "previewDefault"),
     ("disabled", "Disabled", "", '@ui.Textarea(ui.TextareaProps{Disabled: true})', "previewDisabled")],
    [("Placeholder", "string", "Placeholder text"), ("Rows", "int", "Visible row count"), ("Disabled", "bool", "Disables control")],
    [("Input", "/docs/components/input"), ("Form", "/docs/components/form")],
    '''func previewDefault(ctx context.Context, w io.Writer) error { return ui.Textarea(ui.TextareaProps{Placeholder: "Your message"}).Render(ctx, w) }
func previewDisabled(ctx context.Context, w io.Writer) error { return ui.Textarea(ui.TextareaProps{Placeholder: "Disabled", Disabled: true}).Render(ctx, w) }''')

page("label", "Label", "Accessible label for form controls.",
    [("default", "Default", "", '@ui.Label(ui.LabelProps{HTMLFor: "email"}, "Email")', "previewDefault"),
     ("required", "Required hint", "", "Pair with aria-required on control", "previewRequired")],
    [("HTMLFor", "string", "id of associated control"), ("Class", "string", "Tailwind utilities")],
    [("Input", "/docs/components/input"), ("Checkbox", "/docs/components/checkbox")],
    '''func previewDefault(ctx context.Context, w io.Writer) error {
\tif err := ui.Label(ui.LabelProps{HTMLFor: "showcase-email"}, "Email").Render(ctx, w); err != nil { return err }
\treturn ui.Input(ui.InputProps{ID: "showcase-email", Placeholder: "you@example.com"}).Render(ctx, w)
}
func previewRequired(ctx context.Context, w io.Writer) error {
\tif err := ui.Label(ui.LabelProps{HTMLFor: "showcase-name"}, "Name").Render(ctx, w); err != nil { return err }
\treturn ui.Input(ui.InputProps{ID: "showcase-name", Required: true, Placeholder: "Required"}).Render(ctx, w)
}''')

page("checkbox", "Checkbox", "Boolean checkbox input.",
    [("default", "Default", "", '@ui.Checkbox(ui.CheckboxProps{Name: "terms"})', "previewDefault"),
     ("checked", "Checked", "", '@ui.Checkbox(ui.CheckboxProps{Checked: true})', "previewChecked")],
    [("Name", "string", "Form field name"), ("Checked", "bool", "Initial checked state"), ("Disabled", "bool", "Disables control")],
    [("Radio", "/docs/components/radio"), ("Switch", "/docs/components/switch")],
    '''func previewDefault(ctx context.Context, w io.Writer) error { return ui.Checkbox(ui.CheckboxProps{Name: "terms", AriaLabel: "Accept terms"}).Render(ctx, w) }
func previewChecked(ctx context.Context, w io.Writer) error { return ui.Checkbox(ui.CheckboxProps{Name: "terms", Checked: true, AriaLabel: "Accepted"}).Render(ctx, w) }''')

page("radio", "Radio", "Single choice within a group.",
    [("default", "Default", "", '@ui.Radio(ui.RadioProps{Name: "plan", Value: "free"})', "previewDefault"),
     ("checked", "Selected", "", '@ui.Radio(ui.RadioProps{Checked: true})', "previewChecked")],
    [("Name", "string", "Group name"), ("Value", "string", "Option value"), ("Checked", "bool", "Selected state")],
    [("Checkbox", "/docs/components/checkbox"), ("Select", "/docs/components/select")],
    '''func previewDefault(ctx context.Context, w io.Writer) error { return ui.Radio(ui.RadioProps{Name: "plan", Value: "free", AriaLabel: "Free plan"}).Render(ctx, w) }
func previewChecked(ctx context.Context, w io.Writer) error { return ui.Radio(ui.RadioProps{Name: "plan", Value: "pro", Checked: true, AriaLabel: "Pro plan"}).Render(ctx, w) }''')

page("select", "Select", "Native select dropdown (ui.Select / selectfield).",
    [("default", "Default", "", '@ui.Select(ui.SelectProps{Name: "role", Options: opts})', "previewDefault"),
     ("disabled", "Disabled", "", '@ui.Select(ui.SelectProps{Disabled: true})', "previewDisabled")],
    [("Options", "[]ui.Option", "Value/label pairs"), ("Name", "string", "Form field name"), ("Value", "string", "Selected value")],
    [("Combobox", "/docs/components/combobox"), ("Radio", "/docs/components/radio")],
    '''var selectOpts = []ui.Option{{Value: "viewer", Label: "Viewer"}, {Value: "editor", Label: "Editor"}}
func previewDefault(ctx context.Context, w io.Writer) error { return ui.Select(ui.SelectProps{Name: "role", Options: selectOpts, Value: "viewer"}).Render(ctx, w) }
func previewDisabled(ctx context.Context, w io.Writer) error { return ui.Select(ui.SelectProps{Name: "role", Options: selectOpts, Disabled: true}).Render(ctx, w) }''')

page("form", "Form", "Form landmark with item helpers.",
    [("default", "Login", "", '@ui.Form(ui.FormProps{}) { @ui.FormItem … }', "previewDefault"),
     ("inline", "Inline", "", "Compact horizontal FormItem layout", "previewInline")],
    [("Action", "string", "Form action URL"), ("Method", "string", "GET | POST"), ("FormItem", "component", "Label + control group")],
    [("Input", "/docs/components/input"), ("Button", "/docs/components/button")],
    '''func previewDefault(ctx context.Context, w io.Writer) error {
\treturn showcaseutil.RenderWithChildren(ctx, w, ui.Form(ui.FormProps{Class: "max-w-sm"}), func(ctx context.Context, w io.Writer) error {
\t\titem := func(ctx context.Context, w io.Writer) error {
\t\t\treturn showcaseutil.RenderWithChildren(ctx, w, ui.FormItem(ui.FormItemProps{}), func(ctx context.Context, w io.Writer) error {
\t\t\t\tif err := ui.Label(ui.LabelProps{HTMLFor: "login-email"}, "Email").Render(ctx, w); err != nil { return err }
\t\t\t\treturn ui.Input(ui.InputProps{ID: "login-email", Type: "email", Placeholder: "you@example.com"}).Render(ctx, w)
\t\t\t})
\t\t}
\t\tif err := item(ctx, w); err != nil { return err }
\t\treturn ui.Button(ui.ButtonProps{Type: "submit"}, "Sign in").Render(ctx, w)
\t})
}
func previewInline(ctx context.Context, w io.Writer) error {
\treturn showcaseutil.RenderWithChildren(ctx, w, ui.Form(ui.FormProps{Class: "max-w-md"}), func(ctx context.Context, w io.Writer) error {
\t\treturn showcaseutil.RenderWithChildren(ctx, w, ui.Group(ui.GroupProps{Class: "flex items-end gap-2"}), func(ctx context.Context, w io.Writer) error {
\t\t\tif err := ui.Input(ui.InputProps{Placeholder: "Search"}).Render(ctx, w); err != nil { return err }
\t\t\treturn ui.Button(ui.ButtonProps{}, "Go").Render(ctx, w)
\t\t})
\t})
}''')

page("switch", "Switch", "Toggle switch (formswitch / ui.Switch).",
    [("default", "Default", "", '@ui.Switch(ui.SwitchProps{Name: "airplane"})', "previewDefault"),
     ("checked", "On", "", '@ui.Switch(ui.SwitchProps{Checked: true})', "previewChecked")],
    [("Name", "string", "Form field name"), ("Checked", "bool", "On state"), ("AriaLabel", "string", "Accessible name")],
    [("Checkbox", "/docs/components/checkbox"), ("Toggle", "/docs/components/toggle")],
    '''func previewDefault(ctx context.Context, w io.Writer) error { return ui.Switch(ui.SwitchProps{Name: "airplane", AriaLabel: "Airplane mode"}).Render(ctx, w) }
func previewChecked(ctx context.Context, w io.Writer) error { return ui.Switch(ui.SwitchProps{Name: "airplane", Checked: true, AriaLabel: "Airplane mode on"}).Render(ctx, w) }''')

page("slider", "Slider", "Native range input styled via ui.Input.",
    [("default", "Default", "", '@ui.Input(ui.InputProps{Type: "range", Min: "0", Max: "100"})', "previewDefault"),
     ("value", "With value", "", '@ui.Input(ui.InputProps{Type: "range", Value: "50"})', "previewValue")],
    [("Min", "string", "Minimum value"), ("Max", "string", "Maximum value"), ("Value", "string", "Current value")],
    [("Switch", "/docs/components/switch"), ("Input", "/docs/components/input")],
    '''func previewDefault(ctx context.Context, w io.Writer) error { return ui.Input(ui.InputProps{Type: "range", Min: "0", Max: "100", AriaLabel: "Volume"}).Render(ctx, w) }
func previewValue(ctx context.Context, w io.Writer) error { return ui.Input(ui.InputProps{Type: "range", Min: "0", Max: "100", Value: "50", AriaLabel: "Brightness"}).Render(ctx, w) }''')

exec(open(ROOT / "scripts" / "generate_showcases_c.py").read(), globals())
