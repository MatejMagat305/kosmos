package main_drawing

import (
	"fmt"
	"github.com/gonutz/prototype/draw"
	"retardation/draw_control/common"
	"retardation/draw_control/main_drawing/choose"
	"retardation/draw_control/main_drawing/saved"
	h "retardation/help"
	eu "retardation/structures/euler"
	v "retardation/variables"
)

func DrawControlState(window draw.Window) {
	if saveQuestion {
		saveQuestion = saved.DrawControlContinue(window)
		chooseQuestion = false
		return
	}
	if chooseQuestion {
		chooseQuestion = choose.DrawControlContinue(window)
		saveQuestion = false
		return
	}
	drawControlAllAroundPlanets(window)
}

func drawControlAllAroundPlanets(window draw.Window) {
	window.FillRect(-1, -1, int(v.Width)+1, int(v.Height)+1, draw.Black)
	common.DrawControlAll(window, &shapes)
	drawsBothOrbits(window)
	drawsBothBegin(window)
	drawAnalyticIfExistAndShow(window)
	controlKeyboard(window)
}

func drawsBothBegin(window draw.Window) {
	planet1, planet2 := planetsPositions.floats[v.FirstElement], planetsPositions.floats[v.SecondElement]
	sumMasses, mass1, mass2, color1, color2 := eu.MEuler.Planets.SumMass(),
		eu.MEuler.Planets[0].Mass, eu.MEuler.Planets[1].Mass, colors[v.FirstElement], colors[v.SecondElement]
	orbit1, orbit2 := planet1[0], planet2[0]
	x1, y1 := orbit1[0], -orbit1[1]
	x2, y2 := orbit2[0], -orbit2[1]
	momentumX := (mass1/sumMasses)*x1 + mass2/sumMasses*x2
	momentumY := (mass1/sumMasses)*y1 + mass2/sumMasses*y2
	if v.ShowPseudoCentrumOfMomentum {
		x1, y1, x2, y2 = x1-momentumX, y1-momentumY, x2-momentumX, y2-momentumY
		momentumX, momentumY =0,0
	}
	proxyColor := draw.Color{R: color1.R, G: color1.G + 0.65, B: color1.B, A: 1}
	prepareDrawOneBegin(x1, y1)(proxyColor, fmt.Sprint("body ", 1))(window, 1)
	proxyColor = draw.Color{R: color2.R, G: color2.G + 0.65, B: color2.B, A: 1}
	prepareDrawOneBegin(x2, y2)(proxyColor, fmt.Sprint("body ", 2))(window, 1)
	prepareDrawOneBegin(momentumX, momentumY)(draw.LightYellow, "mass center")(window, -2)

}

func prepareDrawOneBegin(x float64, y float64) func(color draw.Color, label string) func(window draw.Window, sign int) {
	return func(color draw.Color, label string) func(window draw.Window, sign int) {
		return func(window draw.Window, sign int) {
			realX, realY := getRealXYFromFloat(x, y)
			window.FillEllipse(realX-2, realY-2, 4, 4, color)
			xL, _ := window.GetScaledTextSize(label, 1)
			window.DrawScaledText(label, realX-xL/2, realY+sign*10, 1, color)
		}
	}
}

var (
	debug0 = 0
)

func drawsBothOrbits(window draw.Window) {
	planet1, planet2 := planetsPositions.floats[0], planetsPositions.floats[1]
	sumMasses, mass1, mass2, color1, color2 := eu.MEuler.Planets.SumMass(),
		eu.MEuler.Planets[0].Mass, eu.MEuler.Planets[1].Mass,
		colors[0], colors[1]
	var (
		ratio1, ratio2 = mass1 / sumMasses, mass2 / sumMasses
	)
	lengthObits := h.Min(len(planet1), len(planet2))
	for i := 1; i < lengthObits; i++ {
		if IsSelectMode() && !IsInSelected(i) {
			continue
		}
		orbit1, orbit2 := planet1[i], planet2[i]
		lengthOne := h.Min(len(orbit1)/2, len(orbit2))
		for j := 0; j < lengthOne; j += 1 + eu.MEuler.Offset {
			k2 := j * 2
			k2Plus1 := k2 + 1
			x1, y1 := orbit1[k2], -orbit1[k2Plus1]
			x2, y2 := orbit2[k2], -orbit2[k2Plus1]
			momentumX := ratio1*x1 + ratio2*x2
			momentumY := ratio1*y1 + ratio2*y2
			if v.ShowPseudoCentrumOfMomentum {
				x1, y1, x2, y2 = x1-momentumX, y1-momentumY, x2-momentumX, y2-momentumY
				momentumX, momentumY = 0, 0
			}
			realX, realY := getRealXYFromFloat(x1, y1)
			window.DrawPoint(realX, realY, color1)
			realX, realY = getRealXYFromFloat(x2, y2)
			window.DrawPoint(realX, realY, color2)
			realX, realY = getRealXYFromFloat(momentumX, momentumY)
			window.DrawPoint(realX, realY, draw.LightYellow)
		}
	}
}
func drawAnalyticIfExistAndShow(window draw.Window) {
	if !v.ShowAnalytic || len(planetsPositions.floats) <= 2 {
		return
	}
	first, second := v.FirstElement+2, v.SecondElement+2
	drawOnePlanetOrbits(window, planetsPositions.floats[first], colors[first])
	drawOnePlanetOrbits(window, planetsPositions.floats[second], colors[second])
}

func drawOnePlanetOrbits(window draw.Window, planet [][]float64, color draw.Color) {
	for i := 0; i < len(planet); i++ {
		drawOneOrbit(window, planet[i], color)
	}
}

func drawOneOrbit(window draw.Window, planet []float64, color draw.Color) {
	for i := 1; i < len(planet)/2; i += 1 + eu.MEuler.Offset {
		x, y := getRealXYInt(planet, i)
		window.DrawPoint(x, y, color)
	}
}

func getRealXYInt(planet []float64, i int) (int, int) {
	return getRealXYFromFloat(planet[i*2], -planet[i*2+1])
}

func getRealXYFromFloat(x, y float64) (int, int) {
	ScaleScreen := v.ScaleScreen * eu.MEuler.Scale
	return int(x*ScaleScreen + v.CenterX + eu.MEuler.MoveX),
		int(y*ScaleScreen + v.CenterY + eu.MEuler.MoveY)
}
