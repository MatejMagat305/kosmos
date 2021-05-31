package main

import (
	"math"
)


func calculateEuler2(chanel chan position, chanel2 chan bool, numberFile int) {
	var(
		r,rPow3 float64
		color = colors[numberFile]
		e = eulers[numberFile]
	)
	if play {
		chanel<-position{x:e.PositionX,	y:e.PositionY,c: color}
	}
	title := "t PositionX PositionY VelocityX VelocityY AccelerationX AccelerationY"
	setExcelTitle(title, numberFile)
	r = math.Sqrt(math.Pow(e.PositionX,2)+math.Pow(e.PositionY,2))
	rPow3=math.Pow(r, 3)
	writeToExcel2_1(2,0, e.PositionX, e.AccelerationX, e.PositionY, e.AccelerationY,e.VelocityX, e.VelocityY, r, numberFile)
	e.AccelerationX = -e.PositionX/rPow3
	e.AccelerationY=-e.PositionY/rPow3
	e.VelocityX+=e.AccelerationX*e.Epsilon/2
	e.VelocityY+=e.AccelerationY*e.Epsilon/2
	writeToExcel2_2(3,0.5*e.Epsilon,e.VelocityX, e.VelocityY, numberFile)

	var i = 0
mainLoop:for i= 1;i<loop||play;i++ {
		select {
		case <-Quick:
			break mainLoop
		default:
		}
		e.PositionX+=e.VelocityX*e.Epsilon
		e.PositionY+=e.VelocityY*e.Epsilon
		if play {
			chanel<-position{x:e.PositionX,	y:e.PositionY,	c: color}
		}
		r = math.Sqrt(math.Pow(e.PositionX,2)+math.Pow(e.PositionY,2))
		rPow3=math.Pow(r, 3)
		e.AccelerationX=-e.PositionX/rPow3
		e.AccelerationY = -e.PositionY/rPow3
		writeToExcel2_1(i*2+2,float64(i)*e.Epsilon, e.PositionX, e.AccelerationX,
			e.PositionY, e.AccelerationY,e.VelocityX, e.VelocityY, r, numberFile)
		e.VelocityX+=e.Epsilon*e.AccelerationX/2
		e.VelocityY+=e.Epsilon*e.AccelerationY/2
		writeToExcel2_2(i*2+3,(float64(i)+0.5)*e.Epsilon,e.VelocityX, e.VelocityY, numberFile)

	}
	if i<loop+1{
		chanel2<-true
	}
}
