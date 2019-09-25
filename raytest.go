package main

import (
	"github.com/gen2brain/raylib-go/raylib"
)

func newWindow() {
	rl.InitWindow(800, 600, "raylib [core] window")

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.DrawText("Makin muh gamus", 190, 200, 20, rl.Lime)
		rl.DrawCircle(400, 300, 50, rl.Black)
		rl.EndDrawing()
	}

	rl.CloseWindow()
}
