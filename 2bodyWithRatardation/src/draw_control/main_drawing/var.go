package main_drawing

import (
	"github.com/gonutz/prototype/draw"
	sh "retardation/draw_control/shapes"
	"sync"
)

type PlanetsPositions struct {
	bytes  [][][]byte
	floats [][][]float64
}

type Momentum struct {
	Vx,Vy, X, Y float64
}

var (
	shapes               = make([]sh.Shape, 0, 7)
	colors               = []draw.Color{draw.Red, draw.LightGreen, draw.LightBlue, draw.LightPurple}
	mux                  sync.Mutex
	planetsPositions     *PlanetsPositions
	selectMode, selected = false, make(map[int]bool)
	saveQuestion         = false
	chooseQuestion       = false
	ChStop = make(chan bool)
	initMomentum *Momentum
)

const (
	showAnalytic  = "show analytic"
	showCenterOfMomentum = "show center of momentum"
)

func IsInSelected(i int) bool {
	val, ok := selected[i]
	return ok && val
}

func IsSelectMode() bool {
	return selectMode
}

