# Part 3 — composites, wireframe patterns, app components, blocks

# Composites (templ/components)
page("alert", "Alert", "Callout for important messages.",
    [("default", "Default", "", '@cmp.Alert(cmp.AlertProps{}) { … }', "previewDefault"),
     ("destructive", "Destructive", "", '@cmp.Alert(cmp.AlertProps{Variant: "destructive"}) { … }', "previewDestructive")],
    [("Variant", "string", "default | destructive"), ("Class", "string", "Extra utilities")],
    [("Card", "/docs/components/card"), ("Badge", "/docs/components/badge")],
    '''func alertBody(ctx context.Context, w io.Writer, p cmp.AlertProps, title, body string) error {
\treturn showcaseutil.RenderWithChildren(ctx, w, cmp.Alert(p), func(ctx context.Context, w io.Writer) error {
\t\tif err := ui.Title(ui.TitleProps{Order: 4, Class: "text-sm font-semibold"}, title).Render(ctx, w); err != nil { return err }
\t\treturn ui.Text(ui.TextProps{Class: "text-sm"}, body).Render(ctx, w)
\t})
}
func previewDefault(ctx context.Context, w io.Writer) error { return alertBody(ctx, w, cmp.AlertProps{}, "Heads up", "You can add components from the gallery.") }
func previewDestructive(ctx context.Context, w io.Writer) error { return alertBody(ctx, w, cmp.AlertProps{Variant: "destructive"}, "Error", "Something went wrong.") }''',
    imports=CMP)

page("breadcrumb", "Breadcrumb", "Hierarchy navigation trail.",
    [("default", "Default", "", '@cmp.Breadcrumb(cmp.BreadcrumbProps{Items: items})', "previewDefault"),
     ("current", "Current page", "", "Last item with Current: true", "previewCurrent")],
    [("Items", "[]BreadcrumbItem", "Label, Href, Current, Disabled"), ("Class", "string", "Nav wrapper utilities")],
    [("Navigation Menu", "/docs/components/navigation-menu"), ("Tabs", "/docs/components/tabs")],
    '''var crumbItems = []cmp.BreadcrumbItem{{Label: "Home", Href: "/"}, {Label: "Docs", Href: "/docs"}, {Label: "Button", Current: true}}
func previewDefault(ctx context.Context, w io.Writer) error { return cmp.Breadcrumb(cmp.BreadcrumbProps{Items: crumbItems}).Render(ctx, w) }
func previewCurrent(ctx context.Context, w io.Writer) error {
\titems := []cmp.BreadcrumbItem{{Label: "Home", Href: "/"}, {Label: "Settings", Current: true}}
\treturn cmp.Breadcrumb(cmp.BreadcrumbProps{Items: items}).Render(ctx, w)
}''', imports=CMP)

page("card", "Card", "Grouped content with header, body, and footer.",
    [("default", "Default", "", '@cmp.Card(cmp.CardProps{}) { @cmp.CardHeader … }', "previewDefault"),
     ("footer", "With footer", "", '@cmp.CardFooter …', "previewFooter")],
    [("Variant", "string", "Surface variant"), ("CardHeader", "component", "Title area"), ("CardContent", "component", "Main body")],
    [("Alert", "/docs/components/alert"), ("Dialog", "/docs/components/dialog")],
    '''func previewDefault(ctx context.Context, w io.Writer) error { return cardPreview(ctx, w, false) }
func previewFooter(ctx context.Context, w io.Writer) error { return cardPreview(ctx, w, true) }
func cardPreview(ctx context.Context, w io.Writer, footer bool) error {
\treturn showcaseutil.RenderWithChildren(ctx, w, cmp.Card(cmp.CardProps{Class: "max-w-sm"}), func(ctx context.Context, w io.Writer) error {
\t\th := func(ctx context.Context, w io.Writer) error {
\t\t\treturn showcaseutil.RenderWithChildren(ctx, w, cmp.CardHeader(cmp.CardHeaderProps{}), func(ctx context.Context, w io.Writer) error {
\t\t\t\tif err := cmp.CardTitle(cmp.CardTitleProps{}, "Card title").Render(ctx, w); err != nil { return err }
\t\t\t\treturn cmp.CardDescription(cmp.CardDescriptionProps{}, "Wireframe card description.").Render(ctx, w)
\t\t\t})
\t\t}
\t\tc := func(ctx context.Context, w io.Writer) error {
\t\t\treturn showcaseutil.RenderWithChildren(ctx, w, cmp.CardContent(cmp.CardContentProps{}), func(ctx context.Context, w io.Writer) error {
\t\t\t\treturn ui.Text(ui.TextProps{}, "Card body copy.").Render(ctx, w)
\t\t\t})
\t\t}
\t\tif err := h(ctx, w); err != nil { return err }
\t\tif err := c(ctx, w); err != nil { return err }
\t\tif footer {
\t\t\treturn showcaseutil.RenderWithChildren(ctx, w, cmp.CardFooter(cmp.CardFooterProps{}), func(ctx context.Context, w io.Writer) error {
\t\t\t\treturn ui.Button(ui.ButtonProps{Size: "sm"}, "Action").Render(ctx, w)
\t\t\t})
\t\t}
\t\treturn nil
\t})
}''', imports=CMP)

# Wireframe helpers
WF_STACK = '''func wfStack(ctx context.Context, w io.Writer, attrs templ.Attributes, body func(context.Context, io.Writer) error) error {
\tprops := ui.StackProps{Class: "gap-2 max-w-md"}
\tif attrs != nil { props.Attrs = attrs }
\treturn showcaseutil.RenderWithChildren(ctx, w, ui.Stack(props), body)
}'''

def wf_page(slug, title, desc, code, fn_body, api, related, variants_extra=None, ui8kit=None):
    attrs = (A % ui8kit) if ui8kit else "nil"
    body = WF_STACK + "\n" + fn_body.replace("ATTRS", attrs)
    v = [("default", "Default", "Wireframe composition from ui primitives.", code, "previewDefault")]
    if variants_extra:
        v.extend(variants_extra)
    page(slug, title, desc, v, api, related, body)

wf_page("accordion", "Accordion", "Vertically stacked expandable sections (wireframe; data-ui8kit accordion).",
    '@ui.Stack(' + A % "accordion" + ') { @ui.Button[data-ui8kit-trigger] … }',
    '''func previewDefault(ctx context.Context, w io.Writer) error {
\treturn wfStack(ctx, w, ATTRS, func(ctx context.Context, w io.Writer) error {
\t\tif err := ui.Button(ui.ButtonProps{Attrs: templ.Attributes{"data-ui8kit-trigger": "item-1", "aria-expanded": "false"}}, "Section 1").Render(ctx, w); err != nil { return err }
\t\treturn ui.Box(ui.BoxProps{Class: "rounded border border-border p-3 text-sm", Attrs: templ.Attributes{"data-ui8kit-panel": "item-1", "hidden": true}}, ).Render(templ.WithChildren(ctx, showcaseutil.Child(func(ctx context.Context, w io.Writer) error {
\t\t\treturn ui.Text(ui.TextProps{}, "Panel content.").Render(ctx, w)
\t\t})), w)
\t})
}''',
    [("Attrs", "templ.Attributes", "data-ui8kit on root"), ("Trigger", "Button", "data-ui8kit-trigger"), ("Panel", "Box", "data-ui8kit-panel")],
    [("Collapsible", "/docs/components/collapsible"), ("Tabs", "/docs/components/tabs")], ui8kit="accordion")

wf_page("tabs", "Tabs", "Tabbed interface (wireframe; data-ui8kit tabs).",
    '@ui.Stack { tablist + panels }',
    '''func previewDefault(ctx context.Context, w io.Writer) error {
\treturn wfStack(ctx, w, ATTRS, func(ctx context.Context, w io.Writer) error {
\t\tif err := ui.Group(ui.GroupProps{Class: "flex gap-1", Attrs: templ.Attributes{"role": "tablist"}}).Render(templ.WithChildren(ctx, showcaseutil.Child(func(ctx context.Context, w io.Writer) error {
\t\t\tif err := ui.Button(ui.ButtonProps{Variant: "secondary", Size: "sm", Attrs: templ.Attributes{"data-ui8kit-tab": "a", "aria-selected": "true"}}, "Tab A").Render(ctx, w); err != nil { return err }
\t\t\treturn ui.Button(ui.ButtonProps{Variant: "ghost", Size: "sm", Attrs: templ.Attributes{"data-ui8kit-tab": "b"}}, "Tab B").Render(ctx, w)
\t\t})), w); err != nil { return err }
\t\treturn ui.Box(ui.BoxProps{Class: "rounded border border-border p-3 text-sm", Attrs: templ.Attributes{"data-ui8kit-panel": "a"}}, ).Render(templ.WithChildren(ctx, showcaseutil.Child(func(ctx context.Context, w io.Writer) error {
\t\t\treturn ui.Text(ui.TextProps{}, "Tab A panel.").Render(ctx, w)
\t\t})), w)
\t})
}''',
    [("Attrs", "templ.Attributes", "data-ui8kit=tabs on root"), ("Tab", "Button", "data-ui8kit-tab"), ("Panel", "Box", "data-ui8kit-panel")],
    [("Accordion", "/docs/components/accordion"), ("Navigation Menu", "/docs/components/navigation-menu")], ui8kit="tabs")

wf_page("dialog", "Dialog", "Modal dialog wireframe (data-ui8kit dialog).",
    '@ui.Box[data-ui8kit=dialog] { title + actions }',
    '''func previewDefault(ctx context.Context, w io.Writer) error {
\treturn showcaseutil.RenderWithChildren(ctx, w, ui.Box(ui.BoxProps{Class: "max-w-sm rounded-lg border border-border bg-card p-4 shadow-lg", Attrs: ATTRS}), func(ctx context.Context, w io.Writer) error {
\t\tif err := ui.Title(ui.TitleProps{Order: 3, Class: "text-base font-semibold", Attrs: templ.Attributes{"id": "dialog-title"}}, "Dialog title").Render(ctx, w); err != nil { return err }
\t\tif err := ui.Text(ui.TextProps{Class: "text-sm text-muted-foreground"}, "Dialog description copy.").Render(ctx, w); err != nil { return err }
\t\treturn ui.Group(ui.GroupProps{Class: "mt-4 flex justify-end gap-2"}).Render(templ.WithChildren(ctx, showcaseutil.Child(func(ctx context.Context, w io.Writer) error {
\t\t\tif err := ui.Button(ui.ButtonProps{Variant: "outline", Attrs: templ.Attributes{"data-ui8kit-close": ""}}, "Cancel").Render(ctx, w); err != nil { return err }
\t\t\treturn ui.Button(ui.ButtonProps{Attrs: templ.Attributes{"data-ui8kit-close": ""}}, "Confirm").Render(ctx, w)
\t\t})), w)
\t})
}'''.replace("ATTRS", A % "dialog"),
    [("Attrs", "templ.Attributes", "data-ui8kit=dialog"), ("Title", "Title", "aria-labelledby target")],
    [("Alert Dialog", "/docs/components/alert-dialog"), ("Sheet", "/docs/components/sheet")])

exec(open(ROOT / "scripts" / "generate_showcases_d.py").read(), globals())
