package main

import (
	"github.com/gonutz/prototype/draw"
	"time"
)


func findSizeScreen() {
	setPlay(true)
	_ = draw.RunWindow("",0,0,
		func(window draw.Window) {
			window.SetFullscreen(true)
			width0, height0 := window.Size()
			width=float64(width0)
			height=float64(height0)
			window.Close()
		})
	centerX, centerY = width/2,height/2
	centerX2, centerY2 = centerX, centerY
}

func simulation() {

	chanelPositons, chanelBools := makeChanels()
	_ = draw.RunWindow("Kepler's problem ", int(width), int(height), func(window draw.Window) {

		if window.WasKeyPressed(draw.KeyEscape) {
			window.Close()
		}

		if isInit {
			initWindow(window, chanelPositons, chanelBools)
			return
		}
		if computePhase {
			lazyComputeWindow(window)
			return
		}
		mainWindow(window, chanelBools )
	})
}

func lazyComputeWindow(window draw.Window) {
	time.Sleep(time.Second/2)
	drawComputeScreen(window)
}

func mainWindow(window draw.Window, chanelBools []chan bool) {
	mux2.Lock()
	indexes := findSizesPosition()
	drawAll(indexes, window)
	mux2.Unlock()
	keyControl(window,chanelBools)
}

func initWindow(window draw.Window, chanelPositons []chan position, chanelBools []chan bool) {
	initControl(window)
	initDraw(window)
	if !isInit {
		CheckCircular()
		startCalculations(chanelPositons, chanelBools)
		ComputePhase(chanelBools)
	}
}