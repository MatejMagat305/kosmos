package init

import (
	"retardation/draw_control/common"
	sh "retardation/draw_control/shapes"
	h "retardation/help"
	"retardation/structures/euler"
	v "retardation/variables"
	"fmt"
	"github.com/gonutz/prototype/draw"
)

var (
	nameAtributEuler  []string
	nameAtributPlanet []string
	formEuler         *h.FormData
	formPlanet1       *h.FormData
	formPlanet2       *h.FormData
	x, y              int
	sizeMain          = 0.5
)

func controlEnter(window draw.Window) {
	if window.WasKeyPressed(draw.KeyEnter) || window.WasKeyPressed(draw.KeyNumEnter) {
		loadForm(window)
	}
}

func loadForm(draw.Window) {
	if v.IsWarned {
		finishInitSignalToStartMain()
		return
	}
	if euler.LoadForms(formEuler, formPlanet1, formPlanet2) {
		if ok, period := needWarning(); ok {
			v.Warning = fmt.Sprint("your orbit is bigger than standard:", period, " > ", v.StandardOrbit, " if you want continue click 'done'")
			v.IsWarned = true
			return
		}
		finishInitSignalToStartMain()
	}
}

func needWarning() (bool, int) {
	period := euler.MEuler.CalculatePeriod()
	eInverse := 1 / euler.MEuler.Epsilon
	floatPeriod := period * eInverse
	return int(floatPeriod) > v.StandardOrbit, int(floatPeriod)
}

func finishInitSignalToStartMain() {
	euler.RmOldBinByName(euler.MEuler.Name)
	v.IsLoad = false
	v.IsInit = false
}

func prepareFormState(window draw.Window) {
	initFormData()
	shapes = sh.MakeEmptyShapesArrayWithCapResetId(21)
	loadAtributeIfDoNotExist()
	createForm(window)
	createOther(window)

}

func initFormData() {
	formEuler = &h.FormData{Data: sh.MakeEmptyTextFieldArrayWithCapResetId(7)}
	formPlanet1 = &h.FormData{Data: sh.MakeEmptyTextFieldArrayWithCapResetId(7)}
	formPlanet2 = &h.FormData{Data: sh.MakeEmptyTextFieldArrayWithCapResetId(7)}
}

func createForm(window draw.Window) {
	x, y = 10, 10
	createEulerForm(window)
	createPlanets(window)
}

func createOther(window draw.Window) {
	createDoneBack(window)
	createGenerate(window)
	createWarningLabel()
}

func createDoneBack(window draw.Window) {
	prepareButton := common.AddButtonPrepare(v.CenterX, v.CenterY, &shapes)
	prepareButton(window, loadForm, "Done")
	prepareButton = common.AddButtonPrepare(v.CenterX/2, v.CenterY, &shapes)
	prepareButton(window, func(window draw.Window) {
		SetToStateInit()
		v.IsLoad = true
		euler.InitEuler(h.FindPath(v.Configs, v.Config))
	}, "Back")
}

func createWarningLabel() {
	textfield := sh.NewTextField("warning")
	textfield.SetXY(10, (v.CenterY+v.Height)/2)
	textfield.SetFunc(func(window draw.Window) {
		if v.IsWarned {
			x, y := textfield.GetXY()
			window.DrawScaledText(v.Warning, int(x), int(y), 2, draw.LightRed)
		}
	})
	shapes = append(shapes, textfield)
}

func createPlanets(window draw.Window) {
	planets := []*h.FormData{formPlanet1, formPlanet2}
	for i := 0; i < len(planets); i++ {
		createPlanetLabel(window)
		createPlanetTextField(window, planets[i], i)
	}
}

func createPlanetTextField(window draw.Window, planetData *h.FormData, j int) {
	floatX, floatY := float64(x), float64(y)
	maxX := floatX
	for i := 0; i < len(nameAtributPlanet); i++ {
		name := nameAtributPlanet[i]
		_, height := window.GetScaledTextSize(name, scaledText)
		floatWidth, floatHeight := getWidthToPercent(floatX, sizeMain+float64(j+1)*sizeMain/2), float64(height)
		textfield := sh.NewTextField(name)
		textfield.SetXY(floatX, floatY)
		textfield.SetSize(floatWidth, floatHeight)
		addToShapesAndForm(planetData, textfield)
		textfield.SetValue(euler.GetValuePlanetByNameDefault(name, "", j))
		textfield.SetFunc(clickOn)
		floatY += floatHeight
		tmp := floatWidth + floatX
		if tmp > maxX {
			maxX = tmp
		}
	}
	x = int(maxX) + 5

}

func createPlanetLabel(window draw.Window) {
	floatX, floatY := float64(x), float64(y)
	maxX := floatX
	for i := 0; i < len(nameAtributPlanet); i++ {
		name := nameAtributPlanet[i]
		width, height := window.GetScaledTextSize(name, scaledText)
		floatWidth, floatHeight := float64(width), float64(height)
		label := sh.NewLabel(name)
		label.SetXY(floatX, floatY)
		label.SetSize(floatWidth, floatHeight)
		shapes = append(shapes, label)
		floatY += floatHeight
		tmp := floatWidth + floatX
		if tmp > maxX {
			maxX = tmp
		}
	}
	x = int(maxX) + 5
}

func createEulerForm(window draw.Window) {
	createEulerLabel(window)
	createEulerTextField(window)
}

func createEulerTextField(window draw.Window) {
	floatX, floatY := float64(x), float64(y)
	maxX := floatX
	for i := 0; i < len(nameAtributEuler); i++ {
		name := nameAtributEuler[i]
		_, height := window.GetScaledTextSize(name, scaledText)
		floatWidth, floatHeight := getWidthToPercent(floatX, sizeMain), float64(height)
		textfield := sh.NewTextField(name)
		textfield.SetXY(floatX, floatY)
		textfield.SetSize(floatWidth, floatHeight)
		textfield.SetValue(euler.GetValueEulerByNameDefault(name, ""))
		addToShapesAndForm(formEuler, textfield)
		textfield.SetFunc(clickOn)
		floatY += floatHeight
		tmp := floatWidth + floatX
		if tmp > maxX {
			maxX = tmp
		}
	}
	x = int(maxX) + 5
}

func clickOn(draw.Window) {
	v.IsGenerate = false
	v.IsWarned = false
	v.IsInitEulers = false
}

func getWidthToPercent(width, f float64) float64 {
	return v.Width*f - width

}

func createEulerLabel(window draw.Window) {
	floatX, floatY := float64(x), float64(y)
	maxX := floatX
	for i := 0; i < len(nameAtributEuler); i++ {
		name := nameAtributEuler[i]
		width, height := window.GetScaledTextSize(name, scaledText)
		floatWidth, floatHeight := float64(width), float64(height)
		label := sh.NewLabel(name)
		label.SetXY(floatX, floatY)
		label.SetSize(floatWidth, floatHeight)
		shapes = append(shapes, label)
		floatY += floatHeight
		tmp := floatWidth + floatX
		if tmp > maxX {
			maxX = tmp
		}
	}
	x = int(maxX) + 5
}

func addToShapesAndForm(data *h.FormData, label *sh.TextField) {
	data.Data = append(data.Data, label)
	shapes = append(shapes, label)
}

func loadAtributeIfDoNotExist() {
	loadEulerAtribute()
	loadPlanetAtriobute()
}

func loadPlanetAtriobute() {
	if nameAtributPlanet != nil {
		return
	}
	nameAtributPlanet = euler.GetJsonNamePlanet()
	h.Sort(nameAtributPlanet)
}

func loadEulerAtribute() {
	if nameAtributEuler != nil {
		return
	}
	nameAtributEuler = euler.GetJsonNameEuler()
	h.Sort(nameAtributEuler)
}
