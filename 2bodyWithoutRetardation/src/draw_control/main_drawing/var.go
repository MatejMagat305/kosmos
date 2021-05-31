package main_drawing

import (
	sh "2bodyBinary/draw_control/shapes"
	"github.com/gonutz/prototype/draw"
	"sync"
)

type PlanetsPositions struct {
	bytes  [][][]byte
	floats [][][]float64
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

