package init

import (
	drawCommon "retardation/draw_control/common"
	sh "retardation/draw_control/shapes"
	"github.com/gonutz/prototype/draw"
)

func prepareFirstOptionState(window draw.Window) {
	shapes = sh.MakeEmptyShapesArrayWithCapResetId(2)

	load, startNew := "load old calculation", "start new calculation"

	loadPrepareButton := drawCommon.AddButtonFirstOption(load, window, &shapes)
	loadPrepareButton(-1, scaledText, changeStateFunc(loadOptionState))
	newPrepareButton := drawCommon.AddButtonFirstOption(startNew, window, &shapes)
	newPrepareButton(1, scaledText, changeStateFunc(formState))
}
