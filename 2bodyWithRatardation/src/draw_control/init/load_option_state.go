package init

import (
	"retardation/compute"
	sh "retardation/draw_control/shapes"
	h "retardation/help"
	"retardation/structures/euler"
	"retardation/variables"
	v "retardation/variables"
	"fmt"
	"github.com/gonutz/prototype/draw"
	"os"
	"path/filepath"
)

var dirs []string

func prepareLoadOptionState(window draw.Window) {
	var (
		x, y float64 = 0, 0
	)
	loadDirs()
	shapes = sh.MakeEmptyShapesArrayWithCapResetId(len(dirs))
	addButtonFunkcion := addButton("Back", x, y)
	x, y = addButtonFunkcion(window, func(window draw.Window) {
		SetToStateInit()
		v.IsLoad = true
		euler.InitEuler(h.FindPath(v.Configs, v.Config))
	})
	if len(dirs)==0 {
		initLabelEmpty(window)
		return
	}
	for i := 0; i < len(dirs); i++ {
		name := dirs[i]
		x, y = addButton(name, x, y)(window, startLoad(name))
	}
}

func initLabelEmpty(windows draw.Window) {
	str := "I have not found any folder previous computations"
	x,y := windows.GetScaledTextSize(str,3)
	label := sh.NewLabel(str)
	label.SetXY(v.CenterX-float64(x)/2, v.CenterY-float64(y)/2)
	shapes = append(shapes, label)
}

func addButton(name string, x float64, y float64) func(window draw.Window, f func(window2 draw.Window)) (float64, float64) {
	return func(window draw.Window, f func(window2 draw.Window)) (float64, float64) {
		newButton := sh.NewButtonRedWhiteGreen(name)
		shapes = append(shapes, newButton)
		x2, y2 := window.GetScaledTextSize(name, scaledText)
		floatX2, floatY2 := float64(x2), float64(y2)
		floatX2 *= 1.07
		floatY2 *= 1.07
		if x+float64(x2) > variables.Width {
			x = 0
			y += floatY2
		}
		newButton.SetXY(x, y)
		newButton.SetSize(floatX2, floatY2)
		newButton.SetFunc(f)
		x += floatX2
		return x, y
	}
}

func loadDirs() {
	dirs = make([]string, 0, 10)
	node := "./bin/"
	if h.FileExist(node) {
		err := filepath.Walk(node, h.Visit(&dirs, func(info os.FileInfo) bool {
			dir := fmt.Sprint("./bin/",info.Name(),"/")
			return info.IsDir() &&
				h.FileExist(fmt.Sprint(dir,v.EulerConfig)) &&
				h.FileExist(fmt.Sprint(dir, v.EulerOriginConfig))
		}))
		if err == nil {
			return
		}
	}
	dirs = make([]string, 0, 0)
}

func startLoad(name string) func(window draw.Window) {
	return func(window draw.Window) {

		compute.SetDir(name)
		variables.IsLoad = true
		variables.IsInit = false
	}
}
