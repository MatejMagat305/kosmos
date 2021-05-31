package main

import (
	"github.com/gonutz/prototype/draw"
	"math"
	"time"
)

func waitAll(bools []chan bool) {
	for i := 0; i < len(bools); i++ {
		_=<-bools[i]
	}
}

func makeChanels() (chanPosotin []chan position,chanBool []chan bool) {
	chanPosotin=make([]chan position,0,numberObject)
	chanBool = make([]chan bool,0, numberObject)
	for i := 0; i < numberObject; i++ {
		chanPosotin = append(chanPosotin, make(chan position))
		chanBool = append(chanBool, make(chan bool))
	}
	return
}

func setPlay(b bool) {
	play=b

}

func findSizesPosition() int {
	defer func() {mux.Unlock()}()
	mux.Lock()
	result:= len(orbit[which])
	if OrbitFull[which] {
		return int(eu.HowManyPositionsBack)
	}
	return result
}

func keyControl(window draw.Window, chanelBools []chan bool) {
	if window.WasKeyPressed(draw.KeySpace) {
		drawLine=!drawLine
	}
	if window.WasKeyPressed(draw.KeyP) {
		if step==0 {
			pause=!pause
		}
	}
	if window.WasKeyPressed(draw.KeyNumAdd) {
		addStep()
	}

	if window.WasKeyPressed(draw.KeyNumSubtract) {
		subStep()
	}

	if window.WasKeyPressed(draw.KeyS) {
		isSelect=!isSelect
	}

	if window.WasKeyPressed(draw.KeyK){
		if pause && step==0 {
			step+=stepAdd
		}
	}
	if window.WasKeyPressed(draw.KeyE) {
		which=(which+1)%numberObject
	}
	if window.WasKeyPressed(draw.KeyC) {
		ComputePhase(chanelBools)
	}
	ScaleControl(window)

}

func ComputePhase(chanelBools[]chan bool) {
	computePhase=true
	ComputeNumber= eu.StepsForward
	pause=false
	go AfterComputeSetNormal(chanelBools)
}

func AfterComputeSetNormal(chanelBools []chan bool) {
	waitAll(chanelBools)
	computePhase=false
	pause=true
}

var (
	actionStep = stepAdd
	howManyStep =1
)

func subStep() {
	if actionStep<=1 && howManyStep==1 {
		return
	}
	howManyStep--
	if howManyStep<=0 {
		howManyStep=9
		actionStep/=10
	}
	stepAdd-=actionStep

}

func addStep() {
	howManyStep++
	stepAdd+=actionStep
	if howManyStep>=10 {
		howManyStep=1
		actionStep*=10
	}
}

var (
	Stop = make([]bool,4,4)
)

func initControl(window draw.Window) {
	if window.WasKeyPressed(draw.KeyEnter) {
		isInit=false
	}
	if window.WasKeyPressed(draw.KeyUp) && !Stop[0] {
		carryKlik(0)
		ChangeChoose(1)
		return
	}
	if window.WasKeyPressed(draw.KeyLeft) && !Stop[1] {
		carryKlik(1)
		SubChose()
		return
	}
	if window.WasKeyPressed(draw.KeyDown) && !Stop[2] {
		carryKlik(2)
		ChangeChoose(-1)
		return
	}
	if window.WasKeyPressed(draw.KeyRight) && !Stop[3] {
		carryKlik(3)
		AddChoose()
		return
	}
}

func carryKlik(i int) {
	Klik[i]=true
	Stop[i]=true
	go KlikGo(i)
}

func ChangeChoose(i int) {
	if i==0 {
		i=1
	}
	WhichName+=i
	checkOverFloor()
	for IsImutable(NameAtributEuler[WhichName]) {
		WhichName+=i
		checkOverFloor()
	}
}

func checkOverFloor() {
	if WhichName==len(NameAtributEuler) {
		WhichName=0
	}
	if WhichName<0 {
		WhichName=len(NameAtributEuler)-1
	}
}

func KlikGo(i int) {
	time.Sleep(time.Second/3)
	Stop[i]=false
	Klik[i]=false
}
var (
	digitToSee = 3
	action = 0.001
	howMany =1
)

func epsilonSub() {
	epsilonSkip()
	howMany--
	if howMany<=0 {
		howMany=9
		action*=0.1
		digitToSee++
		akceptSize =math.Pow(0.1, float64(digitToSee+4))
	}
	eu.Epsilon-=action
}

func epsilonAdd() {
	epsilonSkip()
	howMany++
	eu.Epsilon+=action
	if howMany==10{
		howMany=1
		action*=10
		digitToSee--
		akceptSize =math.Pow(0.1, float64(digitToSee+4))
	}
}

func epsilonSkip() {
	if !isInitEpsilon {
		return
	}
	isInitEpsilon=false
	howMany= int(firsrtDigit)
	digitToSee=afterHowmany0
	action=math.Pow10(-digitToSee)
	eu.Epsilon= float64(howMany) * action
}