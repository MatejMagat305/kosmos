package main

import (
	"github.com/gonutz/prototype/draw"
)

func computeRealXY(p *position) (int,int) {
	x1:=   1*p.x*(centerY-float64(r))*eu.Scale + centerX2
	y1:=  -1*p.y*(centerY-float64(r))*eu.Scale + centerY2
	return int(x1), int(y1)
}

func ScaleControl(window draw.Window) {
	if window.WasKeyPressed(draw.KeyUp) && !Stop[0] {
		carryKlikPositionXY(0,1)
		return
	}
	if window.WasKeyPressed(draw.KeyLeft) && !Stop[1] {
		carryKlikPositionXY(1,0)
		return
	}
	if window.WasKeyPressed(draw.KeyDown) && !Stop[2] {
		carryKlikPositionXY(0,-1)
		return
	}
	if window.WasKeyPressed(draw.KeyRight) && !Stop[3] {
		carryKlikPositionXY(-1,0)
		return
	}
	if window.WasKeyPressed(draw.KeyM) {
		eu.Zoom(+1.0)
	}
	if window.WasKeyPressed(draw.KeyN) {
		eu.Zoom(-1.0)
	}
}

func carryKlikPositionXY(i int64, i2 int64) {
	if eu.FixedCentre {
		return
	}
	centerX2+=float64(i*eu.SpeedMovePixel)
	centerY2+=float64(i2*eu.SpeedMovePixel)
}

func getRealRHR(half int, rrr int) (int, int) {
	h := float64(half)*eu.Scale
	r := float64(rrr)*eu.Scale
	hh, rr := int(h), int(r)
	if hh==0 {
		return 1,2
	}
	if hh>half*3 {
		return half*3,rrr*3
	}
	return hh,rr
}

func CaryScale(indexes int) {
	if !eu.AutoScale|| eu.FixedScale {
		return
	}
		p := orbit[which][indexes-1]
		if needSmaller(&p) {
			eu.ZoomAutoScale(-1.0)
			return
		}
		if needBigger(&p) {
			eu.ZoomAutoScale(1.0)
			return
		}
}

func needBigger(p *position) bool {
	x, y := computeRealXY(p)
	vectorX := NewVector2D(float64(x)-centerX2,0)
	vectorY := NewVector2D(0,float64(y)-centerY2)
	if vectorX.size()<centerX*0.40 &&
		vectorY.size()<centerY*0.40 {
		return true
	}
	return false
}

func needSmaller(p *position) bool {
	x, y := computeRealXY(p)
	vectorX := NewVector2D(float64(x)-centerX2,0)
	vectorY := NewVector2D(0,float64(y)-centerY2)
	if vectorX.size()>centerX*0.90 ||
		vectorY.size()>centerY*0.90 {
		return true
	}
	return false
}