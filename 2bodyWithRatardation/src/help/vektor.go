package help

import (
	"math"
)

type Vektor2D struct {
	x, y float64
}

func (v *Vektor2D) Size() float64 {
	return math.Sqrt(math.Pow(v.x, 2) + math.Pow(v.y, 2))
}

func (v *Vektor2D) GetX() float64 {
	return v.x
}

func (v *Vektor2D) GetY() float64 {
	return v.y
}

func NewVector2DFrom4(X1, Y1, X2, Y2 float64) *Vektor2D {
	return &Vektor2D{
		x: X1 - X2,
		y: Y1 - Y2,
	}
}

func NewVector2DFrom2(X1, Y1 float64) *Vektor2D {
	return &Vektor2D{
		x: X1,
		y: Y1,
	}
}

func (v *Vektor2D) ScalarMultiply(d *Vektor2D) float64 {
	return v.GetX()*d.GetX()+v.GetY()*d.GetY()
}

func (v *Vektor2D) IsPerpendicular(d *Vektor2D) bool {
	product := math.Abs(v.ScalarMultiply(d))
	return product==0
}

func (v Vektor2D) DivideByScalar(sc float64) Vektor2D {
	v.x/=sc
	v.y/=sc
	return v
}