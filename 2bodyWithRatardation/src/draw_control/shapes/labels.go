package shapes

import "github.com/gonutz/prototype/draw"

type Label struct {
	Basic
	value     string
	sizeScale float32
}

func NewLabel(name string) *Label {
	return &Label{newShape(name), name, 3}
}

func (l *Label) CarryEvent(window draw.Window) {
	if l.function != nil {
		l.function(window)
	}
}

func (l *Label) Paint(window draw.Window, scaledText float32) {
	if HideNames.IsIn(l.name) {
		return
	}
	l.Basic.Paint(window, scaledText)
	x, y := l.GetXY()
	window.FillRect(int(x), int(y), int(l.width), int(l.height), draw.White)
	window.DrawScaledText(l.value, int(x), int(y), l.sizeScale, draw.Red)
}

func (l *Label) SetValue(s string) {
	l.value = s
}

func (l *Label) SetSizeScale(s float32)  {
	l.sizeScale = s
}