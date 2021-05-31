package windows

import (
	"github.com/gonutz/prototype/draw"
	"math"
	v "retardation/variables"
)

func FindSizeScreen() {
	_ = draw.RunWindow("", 0, 0,
		func(window draw.Window) {
			window.SetFullscreen(true)
			width0, height0 := window.Size()
			v.Width = float64(width0)*0.99
			v.Height = float64(height0)*0.97
			window.Close()
		})
	v.CenterX = v.Width / 2
	v.CenterY = v.Height / 2
	v.ScaleScreen = math.Min(v.Height, v.Width)
}
