package compute

import (
	h "retardation/help"
	eu "retardation/structures/euler"
)

func Start() {
	NumPeriod = 1
	loadIfIsloadExist()
	compute()
}

func InitLocalVar(e0 *eu.Euler) {
	inetEuler(e0)
	period = GetPeriod()
	epsilon = localE.Epsilon
	PeriodProgres = 0
	speedLight   = localE.C
	gravity      = localE.G
	update1 = doMoveBySelection(Planet1, Planet2)
	update2 = doMoveBySelection(Planet2, Planet1)
	PlanetData1 = &h.FloatData{
		PositionsXY: make([]float64, 0,  period),
		VelocityXY:  make([]float64, 0, const0),
	}
	PlanetData2 = &h.FloatData{
		PositionsXY: make([]float64, 0, period),
		VelocityXY:  make([]float64, 0, const0),
	}
	initChans()
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