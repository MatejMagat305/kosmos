package saved

import (
	drawCommon "2bodyBinary/draw_control/common"
	sh "2bodyBinary/draw_control/shapes"
	"2bodyBinary/structures/euler"
	v "2bodyBinary/variables"
	"github.com/gonutz/prototype/draw"
)

func Begin(window draw.Window) {
	initVar()
	InitSaveCancel(window)
	initFormName()
}

func initFormName() {
	textfield = sh.NewTextField("Name                  ")
	size := 300.0
	textfield.SetXY(v.CenterX-size/2, 10)
	textfield.SetSize(size, 60)
	textfield.SetValue(euler.MEuler.Name)
	textfield.SetFunc(func(draw.Window) {
		IsWarning = false
		v.Warning=""
	})

	shapes = append(shapes, textfield)
}

func InitSaveCancel(window draw.Window) {
	load, startNew := "cancel", "save"
	loadPrepareButton := drawCommon.AddButtonFirstOption(load, window, &shapes)
	loadPrepareButton(-1, 3, cancel)
	newPrepareButton := drawCommon.AddButtonFirstOption(startNew, window, &shapes)
	newPrepareButton(1, 3, save)
}

func initVar() {
	continue0, IsWarning = true, false
	shapes = sh.MakeEmptyShapesArrayWithCapResetId(3)
}
