package euler

import (
	h "2bodyBinary/help"
	v "2bodyBinary/variables"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"
)

func InitEuler(configString string) {
	ch := make(chan *Euler)
	initJsonNameToRealAndContrast()
	MEuler = NewEulerDefault()
	MetaMEuler = reflect.ValueOf(MEuler).Elem()
	go ReadEulerFromFile(ch, configString)
	MEuler = <-ch
	MetaMEuler = reflect.ValueOf(MEuler).Elem()
	if h.FileExist(fmt.Sprint("./bin/", MEuler.Name )) {
		SetFree()
	}
	if len(MEuler.Planets)<2 {
		MEuler.Planets = NewDefaultPlanets()
	}

}

func SetFree() {
	filterString := regexp.MustCompile("\\([0-9]\\)").ReplaceAllString(MEuler.Name, "")
	newLocation := filterString
	binLocation := fmt.Sprint("./bin/",newLocation)
	if !h.FileExist(binLocation) || !h.FileExist(fmt.Sprint(binLocation,"/", v.EulerConfig)) {
		MEuler.Name=newLocation
		return
	}
	i := 1
	for ;; i++ {
		newLocation = fmt.Sprint(filterString,"(",i,")")
		binLocation = fmt.Sprint("./bin/",newLocation)
		if !h.FileExist(binLocation) || !h.FileExist(fmt.Sprint(binLocation,"/", v.EulerConfig)) {
			MEuler.Name=newLocation
			break
		}
	}
}

func ReadEulerFromFile(ch chan *Euler, configString string) {
	c, err := ioutil.ReadFile(configString)
	defer func() {
		if r := recover(); r != nil {
			v.AddWarning(fmt.Sprint(r))
			ch <- NewEulerDefault()
		}
	}()
	if err != nil {
		panic("file config.txt can not open, program will run with default value")
	}
	str := string(c)
	array := strings.Split(str, fmt.Sprintln())
	r, _ := regexp.Compile("\\[.*]")
	num = 0
	for i := 0; i < len(array); i++ {
		row := array[i]
		if r.Match([]byte(row)) {
			filterString := r.FindString(row)
			MEuler.AddPlanet(removeParantess(filterString))
			continue
		}
		MEuler.SetValueCatchPanic(row, i)
	}
	if MEuler.IsEqualPositision() {
		v.Warning = "the load values were only zero or objects were too close, program run by default"
		MEuler.Planets = NewDefaultPlanets()
	}
	carryEpsilon(MEuler)
	ch <- MEuler
}

func removeParantess(filterString string) string {
	return filterString[1:][:len(filterString)-2]
}

func (e *Euler) SetValueCatchPanic(row string, i int) {
	defer func() {
		if r := recover(); r != nil {
			v1 := fmt.Sprint("row ", i+1, ": ")
			v2 := fmt.Sprint(" was ignored(", r, ")")
			W := fmt.Sprint(v1, "\"", strings.TrimSpace(row), "\"", v2)
			v.AddWarning(W)
		}
	}()
	e.SetValue(strings.ToLower(row))
}
func carrySeeDigit(name1 string, name0 string, result *Euler) {
	if strings.EqualFold(name0, "Epsilon") {
		carryEpsilon(result)
	}
}

func carryEpsilon(result *Euler) {
	if result.Epsilon < math.Pow(10, -10) || result.Epsilon > 1.00 {
		v.AddWarning("epsilon was too small or too big, it is set to default")
		result.Epsilon = epsilon
		HowMany = 1
		DigitToSee = 3
		IsEpsDefault = true

	}
}

func (e *Euler) SetValue(row string) {
	nameValue := strings.Split(row, "=")
	if len(nameValue) < 2 {
		panic("wrong structure")
	}
	name0 := strings.TrimSpace(nameValue[0])
	value := strings.TrimSpace(nameValue[1])
	name0 = fmt.Sprint(strings.ToUpper(name0[:1]), name0[1:])
	if e.IsBoolType(name0) {
		e.SetBool(value, name0)
		return
	}
	if e.IsStringType(name0) {
		e.SetString(value, name0)
		return
	}
	if e.IsInteger64Type(name0) {
		e.SetInteger64(value, name0)
		return
	}
	if e.IsIntegerType(name0) {
		e.SetInteger(value, name0)
		return
	}
	e.SetFloat(value, name0)
}
func initJsonNameToRealAndContrast() {
	jsonToNameEuler, nameToJsonEuler, jsonToNamePlanet, nameToJsonPlanet = make(map[string]string), make(map[string]string), make(map[string]string), make(map[string]string)
	initJsonNameToRealAndContrastFeatures()
	initJsonNameToRealAndContrastPlanet()
}

func initJsonNameToRealAndContrastFeatures() {
	metaFeatures := reflect.Indirect(reflect.ValueOf(Features{}))
	initJsonNameToRealAndContrastBy(metaFeatures, jsonToNameEuler, nameToJsonEuler)
}

func initJsonNameToRealAndContrastPlanet() {
	metaPlanet := reflect.Indirect(reflect.ValueOf(Planet{}))
	initJsonNameToRealAndContrastBy(metaPlanet, jsonToNamePlanet, nameToJsonPlanet)
}

func GetEulerFromName(config string) (*Euler, error) {
	b, e := ioutil.ReadFile(config)
	if e != nil {
		return nil, e
	}
	var euler Euler
	err := json.Unmarshal(b, &euler)
	if err != nil {
		fmt.Println(e)
		return nil, e
	}
	return &euler, nil
}

func RmOldBinByName(name string) {
	dir := fmt.Sprint("./bin/", name)
	e := os.RemoveAll(dir)
	if e != nil {
		fmt.Println(e)
	}
}

func initJsonNameToRealAndContrastBy(meta reflect.Value, jName, nJson map[string]string) {
	for i := 0; i < meta.NumField(); i++ {
		metaField := meta.Type().Field(i)
		jsonTag := metaField.Tag.Get("json")
		name := metaField.Name
		if jsonTag != "" && jsonTag != "-" && ! strings.HasPrefix(jsonTag, "_") {
			jName[jsonTag] = name
			nJson[name] = jsonTag
		}
	}
}

func LoadOriginMEuler() {
	nameDir := MEuler.Name
	euler, err := GetEulerFromName(filepath.Join("./bin", nameDir, "origin_euler.json"))
	if err==nil {
		MEuler=euler
		RmOldBinByName(nameDir)
	}
}
