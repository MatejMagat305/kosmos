package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"reflect"
	"strconv"
	"strings"
)

type euler struct {
	PositionX, PositionY ,
	VelocityX, VelocityY, Scale,
	AccelerationX,AccelerationY, Epsilon,SpeedScaledPercent float64
	AutoScale,FixedScale, ForceCircularOrbit,
	FixedCentre,WriteExcelFiles, RememberWholeOrbit bool
	StepsForward, HowManyPositionsBack, SpeedMovePixel int64
}

func (e *euler) clone() *euler {
	return &euler{
		PositionX:          e.PositionX,
		PositionY:          e.PositionY,
		VelocityX:          e.VelocityX,
		VelocityY:          e.VelocityY,
		Scale:              e.Scale,
		AccelerationX:      e.AccelerationX,
		AccelerationY:      e.AccelerationY,
		Epsilon:            e.Epsilon,
		AutoScale:          e.AutoScale,
		FixedScale:         e.FixedScale,
		ForceCircularOrbit: e.ForceCircularOrbit,
		StepsForward:          e. StepsForward,
		FixedCentre: e.FixedCentre,
		WriteExcelFiles: e.WriteExcelFiles,
	}
}


func (e *euler) isOnlyZeros() bool {
	if (math.Abs(e.VelocityY)>akceptSize  || math.Abs(e.VelocityX)>akceptSize) &&
		(math.Abs(e.PositionY)>akceptSize || math.Abs(e.PositionX)>akceptSize){
		return false
	}
	return true
}


var(
	epsilon  = .0010
	numberSizeOrbit = 10000
)
func newEulerDefault() *euler {
	var(
		PositionX          =  1.6
		PositionY          =  0.1
		VelocityX          =  0.0
		VelocityY          =  0.2
		Epsilon            =  0.0001
		Scale              =  3.00
		StepsForward          =  int64(10000)
	)
	digitToSee=4
	action=Epsilon
	return &euler{
		PositionX:            PositionX,
		PositionY:            PositionY,
		VelocityX:            VelocityX,
		VelocityY:            VelocityY,
		Scale:                Scale,
		AccelerationX:        0,
		AccelerationY:        0,
		Epsilon:              Epsilon,
		SpeedScaledPercent:   0.01,
		AutoScale:            true,
		FixedScale:           false,
		ForceCircularOrbit:   false,
		FixedCentre:          false,
		WriteExcelFiles:      true,
		RememberWholeOrbit:   true,
		StepsForward:         StepsForward,
		HowManyPositionsBack: int64(numberSizeOrbit),
		SpeedMovePixel:       5,
	}
}

func newEuler() *euler {
	ch := make(chan *euler)
	go readEulerFromFile(ch)
	return <-ch
}

func readEulerFromFile(ch chan *euler) {
	c, err := ioutil.ReadFile("config.txt")
	defer func() {
		if r := recover(); r != nil {
			AddWarning(fmt.Sprint(r))
			ch <- newEulerDefault()
		}
	}()
	defaultV:=", program will run with default value"
	if err != nil {
		panic("file config.txt can not open"+defaultV)
	}
	str := string(c)
	lenStr := len(str)
	if lenStr<5 {
		panic("file is empty or too short"+", program will run with default value")
	}
	array := strings.Split(str, fmt.Sprintln())
	result := newEulerDefault()
	y :=  reflect.ValueOf(result).Elem()
	for i := 0; i < len(array); i++ {
		row := array[i]
		result.SetValue(row, i, y)
	}
	carryEpsilon(result)
	if result.isOnlyZeros() {
		panic("read values were only zeros, program will run with default value")
	}
	ch <- result
}

func (e *euler) SetValue(row string, i int, y reflect.Value) {
	nameValue := strings.Split(row, "=")
	defer func() {
		if r := recover(); r != nil {
			v1 := fmt.Sprint("row ",i+1, ": ")
			v2 :=fmt.Sprint( " was ignored(", r, ")")
			W:=fmt.Sprint(v1,"\"", strings.TrimSpace(row),"\"", v2)
			AddWarning(W)
		}
	}()
	if len(nameValue)<2 {
		panic("wrong structure")
	}
	name0 := strings.TrimSpace(nameValue[0])
	value := strings.TrimSpace(nameValue[1])
	if e.IsBoolType(name0) {
		e.SetBool(value,name0, y )
		return
	}
	if e.IsIntegerType(name0) {
		e.SetInteger(value, name0, y)
		return
	}
	e.SetFloat(value, name0,y)
}
func (e *euler) SetInteger(value string, name0 string, y reflect.Value) {
	b, err := strconv.ParseInt(value,10,64)
	if err != nil {
		panic("can not read value!(integer)")
	}
	x:=y.FieldByName(name0)
	if !x.IsValid() {
		panic("wrong name")
	}
	x.SetInt(b)
}
func (e *euler) SetFloat(value string, name0 string, y reflect.Value) {
	X, err := strconv.ParseFloat(value,64)
	if err != nil {
		panic("wrong value for name(float)")
	}
	x:=y.FieldByName(name0)
	if !x.IsValid() {
		panic("wrong name")
	}
	x.SetFloat(X)
	carrySeeDigit(value,name0, e)
}
func (e *euler) SetBool(value string, name0 string, y reflect.Value) {
	b, err := strconv.ParseBool(value)
	if err != nil {
		panic("can not read value!(true or false)")
	}
	x:=y.FieldByName(name0)
	if !x.IsValid() {
		panic("wrong name")
	}
	x.SetBool(b)
}

func (e *euler)IsBoolType(name0 string) bool {
	for i := 0; i < len(BoolType); i++ {
		if BoolType[i]==name0 {
			return true
		}
	}
	return false
}

func (e *euler) IsIntegerType(name0 string) bool {
	for i := 0; i < len(IntegerType); i++ {
		if IntegerType[i]==name0 {
			return true
		}
	}
	return false
}
var (
	isInitEpsilon = true
	firsrtDigit =int64(0)
	afterHowmany0=1
	isEpsDefault =false
)

func carrySeeDigit(value string, name0 string, result *euler) {
	if strings.EqualFold(name0, "Epsilon") {
		isInitEpsilon=true
		carryEpsilon(result)
		checkFirstDigit(value)
	}
}

func checkFirstDigit(value string) {
	if isEpsDefault {
		return
	}
	Dot := strings.Split(value, ".")
	if len(Dot)!=2  {
		isInitEpsilon=false
		return
	}
	afterDot := Dot[1]
	temp:=len(afterDot)
	if temp==0 {
		isInitEpsilon=false
	}

	digitToSee=temp
	for i := 0; i < temp; i++ {
		if afterDot[i]=='0' {
			afterHowmany0++
		}else {
			firsrtDigit, _=strconv.ParseInt(string(afterDot[i]), 10, 32)
			return
		}
	}
}

func carryEpsilon(result *euler) {
	if result.Epsilon<math.Pow(10,-10) || result.Epsilon>1.00{
		AddWarning("epsilon was too small or too big, it is set to default")
		result.Epsilon=epsilon
		howMany=1
		digitToSee=3
		isEpsDefault=true
		isInitEpsilon=false
	}
}


func (e *euler) Zoom(f float64) {
	if e.FixedScale || e.AutoScale {
		return
	}
	e.Scale+=f*e.Scale*e.SpeedScaledPercent/100
}
func (e *euler) ZoomAutoScale(f float64) {
	if e.FixedScale {return}
	e.Scale+=f*e.Scale*e.SpeedScaledPercent/100
}

func (e *euler) EscapeSpeed() bool {
	r := e.GetR()
	EscapeSpeed := math.Sqrt(2.0/r)
	TotalS := e.GetSpeed()
	if TotalS>=EscapeSpeed {
		return true
	}
	return false
}

func (e *euler) GetR() float64 {
	v := NewVector2D(e.PositionX,e.PositionY)
	return v.size()
}

func (e *euler) GetSpeed() float64 {
	v := NewVector2D(e.VelocityX, e.VelocityY)
	return v.size()
}

func (e *euler) CheckNumberOrbit() {
	if int(e.HowManyPositionsBack+e.HowManyPositionsBack/100)<=0 {
		e.HowManyPositionsBack=int64(numberSizeOrbit)
		Warning+=" HowManyOrbitBack was too big, negative or missing, it was ignored and got default"
	}
}

func (e *euler) InSun() bool {
	 v:=NewVector2D(e.VelocityX, e.PositionY)
	if v.size()<math.Pow(0.1,6) {
		return true
	}
	return false
}

func (e *euler) CheckPercentScale() {
	if e.SpeedScaledPercent>50 {
		AddWarning("SpeedScaledPercent was bigger than 50, it will run default(50%)")
		e.SpeedScaledPercent=50
	}
	if e.SpeedScaledPercent<0 {
		AddWarning("SpeedScaledPercent was smaller than 0, it will run default(50%)")
		e.SpeedScaledPercent=50
	}
}