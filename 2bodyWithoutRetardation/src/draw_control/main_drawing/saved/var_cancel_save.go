package saved

import (
	sh "2bodyBinary/draw_control/shapes"
	eu "2bodyBinary/structures/euler"
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
