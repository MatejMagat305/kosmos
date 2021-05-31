package main

import (
	"math"
)

//euler-croner

func calculateEuler3(chanel chan position, chanel2 chan bool, numberFile int) {
	var(
		r, rPow3 float64
		color = colors[numberFile]
		e = eulers[numberFile]
	)
	if play {
		chanel<-position{x:e.PositionX,	y:e.PositionY,c: color}
	}
	title := "t PositionX PositionY VelocityX VelocityY AccelerationX AccelerationY"
	setExcelTitle(title, numberFile)
	var i = 0
mainLoop:for i=0;i<loop||play;i++ {
		select {
		case <-Quick:
			break mainLoop
		default:
		}
		rPow3 = math.Pow(r,3)
		r = math.Sqrt(math.Pow(e.PositionX, 2)+math.Pow(e.PositionY, 2))
		rPow3 = math.Pow(r,3)
		e.AccelerationX=-e.PositionX/rPow3
		e.AccelerationY=-e.PositionY/rPow3
		e.VelocityX+=e.Epsilon*e.AccelerationX
		e.VelocityY+=e.Epsilon*e.AccelerationY
		e.PositionX+=e.VelocityX*e.Epsilon
		e.PositionY+=e.VelocityY*e.Epsilon
		t := e.Epsilon*float64(i+1)
		writeToExcel1(t,e.PositionX, e.PositionY, e.VelocityX,
			e.VelocityY, e.AccelerationX, e.AccelerationY,i+2, numberFile)
		if play {
			chanel<-position{x:e.PositionX,	y:e.PositionY,	c: color}
		}
	}
	if i<loop+1{
		chanel2<-true
	}
}
