package main_drawing

import (
	"2bodyBinary/compute"
	initDraw "2bodyBinary/draw_control/init"
	"2bodyBinary/draw_control/main_drawing/choose"
	"2bodyBinary/draw_control/main_drawing/saved"
	shapes2 "2bodyBinary/draw_control/shapes"
	h "2bodyBinary/help"
	eu "2bodyBinary/structures/euler"
	v "2bodyBinary/variables"
	"fmt"
	"github.com/gonutz/prototype/draw"
	"math"
)

func prepareOffsetFuncs(label *shapes2.Label) (func(window draw.Window), func(window draw.Window)) {
	plus := func(window draw.Window) {
		if eu.MEuler.Offset < 10000 {
			eu.MEuler.Offset++
		}
		label.SetValue(fmt.Sprint("offset = ", eu.MEuler.Offset))
	}
	minus := func(window draw.Window) {
		if eu.MEuler.Offset > 0 {
			eu.MEuler.Offset--
		}
		label.SetValue(fmt.Sprint("offset = ", eu.MEuler.Offset))
	}
	return plus, minus
}

func giveInOrderNextTurn( draw.Window) {
	go func() {
		for i := 0; i < eu.MEuler.Howmanyturn; i++ {
			compute.ChStart <- true
		}
	}()
}

func prepareScaleFunc() (func(window draw.Window), func(window draw.Window)) {
	plus := func(window draw.Window) {
		eu.MEuler.Scale *= 1.1
	}
	minus := func(window draw.Window) {
		eu.MEuler.Scale /= 1.1
	}
	return plus, minus
}

func chooseOrbit(window draw.Window) {
	chooseQuestion = true
	choose.InitSelect(&selectMode, &selected)
	lengthOrbit := int(math.Min(float64(len(planetsPositions.floats[v.FirstElement])),
		float64(len(planetsPositions.floats[v.SecondElement]))))
	choose.InitLength(lengthOrbit)
	choose.Begin(window)
}

func savedState(window draw.Window) {
	saveQuestion = true
	saved.Begin(window)
}

func begin( draw.Window) {
	initDraw.SetToStateInit()
	eu.InitEuler(h.FindPath(v.Configs, v.Config))
	v.IsInit = true
	shapes = nil
	v.ShowAnalytic = false
	v.IsLoad=true
	v.ShowPseudoCentrumOfMomentum = false
	shapes2.HideNames.DeleteAll()
	compute.ChStart <- false
}

func prepareShowProgres(temp string, label *shapes2.Label) func(window draw.Window) {
	myFunc := func(window draw.Window) {
		w, he := window.GetScaledTextSize(temp, 3)
		label.SetSize(float64(w), float64(he))
		period :=  "one period(in days): "
		if eu.MEuler.State != eu.Elliptic && eu.MEuler.State != eu.Circle {
			period="total steps(in days): "
		}
		day := fmt.Sprintf("%.2f",float64(compute.GetPeriod())*eu.MEuler.Epsilon)
		info := fmt.Sprint("shape: ", eu.GetNameOrbit(eu.MEuler), "\n",
			period, day)
		preCounter := fmt.Sprint("count: ", day, "/")
		label.SetValue(
			fmt.Sprint(info, "\n", preCounter,  fmt.Sprintf("%.2f",float64(compute.PeriodProgres)*eu.MEuler.Epsilon)))
	}
	return prepareExecuteOnePerNFrame(myFunc, 60)
}

func prepareShowTurn(label *shapes2.Label) func(window draw.Window) {
	myFunc := func(window draw.Window) {
		label.SetValue(fmt.Sprint("turn = ", compute.NumPeriod))
	}
	return prepareExecuteOnePerNFrame(myFunc, 1)
}

func prepareExecuteOnePerNFrame(myFunc func(window draw.Window), N int) func(window draw.Window) {
	frame := 0
	return func(window draw.Window) {
		if frame++; frame == 1 {
			myFunc(window)
		}
		if frame >= N {
			frame = 0
		}
	}
}

func prepareAnalyticOnOff(chc *shapes2.CheckBox) func(window draw.Window) {
	chc.SetActive(false)
	return func(window draw.Window) {
		chc.SetActive(!chc.IsActive())
		v.ShowAnalytic = chc.IsActive()
		tryLoadAnalytic()
		enableDisableCentrumOfMomentum(chc)
	}
}

func enableDisableCentrumOfMomentum(chc *shapes2.CheckBox) {
	if chc.IsActive() {
		shapes2.HideNames.Add(showCenterOfMomentum)
	}else {
		shapes2.HideNames.Delete(showCenterOfMomentum)
	}
}

func tryLoadAnalytic() {
	if len(planetsPositions.floats) <= 2 {
		planet1, planet2, err := eu.GetAnalyticOrbits()
		if err != nil {
			return
		}
		planetsPositions.floats = append(planetsPositions.floats, [][]float64{planet1})
		planetsPositions.floats = append(planetsPositions.floats, [][]float64{planet2})
	}
}

func preparePseudoCenterOfMomentumOnOff(chc *shapes2.CheckBox) func(window draw.Window) {
	chc.SetActive(false)
	shapes2.HideNames.Add(showAnalytic)
	return func(window draw.Window) {
		chc.SetActive(!chc.IsActive())
		v.ShowPseudoCentrumOfMomentum = chc.IsActive()
		enableDisableAnalytic(chc)
	}
}

func enableDisableAnalytic(chc *shapes2.CheckBox) {
	if chc.IsActive() {
		shapes2.HideNames.Delete(showAnalytic)
	}else {
		shapes2.HideNames.Add(showAnalytic)
	}
}