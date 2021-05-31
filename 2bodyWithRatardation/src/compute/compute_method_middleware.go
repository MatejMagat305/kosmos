package compute

import (
	"math"
	h "retardation/help"
	eu "retardation/structures/euler"
	v "retardation/variables"
	"runtime"
)

func doMoveBySelection(planet1, planet2 *eu.Planet)func( data2, data1 *h.FloatData) {
	switch localE.Method {
	case v.WithoutRetardation:
		return computeWithoutRetardation(planet1,planet2)
	default:
		return computeFirstMethod(planet1,planet2)
	}
}

func computeWithoutRetardation(planet1, planet2 *eu.Planet) func(data2, data1 *h.FloatData) {
	return func(data2, data1 *h.FloatData) {
		vec := h.NewVector2DFrom4(planet1.PositionX, planet1.PositionY, planet2.PositionX, planet2.PositionY)
		r := vec.Size()
		rPow3 := math.Pow(r, 3)
		massByGDivRPow3 := (gravity )/rPow3
		planet1.AccelerationX = -vec.GetX() * massByGDivRPow3 * planet2.Mass
		planet1.AccelerationY = -vec.GetY() * massByGDivRPow3 * planet2.Mass
		planet1.VelocityX += epsilon * planet1.AccelerationX
		planet1.VelocityY += epsilon * planet1.AccelerationY
		chOut1<-true
		runtime.Goexit()
	}
}
func computeFirstMethod(planet1, planet2 *eu.Planet)func( data, data1 *h.FloatData) {
	return func(data2, data1 *h.FloatData) {
		rdr := computeRDR(planet1, planet2, data2,  data1)
		massByGDivRPow3 := (gravity )/math.Pow( rdr.Size(), 3)
		planet1.AccelerationX = -rdr.GetX() * massByGDivRPow3 * planet2.Mass
		planet1.AccelerationY = -rdr.GetY() * massByGDivRPow3 * planet2.Mass
		planet1.VelocityX += epsilon * planet1.AccelerationX
		planet1.VelocityY += epsilon * planet1.AccelerationY
		chOut1<-true
		runtime.Goexit()
	}
}

func computeRDR(planet1, planet2 *eu.Planet, data1, data2 *h.FloatData) *h.Vektor2D {
	len2 := len(data2.VelocityXY)
	V := h.NewVector2DFrom2(data2.VelocityXY[len2-2], data2.VelocityXY[len2-1])
	R := h.NewVector2DFrom4(planet2.PositionX, planet2.PositionY, planet1.PositionX, planet1.PositionY )
	RxD := -R.GetX()- R.Size()*V.GetX()/localE.C
	RyD := -R.GetY() + R.Size()*V.GetY()/localE.C
	return h.NewVector2DFrom2(RxD, RyD)
}