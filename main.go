package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	gWindowWidth  = 800
	gWindowHeight = 600
)

func loadLevel(levelID uint16) {
	// setup game -- these are globally accessible
	gPositionCompMap[0] = positionComp{x: 100, y: 100, width: 20, height: 20}
	gRenderCompMap[0] = renderComp{color: rl.Red}
	gPhysicsCompMap[0] = physicsComp{x: 0.01, y: 0.01}
	// since this entity is the only one with a physicsComp, it's the only one that the physicsSystem will update

	gPositionCompMap[56] = positionComp{x: 300, y: 40, width: 20, height: 50}
	gRenderCompMap[56] = renderComp{color: rl.Blue}

	gPositionCompMap[404] = positionComp{x: 30, y: 400, width: 50, height: 20}
	gRenderCompMap[404] = renderComp{color: rl.Green}

	gPositionCompMap[76] = positionComp{x: 30, y: 400, width: 50, height: 20}
	// since this doesn't have a renderComp associated with this entity ID, the renderSystem will skip it
}

func main() {

	loadLevel(0)

	rl.InitWindow(gWindowWidth, gWindowHeight, "Ludum Dare #45 Game")
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		physicsSystem(gPositionCompMap, gPhysicsCompMap)
		renderSystem(gPositionCompMap, gRenderCompMap)
		rl.EndDrawing()
	}
	rl.CloseWindow()
}
