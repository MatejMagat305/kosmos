package euler

import (
	"bytes"
	"fmt"
	"reflect"
	v "retardation/variables"
	"strings"
)

type Features struct {
	Epsilon               float64 `json:"epsilon"`
	Centerofmomentumframe bool    `json:"centerOfMomentumFrame"`
	Beginsuitablescale    bool    `json:"beginSuitableScale"`
	Scale                 float64 `json:"scale"`
	G                     float64 `json:"gravitationalConstant"`
	Name                  string  `json:"name"`
	Offset                int     `json:"offset"`
	Howmanyturn                int     `json:"How many turn next"`
	C                     float64 `json:"speed of light"`
	Method                int     `json:"method"`
	Periodlenght          float64 `json:"_periodLenght"`
	State                 int     `json:"_state"`
	MoveX                 float64 `json:"_moveX"`
	MoveY                 float64 `json:"_moveY"`
}

func (feature *Features) clone() *Features {
	return &Features{
		Epsilon:               feature.Epsilon,
		Centerofmomentumframe: feature.Centerofmomentumframe,
		Beginsuitablescale:    feature.Beginsuitablescale,
		Scale:                 feature.Scale,
		G:                     feature.G,
		Name:                  feature.Name,
		Offset:                feature.Offset,
		C:                     feature.C,
		Method:                feature.Method,
		Periodlenght:          feature.Periodlenght,
		State:                 feature.State,
		MoveX:                 feature.MoveX,
		MoveY:                 feature.MoveY,
		Howmanyturn: feature.Howmanyturn,
	}
}
func newDefaultMutable() *Features {
	var (
		Epsilon               = 0.0001
		Scale                 = 3.00
		gravitationalConstant = 2.9591220823e-4
		C                     = 3e08
	)
	DigitToSee = 4
	Action = Epsilon
	return &Features{
		Epsilon:               Epsilon,
		Centerofmomentumframe: false,
		Scale:                 Scale,
		G:                     gravitationalConstant,
		C:                     C,
		Method:                v.FirstMethod,
		Howmanyturn: 1,
	}
}

func (feature *Features) SetName(s string) {
	feature.Name = s
}

func (feature *Features) FormatFeaturesToBuf(buf *bytes.Buffer) {
	FormatToBufGeneral(feature, buf, "\n")
	buf.WriteString("\n")
}

func FormatToBufGeneral(object interface{}, buf *bytes.Buffer, separator string) {
	s := reflect.ValueOf(object).Elem()
	typeOfT := s.Type()
	arrayNameValue := make([]string, 0, s.NumField())
	for i := 0; i < s.NumField(); i++ {
		field := typeOfT.Field(i)
		jsonTag := field.Tag.Get("json")
		name := field.Name
		if jsonTag == "" || jsonTag == "-" || jsonTag == "id" || strings.HasPrefix(jsonTag, "_") {
			continue
		}
		f := s.Field(i)
		str := fmt.Sprintf("%s = %v", name, f.Interface())
		arrayNameValue = append(arrayNameValue, str)
	}
	buf.WriteString(strings.Join(arrayNameValue, separator))
}
