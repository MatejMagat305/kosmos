package shapes

import "github.com/gonutz/prototype/draw"

type CheckBox struct {
	Basic
}

func (c *CheckBox) CarryEvent(window draw.Window) {
	if HideNames.IsIn(c.name) {
		return
	}
	click := window.Clicks()
	for _, mouseClick := range click {
		if mouseClick.Button==draw.LeftButton == false {
			continue
		}
		if c.In(mouseClick.X, mouseClick.Y) {
			c.GetFunc()(window)
		}
	}
}

func (c *CheckBox) Paint(window draw.Window, scaledText float32) {
	if HideNames.IsIn(c.name) {
		return
	}
	c.Basic.Paint(window, scaledText)
	color := draw.LightRed
	if c.active {
		color = draw.LightGreen
	}
	window.FillRect(int(c.x), int(c.y), int(c.width), int(c.height), draw.White)
	window.FillEllipse(int(c.x), int(c.y), int(c.width), int(c.height), color)
}

func NewCheckBox(s string) *CheckBox {
	return &CheckBox{
		Basic: newShape(s),
	}
}
