package compute

import (
	h "2bodyBinary/help"
	e "2bodyBinary/structures/euler"
	v "2bodyBinary/variables"
	"math"
)


func compute() {
	for ; ; NumPeriod++ {
		if <-ChStart {
			go func() {<-chOut1}()
			createFile(NumPeriod)
			send()
		} else {goto end}
		for PeriodProgres = 0; PeriodProgres < period; PeriodProgres++ {
			Time++
			upgrade()
		}
		saveBin()
		go sendNumPeriod(NumPeriod)

	}
end:
}

func GetPeriod() int {
	if localE.State != e.Elliptic && localE.State != e.Circle {
		return v.StandardOrbit
	}
	epsilon = localE.Epsilon
	length := localE.Periodlenght
	eInverse := 1 / epsilon
	floatPeriod := length * eInverse
	return int(floatPeriod)
}


func sendNumPeriod(period int) {
	ChDone <- period
}

func upgrade() {
	go func() {
		Planet1.PositionX += Planet1.VelocityX * epsilon
		Planet1.PositionY += Planet1.VelocityY * epsilon
		chOut1<-true
	}()
	go func() {
		Planet2.PositionX += Planet2.VelocityX * epsilon
		Planet2.PositionY += Planet2.VelocityY * epsilon
		chOut1<-true
	}()
	<-chOut1; <-chOut1
	go write(Planet1)
	go write(Planet2)
	vec = h.NewVector2DFrom4(Planet1.PositionX, Planet1.PositionY, Planet2.PositionX, Planet2.PositionY)
	r = vec.Size()
	rPow3 = math.Pow(r, 3)
	massByGDivRPow3 = (gravity )/rPow3
	x = vec.GetX() * massByGDivRPow3
	y = vec.GetY() * massByGDivRPow3
	go func() {
		Planet1.AccelerationX = -x * Planet2.Mass
		Planet1.AccelerationY = -y * Planet2.Mass
		Planet1.VelocityX += epsilon * Planet1.AccelerationX
		Planet1.VelocityY += epsilon * Planet1.AccelerationY
		chOut1<-true
	}()
	go func() {
		Planet2.AccelerationX = x * Planet1.Mass
		Planet2.AccelerationY = y * Planet1.Mass
		Planet2.VelocityX += epsilon * Planet2.AccelerationX
		Planet2.VelocityY += epsilon * Planet2.AccelerationY
		chOut1<-true
	}()
	<-chOut1;<-chOut1; <-chOut2; <-chOut2
}