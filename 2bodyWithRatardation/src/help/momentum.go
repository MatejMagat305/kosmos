package help

import "math"

type Momentum struct {
	X, Y float64
}

// 1/|r|
func (r *Momentum) ReverseValue() float64 {
	value := r.Value()
	return 1 / value
}

// |r|
func (r *Momentum) Value() float64 {
	x := math.Pow(r.X, 2)
	y := math.Pow(r.Y, 2)
	temp := x + y
	return math.Sqrt(temp)
}
