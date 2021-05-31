package common

import (
	sh "2bodyBinary/draw_control/shapes"
	v "2bodyBinary/variables"
	"github.com/gonutz/prototype/draw"
)

func DrawControlAll(window draw.Window, shapes *[]sh.Shape) {
	if v.IsAllHide {
		return
	}
	for i := 0; i < len(*shapes); i++ {
		shape := (*shapes)[i]
		shape.CarryEvent(window)
		shape.Paint(window, 3)
	}
}
