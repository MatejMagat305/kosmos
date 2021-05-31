package shapes

import (
	"fmt"
	"github.com/gonutz/prototype/draw"
	"unicode"
	"unicode/utf8"
)

type TextField struct {
	Basic
	Value        string
	frame        int
	isErrorValue bool
}

func MakeEmptyTextFieldArrayWithCapResetId(howMany int) []*TextField {
	resetId()
	return make([]*TextField, 0, howMany)
}

func NewTextField(name string) *TextField {
	return &TextField{
		Basic:        newShape(name),
		Value:        "",
		frame:        0,
		isErrorValue: false,
	}
}

func (t *TextField) CarryEvent(window draw.Window) {
	if HideNames.IsIn(t.name) {
		return
	}
	t.carryClick(window)
	if t.name == "warning" {
		t.GetFunc()(window)
	}
	if t.active == false {
		return
	}
	t.backSpace(window)

	for _, r := range window.Characters() {
		t.execSelfFuncIfNotNil(window)
		if unicode.IsGraphic(r) {
			t.setTryString(r)
		}
	}
	t.careMouseEvent(window)
}
func (t *TextField) SetValue(value string) {
	t.Value = value
}

func (t *TextField) setTryString(char int32) {
	t.Value = fmt.Sprint(t.Value, string(char))
}

func (t *TextField) backSpace(window draw.Window) {

	if window.WasKeyPressed(draw.KeyBackspace) && t.active && t.Value != "" {
		_, size := utf8.DecodeLastRuneInString(t.Value)
		t.Value = t.Value[:len(t.Value)-size]
	}
}

func (t *TextField) careMouseEvent(window draw.Window) {
	click := window.Clicks()
	length := len(click)
	if length == 0 {
		return
	}
	last := length - 1
	lastClick := click[last]
	if t.In(lastClick.X, lastClick.Y) {
		t.active = true
		t.SetIsErrorValue(false)
	} else {
		t.active = false
	}
}

func (t *TextField) Paint(window draw.Window, scaledText float32) {
	if HideNames.IsIn(t.name) {
		return
	}
	x, y := t.GetXY()
	color := draw.White
	if t.isErrorValue {
		color = draw.Red
	}
	window.FillRect(int(x), int(y), int(t.width), int(t.height), color)
	text := t.getTextCursor()
	t.Basic.Paint(window, scaledText)
	window.DrawScaledText(text, int(t.x), int(t.y), 3, draw.Blue)
}

func (t *TextField) SetIsErrorValue(b bool) {
	t.isErrorValue = b
}

func (t *TextField) GetValue() string {
	return t.Value
}

func (t *TextField) carryClick(window draw.Window) {
	x, y := window.MousePosition()
	if window.IsMouseDown(draw.LeftButton) {
		if t.In(x, y) {
			t.SetActive(true)
		} else {
			t.SetActive(false)
		}
	}
}

func (t *TextField) getTextCursor() string {
	result := t.Value
	if t.active {
		t.frame++
		if (t.frame)<30 {
			result = fmt.Sprint(result, "|")
		}else {
			if t.frame>60 {
				t.frame = 0
			}
		}
	}else {
		t.frame = 0
	}
	return result
}

func (t *TextField) execSelfFuncIfNotNil(window draw.Window) {
	if t.function != nil {
		t.function(window)
	}
}
