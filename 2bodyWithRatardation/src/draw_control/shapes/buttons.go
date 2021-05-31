package shapes

import "github.com/gonutz/prototype/draw"

type Button struct {
	Basic
	aktiv, inaktiv, pressed *draw.Color
	interac bool
	frame int
}

func (b *Button) SetColours(aktiv, inaktiv, pressed *draw.Color) {
	b.aktiv = aktiv
	b.inaktiv = inaktiv
	b.pressed = pressed
}

func (b *Button) GetColour() draw.Color {
	if b.WasPressed() {
		return *b.pressed
	}
	if b.IsActive() {
		return *b.aktiv
	}
	return *b.inaktiv
}

func (b *Button) CarryEvent(window draw.Window) {
	if HideNames.IsIn(b.name) {
		return
	}
	x, y := window.MousePosition()
	if b.In(x, y) {
		b.SetActive(true)
		if window.IsMouseDown(draw.LeftButton) {
			b.SetPressed(true)
			b.executeIfFameWorkInterac(window)
		} else {
			b.executeIfWasRealaseDontInterac(window)
			b.SetPressed(false)
		}
	} else {
		b.SetPressed(false)
		b.SetActive(false)
	}
}

func (b *Button) executeIfFameWorkInterac(window draw.Window) {
	if b.IsInterac()==false {
		return
	}
	b.frame++
	if b.frame>=0 {
		b.frame=0
		b.GetFunc()(window)
	}
}

func (b *Button) Paint(window draw.Window, scaledText float32) {
	if HideNames.IsIn(b.name) {
		return
	}
	b.Basic.Paint(window, scaledText)
	x, y := b.GetXY()
	width, height := b.GetSize()
	c := b.GetColour()
	intX, intY, intWidth, intHeigth := int(x), int(y), int(width), int(height)
	window.FillRect(intX, intY, intWidth, intHeigth, c)
	window.DrawScaledText(b.name, intX-2, intY-2, scaledText, draw.LightBlue)
}

func (b *Button) IsInterac() bool {
	return b.interac
}

func (b *Button) SetInterac(bool bool) {
	b.interac=bool
}

func (b *Button) executeIfWasRealaseDontInterac(window draw.Window) {
	if b.IsInterac() {
		return
	}
	if b.WasPressed() {
		b.GetFunc()(window)
	}
}

func NewButton(name string) *Button {
	return &Button{
		Basic: newShape(name),
	}
}

func NewButtonRedWhiteGreen(name string) *Button {
	result := NewButton(name)
	result.SetColours(&draw.LightRed, &draw.White, &draw.LightGreen)
	return result
}
