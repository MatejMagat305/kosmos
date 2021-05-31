package saved

import (
	sh "retardation/draw_control/shapes"
	eu "retardation/structures/euler"
	"github.com/gonutz/prototype/draw"
)

var (
	IsWarning = false
	shapes []sh.Shape
	continue0 = false
	textfield *sh.TextField
)


func save( draw.Window) {
	name := textfield.GetValue()
	done := eu.MEuler.TryToSaveIfNoChangeWarningReturnDone(name)
	if done {
		continue0 = false
	}else {
		IsWarning = true
	}
}

func cancel( draw.Window) {
	continue0 = false
}
