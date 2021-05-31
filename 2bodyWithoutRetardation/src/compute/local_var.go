package compute

import (
	h "2bodyBinary/help"
	eu "2bodyBinary/structures/euler"
	v "2bodyBinary/variables"
	"fmt"
	"os"
	"path/filepath"
)

var (
	localE                         *eu.Euler
	Time                           = 0
	Planet2                        *eu.Planet
	Planet1                        *eu.Planet
	ChStart                        = make(chan bool)
	ChDone                         = make(chan int)
	chSave, chOut1, chOut2         = make(chan bool), make(chan bool), make(chan bool)
	vec                            *h.Vektor2D
	gravity                        float64
	r                              float64
	rPow3                          float64
	epsilon, x, y, massByGDivRPow3 float64
	period                         int
	f1, f2                         *os.File
	NumPeriod                      = 1
	PeriodProgres                  = 0
	dirName                        string
	PlanetData1                    *h.FloatData
	PlanetData2                    *h.FloatData
)

func GetOriginByLocalEuler() (*eu.Euler, error) {
	e, err := eu.GetEulerFromName(filepath.Join("./bin", localE.Name, v.EulerOriginConfig))
	if err != nil {
		return nil, fmt.Errorf("unreache bin")
	}
	return e, nil
}
