package choose

import (
	sh "retardation/draw_control/shapes"
	v "retardation/variables"
	"fmt"
	"github.com/gonutz/prototype/draw"
	"strconv"
)

func prepareShowAll() func(chc *sh.CheckBox) func(window draw.Window) {
	return func(chc *sh.CheckBox) func(window draw.Window) {
		return func(window draw.Window) {
			*selectMode = chc.IsActive()
			chc.SetActive(!(*selectMode))
		}
	}
}

func prepareInterval() func(chc *sh.CheckBox) func(window draw.Window) {
	return func(chc *sh.CheckBox) func(window draw.Window) {
		return func(window draw.Window) {
			isInterval = !chc.IsActive()
			chc.SetActive(isInterval)
		}
	}
}

func doneFunc(draw.Window) {
	from, to, err := loadFromTextBox()
	if err != nil {
		v.Warning = err.Error()
		IsWarning = true
		return
	}
	war := ""
	if from < 1 {
		war = fmt.Sprint("'from' is ", "smaller than zero: ", from)
	}
	if to>length {
		war = fmt.Sprint(war, "'to' is bigger than length of orbits: ", to)
	}
	if to<from {
		war = fmt.Sprint(war, "'to' is bigger than 'from' of orbits: ", to,"<", from)
	}

	if war == "" || IsWarning && war == v.Warning {
		loadToMap(from, to)
		continue0 = false
	}else {
		IsWarning = true
		v.Warning = war
	}
}

func loadToMap(from, to int) {
	for b := range *selected {
		delete(*selected, b)
	}
	if !*selectMode {
		return
	}
	if to<from {
		to, from = from, to
	}
	if isInterval {
		for i := from; i <= to; i++ {
			(*selected)[i-1] = true
		}
		return
	}
	(*selected)[from-1] = true
	(*selected)[to-1]   = true
}

func loadFromTextBox() (int, int, error) {
	from, err := strconv.Atoi(numFrom.GetValue())
	if err!=nil {
		numFrom.SetIsErrorValue(true)
		return 0, 0, fmt.Errorf(notRead,"first")
	}
	to, err := strconv.Atoi(numTo.GetValue())
	if err!=nil {
		numTo.SetIsErrorValue(true)
		return 0, 0, fmt.Errorf(notRead, "second")
	}
	return from,to,nil
}

func cancelFunc(draw.Window) {
	continue0 = false
}
