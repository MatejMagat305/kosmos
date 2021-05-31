package choose

import (
	sh "2bodyBinary/draw_control/shapes"
	v "2bodyBinary/variables"
	"fmt"
	"github.com/gonutz/prototype/draw"
)

func Begin(window draw.Window) {
	start()
	loadShapes(window)
}

func start() {
	continue0 = true
	IsWarning = false
	v.Warning=""
	v.IsWarned = false
}

func loadShapes(window draw.Window) {
	y := int(v.Height)/3
	initField()
	y, chc := addLabelCheckBox(y,window)("show all:", prepareShowAll())
	chc.SetActive(!*selectMode)
	y, chc = addLabelCheckBox(y,window)("interval of orbit:", prepareInterval())
	chc.SetActive(isInterval)
	y, t := addLabelTextfield(y, window)("from:", 1)
	numFrom = t
	y, t = addLabelTextfield(y, window)("to:  ", length)
	numTo = t
	y = addButton(window,y)(doneFunc, "done")
	y = addButton(window,y)(cancelFunc, "cancel")
}

func addButton(window draw.Window, y int) func(f func(window2 draw.Window), name string) int{
	return func(f func(window2 draw.Window), name string) int {
		btn := sh.NewButtonRedWhiteGreen(name)
		w,h := window.GetScaledTextSize(name, 3)
		btn.SetXY(v.CenterX-float64(w)/2, float64(y))
		btn.SetSize(float64(w), float64(h))
		btn.SetFunc(f)
		shapes = append(shapes, btn)
		return h+y
	}
}

func addLabelCheckBox(y int, window draw.Window) func(string, func(*sh.CheckBox)func(draw.Window)) (int, *sh.CheckBox) {
	return func(nameLabel string, f func(box *sh.CheckBox) func(window2 draw.Window)) (int,*sh.CheckBox) {
		_,h := addLabelReturnWidthHeight(window, nameLabel)(y)
		chc := sh.NewCheckBox(nameLabel)
		ff := f(chc)
		setXYSizeBehinCenter(chc, y)(h, h)
		chc.SetFunc(ff)
		shapes = append(shapes, chc)
		return h+y+10, chc
	}
}

func addLabelTextfield(y int, window draw.Window) func(string, int) (int,*sh.TextField) {
	return func(nameLabel string, num int) (int,*sh.TextField) {
		w,h := addLabelReturnWidthHeight(window, nameLabel)(y)
		text := sh.NewTextField(nameLabel)
		text.SetValue(fmt.Sprint(num))
		text.SetXY(v.CenterX, float64(y))
		text.SetSize(float64(w)+50, float64(h))
		shapes = append(shapes, text)
		return h+y+10, text
	}
}

func addLabelReturnWidthHeight(window draw.Window, nameLabel string) func(y int) (int,int){
	return func(y int) (int, int) {
		w,h := window.GetScaledTextSize(nameLabel, 3)
		l:= sh.NewLabel(nameLabel)
		l.SetXY(v.CenterX-float64(w)-10, float64(y))
		l.SetSize(float64(w), float64(h))
		shapes = append(shapes, l)
		return w,h
	}
}

func setXYSizeBehinCenter(shape sh.Shape, y int) func(w int,h int) {
	return func(w int, h int) {
		shape.SetXY(v.CenterX, float64(y))
		shape.SetSize(float64(w)+50, float64(h))
	}
}

func initField() {
	shapes = sh.MakeEmptyShapesArrayWithCapResetId(3)
}

func InitSelect(sMode *bool, mapSelect *map[int]bool) {
	selectMode = sMode
	selected = mapSelect
}

func InitLength(inputLength int)  {
	length = inputLength
}