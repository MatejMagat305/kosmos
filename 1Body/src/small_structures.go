package main

import (
	"github.com/gonutz/prototype/draw"
	"math"
)

type position struct {
	x, y float64
	isSelect, isFake bool
	whichSelect int
	howMany int
	c, area draw.Color
}

func (p position) IsFake() bool {
	return p.isFake
}

type calcul = func( chan position, chan bool, int )

func newZeroPosition() *position {
	return &position{
		x:        0,
		y:        0,
		isSelect: false,
		c:        draw.Color{},
		area:     draw.Color{},
	}
}

type tuple struct {
	areaFirst float64
	areas []float64
	wasVerified,was, verify bool

}

func (t *tuple) doVerify() {
	t.wasVerified=true
	last := t.areas[len(t.areas)-1]
	if math.Abs(t.areaFirst-last)<=akceptSize {
		t.verify=true
	}
}
