package compute

import (
	e "retardation/structures/euler"
	v "retardation/variables"
	"runtime"
)


func compute() {
	send()
	for ; ; NumPeriod++ {
		if <-ChStart {
			createFile(NumPeriod)
		} else {
			goto end
		}
		for PeriodProgres = 0; PeriodProgres < period; PeriodProgres++ {
			Time++
			upgrade()
		}
		saveBin()
		localE.ResizePeriod()
		period = GetPeriod()
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
	runtime.Goexit()
}

func upgrade() {
	go func() {
		Planet1.PositionX += Planet1.VelocityX * epsilon
		Planet1.PositionY += Planet1.VelocityY * epsilon
		chOut1<-true
		runtime.Goexit()
	}()
	go func() {
		Planet2.PositionX += Planet2.VelocityX * epsilon
		Planet2.PositionY += Planet2.VelocityY * epsilon
		chOut1<-true
		runtime.Goexit()
	}()
	<-chOut1; <-chOut1
	go write(Planet1)
	go write(Planet2)
	<-chOut2; <-chOut2
	go update1(PlanetData2, PlanetData1)
	go update2(PlanetData1, PlanetData2)
	<-chOut1;<-chOut1
}