package common

import (
	sh "retardation/draw_control/shapes"
	v "retardation/variables"
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
