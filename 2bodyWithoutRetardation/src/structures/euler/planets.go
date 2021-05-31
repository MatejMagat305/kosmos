package euler

import (
	h "2bodyBinary/help"
	"bytes"
)

type Planets []*Planet

func (p Planets) clone() Planets {
	result := make(Planets, 0, len(p))
	for i := 0; i < len(p); i++ {
		result = append(result, p[i].clone())
	}
	return result
}

func (p Planets) IsEqualPositision() bool {
	if len(p) <= 1 {
		return true
	}
	for i := 0; i < len(p); i++ {
		for j := 1 + i; j < len(p); j++ {
			if p[i].IsIOverlaping(p[j]) {
				return true
			}
		}
	}
	return false
}

func NewDefaultPlanets() Planets {
	return Planets{
		NewPlanet(defaultPositionX, defaultPositionY, defaultVelocityX, defaultVelocityY, defaultSize),
		NewPlanet(-defaultPositionX, -defaultPositionY, -defaultVelocityX, -defaultVelocityY, defaultSize)}
}

// V_x = v_x_1 * m_1 + v_x_2 * m_2
// V_y = v_y_1 * m_1 + v_y_2 * m_2
// v_x_1 = v_x_1-V_x, v_y_1 = v_y_1-V_y
// v_x_2 = v_x_2-V_x, v_y_2 = v_y-V_y
func (p Planets) MakeCenterOfMomentumFrame() {
	speedOfMassCenterX, speedOfMassCenterY := p.computeSpeedOfMassCenter()
	momentum0 := p.computeLocationOfMassCenter()
	momentum = momentum0
	p[0].RefixVelocity(speedOfMassCenterX, speedOfMassCenterY)
	p[1].RefixVelocity(speedOfMassCenterX, speedOfMassCenterY)
}

// (v_x_1*mass_1+v_x_2*mass_2)/(mass_1+mass_2)
// (v_y_1*mass_1+v_y_2*mass_2)/(mass_1+mass_2)
func (p Planets) computeSpeedOfMassCenter() (resultVX float64, resultVY float64) {
	resultVX, resultVY = 0, 0
	vX1, vY1 := p[0].GetMomentumXY()
	vX2, vY2 := p[1].GetMomentumXY()
	mass := p.SumMass()
	resultVX = (vX1 + vX2) / mass
	resultVY = (vY1 + vY2) / mass
	return
}

// (x1*mass-x2)/(mass1+mass2), (y1+y2)/(mass1+mass2)
func (p Planets) computeLocationOfMassCenter() *h.Momentum {
	result := &h.Momentum{
		X: 0,
		Y: 0,
	}
	x1, y1 := p[0].GetMassXY()
	x2, y2 := p[1].GetMassXY()
	sumMass := p.SumMass()
	result.X = (x1 + x2) / sumMass
	result.Y = (y1 + y2) / sumMass
	return result
}

func (p Planets) SumMass() float64 {
	var result float64 = 0
	for i := 0; i < len(p); i++ {
		result += p[i].Mass
	}
	return result
}

func (p Planets) FormatPlanetsToBuf(buf *bytes.Buffer) {
	p[0].FormatToBuf(buf)
	p[1].FormatToBuf(buf)
}
