package docsstatic

import (
	"context"
	"io"
)

const (
	IllusSectionPrimitives  = "primitives"
	IllusSectionComponents = "components"
)

type indexIllustrationSpec struct {
	label   string
	section string
	suffix  string
	render  func(context.Context, io.Writer) error
}

func illusBars(count, accentIndex int) func(context.Context, io.Writer) error {
	return func(ctx context.Context, w io.Writer) error {
		return renderStackBars(ctx, w, count, accentIndex)
	}
}

func illusGrid(cols, rows int) func(context.Context, io.Writer) error {
	return func(ctx context.Context, w io.Writer) error {
		return renderCellGrid(ctx, w, cols, rows)
	}
}

func illusFieldStack(fields int) func(context.Context, io.Writer) error {
	return func(ctx context.Context, w io.Writer) error {
		return renderFieldStack(ctx, w, fields)
	}
}

var indexIllustrationSpecs = []indexIllustrationSpec{
	// Primitives
	{label: "Badge", section: IllusSectionPrimitives, suffix: "/primitives/badge/", render: renderMarkLine},
	{label: "Block", section: IllusSectionPrimitives, suffix: "/primitives/block/", render: renderBlockCaption},
	{label: "Box", section: IllusSectionPrimitives, suffix: "/primitives/box/", render: renderBlockCaption},
	{label: "Button", section: IllusSectionPrimitives, suffix: "/primitives/button/", render: renderSingleAccent},
	{label: "Checkbox", section: IllusSectionPrimitives, suffix: "/primitives/checkbox/", render: renderToggleControl},
	{label: "Container", section: IllusSectionPrimitives, suffix: "/primitives/container/", render: renderBlockCaption},
	{label: "Controls", section: IllusSectionPrimitives, suffix: "/primitives/controls/", render: illusFieldStack(2)},
	{label: "Dialog", section: IllusSectionPrimitives, suffix: "/primitives/dialog/", render: renderPanel},
	{label: "Disclosure", section: IllusSectionPrimitives, suffix: "/primitives/disclosure/", render: renderMarkLine},
	{label: "Grid", section: IllusSectionPrimitives, suffix: "/primitives/grid/", render: illusGrid(3, 3)},
	{label: "Group", section: IllusSectionPrimitives, suffix: "/primitives/group/", render: illusFieldStack(2)},
	{label: "Icon", section: IllusSectionPrimitives, suffix: "/primitives/icon/", render: renderIconMark},
	{label: "Image", section: IllusSectionPrimitives, suffix: "/primitives/image/", render: renderBlockCaption},
	{label: "Input", section: IllusSectionPrimitives, suffix: "/primitives/input/", render: renderFieldLines},
	{label: "Label", section: IllusSectionPrimitives, suffix: "/primitives/label/", render: illusBars(1, -1)},
	{label: "Linebreak", section: IllusSectionPrimitives, suffix: "/primitives/linebreak/", render: renderSeparatorLine},
	{label: "Link", section: IllusSectionPrimitives, suffix: "/primitives/link/", render: renderLinkPill},
	{label: "List", section: IllusSectionPrimitives, suffix: "/primitives/list/", render: illusBars(3, -1)},
	{label: "Radio", section: IllusSectionPrimitives, suffix: "/primitives/radio/", render: renderToggleControl},
	{label: "Select", section: IllusSectionPrimitives, suffix: "/primitives/select/", render: renderFieldLines},
	{label: "Separator", section: IllusSectionPrimitives, suffix: "/primitives/separator/", render: renderSeparatorLine},
	{label: "Stack", section: IllusSectionPrimitives, suffix: "/primitives/stack/", render: illusBars(2, -1)},
	{label: "Switch", section: IllusSectionPrimitives, suffix: "/primitives/switch/", render: renderSwitchControl},
	{label: "Text", section: IllusSectionPrimitives, suffix: "/primitives/text/", render: illusBars(2, -1)},
	{label: "Textarea", section: IllusSectionPrimitives, suffix: "/primitives/textarea/", render: illusFieldStack(3)},
	{label: "Title", section: IllusSectionPrimitives, suffix: "/primitives/title/", render: illusBars(1, -1)},

	// Components
	{label: "Accordion", section: IllusSectionComponents, suffix: "/components/accordion/", render: illusBars(2, -1)},
	{label: "Alert", section: IllusSectionComponents, suffix: "/components/alert/", render: renderMarkLine},
	{label: "Alert Dialog", section: IllusSectionComponents, suffix: "/components/alert-dialog/", render: renderPanel},
	{label: "Aspect Ratio", section: IllusSectionComponents, suffix: "/components/aspect-ratio/", render: renderBlockCaption},
	{label: "Avatar", section: IllusSectionComponents, suffix: "/components/avatar/", render: renderIconMark},
	{label: "Blog Card", section: IllusSectionComponents, suffix: "/components/blog-card/", render: renderBlockCaption},
	{label: "Breadcrumb", section: IllusSectionComponents, suffix: "/components/breadcrumb/", render: renderMenuList},
	{label: "Card", section: IllusSectionComponents, suffix: "/components/card/", render: renderBlockCaption},
	{label: "Carousel", section: IllusSectionComponents, suffix: "/components/carousel/", render: renderBlockCaption},
	{label: "Dialog", section: IllusSectionComponents, suffix: "/components/dialog/", render: renderPanel},
	{label: "Drawer", section: IllusSectionComponents, suffix: "/components/drawer/", render: renderPanel},
	{label: "Dropdown Menu", section: IllusSectionComponents, suffix: "/components/dropdown-menu/", render: renderMenuList},
	{label: "Form", section: IllusSectionComponents, suffix: "/components/form/", render: illusBars(4, 3)},
	{label: "Icon", section: IllusSectionComponents, suffix: "/components/icon/", render: renderIconMark},
	{label: "Icon Badge", section: IllusSectionComponents, suffix: "/components/iconbadge/", render: renderMarkLine},
	{label: "Language Toggle", section: IllusSectionComponents, suffix: "/components/language-toggle/", render: renderTabButtons},
	{label: "Menubar", section: IllusSectionComponents, suffix: "/components/menubar/", render: renderMenuList},
	{label: "Nav", section: IllusSectionComponents, suffix: "/components/nav/", render: renderMenuList},
	{label: "Navigation Menu", section: IllusSectionComponents, suffix: "/components/navigation-menu/", render: renderMenuList},
	{label: "Pagination", section: IllusSectionComponents, suffix: "/components/pagination/", render: renderPaginationButtons},
	{label: "Progress", section: IllusSectionComponents, suffix: "/components/progress/", render: renderProgressLine},
	{label: "Separator", section: IllusSectionComponents, suffix: "/components/separator/", render: renderSeparatorLine},
	{label: "Sheet", section: IllusSectionComponents, suffix: "/components/sheet/", render: renderPanel},
	{label: "Slider", section: IllusSectionComponents, suffix: "/components/slider/", render: renderProgressLine},
	{label: "Table", section: IllusSectionComponents, suffix: "/components/table/", render: illusGrid(3, 3)},
	{label: "Tabs", section: IllusSectionComponents, suffix: "/components/tabs/", render: renderTabButtons},
}

var indexIllustrationBySuffix map[string]indexIllustration

func init() {
	indexIllustrationBySuffix = make(map[string]indexIllustration, len(indexIllustrationSpecs))
	for _, spec := range indexIllustrationSpecs {
		indexIllustrationBySuffix[spec.suffix] = indexIllustration{render: spec.render}
	}
}
