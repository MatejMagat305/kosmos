package compute

import (
	h "2bodyBinary/help"
	eu "2bodyBinary/structures/euler"
	v "2bodyBinary/variables"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func Create() {
	dirName = filepath.Join("./bin", eu.MEuler.Name)
	myMkdDirIfNotExist()
}

func SetDir(name string)  {
	dirName = filepath.Join("./bin", name)

}

func createFile(n int) {
	s := fmt.Sprint(n, ".bin")
	ss := fmt.Sprint(dirName, "/planet1_period_", s)
	f1, _ = os.Create(ss)
	ss = fmt.Sprint(dirName, "/planet2_period_", s)
	f2, _ = os.Create(ss)
	PlanetData1= &h.FloatData{
		PositionsXY: make([]float64,0,period),
	}
	PlanetData2= &h.FloatData{
		PositionsXY: make([]float64,0,period),
	}
}

func saveBin() {
	go func() {	saveOneFile(f1, PlanetData1); chSave<-true}()
	go func() {	saveOneFile(f2, PlanetData2); chSave<-true}()
	go func() {	saveEulerToBin(localE); chSave<-true}()
	<-chSave; <-chSave; <-chSave
}

func saveOneFile(file *os.File, data *h.FloatData) {
	byte0 := h.Float64PointerToBytePointer(data.PositionsXY)
	_, _ = file.Write(byte0)
	_ = file.Sync()
	_ = file.Close()
}

func saveEulerToBin(e *eu.Euler) {
	jsonEuler, _ := json.Marshal(&e)
	err := writeFile(fmt.Sprint(".", string(filepath.Separator),"bin", string(filepath.Separator),
		e.Name,  string(filepath.Separator), v.EulerConfig), jsonEuler, 0666)
	if err != nil {
		fmt.Println(err)
	}
}
func writeFile(name string, data []byte, perm os.FileMode) error {
	f, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
	if err != nil {
		return err
	}
	_, err = f.Write(data)
	_ = f.Sync()
	if err1 := f.Close(); err1 != nil && err == nil {
		err = err1
	}
	return err
}
func myMkdDirIfNotExist() {
	if _, err := os.Stat("./bin"); os.IsNotExist(err) {
		_ = os.Mkdir("./bin", 0777)
	}
	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		_ = os.Mkdir(dirName, 0777)
	}
}

func send() {
	go write(Planet1)
	go write(Planet2)
	<-chOut2
	chOut1<-<-chOut2
}

func write(p *eu.Planet) {
	planetData := PlanetData1
	if p.Id == 2 {
		planetData = PlanetData2
	}
	planetData.PositionsXY = append(planetData.PositionsXY, p.PositionX)
	planetData.PositionsXY = append(planetData.PositionsXY, p.PositionY)
	chOut2 <- true
}
