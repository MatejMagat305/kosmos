package main_drawing

import (
	"2bodyBinary/draw_control/common"
	shapes2 "2bodyBinary/draw_control/shapes"
	h "2bodyBinary/help"
	eu "2bodyBinary/structures/euler"
	v "2bodyBinary/variables"
	"fmt"
	"github.com/gonutz/prototype/draw"
)

func InitVar(euler *eu.Euler, window draw.Window) {
	initPositionStruct()
	shapes = shapes2.MakeEmptyShapesArrayWithCapResetId(8)
	loadFistPositionPlanets(euler)
	loadShapes(euler, window)
}

func initPositionStruct() {
	planet1 := make([][]byte, 0, 10)
	planet2 := make([][]byte, 0, 10)
	planetsPositions = &PlanetsPositions{
		bytes:  make([][][]byte, 0, 2),
		floats: make([][][]float64, 2, 2),
	}
	planetsPositions.bytes = append(planetsPositions.bytes, planet1)
	planetsPositions.bytes = append(planetsPositions.bytes, planet2)
}

func loadFistPositionPlanets(euler *eu.Euler) {
	loadFistPositionOne(euler, v.FirstElement)
	loadFistPositionOne(euler, v.SecondElement)
	controlScale()
}

func controlScale() {
	if eu.MEuler.Beginsuitablescale == false {
		return
	}
	planet1,planet2 := planetsPositions.floats[v.FirstElement][0], planetsPositions.floats[v.SecondElement][0]
	getBigger(planet1,planet2)
	getSmaller(planet1,planet2)
}

func getBigger(planet1 []float64, planet2 []float64) {
	doOperation(planet1,planet2)(tooClose, func() {
		eu.MEuler.Scale*=1.1
	})

}

func tooClose(x1, y1, x2, y2 int) bool {
	x1, x2 = h.SwapIfNeed(x1, x2)
	y1,y2 = h.SwapIfNeed(y1,y2)
	w := int(v.Width/4)
	height := int(v.Height/4)
	return x1-x2>w || y1-y2> height
}

func getSmaller(planet1 []float64, planet2 []float64) {
	doOperation(planet1,planet2)(func(x1, y1, x2, y2 int) bool {
		return isInScreen(x1, y1) && isInScreen(x2, y2)
	}, func() {
		eu.MEuler.Scale/=1.1
	})
}

func doOperation(planet1 []float64, planet2 []float64)func(condition func(x1,y1,x2,y2 int)bool, execute func()) {
	return func(condition func(x1,y1,x2,y2 int)bool, execute func()) {
		for{
			x1,y1 := getRealXYFromFloat(planet1[0],planet1[1])
			x2,y2 := getRealXYFromFloat(planet2[0],planet2[1])
			if condition(x1,y1,x2,y2){
				return
			}
			execute()
		}
	}
}

func isInScreen(x int, y int) bool {
	return x>0 && x<int(v.Width) && y>0 && y<int(v.Height)
}

func loadFistPositionOne(euler *eu.Euler, i int) {
	planetArrayPositions := make([]byte, 0, 2)
	planetStruct := euler.Planets[i]
	byteX, byteY, err := h.GetBytesPosition(planetStruct.PositionX, planetStruct.PositionY)
	if err != nil {
		return
	}
	planetArrayPositions = append(planetArrayPositions, byteX...)
	planetArrayPositions = append(planetArrayPositions, byteY...)
	LoadPositions(planetArrayPositions, i)
}

func loadShapes(euler *eu.Euler, window draw.Window) {
	loadButtons(window)
	loadLabelsCheckBox(euler, window)
	loadOffsetButtonLabel(window)
}

func loadOffsetButtonLabel(window draw.Window) {
	ofsetName, plus, minus, x := " offset = 1000", "+", "-", v.Width
	width, height := window.GetScaledTextSize(ofsetName, 3)
	addFuntionPlus := common.AddButtonPrepare(x-float64(width)/10, v.Height-float64(height), &shapes)

	l := shapes2.NewLabel(ofsetName)
	l.SetXY(x-float64(width)-1.5*float64(width)/10, v.Height-1.5*float64(height))
	l.SetSize(float64(width), float64(height))
	fPlus, fMinus := prepareOffsetFuncs(l)
	l.SetValue(fmt.Sprint("offset = ", eu.MEuler.Offset))
	addFuntionMinus := common.AddButtonPrepare(x-float64(width)-float64(width)/10-float64(width)/10,
		v.Height-float64(height), &shapes)
	addFuntionPlus(window, fPlus, plus)
	addFuntionMinus(window, fMinus, minus)
	shapes[len(shapes)-1].(*shapes2.Button).SetInterac(true)
	shapes[len(shapes)-2].(*shapes2.Button).SetInterac(true)
	shapes = append(shapes, l)
}

func loadButtons(window draw.Window) {
	addButtonNext(window)
	addButtonSubSumScale(window)
	loadSave(window)
	loadChooseOrbit(window)
	loadBegin(window)
}

func loadChooseOrbit(window draw.Window) {
	clear := "choose"
	width, height := window.GetScaledTextSize(clear, 3)
	x := v.Width
	addFuntion := common.AddButtonPrepare(x-float64(width)/2, v.Height/(3.0/2.0)-float64(height)/2, &shapes)
	addFuntion(window, chooseOrbit, clear)
}

func loadSave(window draw.Window) {
	clear := "save"
	width, height := window.GetScaledTextSize(clear, 3)
	x := v.Width
	addFuntion := common.AddButtonPrepare(x-float64(width)/2, v.Height/3.0-float64(height)/2, &shapes)
	addFuntion(window, savedState, clear)
}

func addButtonSubSumScale(window draw.Window) {
	plus, minus := "+", "-"
	fPlus, fMinus := prepareScaleFunc()
	addFuntionPlus := common.AddButtonPrepare(20, v.Height/2, &shapes)
	addFuntionMinus := common.AddButtonPrepare(v.Width-20, v.Height/2, &shapes)
	addFuntionPlus(window, fPlus, plus)
	addFuntionMinus(window, fMinus, minus)
}

func addButtonNext(window draw.Window) {
	next := "compute next"
	addFuntion := common.AddButtonPrepare(v.Width/2, 50, &shapes)
	addFuntion(window, giveInOrderNextTurn, next)
}

func loadLabelsCheckBox(euler *eu.Euler, window draw.Window) {
	loadNumberLabel(window)
	loadEpsilonLabel(window, euler)
	loadOptionsLabelCheckBox(window)
	loadOrbitCounter(window)
}

func loadBegin(window draw.Window) {
	_, height := window.GetScaledTextSize("l", 3)
	prepareButton := common.AddButtonPrepare(v.CenterX, v.Height-5-float64(height), &shapes)
	prepareButton(window, begin, "Begin")
}

func loadOrbitCounter(window draw.Window) {
	temp := fmt.Sprint("\n")
	label := shapes2.NewLabel(temp)
	_, he := window.GetScaledTextSize("temp", 3)
	label.SetXY(10, float64(he)*2)
	label.SetFunc(prepareShowProgres(temp,label))
	shapes = append(shapes, label)
}


func loadEpsilonLabel(window draw.Window, euler *eu.Euler) {
	epsilon := euler.FormatEpsilon()
	width, height := window.GetScaledTextSize(epsilon, 3)
	l := shapes2.NewLabel(epsilon)
	l.SetXY(v.Width-float64(width)-10, 10)
	l.SetSize(float64(width), float64(height))
	shapes = append(shapes, l)
}

func loadNumberLabel(window draw.Window) {
	name := "number period   "
	label := shapes2.NewLabel(name)
	label.SetXY(10, 10)
	width, height := window.GetScaledTextSize(name, 3)
	label.SetSize(float64(width), float64(height))
	label.SetFunc(prepareShowTurn(label))
	shapes = append(shapes, label)
}

func loadOptionsLabelCheckBox(window draw.Window) {
	loadAnaliticMenu(window)
	pseudoCenterOfMomentum(window)
}

func loadAnaliticMenu(window draw.Window) {
	if eu.MEuler.State != eu.Elliptic && eu.MEuler.State != eu.Circle {
		return
	}
	l := shapes2.NewLabel(showAnalytic)
	l.SetSizeScale(2)
	x, y := window.GetScaledTextSize(showAnalytic, 2)
	beginY := v.Height - float64(y)*3 - 10
	l.SetXY(10, beginY)
	l.SetSize(0, 0)

	shapes = append(shapes, l)
	chc := shapes2.NewCheckBox(showAnalytic)
	chc.SetActive(v.ShowAnalytic)
	chc.SetXY(20+float64(x), beginY-10)
	chc.SetSize(float64(y)+10, float64(y)+10)
	chc.SetFunc(prepareAnalyticOnOff(chc))
	shapes = append(shapes, chc)
}

func pseudoCenterOfMomentum(window draw.Window) {
	if eu.MEuler.Centerofmomentumframe == true  {
		return
	}
	l := shapes2.NewLabel(showCenterOfMomentum)
	l.SetSizeScale(2)
	x, y := window.GetScaledTextSize(showCenterOfMomentum, 2)
	beginY := v.Height - float64(y)*2 - 10
	l.SetXY(10, beginY)
	l.SetSize(0, 0)
	shapes = append(shapes, l)
	chc := shapes2.NewCheckBox(showCenterOfMomentum)
	chc.SetActive(v.ShowPseudoCentrumOfMomentum)
	chc.SetXY(20+float64(x), beginY-10)
	chc.SetSize(float64(y)+10, float64(y)+10)
	chc.SetFunc(preparePseudoCenterOfMomentumOnOff(chc))
	shapes = append(shapes, chc)
}