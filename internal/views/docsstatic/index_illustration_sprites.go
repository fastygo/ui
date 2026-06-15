package docsstatic

import (
	"fmt"
	"strings"
)

const (
	illusSpriteCellW   = 144
	illusSpriteCellH   = 80
	illusSpriteSheetW  = 864
	illusSpriteSheetH  = 400
	illusSpriteSheetURL = "/static/img/docs-illus"
)

type indexIllustrationSprite struct {
	section string
	index   int
}

func indexIllustrationSpriteMeta(href string) (indexIllustrationSprite, bool) {
	for i, spec := range indexIllustrationSpecs {
		if !strings.HasSuffix(href, spec.suffix) {
			continue
		}
		idx := 0
		for j := 0; j < i; j++ {
			if indexIllustrationSpecs[j].section == spec.section {
				idx++
			}
		}
		return indexIllustrationSprite{section: spec.section, index: idx}, true
	}
	return indexIllustrationSprite{}, false
}

func (s indexIllustrationSprite) sheetClass() string {
	return "docs-index-card-illus--" + s.section
}

func (s indexIllustrationSprite) backgroundPositionStyle() string {
	col := s.index % 6
	row := s.index / 6
	return fmt.Sprintf(
		"background-position:%dpx %dpx",
		-col*illusSpriteCellW,
		-row*illusSpriteCellH,
	)
}

func (s indexIllustrationSprite) sheetURL() string {
	return fmt.Sprintf("%s/%s.png", illusSpriteSheetURL, s.section)
}
