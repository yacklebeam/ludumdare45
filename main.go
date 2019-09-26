package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	gWindowWidth  = 800
	gWindowHeight = 600
)

func main() {
	// setup game
	gPositionCompMap[0] = positionComp{x: 100, y: 100, width: 20, height: 20}
	gRenderCompMap[0] = renderComp{color: rl.Red}

	gPositionCompMap[56] = positionComp{x: 300, y: 40, width: 20, height: 50}
	gRenderCompMap[56] = renderComp{color: rl.Blue}

	gPositionCompMap[404] = positionComp{x: 30, y: 400, width: 50, height: 20}
	gRenderCompMap[404] = renderComp{color: rl.Green}

	gPositionCompMap[76] = positionComp{x: 30, y: 400, width: 50, height: 20}
	// since this doesn't have a renderComp associated with this entity ID, the renderSystem will skip it

	rl.InitWindow(gWindowWidth, gWindowHeight, "Ludum Dare #45 Game")
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		renderSystem(gPositionCompMap, gRenderCompMap)
		rl.EndDrawing()
	}
	rl.CloseWindow()
}
