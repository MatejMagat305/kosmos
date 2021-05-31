package choose

import (
	drawCommon "retardation/draw_control/common"
	v "retardation/variables"
	"github.com/gonutz/prototype/draw"
)

func DrawControlContinue(window draw.Window) bool {
	window.FillRect(0, 0, int(v.Width), int(v.Height), draw.Black)
	drawCommon.DrawControlAll(window, &shapes)
	if IsWarning {
		drawWarning(window)
	}
	if !continue0 {
		clear()
	}
	return continue0
}

func drawWarning(window draw.Window) {
	x,y := window.GetScaledTextSize(v.Warning,3)
	window.DrawScaledText(v.Warning, int(v.CenterX)-x/2, int(v.Height)-2*y, 2,draw.LightRed)
}

func clear() {
	shapes=nil
	numFrom=nil
	numTo=nil
}