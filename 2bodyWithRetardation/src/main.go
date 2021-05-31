package main

import (
	"fmt"
	h "retardation/help"
	e "retardation/structures/euler"
	v "retardation/variables"
	win "retardation/windows"
	"runtime"
)
func main() {
	fmt.Println( "first method:",v.FirstMethod, ", without retardation:", v.WithoutRetardation)
	runtime.GOMAXPROCS(runtime.NumCPU())
	e.InitEuler(h.FindPath(v.Configs, v.Config))
	win.Simulation()
	fmt.Println("terminate")
}

