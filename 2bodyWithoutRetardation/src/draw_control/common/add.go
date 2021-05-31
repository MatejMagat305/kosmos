package common

import (
	sh "2bodyBinary/draw_control/shapes"
	"2bodyBinary/variables"
	"github.com/gonutz/prototype/draw"
)

func AddButtonPrepare(x, y float64, shapes *[]sh.Shape) func(window draw.Window, f func(window2 draw.Window), name string) {
	return func(window draw.Window, f func(window2 draw.Window), name string) {
		buttonNext := sh.NewButtonRedWhiteGreen(name)
		w, h := window.GetScaledTextSize(name, 3)
		buttonNext.SetXY(x-float64(w)/2, y-float64(h)/2)
		buttonNext.SetFunc(f)
		buttonNext.SetSize(float64(w), float64(h))
		*shapes = append(*shapes, buttonNext)
	}
}

func AddButtonFirstOption(name string, window draw.Window, shape *[]sh.Shape) func(negativ float64, scale float32, f func(window2 draw.Window)) {
	return func(negativ float64, scale float32, f func(window2 draw.Window)) {
		width, _ := window.GetScaledTextSize(name, scale)
		x, y := variables.CenterX+negativ*float64(width)/2+negativ*2, variables.CenterY
		addFuntionNext := AddButtonPrepare(x, y, shape)
		addFuntionNext(window, f, name)
	}
}