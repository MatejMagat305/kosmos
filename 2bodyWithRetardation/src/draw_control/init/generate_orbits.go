package init

import (
	sh "retardation/draw_control/shapes"
	"retardation/structures/euler"
	v "retardation/variables"
	"github.com/gonutz/prototype/draw"
)

func createGenerate(window draw.Window) {
	generate, elliptic,parabolic,   hyperbolic, circle :=  "generate random orbit: ",
	"elliptic", "parabolic", "hyperbolic", "elliptic(circle)"
	stringsButtons := []string{elliptic, parabolic, hyperbolic, circle}
	funcsButtons := prepareOrbitFunc()
	x,y := window.GetScaledTextSize(generate,3)
	floatX, floatY := float64(x), float64(y)
	label := sh.NewLabel(generate)
	label.SetXY(10, v.Height-floatY*2-5)
	label.SetSize(floatX,floatY)
	shapes = append(shapes, label)
	floatX+=30
	for i := 0; i < len(stringsButtons); i++ {
		name := stringsButtons[i]
		w, h := window.GetScaledTextSize(name,3)
		floatW, floatH := float64(w), float64(h)
		button := sh.NewButtonRedWhiteGreen(name)
		button.SetXY(floatX, v.Height-floatH*2)
		button.SetSize(floatW, floatH)
		floatX+=floatW+10
		button.SetFunc(funcsButtons[i])
		shapes = append(shapes, button)
	}
}

func prepareOrbitFunc() []func(window2 draw.Window) {
	result := make([]func(window2 draw.Window), 0, 4)
	typesOrbit := []int{euler.Elliptic, euler.Parabolic, euler.Hyperbolic,euler.Circle}//euler.Parabolic,
	for i:=0;i<len(typesOrbit);i++ {
		result = append(result, prepareOneButtonOrbit(typesOrbit[i]))
	}
	return result
}

func prepareOneButtonOrbit(typeOrbit int) func(window2 draw.Window) {
	fn := changeStateFunc(formState)
	return func(window draw.Window) {
		ok := euler.MEuler.LoadFeatures(formEuler)
		if !ok {
			return
		}
		euler.MEuler.GenerateOrbitByType(typeOrbit)
		fn(window)
		v.IsGenerate = true
		v.IsWarned = false
	}
}