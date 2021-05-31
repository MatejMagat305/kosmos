package compute

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	h "retardation/help"
	eu "retardation/structures/euler"
	v "retardation/variables"
)

var validName = regexp.MustCompile(`planet[0-9]+_period_[0-9]+\.bin`)

func loadIfIsloadExist() {
	if !v.IsLoad {
		return
	}
	euler, err := eu.GetEulerFromName(filepath.Join(dirName, v.EulerConfig))
	if err != nil {
		LoadOrigin()
		return
	}
	localE = euler
	n := 0
	ff := func(path string, info os.FileInfo, err error) error {
		if info == nil {
			return err
		}
		if validName.MatchString(info.Name()) {
			n++
		}
		return nil
	}
	_ = filepath.Walk(dirName, ff)
	NumPeriod = n/2 + 1
	loadNeed()
	ChDone <- NumPeriod - 1
}

func loadNeed() {
	go func() {loadPrevVelocity(PlanetData1, 1); chOut1<-true}()
	go func() {loadPrevVelocity(PlanetData2, 2); chOut1<-true}()
	<-chOut1;<-chOut1
}

func loadPrevVelocity(data *h.FloatData, i int) {
	nameDir := fmt.Sprintf(neddVelocity, dirName, i)
	b, err := ioutil.ReadFile(nameDir)
	if err != nil {
		return
	}
	data.VelocityXY = append(data.VelocityXY, h.BytePointerToFloat64Pointer(b)...)
}

func LoadOrigin() {
	eur, err := eu.GetEulerFromName(filepath.Join(dirName, v.EulerOriginConfig))
	if err != nil {
		return
	}
	eu.SetEuler(eur)
	eu.SetState()
}
func SaveOrigin(fileEuler string, euler *eu.Euler) {
	Create()
	jsonEuler, _ := json.Marshal(&euler)
	_ = ioutil.WriteFile(fileEuler, jsonEuler, 7777)
}