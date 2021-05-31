package euler

import (
	h "retardation/help"
	v "retardation/variables"
	"bytes"
	"fmt"
	"os"
	"regexp"
)

func (e *Euler) TryToSaveIfNoChangeWarningReturnDone(name string) bool {
	if _, err := os.Stat(v.SavedDir); os.IsNotExist(err) {
		_ = os.MkdirAll(v.SavedDir, os.ModePerm)
	}
	n := fmt.Sprint("name: ", name, " already exist, do you want save with number? click 'save'")
	_, err := os.Stat(fmt.Sprint(v.SavedDir, "/", name, ".txt"))
	if !os.IsNotExist(err) && v.Warning != n {
		v.Warning = n
		return false
	}
	return e.saveConfig(name)
}

func (e *Euler) saveConfig(name string) bool {
	dirName := fmt.Sprint(v.SavedDir, "/", name, ".txt")
	i := 1
	if h.FileExist(dirName) {
		name = regexp.MustCompile("\\([0-9]\\)").ReplaceAllString(name, "")
		for ; ; i++ {
			dirName = fmt.Sprint(v.SavedDir, "/", name, "(", i, ")", ".txt")
			if !h.FileExist(dirName) {
				name = fmt.Sprint(name, "(", i, ")")
				break
			}
		}
	}
	f, err := os.OpenFile(dirName,
		os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		v.Warning = h.GetSytemError()
		return false
	}
	defer func() {
		_ = f.Sync()
		_ = f.Close()
	}()
	path := fmt.Sprint("./bin/", e.Name, "/", v.EulerOriginConfig)
	Origin, err := GetEulerFromName(path)
	if err != nil {
		if v.Warning != h.DoNotFound(path) {
			v.Warning = h.DoNotFound(path)
			return false
		} else {
			Origin = e.Clone()
		}
	}
	Origin.SetName(name)
	format := Origin.FormatAll()
	_, err = f.WriteString(format)
	if err != nil {
		v.Warning = h.GetSytemError()
		return false
	}
	return true
}

func (e *Euler) FormatAll() string {
	var buf bytes.Buffer
	e.Planets.FormatPlanetsToBuf(&buf)
	e.Features.FormatFeaturesToBuf(&buf)
	return buf.String()
}
