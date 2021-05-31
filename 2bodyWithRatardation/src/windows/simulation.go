package windows

import (
	"fmt"
	"github.com/gonutz/prototype/draw"
	"retardation/compute"
	initDrawControl "retardation/draw_control/init"
	mainDrawControl "retardation/draw_control/main_drawing"
	eu "retardation/structures/euler"
	v "retardation/variables"
)

func Simulation() {
	FindSizeScreen()
	runWindow()
}

func runWindow() {
	_ = draw.RunWindow("2D", int(v.Width), int(v.Height), func(window draw.Window) {
		if v.IsInit {
			window.FillRect(0, 0, int(v.Width), int(v.Height), draw.Black)
			initWindow(window)
			return
		}
		mainWindow(window)
	})
}

func mainWindow(window draw.Window) {
	mainDrawControl.DrawControlState(window)
}

func initWindow(window draw.Window) {
	initDrawControl.DrawControlState(window)
	controlPostInit(window)
}

func controlPostInit(window draw.Window) {
	if v.IsInit == false {
		prepareMain(window)
	}
}

func prepareMain(window draw.Window) {
	if v.IsLoad {
		compute.LoadOrigin()
	}else {
		compute.SaveOrigin(fmt.Sprint("./bin/",eu.MEuler.Name,"/", v.EulerOriginConfig), eu.MEuler)
	}
	eu.MEuler.MakeChangesCalculateShapePeriodLenght()
	mainDrawControl.InitVar(eu.MEuler, window)
	compute.InitLocalVar(eu.MEuler)
	go compute.Start()
	go mainDrawControl.LoadDone()
}
