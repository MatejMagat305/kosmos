package euler

import (
	h "2bodyBinary/help"
	v "2bodyBinary/variables"
	"math"
)

func (e *Euler)GenerateOrbitByType(orbitType int) {
	switch orbitType {
	case Parabolic:
		e.generateParabolicOrbit()
	case Circle:
		e.generateCircleOrbit()
		orbitType = Elliptic
	default:
		e.generateEllipticHyperbilicOrbit(getCondition(orbitType))
	}
	e.SetCenterOfMassZero()
	e.State=orbitType
	e.SetName(GetNameOrbit(e))
}

func (e *Euler) generateEllipticHyperbilicOrbit(condition func(energy float64) bool, n float64) {
	for {
		e.generateRandomPositions()
		e.generateRandomVelocity(n)
		energy, _ := getCloneEulerEnergyFromEuler(e)
		if condition(energy) {
			return
		}
	}
}

func getCondition(orbitType int) (func(energy float64) bool,  float64 ) {
	if orbitType==Hyperbolic {
		return func(energy float64) bool {
			return energy>0 && energy!=math.NaN() && energy!=math.Inf(1)
		}, 5
	}
	return func(energy float64) bool {
		return energy<0 && energy!=math.Inf(-1) && energy!=math.NaN()
	},1
}

func (e *Euler)generateParabolicOrbit() {
	for{
		e.generateRandomPositions()
		e.RandomPlanetsVelocityParabolic()

		energy := e.calculateEnergy()
		//fmt.Println(energy)
		if e.GetStateNew(energy)==Parabolic {
			return
		}
	}
}

func (e *Euler) RandomPlanetsVelocityParabolic() {
	planet1, planet2 := e.Planets[v.FirstElement], e.Planets[v.SecondElement]
	sizeVector := (2/(1+planet1.Mass/planet2.Mass))*e.G*planet2.Mass*reverseR(planet1, planet2)
	v1XPow2 := sizeVector*h.RandomFromZeroToOne()
	v1YPow2 := sizeVector - v1XPow2
	planet1.VelocityX = math.Sqrt(v1XPow2)
	planet1.VelocityY = math.Sqrt(v1YPow2)

	planet2.VelocityX = -planet1.Mass*planet1.VelocityX/planet2.Mass
	planet2.VelocityY = -planet1.Mass*planet1.VelocityY/planet2.Mass

}

func (e *Euler) getEnergyNeedToSecond() float64 {
	gMassMass := e.GetGMMR()
	e1 := 0.5 * e.Planets[v.FirstElement].GetM_VPow2()
	return((gMassMass - e1)*2)/e.Planets[v.SecondElement].Mass
}

func (e *Euler) generateCircleOrbit() {
	planet1 := e.Planets[v.FirstElement]
	planet1.PositionX = (h.RandomFromMinusToPlusOne()*1)/2
	e.copyMinusXToSecondSetZeroToBothY()
	e.setVelocity(v.FirstElement,1)
	e.setVelocity(v.SecondElement,-1)
	e.randomSwapXYVelocityXVelocityY()
}

func (e *Euler) randomSwapXYVelocityXVelocityY() {
	if h.RandomFromZeroToOne()>0.5 {
		planet1, planet2 := e.Planets[v.FirstElement], e.Planets[v.SecondElement]
		swapXYVelocityXVelocityYOnePlanet(planet1)
		swapXYVelocityXVelocityYOnePlanet(planet2)
	}
}

func swapXYVelocityXVelocityYOnePlanet(planet *Planet) {
	x,y,vX,vY := planet.PositionX, planet.PositionY, planet.VelocityX, planet.VelocityY
	planet.PositionX, planet.PositionY, planet.VelocityX, planet.VelocityY = y,x,vY,vX
}

func (e *Euler) setVelocity(element int, sign float64) {
	planet1 := e.Planets[element]
	energy := e.getEnergyCircle(element, sign)
	planet1.VelocityY = energy
	planet1.VelocityX=0
}

func (e Euler) getEnergyCircle(element int, sign float64) float64 {
	planet1, planet2 := e.Planets[element], e.Planets[v.SecondElement-element]
	GM2 := e.Features.G * math.Pow(planet2.Mass, 2)
	r := R(planet1, planet2)
	MM := planet1.Mass+planet2.Mass
	return sign * math.Sqrt((GM2)/(r*MM))
}

func (e *Euler) copyMinusXToSecondSetZeroToBothY() {
	planet1, planet2 := e.Planets[v.FirstElement], e.Planets[v.SecondElement]
	planet2.PositionX=-(planet1.PositionX*planet1.Mass)
	planet1.PositionX/=planet2.Mass
	if planet1.PositionX < planet2.PositionX  {
		planet2,planet1 = planet1,planet2
	}
	for planet1.PositionX>0.3  || planet2.PositionX<(-0.3) {
		planet1.PositionX/=1.1
		planet2.PositionX/=1.1
	}
	planet2.PositionY=0
	planet1.PositionY=0
}

func (e *Euler)generateRandomPositions() {
	e.setOneRandomPositionPlanetByIndex(v.FirstElement)
	e.setOneRandomPositionPlanetByIndex(v.SecondElement)
}

func (e *Euler) setOneRandomPositionPlanetByIndex(index int) {
	planet := e.Planets[index]
	planet.PositionX = h.RandomFromMinusToPlusOne()
	planet.PositionY = h.RandomFromMinusToPlusOne()
}
func (e *Euler) generateRandomVelocity(scale float64) {
	e.setOneRandomVelocityPlanetByIndex(v.FirstElement, scale)
	e.setOneRandomVelocityPlanetByIndex(v.SecondElement, scale)

}

func (e *Euler) setOneRandomVelocityPlanetByIndex(index int, scale float64) {
	planet := e.Planets[index]
	planet.VelocityX = h.RandomFromMinusToPlusOne() * scale * planet.Mass
	planet.VelocityY = h.RandomFromMinusToPlusOne() * scale * planet.Mass

}