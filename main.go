package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	gWindowWidth  = 800
	gWindowHeight = 600
)

func main() {
	rl.InitWindow(gWindowWidth, gWindowHeight, "Ludum Dare #45 Game")

	textColor := rl.Gray

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		if rl.IsKeyDown(rl.KeyQ) {
			textColor = rl.Red
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.DrawText("LUDUM DARE 45 GAME", 190, 200, 20, textColor)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
