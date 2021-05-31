package shapes

import "github.com/gonutz/prototype/draw"

var (
	HideNames = Names(make(map[string]bool))
)

type Basic struct {
	x, y, endX, endY, width, height float64
	id                              int64
	name                            string
	pressed, active                 bool
	function                        func(draw.Window)
}

type Shape interface {
	GetXY() (float64, float64)
	SetXY(x, y float64)

	GetSize() (float64, float64)
	SetSize(width, height float64)

	GetIdName() (int64, string)

	WasPressed() bool
	WasActive() bool
	In(x int, y int) bool

	GetFunc() func(window draw.Window)
	SetFunc(function func(window draw.Window))

	Paint(window draw.Window, scaledText float32)
	CarryEvent(window draw.Window)
}

var (
	id int64 = 0
)

func resetId() {
	id = 0
}

func MakeEmptyShapesArrayWithCapResetId(howMany int) []Shape {
	resetId()
	return make([]Shape, 0, howMany)
}
func (s *Basic) WasActive() bool {
	return s.active
}

func newShape(name string) Basic {
	id++
	return Basic{name: name, id: id}
}
func (s *Basic) GetFunc() func(draw.Window) {
	return s.function
}
func (s *Basic) WasPressed() bool {
	return s.pressed
}

func (s *Basic) GetXY() (float64, float64) {
	return s.x, s.y
}

func (s *Basic) SetXY(x, y float64) {
	s.x = x
	s.y = y
}

func (s *Basic) GetSize() (float64, float64) {
	return s.width, s.height
}

func (s *Basic) SetSize(width, height float64) {
	s.width = width
	s.height = height
	s.endX = s.x + width
	s.endY = s.y + height
}

func (s *Basic) GetIdName() (int64, string) {
	return s.id, s.name
}
func (s *Basic) SetPressed(b bool) {
	s.pressed = b
}

func (s *Basic) SetActive(b bool) {
	s.active = b
}

func (s *Basic) SetFunc(function func(window draw.Window)) {
	s.function = function
}

func (s *Basic) IsActive() bool {
	return s.active
}

func (s *Basic) In(x int, y int) bool {
	floatX, floatY := float64(x), float64(y)
	return floatX >= s.x &&
		floatX <= s.endX &&
		floatY >= s.y &&
		floatY <= s.endY
}
func (s *Basic) CarryEvent(window draw.Window) {}

func (s *Basic) Paint(window draw.Window, scaledText float32) {
	if HideNames.IsIn(s.name) {
		return
	}
	x, y := s.GetXY()
	width, height := s.GetSize()
	window.DrawRect(int(x), int(y), int(width), int(height), draw.LightRed)
}
