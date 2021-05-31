package main

import (
	h "2bodyBinary/help"
	e "2bodyBinary/structures/euler"
	v "2bodyBinary/variables"
	win "2bodyBinary/windows"
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	e.InitEuler(h.FindPath(v.Configs, v.Config))
	win.Simulation()
	fmt.Println("terminate")
}

