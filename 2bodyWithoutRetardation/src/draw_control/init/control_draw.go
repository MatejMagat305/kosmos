package init

import (
	common "2bodyBinary/draw_control/common"
	sh "2bodyBinary/draw_control/shapes"
	v "2bodyBinary/variables"
	"github.com/gonutz/prototype/draw"
)

const (
	initState = iota
	firstOptionState
	loadOptionState
	formState
	scaledText float32 = 3
)

var (
	state  = initState
	shapes []sh.Shape
)

func DrawControlState(window draw.Window) {
	common.Mux.Lock()
	switch state {
	case initState:
		makeInitState(window)
		v.IsInitEulers=true
		fallthrough
	case formState, loadOptionState, firstOptionState:
		drawControlAll(window)
	default:
		state = firstOptionState
	}
	common.Mux.Unlock()
}

func drawControlAll(window draw.Window) {
	common.DrawControlAll(window, &shapes)
	if state == formState {
		controlEnter(window)
	}
}

func SetToStateInit() {
	setState(initState)
}