package init

import (
	v "retardation/variables"
	"github.com/gonutz/prototype/draw"
)

func makeInitState(window draw.Window) {
	prepareFirstOptionState(window)
	state = firstOptionState
	v.IsWarned = false
}
