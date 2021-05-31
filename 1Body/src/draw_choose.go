package main

import (
	"fmt"
	"github.com/gonutz/prototype/draw"
	"strings"

	"reflect"
)

var (
	ReflE reflect.Value
	NameAtributEuler []string
	WhichNameSelect []reflect.Type
	WhichName int
)

func init0()  {
	ReflE=reflect.ValueOf(eu).Elem()
	NameAtributEuler = make([]string,0,12)
	x := reflect.Indirect(reflect.ValueOf(eu))
	WhichNameSelect=make([]reflect.Type, 0,cap(NameAtributEuler))
	for i := 0; i < x.NumField(); i++ {
		y := x.Type().Field(i)
		NameAtributEuler = append(NameAtributEuler, y.Name)
		WhichNameSelect = append(WhichNameSelect, y.Type)
	}
	WhichName=0
	ChangeChoose(0)
}
func DrawChoose(window draw.Window) {
	DrawPictures(window)
	drawSelectValue(window)
}

var (
	xs = []int{0,-125,0,125}
	ys = []int{-125,0,125,0}
	wasPanicPiture = false
)

func DrawPictures(window draw.Window) {
	defer checkPanicPicture(window)
	if wasPanicPiture {
		panic("")
	}
	f := []int{0,270,180,90}
	source := []string{"redArrow.png","greenArrow.png"}
	for i:=0;i<len(f);i++{
		temp :=source[0]
		if Klik[i] {
			temp=source[1]
		}
		err := window.DrawImageFilePart(temp,0,0,696,
			564,	int(centerX)+xs[i], int(centerY)+ys[i],100,100,
			f[i])
		if err != nil {
			panic("")
		}
	}
}

func checkPanicPicture(window draw.Window) {
	if err:=recover();err!=nil {
		wasPanicPiture=true
		drawPicturesWarning(window, "redArrow.png","greenArrow.png")
		drawAlternative(window)
	}
}

func drawAlternative(window draw.Window) {
	source := []draw.Color{draw.DarkRed, draw.Green}
	for i:=0;i<len(xs);i++{
		temp :=source[0]
		if Klik[i] {
			temp=source[1]
		}
		window.FillEllipse(int(centerX)+xs[i],
			int(centerY)+ys[i],100,100,temp	)
	}
}

func drawPicturesWarning(window draw.Window, warnings ...string) {
	text:= fmt.Sprint("Warning: ",strings.Join(warnings, " or "), " not found")
	_, y := window.GetScaledTextSize(text, 2)
	window.DrawScaledText(text, 0,int(height)-y-y, 2,
		draw.RGB(255.0/255.0,165.0/255.0,0))
}

func drawSelectValue(window draw.Window) {
	text := "change: "+formatValue()
	x,_ := window.GetScaledTextSize(text,4)
	window.DrawScaledText(text,
		int(centerX)-x/2,0,4,draw.Green)
}

func formatValue() string {
	name:= NameAtributEuler[WhichName]
	if strings.EqualFold(name, "Epsilon") {
		return formatEuler()
	}
	return fmt.Sprint(name,
		" = ", ReflE.FieldByName(name).Interface())
}

func formatEuler() string {
	return fmt.Sprintf(fmt.Sprint("Epsilon = %.", digitToSee,"f"), eu.Epsilon)
}
