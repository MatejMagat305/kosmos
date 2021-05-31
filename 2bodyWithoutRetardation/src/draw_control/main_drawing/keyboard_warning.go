package main_drawing

import (
	eu "2bodyBinary/structures/euler"
	v "2bodyBinary/variables"
	"github.com/gonutz/prototype/draw"
)

func controlKeyboard(window draw.Window) {
	if window.WasKeyPressed(draw.KeyUp) {
		eu.MEuler.MoveY+=v.CenterY/5
	}
	if window.WasKeyPressed(draw.KeyDown) {
		eu.MEuler.MoveY-=v.CenterY/5
	}
	if window.WasKeyPressed(draw.KeyLeft) {
		eu.MEuler.MoveX+=v.CenterX/5
	}
	if window.WasKeyPressed(draw.KeyRight) {
		eu.MEuler.MoveX-=v.CenterX/5
	}
	if window.WasKeyPressed(draw.KeySpace) {
		eu.MEuler.MoveX = 0
		eu.MEuler.MoveY = 0
	}
	if window.WasKeyPressed(draw.KeyH) {
		v.IsAllHide = !v.IsAllHide
	}
}
