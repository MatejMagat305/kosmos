package euler

import (
	sh "2bodyBinary/draw_control/shapes"
	h "2bodyBinary/help"
	v "2bodyBinary/variables"
	"fmt"
)

func LoadForms(formEuler, formPlanet1, formPlanet2 *h.FormData) bool {
	ok := MEuler.LoadFeatures(formEuler)
	if v.IsGenerate || v.IsInitEulers{
		return ok
	}
	ok2 := MEuler.loadPlanets(formPlanet1, formPlanet2)
	return ok && ok2
}

func (e *Euler) loadPlanets(planet1 *h.FormData, planet2 *h.FormData) bool {
	ok1 := e.loadOnePlanet(planet1, 1)
	ok2 := e.loadOnePlanet(planet2, 2)
	return ok1 && ok2
}

func (e *Euler) loadOnePlanet(planetData *h.FormData, number int) bool {
	result := true
	for i := 0; i < len(planetData.Data); i++ {
		field := planetData.Data[i]
		row, ok := getNameValuePlanet(field)
		if !ok {
			field.SetIsErrorValue(true)
			result = false
			continue
		}
		err := MEuler.Planets[number-1].SetByName(row)
		if err != nil {
			field.SetIsErrorValue(true)
			result = false
		}
	}
	return result
}

func getNameValuePlanet(field *sh.TextField) (string, bool) {
	_, name := field.GetIdName()
	val, ok := jsonToNamePlanet[name]
	if !ok {
		return "", ok
	}
	return fmt.Sprint(val, "=", field.GetValue()), true
}

func (e *Euler) LoadFeatures(features *h.FormData) bool {
	ok := true
	for i := 0; i < len(features.Data); i++ {
		ok = ok && e.setByTextFieldReturnSuccessEuler(features.Data[i])
	}
	return ok
}

func (e *Euler) setByTextFieldReturnSuccessEuler(field *sh.TextField) bool {
	ch := make(chan bool)
	go e.setByTextFieldEuler(field, ch)
	return <-ch
}

func (e *Euler) setByTextFieldEuler(field *sh.TextField, ch chan bool) {
	defer func() {
		if err := recover(); err != nil {
			ch <- false
			field.SetIsErrorValue(true)
		} else {
			ch <- true
		}
	}()
	if field.Value=="" {
		panic("empty")
	}
	_, name := field.GetIdName()
	pseudoRow := fmt.Sprint(jsonToNameEuler[name], "=", field.Value)
	e.SetValue(pseudoRow)
}
