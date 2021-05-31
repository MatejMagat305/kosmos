package compute

import (
	"math"
	h "retardation/help"
	e "retardation/structures/euler"
)

func initLeapfrog() {
	vec := h.NewVector2DFrom4(Planet1.PositionX, Planet1.PositionY, Planet2.PositionX, Planet2.PositionY)
	r := vec.Size()
	rPow3 := math.Pow(r, 3)
	x := vec.GetX() * localE.G
	y := vec.GetY() * localE.G
	epsilon = localE.Epsilon
	initLastVelocity()
	go func() {
		Planet1.AccelerationX = -((x * Planet2.Mass) / rPow3)
		Planet1.AccelerationY = -((y * Planet2.Mass) / rPow3)
		Planet1.VelocityX += (epsilon /2) * Planet1.AccelerationX
		Planet1.VelocityY += (epsilon /2) * Planet1.AccelerationY
		Planet1.PositionX += Planet1.VelocityX * (epsilon /2)
		Planet1.PositionY += Planet1.VelocityY * (epsilon /2)
		chOut1 <- true
	}()
	go func() {
		Planet2.AccelerationX = (x * Planet1.Mass) / rPow3
		Planet2.AccelerationY = (y * Planet1.Mass) / rPow3
		Planet2.VelocityX += (epsilon /2) * Planet2.AccelerationX
		Planet2.VelocityY += (epsilon /2) * Planet2.AccelerationY
		Planet2.PositionX += Planet2.VelocityX * (epsilon /2)
		Planet2.PositionY += Planet2.VelocityY * (epsilon /2)
		chOut1 <- true
	}()
	<-chOut1
	<-chOut1
}

func initLastVelocity() {
	for j, planet := range []*e.Planet{Planet1, Planet2} {
		data := PlanetData2
		if j==0 {
			data=PlanetData1
		}
		for i := 0; i < const0; i++ {
			data.VelocityXY = append(data.VelocityXY, planet.VelocityX)
			data.VelocityXY = append(data.VelocityXY, planet.VelocityY)
		}
	}
}
