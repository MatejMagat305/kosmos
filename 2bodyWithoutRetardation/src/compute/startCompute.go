package compute

import (
	h "2bodyBinary/help"
	eu "2bodyBinary/structures/euler"
)

func Start() {
	NumPeriod = 1
	gravity = localE.G
	epsilon = localE.Epsilon
	period = GetPeriod()
	PeriodProgres = 0
	loadIfIsloadExist()
	compute()
}

func InitLocalVar(e0 *eu.Euler) {
	PlanetData1 = &h.FloatData{
		PositionsXY: make([]float64, 0, 2500),
	}
	PlanetData2 = &h.FloatData{
		PositionsXY: make([]float64, 0, 2500),
	}
	initChans()
	inetEuler(e0)
	initLeapfrog()
}

func initChans() {
	ChStart = make(chan bool)
	ChDone = make(chan int)
	chOut1, chOut2 = make(chan bool), make(chan bool)
}

func inetEuler(e0 *eu.Euler) {
	localE = e0
	Planet1 = e0.Planets[0]
	Planet2 = e0.Planets[1]
}
