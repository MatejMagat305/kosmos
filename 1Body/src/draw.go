package main

import (
	"fmt"
	"github.com/gonutz/prototype/draw"
	"strconv"
)

func drawAll(indexes int, window draw.Window) {
	CaryScale(indexes)
	drawOrbit(window, indexes)
	drawSelect(window)
	drawData(window, indexes)
	drawLegend(window)
	drawObjects(window, indexes)
}

func initDraw(window draw.Window) {
	DrawChoose(window)
	DrawWarning(window)
	DrawWarningEscapeSpeed(window)
}

func drawComputeScreen(window draw.Window) {
	if eu.StepsForward<500000 {
		return
	}
	wait := "wait please, the calculation is in progress:"
	per := getPercent()
	f := fmt.Sprintf("%2.2f%v",per*100,"%")
	wxf, _ := window.GetScaledTextSize(f,2)
	wx,wy := window.GetScaledTextSize(wait,2)
	y := int(height/3)
	window.DrawScaledText(wait,int(centerX)-wx/2,y,2,draw.LightGreen)
	y2:= y+wy
	end := int(width/50)
	window.DrawRect(end,y2, int(width)-end-wxf, wy,draw.LightGreen )
	window.FillRect(end,y2, int(width*per)-end-wxf, wy,draw.LightGreen)

	window.DrawText(f, int(width*per)-end,y2+wy/4,draw.LightGreen)
}


func getPercent() float64 {
	return float64(eu.StepsForward-ComputeNumber)/float64(eu.StepsForward)
}

var (
	scaled = float32(2.5)
)

func DrawWarning(window draw.Window) {
	if len(Warning)!=0 {
		text := fmt.Sprint("warning: ",Warning)
		for  {
			sizeX, sizeY := window.GetScaledTextSize(text, scaled)
		if float64(sizeY)<centerY && float64(sizeX)<width{
			break
		}
			scaled-=0.00001
		}
		window.DrawScaledText(text,
			0,int(centerY+250),
			scaled,draw.RGB(255.0/255.0,165.0/255.0,0))
	}
}

func drawData(window draw.Window, indexes int) {
	DrawData2(window)
	text:=""
	if !isInit {
		if orbit==nil {
			goto end
		}
		if len(orbit)==0 {
			goto end
		}
		if len(orbit[which])==0 {
			goto end
		}
		p:=orbit[which][indexes-1]
		text=fmt.Sprint("x = ", fmt.Sprintf("%.10f",p.x), ", y = ", fmt.Sprintf("%.10f",p.y))

	}
	end:
		x:=""
	window.DrawScaledText(fmt.Sprint(text,x),0, int(height-height/10),3,draw.LightRed)
}

func DrawData2(window draw.Window) {
	text := formatData()
	window.DrawScaledText(text, 0,0,3, draw.White)
}

func drawOrbit(window draw.Window, indexes int) {
	drawOnePathOrbit(window, orbit[which], indexes)
}

func drawOnePathOrbit(window draw.Window, positions []position, indexes int) {
	if len(positions)<=offset+1 {
		return
	}
	first := positions[offset]
	for i := offset; i < indexes-offset; i++ {
		second := positions[i]
		drawOrbitO(window, &first, &second)
		drawOrbitArea(window, &first)
		first=second
	}
}

func drawOrbitO(window draw.Window, first *position, second *position) {
	x1, y1:= computeRealXY(first)
	if drawLine {
		x2,y2 := computeRealXY(second)
		window.DrawLine(x1, y1,	x2, y2,second.c)
	}else {
		window.DrawPoint(x1,y1,first.c)
	}
}

func drawOrbitArea(window draw.Window, p *position) {
	x, y := computeRealXY(p)
	window.DrawLine(x,y,int(centerX2), int(centerY2), p.area)
}

func drawObjects(window draw.Window, indexes int) {
	rSunHalfr, rSunr := getRealRHR(rSunHalf,rSun)
	window.FillEllipse(int(centerX2)-rSunHalfr, int(centerY2)-rSunHalfr,
		rSunr,rSunr, draw.Yellow)
	if len(orbit[which])==0 {
		return
	}
	p:=orbit[which][indexes-1]
	x1, y1 := computeRealXY(&p)
	rHalfr, rr := getRealRHR(rHalf, r)
	window.FillEllipse(x1-rHalfr, y1-rHalfr, rr,rr,p.c)
	window.DrawLine(int(centerX2),int(centerY2), x1, y1, draw.White )
}

func drawSelect(window draw.Window) {
	text:= "unselect"
	c := draw.Red
	if isSelect {
		text="select"
		c = draw.Green
	}
	window.DrawScaledText(text,0,int(centerY/3),2,c)
	window.FillEllipse(0,int(centerY/3+50),30,30,c)
}

func drawLegend(window draw.Window) {
	if len(areasSelect[which])==0 {
		return
	}
	x, y := int(width-100.0), int(60)
	area := "area:"
	lX, lY := window.GetScaledTextSize(area,2)
	window.DrawScaledText(area, x-lX*2,y,2, draw.LightRed)
	y+=30
	for i := 1; i < len(areasSelect[which])+1; i++ {
		c := colors[(i-1)%len(colors)]
		window.DrawScaledText(strconv.Itoa(i),x, y, 2,c)
		text := formatArea(i-1)
		window.DrawScaledText(text, x-len(text)*20, y, 2, c )
		window.FillEllipse(x+20,y,30,30,c)
		y+=lY+5
	}
}

func formatArea(i int) string {
	return fmt.Sprintf("%.15f", areasSelect[which][i])
}

func formatData() string {
	temp :=float64(index)-1
	if temp==-1 {
		temp=float64(0)
	}
	r :=  fmt.Sprint("method: ",name[which], ", ", formatEuler(),
		fmt.Sprint(", step = ", stepAdd),
		fmt.Sprint(", time = ",
			fmt.Sprintf(fmt.Sprint("%.", digitToSee, "f"), temp*eu.Epsilon)),
	)
	return r
}


func DrawWarningEscapeSpeed(window draw.Window) {
	if eu.EscapeSpeed() && !eu.ForceCircularOrbit{
		text := "Warning: escape velocity!!"
		x,y:=window.GetScaledTextSize(text,4)
		window.DrawScaledText(text, int(centerX)-x/2,int(height)-y*2,
			4,draw.RGB(255.0/255.0,165.0/255.0,0))
	}
}