package compute

import (
	"os"
	h "retardation/help"
	eu "retardation/structures/euler"
)

var (
	localE                 *eu.Euler
	Time                   = 0
	Planet2                *eu.Planet
	Planet1                *eu.Planet
	ChStart                = make(chan bool)
	ChDone                 = make(chan int)
	chSave, chOut1, chOut2 = make(chan bool), make(chan bool), make(chan bool)
	epsilon                float64
	speedLight             float64
	gravity                float64
	period                 int
	f1, f2                 *os.File
	NumPeriod              = 1
	PeriodProgres          = 0
	dirName                string
	PlanetData2            *h.FloatData
	PlanetData1            *h.FloatData
	update1, update2 func( *h.FloatData,  *h.FloatData)
)

const (
	neddVelocity = "%v/actual_need_velocity_%d_planet.bin"
	const0       = 2
)