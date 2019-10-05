package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	eng "github.com/yacklebeam/ludumdare45/engine"
	sys "github.com/yacklebeam/ludumdare45/system"
)

const (
	gWindowWidth  = 800
	gWindowHeight = 600
)

func main() {
	rl.InitWindow(gWindowWidth, gWindowHeight, "Ludum Dare #45 Game")
	sys.LoadDefaults()

	loadLevel()

	for !rl.WindowShouldClose() {
		t := rl.GetFrameTime()
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		eng.Tick(t)
		rl.EndDrawing()
	}

	rl.CloseWindow()
}
