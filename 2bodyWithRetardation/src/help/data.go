package help

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/rand"
	sh "retardation/draw_control/shapes"
	"time"
	"unsafe"
)

type FormData struct {
	Data []*sh.TextField
}

type FloatData struct {
	PositionsXY, VelocityXY []float64
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

const (
	bigCapacity = 1 << 31
	systemError = "ops system error, I do not know save"
)
// change pointer to float64
func BytePointerToFloat64Pointer(src []byte) []float64 {
	if len(src) == 0 {
		return nil
	}
	l := len(src) / 8
	ptr := unsafe.Pointer(&src[0])
	result := ((*[bigCapacity]float64)(ptr))[:l:l]
	return result
}

func GetBytesPosition(x, y float64) ([]byte, []byte, error) {
	var buf1, buf2 bytes.Buffer
	err1 := binary.Write(&buf1, binary.LittleEndian, x)
	err2 := binary.Write(&buf2, binary.LittleEndian, y)
	if err2 != nil || err1 != nil {
		return nil, nil, fmt.Errorf("%v, %v", err1, err2)
	}
	return buf1.Bytes(), buf2.Bytes(), nil
}

func RandomFromMinusToPlusOne() float64 {
	return RandomFromZeroToOne()*2-1
}
func RandomFromZeroToOne() float64 {
	return rand.Float64()
}
// change pointer to Byte
func Float64PointerToBytePointer(src []float64) []byte {
	if len(src) == 0 {
		return nil
	}
	l := len(src) * 8
	ptr := unsafe.Pointer(&src[0])
	result := ((*[bigCapacity*5]byte)(ptr))[:l:l]
	return result
}

func GetSytemError() string {
	return systemError
}

func DoNotFound(path string) string {
	return fmt.Sprintf("the flile %s is not available, if you want save current state klick 'save'", path)
}