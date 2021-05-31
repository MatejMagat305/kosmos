package variables

import "math"

var (
	Height = 500.0
	Width  = 500.0
	CenterY     = Height / 2
	CenterX     = Width / 2
	ScaleScreen = math.Min(Height, Width)
)

const (
	FirstElement  =  iota
	SecondElement
	StandardOrbit = 36525602
)
