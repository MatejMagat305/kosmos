package init

import (
	"retardation/draw_control/common"
	"github.com/gonutz/prototype/draw"
)

func changeStateFunc(st int) func(window draw.Window) {
	return func(window draw.Window) {
		go func() {
			common.Mux.Lock()
			prepareState(st, window)
			setState(st)
			common.Mux.Unlock()
		}()
	}
}

func setState(st int) {
	state = st
}

func prepareState(st int, window draw.Window) {
	switch st {
	case firstOptionState:
		prepareFirstOptionState(window)
	case loadOptionState:
		prepareLoadOptionState(window)
	case formState:
		prepareFormState(window)
	default:
		return
	}
}
