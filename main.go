package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	gWindowHeight = 600
	gWindowWidth  = 800
)

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("ld45 test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, gWindowWidth, gWindowHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}
	surface.FillRect(nil, 0)

	rect := sdl.Rect{X: 0, Y: 0, W: gWindowWidth, H: gWindowHeight}
	surface.FillRect(&rect, 0xff0000) // RGB format
	window.UpdateSurface()

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
			}
		}
	}
}
