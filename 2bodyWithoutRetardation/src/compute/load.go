package compute

import (
	eu "2bodyBinary/structures/euler"
	v "2bodyBinary/variables"
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
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
	ChDone <- NumPeriod - 1
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