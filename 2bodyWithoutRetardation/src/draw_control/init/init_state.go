package init

import (
	v "2bodyBinary/variables"
	"github.com/gonutz/prototype/draw"
)

func makeInitState(window draw.Window) {
	prepareFirstOptionState(window)
	state = firstOptionState
	v.IsWarned = false
}
