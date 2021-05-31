package euler

import (
	h "retardation/help"
	v "retardation/variables"
	"bytes"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
)

type Planet struct {
	Id                                     int     `json:"id"`
	PositionX                              float64 `json:"x"`
	PositionY                              float64 `json:"y"`
	VelocityX                              float64 `json:"v_x"`
	VelocityY                              float64 `json:"v_y"`
	Mass                                   float64 `json:"mass"`
	AccelerationX, AccelerationY float64 `json:"-"`
}

func NewPlanetFromString(row string) *Planet {
	ch := make(chan *Planet)
	go newPlanetFromString(ch, row)
	return <-ch
}

func newPlanetFromString(ch chan *Planet, rows string) {
	array := strings.Split(rows, ",")
	defer func() {
		if r := recover(); r != nil {
			v.Warning += "  warning: error at: " + rows
			ch <- newPlanetDefault()
		}
	}()
	result := newPlanetDefault()
	for i := 0; i < len(array); i++ {
		row := array[i]
		err := result.SetByName(row)
		if err != nil {
			v.AddWarning(err.Error())
		}
	}
	ch <- result
}

func newPlanetDefault() *Planet {
	return &Planet{
		Id:            num+1,
		PositionX:     0,
		PositionY:     0,
		VelocityX:     0,
		VelocityY:     0,
		AccelerationX: 0,
		AccelerationY: 0,
		Mass:          1,
	}
}
func (p *Planet) clone() *Planet {
	return &Planet{
		PositionX:     p.PositionX,
		PositionY:     p.PositionY,
		VelocityX:     p.VelocityX,
		VelocityY:     p.VelocityY,
		AccelerationX: p.AccelerationX,
		AccelerationY: p.AccelerationY,
		Mass:          p.Mass,
		Id:            p.Id,
	}
}

func NewPlanet(positionX, positionY, velocityX, velocityY, size float64) *Planet {
	return &Planet{
		Id:        num,
		PositionX: positionX,
		PositionY: positionY,
		VelocityX: velocityX,
		VelocityY: velocityY,
		Mass:      size,
	}
}

func (p Planet) IsIOverlaping(p2 *Planet) bool {
	vector := h.NewVector2DFrom4(p.PositionX, p.PositionY, p2.PositionX, p2.PositionY)
	if vector.Size() < math.Pow10(-15) {
		return true
	}
	return false
}

func (p *Planet) IsEqual(p2 *Planet) bool {
	return p.Id == p2.Id
}

// vector v_x*m, v_y*m
func (p *Planet) GetMomentumXY() (float64, float64) {
	return p.Mass * p.VelocityX, p.Mass * p.VelocityY
}

// v_x = v_x-V_x
// v_y = v_y-V_y
func (p *Planet) RefixVelocity(speedOfMassCenterX float64, speedOfMassCenterY float64) {
	p.VelocityX -= speedOfMassCenterX
	p.VelocityY -= speedOfMassCenterY
}

// m * v**2
func (p *Planet) GetM_VPow2() float64 {
	vPow2 := p.Mass * p.vPow2()
	return vPow2
}

// v**2
func (p *Planet) vPow2() float64 {
	vXPow2 := math.Pow(p.VelocityX, 2)
	vYPow2 := math.Pow(p.VelocityY, 2)
	return vXPow2 + vYPow2
}

// x*m, y*m
func (p *Planet) GetMassXY() (float64, float64) {
	return p.Mass * p.PositionX, p.Mass * p.PositionY
}

func (p *Planet) SetByName(row string) error {
	ch := make(chan error)
	go p.SetByNameCatchPanic(row, ch)
	return <-ch
}

func (p *Planet) SetByNameCatchPanic(row string, ch chan error) {
	defer h.RecoverSendErrorIfExist(ch)
	nameValue := strings.Split(row, "=")
	if len(nameValue) < 2 {
		ch <- fmt.Errorf("ignored value: " + row)
	}
	trimSpace := strings.TrimSpace(nameValue[1])
	value, err := strconv.ParseFloat(trimSpace, 64)
	if err != nil {
		ch <- fmt.Errorf("ignored value: " + row)
	}
	name0 := strings.TrimSpace(nameValue[0])
	name0 = fmt.Sprint(strings.ToUpper(name0[:1]), name0[1:])
	if name0 == "Id" {
		reflect.ValueOf(p).Elem().FieldByName(name0).SetInt(int64(value))
	} else {
		reflect.ValueOf(p).Elem().FieldByName(name0).SetFloat(value)
	}
	ch <- nil
}

func (p *Planet) DifferencePositionPlanet(planet2 *Planet) (float64, float64) {
	return planet2.PositionX-p.PositionX, planet2.PositionY - p.PositionY
}

func (p *Planet) DifferenceVelocityPlanet(planet2 *Planet) (float64, float64) {
	return planet2.VelocityX-p.VelocityX, planet2.VelocityY - p.VelocityY
}

func (p *Planet) FormatToBuf(buf *bytes.Buffer) {
	buf.WriteString("[ ")
	FormatToBufGeneral(p, buf, ", ")
	buf.WriteString("]\n")
}