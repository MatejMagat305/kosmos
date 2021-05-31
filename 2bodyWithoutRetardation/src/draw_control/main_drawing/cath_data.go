package main_drawing

import (
	"2bodyBinary/compute"
	h "2bodyBinary/help"
	eu "2bodyBinary/structures/euler"
	v "2bodyBinary/variables"
	"fmt"
	"io/ioutil"
)

func LoadDone() {
	if v.IsLoad {
		numberFile := <-compute.ChDone
		loadIfIsCatch(numberFile, true)
	}
	for {
		select {
		case numberFile := <-compute.ChDone:
			if debug0==1 {
				debug0=2
			}
			loadIfIsCatch(numberFile, false)
		case <-ChStop:
			return
		}
	}
}

func loadIfIsCatch(numberPeriod int, bg bool) {
	mux.Lock()
	defer func() {
		mux.Unlock()
	}()
	i := 1
	if !bg {
		i = numberPeriod
	}
	ch := make(chan bool)
	for ; i <= numberPeriod; i++ {
		go loadOnePlanetIfMarkedOrNotSelectMode(i, v.FirstElement, ch)
		go loadOnePlanetIfMarkedOrNotSelectMode(i, v.SecondElement, ch)
		<-ch;<-ch
	}
}

func loadOnePlanetIfMarkedOrNotSelectMode(i int, numberPlanet int, ch chan bool) {
	defer func() {
		_ = recover()
		ch<-true
	}()
	if i>=len(planetsPositions.floats[numberPlanet]) {
		loadChoose(i, numberPlanet)
	}
}

func loadChoose(i, numberPlanet int) {
	loadFromFile(i, numberPlanet)
}

func loadFromFile(i , numberPlanet int) {
	nameDir := fmt.Sprint("./bin/", eu.MEuler.Name, "/",
		"planet",numberPlanet+1,"_period_",i,".bin")
	b, err := ioutil.ReadFile(nameDir)
	if err != nil {
		return
	}
	LoadPositions(b, numberPlanet)
}

func LoadPositions(b []byte, numberPlanet int) {
	planetsPositions.bytes[numberPlanet] = append(planetsPositions.bytes[numberPlanet], b)
	f :=  h.BytePointerToFloat64Pointer(b)
	planetsPositions.floats[numberPlanet] = append(planetsPositions.floats[numberPlanet], f)
}
