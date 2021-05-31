package euler

import (
	"fmt"
	"reflect"
	h "retardation/help"
)

var (
	MetaMEuler                                                           reflect.Value
	MEuler                                                               *Euler
	FirsrtDigit                                                          = int64(0)
	AfterHowmany0                                                        = 1
	IsEpsDefault                                                         = false
	HowMany                                                              = 0
	DigitToSee                                                           = 1
	Action                                                               = 0.001
	epsilon                                                              = .0010
	momentum                                                             *h.Momentum
	num                                                                  = 0
	jsonToNameEuler, nameToJsonEuler, jsonToNamePlanet, nameToJsonPlanet map[string]string
	defaultPositionX                                                     float64 = 0.5
	defaultPositionY                                                     float64 = 0
	defaultVelocityX                                                     float64 = 0
	defaultVelocityY                                                     float64 = 0.630
	defaultSize                                                          float64 = 1
)


func SetEuler(euler *Euler)  {
	MEuler=euler
	MetaMEuler=reflect.ValueOf(MEuler).Elem()
}

func GetJsonNameEuler() []string {
	return getJsonNameFromMap(jsonToNameEuler)
}

func GetJsonNamePlanet() []string {
	return getJsonNameFromMap(jsonToNamePlanet)
}
func getJsonNameFromMap(ma map[string]string) []string {
	result := make([]string, 0, len(ma))
	for k := range ma {
		result = append(result, k)
	}
	return result
}
func GetValuePlanetByNameDefault(name string, defaultString string, j int) string {
	ch := make(chan string)
	go TryGetValuePlanetByNameDefault(name, defaultString, ch, j)
	return <-ch
}

func TryGetValuePlanetByNameDefault(name string, defaultString string, ch chan string, j int) {
	defer func() {
		if erro := recover(); erro != nil {
			ch <- defaultString
		}
	}()
	temp := reflect.ValueOf(MEuler.Planets[j]).Elem().FieldByName(jsonToNamePlanet[name])

	if temp.Kind()==reflect.Float64 {
		ch <- fmt.Sprintf("%.10f", temp.Float())
	}
	ch <- fmt.Sprint(temp.Interface())
}

func GetValueEulerByNameDefault(name string, defaultString string) string {
	ch := make(chan string)
	go TryGetValueEulerByNameDefault(name, defaultString, ch)
	return <-ch
}

func TryGetValueEulerByNameDefault(name string, defaultString string, ch chan string) {
	defer func() {
		if erro := recover(); erro != nil {
			ch <- defaultString
		}
	}()
	refVal := reflect.ValueOf(MEuler).Elem().FieldByName(jsonToNameEuler[name])
	if refVal.Kind()==reflect.Float64 {
		ch <- fmt.Sprintf("%.10f", refVal.Float())
	}
	ch <- fmt.Sprint(refVal.Interface())
}
func GetNameOrbit(e *Euler) string {
	switch e.State {
	case Elliptic:
		return "elliptic"
	case Parabolic:
		return "parabolic"
	case Circle:
		return "circle"
	default:
		return "hyperbolic"
	}
}