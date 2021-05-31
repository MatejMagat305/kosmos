package main

import (
	"math"
	"strings"
)

func AddChoose() {
	checkBool()
	checkInt(1)
	checkFloat(1.0)
}

func SubChose() {
	checkBool()
	checkInt(-1)
	checkFloat(-1.0)
}

func checkBool() {
	n := NameAtributEuler[WhichName]
	if eu.IsBoolType(n) {
		ReflE.FieldByName(n).SetBool(!ReflE.FieldByName(n).Bool())
	}
}

func checkInt(i int) {
	n := NameAtributEuler[WhichName]
	if eu.IsIntegerType(n) {
		e := ReflE.FieldByName(n).Int()
		x :=0
		ee := e
		for e>=10 {
			x++
			e/=10
		}
		k := int64(i)*int64(math.Pow10(x))
		if -k==ee {
			k/=10
		}
		ee = ee+k
		if ee<=1 {
			ee=1
		}
		if ee>=100000000 {
			ee = 100000000
		}
		ReflE.FieldByName(n).SetInt(ee)
	}
}


func checkFloat(f float64) {
	n := NameAtributEuler[WhichName]
	if !eu.IsIntegerType(n)&& !eu.IsBoolType(n) {
		if strings.EqualFold(n,"Epsilon") {
			if f>0 {
				epsilonAdd()
				return
			}
			epsilonSub()
			return
		}
	}
}

func IsImutable(n string) bool {
	return ImutableType[n]
}
