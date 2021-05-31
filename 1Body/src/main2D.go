package main

import (
	"math/rand"
	"time"
)

func main() {
	findSizeScreen()
	initVariable()
	rand.Seed(time.Now().UTC().UnixNano())
	xlsxs = makeExcelFile()
	defer func() {
		closeExcelFiles()
	}()

	simulation()
	play=false
}






