package main

import "math"

type vektor2D struct {
	x,y float64

}

// v/||v|| == u/||u||
func (v *vektor2D) isSameDirection(u *vektor2D) bool {
	sizeV := v.size()
	sizeU := u.size()
	x , y := math.Abs(v.x/sizeV-u.x/sizeU), math.Abs(v.y/sizeV-u.y/sizeU)
	if x<=akceptSize*100 {
		if y<=akceptSize*100 {
			return true
		}
	}
	return false
}

func (v *vektor2D) size() float64 {
	return math.Sqrt(math.Pow(v.x,2)+math.Pow(v.y,2))
}

func (v *vektor2D) GetPerpendicularSizeOf(size float64) vektor2D {
	midleResult := NewVector2D(-v.y, v.x)
	n := midleResult.size()
	midleResult.x/=n
	midleResult.y/=n
	midleResult.x*=size
	midleResult.y*=size
	return vektor2D{midleResult.x,midleResult.y}
}

func NewVector2D(x float64, y float64) *vektor2D {
	return &vektor2D{
		x: x,
		y: y,
	}
}
