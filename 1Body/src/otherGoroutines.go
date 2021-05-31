package main

import (
	"math"
	"time"
)

func startCalculations(chP []chan position, chB []chan bool) {
	for i := 0; i < numberObject; i++ {
		orbit = append(orbit, make([]position,0,1024))
		eulers = append(eulers, eu.clone())
	}
	for i := 0; i < numberObject; i++ {
		go functions[i](chP[i], chB[i], i)
	}
	go catchingData(chP, chB)
}

var (
	timeToSleep = time.Now()
	index = 0
)
func catchingData(p []chan position, b []chan bool) {
	lenght := len(p)
	for play {
		results := make([]position,0, lenght)
		index++
		for i := 0; i < lenght; i++ {
			temp := <-p[i]
			results = append(results, temp)
		}
		howManyFake :=loadPosition(results, lenght)
		if howManyFake == numberObject{
			break
		}
		SleepAndPause()
		CountdownCompute(b)
	}
}

func loadPosition(results []position, lenght int)(result int) {
	result = 0
	mux.Lock()
	for j := 0; j < lenght; j++ {
		p := results[j]
		carrySelect(step, &p, j)
		orbit[j] = append(orbit[j], p)
	}
	carrySelectBegin()
	mux.Unlock()

	if eu.RememberWholeOrbit {
		return
	}
	for j := 0; j < lenght; j++ {
		reduceOrbit(j)
	}
return
}

func reduceOrbit(j int) {
	if int64(len(orbit[j]))>eu.HowManyPositionsBack+eu.HowManyPositionsBack/100 {
		mux2.Lock()
		orbit[j]=orbit[j][eu.HowManyPositionsBack/100:]
		mux2.Unlock()
	}
}

func CountdownCompute(b []chan bool) {
	if computePhase && ComputeNumber>0{
		ComputeNumber--
		if step>0 {
			step--
			temp=true
		}
		if ComputeNumber==0 {
			sendAll(b)
		}
	}
}

func sendAll(b []chan bool) {
	for i := 0; i < len(b); i++ {
		b[i]<-true
	}
}

func SleepAndPause() {
	if computePhase {
		return
	}
	timeToSleep = ifNesseserySleepReturnTime(timeToSleep)
	carryPauseSleepSelectBegin()
}

func carryPauseSleepSelectBegin() {
	for pause {
		if step>0 {
			step--
			temp=true
			break
		}
	time.Sleep(time.Second/6)
	}
}

var (
	temp =  false
)

func carrySelectBegin() {
	if step==0 && isSelect && temp &&  (pause||ComputeNumber>0) {
		whichSelect++
		isSelect=false
		temp=false
		for i := 0; i < numberObject; i++ {
			areasSelect[i] = append(areasSelect[i], aktualArea[i])
			aktualArea[i]=float64(0)
		}
	}
}
func ifNesseserySleepReturnTime(sleep time.Time) time.Time {
	now := time.Now()
	if int(now.Sub(sleep))<myTime {
		time.Sleep(time.Duration(myTime-int(now.Sub(sleep))))
	}
	return time.Now()
}

func carrySelect(step int, p *position, j int) {
	if step>=0 && isSelect && temp && (pause||ComputeNumber>0) {
		p.isSelect=isSelect
		p.howMany = stepAdd
		p.whichSelect=whichSelect
		p.area=colors[p.whichSelect%len(colors)]
		l := len(orbit[j])-1
		p2:=orbit[j][l]
		aktualArea[j]+=computeAreaFromTwoPoint(p,&p2)
	}
}

func computeAreaFromTwoPoint(p *position, p2 *position) float64 {
	return math.Abs((p.x*p2.y-p.y*p2.x)/2)
}
