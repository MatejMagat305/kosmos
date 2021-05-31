package euler

import (
	h "2bodyBinary/help"
	v "2bodyBinary/variables"
	"fmt"
	"math"
	"path/filepath"
)

const (
	Elliptic = iota + 1
	Hyperbolic
	Circle
	Parabolic
)

func (e *Euler) MakeChangesCalculateShapePeriodLenght() {
	eClone := e
	if e.Centerofmomentumframe == false {
		eClone = e.Clone()
		eClone.Centerofmomentumframe = true
	}
	eClone.checkWhetherCenterOfMomentumFrame()
	energy := eClone.calculateEnergy()
	e.State = eClone.GetStateNew(energy)
	e.Periodlenght = eClone.calculatePeriod(energy)
}


func (e *Euler) GetStateNew(energy float64) int {
	return e.getState(energy)
}

func (e *Euler) CalculatePeriod() float64 {
	energy := e.calculateEnergy()
	state := e.GetStateNew(energy)
	if state != Circle && state != Elliptic {
		return 0
	}
	return e.calculatePeriod(energy)
}

func getCloneEulerEnergyFromEuler(e *Euler) (float64, *Euler) {
	eClone := e.Clone()
	eClone.Centerofmomentumframe = true
	eClone.checkWhetherCenterOfMomentumFrame()
	energy := eClone.calculateEnergy()
	eClone.State = eClone.GetStateNew(energy)
	e.State=eClone.State
	return energy, eClone
}

func SetState() {
	_,_=getCloneEulerEnergyFromEuler(MEuler)
}

func (e *Euler) checkWhetherCenterOfMomentumFrame() {
	if e.Centerofmomentumframe == false {
		return
	}
	e.Planets.MakeCenterOfMomentumFrame()
	e.SetCenterOfMassZero()
}

func (e *Euler) GetState() int {
	return e.State
}

func (e *Euler) getState(energy float64) int {
	se := e.Epsilon*0.01
	if energy < 0{
		/*if e.IsCircle() {
			return Circle
		}*/
		return Elliptic
	}
	if energy <= se {
		return Parabolic
	}
	return Hyperbolic
}

func (e *Euler) IsCircle() bool {
	planet1, planet2 := e.GetPlanets()
	vectorPosition := h.NewVector2DFrom4(planet1.PositionX, planet1.PositionY, planet2.PositionX, planet2.PositionY)
	vectorVelocity1, vectorVelocity2 := h.NewVector2DFrom2(planet1.VelocityX*planet1.Mass, planet1.VelocityY*planet1.Mass),
		h.NewVector2DFrom2(planet2.VelocityX*planet2.Mass, planet2.PositionY*planet2.Mass)
	rightSize1, rightSize2 := e.getEnergyCircle(v.FirstElement, 1), e.getEnergyCircle(v.SecondElement, 1)
	accurency := e.Epsilon*0.0001
	return vectorPosition.IsPerpendicular(vectorVelocity1) &&
		vectorPosition.IsPerpendicular(vectorVelocity2) &&
		vectorVelocity1.Size()-rightSize1<accurency &&
		vectorVelocity2.Size()-rightSize2<accurency
}

func (e *Euler) calculatePeriod(energy float64) float64 {
	if e.State == Parabolic || e.State == Hyperbolic {
		return 0
	}
	a := e.computeMainSemiEliptic(energy)
	Pi2_4_a3 := e.Pi2_4_a3(a)
	GM_1M_2 := e.GM_1M_2()
	return math.Sqrt(Pi2_4_a3 / GM_1M_2)
}

func (e *Euler) SetCenterOfMassZero() {
	planet1, planet2 := e.GetPlanets()
	centerX, centerY := e.averagePosition(planet1.PositionX, planet2.PositionX),
		e.averagePosition(planet1.PositionY, planet2.PositionY)
	planet1.SubPosition(centerX, centerY)
	planet2.SubPosition(centerX,centerY)
}
// p_x - x
// p_y - y
func (p *Planet) SubPosition(x float64, y float64) {
	p.PositionX -= x
	p.PositionY -= y
}
// (position_1 * Mass_1 + position_2 * Mass_2)/(Mass_1 + Mass_2)
func (e *Euler) averagePosition(Position1 float64, Position2 float64) float64 {
	planet1, planet2 := e.GetPlanets()
	return (Position1*planet1.Mass+Position2*planet2.Mass)/(planet1.Mass + planet2.Mass)
}

// G * (m_1+m_2)
func (e *Euler) GM_1M_2() float64 {
	planet1, planet2 := e.GetPlanets()
	var m1, m2 = planet1.Mass, planet2.Mass
	return e.G * (m1 + m2)
}

// 4 * PI**2 * a**3
func (e *Euler) Pi2_4_a3(a float64) float64 {
	pi2, a3 := math.Pow(math.Pi, 2), math.Pow(a, 3)
	return 4 * pi2 * a3
}

// G * m_1 * m_2 / (2 * |E| )
func (e *Euler) computeMainSemiEliptic(energy float64) float64 {
	planet1, planet2 := e.GetPlanets()
	g, m1, m2 := e.G, planet1.Mass, planet2.Mass
	absEnergyBy2 := 2 * math.Abs(energy)
	return g * m1 * m2 / absEnergyBy2

}

// G * m_1 * m_2 * (1/|r|)
func (e *Euler) GetGMMR() float64 {
	planet1, planet2 := e.GetPlanets()
	g, m1, m2 := e.G, planet1.Mass, planet2.Mass
	reversedR := reverseR(planet1, planet2)
	return g * m1 * m2 * reversedR
}

func reverseR(planet *Planet, planet2 *Planet) float64 {
	return 1 / R(planet, planet2)
}

// R ((x1-x2)**2 + (y1-y2)**2)**(1/2)
func R(planet *Planet, planet2 *Planet) float64 {
	x, y := planet.PositionX, planet.PositionY
	x2, y2 := planet2.PositionX, planet2.PositionY
	return math.Sqrt(
		math.Pow(x-x2, 2) +
			math.Pow(y-y2, 2))
}


func GetAnalyticOrbits() ([]float64, []float64, error) {
	e, err := GetEulerFromName(filepath.Join("./bin", MEuler.Name, v.EulerOriginConfig))
	if err != nil {
		return nil, nil, fmt.Errorf("unreache bin")
	}
	e.Centerofmomentumframe = true
	e.MakeChangesCalculateShapePeriodLenght()
	planet1, planet2 := e.GetPlanets()
	psi, eps, _, p, reduceMass := e.getPsi()
	forX, forY := prepareXFromFiPEps(p, eps, psi), prepareYFromFiPEps(p, eps, psi)
	convertPlanet1, convertPlanet2 := reduceMass/planet1.Mass, reduceMass/planet2.Mass
	resultPlanet1, resultPlanet2 :=  make([]float64,0,10000), make([]float64,0,10000)

	for angle := 0.; angle < math.Pi*2; angle+=0.00001 {
		momentumX, momentumY := forX(angle), forY(angle)
		x1, y1 :=  momentumX*convertPlanet1, momentumY*convertPlanet1
		x2, y2 :=  -momentumX*convertPlanet2, -momentumY*convertPlanet2
		resultPlanet1 = append(append(resultPlanet1, x1), y1)
		resultPlanet2 = append(append(resultPlanet2, x2),y2 )
	}
	return resultPlanet1,resultPlanet2, nil
}

func (e *Euler)getPsi() (float64, float64, float64, float64, float64) {
	planet1, planet2 := e.GetPlanets()
	Mass1,Mass2:= planet1.Mass, planet2.Mass
	PositionX1,PositionX2 := planet1.PositionX, planet2.PositionX
	PositionY1,PositionY2 := planet1.PositionY, planet2.PositionY
	VelocityX1 , VelocityX2:= planet1.VelocityX, planet2.VelocityX
	VelocityY1, VelocityY2 := planet1.VelocityY, planet2.VelocityY
	a := e.G*planet1.Mass*planet2.Mass  // pot.en. fiktivneho telesa: U = -a/r
	// vypocitame rychlost taziska sustavy
	Vx := (Mass1*VelocityX1 + Mass2*VelocityX2)/(Mass1+Mass2);
	Vy := (Mass1*VelocityY1 + Mass2*VelocityY2)/(Mass1+Mass2);
	// v zadanych rychlostiach prejdeme do taziskovej sustavy
	vx1 := VelocityX1 - Vx; vy1 := VelocityY1 -Vy;
	vx2 := VelocityX2 - Vx; vy2 := VelocityY2 -Vy;
	//	// parametre fiktivneho telesa
	mu := (planet1.Mass * planet2.Mass) / (planet1.Mass + planet2.Mass)
	r := math.Sqrt(math.Pow(PositionX1-PositionX2,2)+math.Pow(PositionY1-PositionY2,2)); // poloha
	L := mu*((vy1-vy2)*(PositionX1-PositionX2) - (PositionY1-PositionY2)*(vx1-vx2));    // moment hybnosti (zlozka z)
	// konstantny vektor Laplace-Runge_Lenz [Landau I (15,17)]  v x L - a \vec {r} /r       (L ma len z-zlozku)
	vecx := (vy1-vy2)*L - a*(PositionX1-PositionX2)/r;
	vecy := -(vx1-vx2)*L - a*(PositionY1-PositionY2)/r;
	// pocitame trajektoriu fiktivneho telesa
	p := math.Pow(L,2)/(a*mu)
	vv := math.Sqrt(math.Pow(vx1-vx2,2)+math.Pow(vy1-vy2,2))  // rychlost
	totalE := mu*math.Pow(vv,2)/2 - a/r
	eps := math.Sqrt(math.Abs(1+2*totalE*math.Pow(L,2)/(mu*math.Pow(a,2))))
	psi := math.Atan2(vecy,vecx)  // vracia hodnotu v intervale (-pi, pi)
	GMM := e.G*planet1.Mass*planet2.Mass
	a2:= math.Pow(GMM,2)*mu
	en2 := 2*math.Pow(L,2)*totalE
	eps = math.Sqrt(math.Max(en2/a2+1, 0))
	return psi, eps, totalE, p,mu
}

// prepare fi -> X
func prepareXFromFiPEps(p, eps, psi float64) func(fi float64) float64 {
	return func(fi float64) float64 {
		return (p * math.Cos(fi+psi)) /
			(1 + eps * math.Cos(fi))
	}
}

// prepare fi -> Y
func prepareYFromFiPEps(p, eps, psi float64) func(fi float64) float64 {
	return func(fi float64) float64 {
		return (p * math.Sin(fi+psi)) /
			(1 + eps * math.Cos(fi))
	}
}

// -a/r
func (e *Euler) vr(a float64) float64 {
	r := R(e.GetPlanets())
	return -a / r
}

// G * m_1 * m_2
func (e *Euler) a() float64 {
	planet1, planet2 := e.GetPlanets()
	return e.G * planet1.Mass * planet2.Mass
}

// redukcion mass - m1*m2/(m1+m2)
func redukcionMassFromPlanet(planet, planet2 *Planet) float64 {
	return planet.Mass * planet2.Mass / (planet.Mass + planet2.Mass)
}

func (e *Euler) calculateEnergy() float64 {
	_, _, totalE, _, _ := e.getPsi()
	return totalE
}
