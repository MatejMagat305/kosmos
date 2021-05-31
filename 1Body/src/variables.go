package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gonutz/prototype/draw"
	"math"
	"sync"
)

const (
	offset = 3
)
var (
	pause =true
	eu                       *euler
	eulers                   []*euler
	play, drawLine           =false, true
	r, rSun                  =4,4
	rHalf, rSunHalf          =r/2, rSun/2
	width, height            =500.0,500.0
	centerX, centerY         = width/2,height/2
	loop                     =100
	step                     = 0
	stepAdd                  = 100
	akceptSize               = math.Pow(0.1,7)
	mux, mux2                  sync.Mutex


	functions    = []calcul{calculateEuler1,calculateEuler2, calculateEuler3}
	numberObject = len(functions)
	colors                   = []draw.Color{draw.Red, draw.Cyan, draw.LightGreen, draw.White, draw.Yellow}

	orbit                    = make([][]position,0,numberObject)
	areas                    = make([]tuple,0,numberObject)
	xlsxs []*excelize.File

	which =0
	myTime = 1000
	name = []string{"euler", "leapfrog", "euler-Cromer"}
	isInit= true
	isSelect = false

	whichSelect =  0
	begins = make([][]int, 0, 0)
	areasSelect = make([][]float64,0,0)
	aktualArea = make([]float64,0)



	// new After 2planet
	Warning = ""
	BoolType = []string{"AutoScale", "WriteExcelFiles", "RememberWholeOrbit","FixedScale", "ForceCircularOrbit","FixedCentre"}
	IntegerType = []string{"StepsForward","HowManyPositionsBack","SpeedMovePixel"}
	ImutableType = map[string]bool{"PositionX":true,"PositionY":true,
		"VelocityX":true,"VelocityY":true, "Scale":true,"SpeedScaledPercent":true,
		"AccelerationX": true,"AccelerationY":true,
		"Epsilon":false, "AutoScale":false,"FixedScale":false,
		"ForceCircularOrbit":false,	"StepsForward":false , "FixedCentre":false,
		"WriteExcelFiles": true,"HowManyPositionsBack":false,
		"RememberWholeOrbit":false, "SpeedMovePixel":false	}

	Klik = make([]bool,4,4)
	centerX2 =centerX
	centerY2 = centerY
	computePhase = true
	ComputeNumber = int64(0)
	OrbitFull = make([]bool, numberObject, numberObject)

)


func initVariable() {
	eu = newEuler()
	eu.CheckNumberOrbit()
	eu.CheckPercentScale()
	init0()
	initAreas()
	for i := 0; i < numberObject; i++ {
		areasSelect = append(areasSelect, make([]float64,0,20))
		begins = append(begins, make([]int,0,20))
		aktualArea = append(aktualArea, 0)
	}

}

func initAreas() {
	for i := 0; i < numberObject; i++ {
		areas = append(areas, tuple{
			areaFirst:   0,
			areas:       make([]float64, 0, 128),
			wasVerified: false,
			was:         false,
			verify:      false,
		})
	}
}

func AddWarning(s string) {
	if len(Warning)>0 {
		Warning=fmt.Sprint(Warning,"\n", s)
		return
	}
	Warning=s

}
func CheckCircular() {
	if !eu.ForceCircularOrbit {
		return
	}
	vec := NewVector2D(eu.PositionX,eu.PositionY)
	oneDivideR:=math.Sqrt( 1/vec.size())
	vec2 := vec.GetPerpendicularSizeOf(oneDivideR)
	eu.VelocityX=vec2.x
	eu.VelocityY=vec2.y
}