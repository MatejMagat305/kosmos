package euler

import (
	"fmt"
	"reflect"
	"strconv"
)

type Euler struct {
	Planets   Planets `json:"planets"`
	*Features `json:"features"`
}

func (e *Euler) Clone() *Euler {
	return &Euler{
		Planets:  e.Planets.clone(),
		Features: e.Features.clone(),
	}
}

func (e *Euler) IsEqualPositision() bool {
	return e.Planets.IsEqualPositision()
}

func (e *Euler) AddPlanet(row string) {
	if num >= 2 {
		return
	}
	pl := NewPlanetFromString(row)
	e.Planets = append(e.Planets, pl)
	num++
}

func NewEulerDefault() *Euler {
	return &Euler{
		Planets:  make(Planets, 0, 2),
		Features: newDefaultMutable(),
	}
}

func (e *Euler) IsBoolType(name0 string) bool {
	field := MetaMEuler.FieldByName(name0)
	return field.Kind() == reflect.Bool
}

func (e *Euler) IsInteger64Type(name0 string) bool {
	field := MetaMEuler.FieldByName(name0)
	return field.Kind() == reflect.Int64
}

func (e *Euler) IsIntegerType(name0 string) bool {
	field := MetaMEuler.FieldByName(name0)
	return field.Kind() == reflect.Int
}

func (e *Euler) SetInteger64(value string, name0 string) {
	b, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		panic("can not read value!(true or false)")
	}
	x := MetaMEuler.FieldByName(name0)
	if !x.IsValid() {
		panic("wrong name")
	}
	x.SetInt(b)
}
func (e *Euler) SetFloat(value string, name0 string) {
	X, err := strconv.ParseFloat(value, 64)
	if err != nil {
		panic("wrong value for name")
	}
	x := MetaMEuler.FieldByName(name0)
	if !x.IsValid() {
		panic("wrong name")
	}
	x.SetFloat(X)
	carrySeeDigit(value, name0, e)
}
func (e *Euler) SetBool(value string, name0 string) {
	b, err := strconv.ParseBool(value)
	if err != nil {
		panic("can not read value!(true or false)")
	}
	x := MetaMEuler.FieldByName(name0)
	if !x.IsValid() {
		panic("wrong name")
	}
	x.SetBool(b)
}

func (e *Euler) SetInteger(value string, name0 string) {
	b, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		panic("can not read value!(true or false)")
	}
	x := MetaMEuler.FieldByName(name0)
	if !x.IsValid() {
		panic("wrong name")
	}
	x.SetInt(b)
}
func (e *Euler) IsStringType(name0 string) bool {
	field := MetaMEuler.FieldByName(name0)
	return field.Kind() == reflect.String
}

func (e *Euler) SetString(value string, name0 string) {
	MetaMEuler.FieldByName(name0).SetString(value)
}

func (e *Euler) FormatEpsilon() string {
	return fmt.Sprintf(fmt.Sprint("epsilon = %", AfterHowmany0, "v "), e.Epsilon)
}

func (e *Euler) GetPlanets() (*Planet, *Planet) {
	return e.Planets[0],e.Planets[1]
}