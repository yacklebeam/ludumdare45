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
	gCollisionCompMap[0] = collisionComp{xOffset: -2.5, yOffset: -2.5, width: 25, height: 25}

	gPositionCompMap[56] = positionComp{x: 300, y: 40, width: 20, height: 50}
	gRenderCompMap[56] = renderComp{color: rl.Blue}
	gCollisionCompMap[56] = collisionComp{xOffset: 5, yOffset: -25, width: 10, height: 100}

	gPositionCompMap[404] = positionComp{x: 30, y: 400, width: 50, height: 20}
	gRenderCompMap[404] = renderComp{color: rl.Green}

	gPositionCompMap[76] = positionComp{x: 30, y: 400, width: 50, height: 20}
}

func main() {

	loadLevel(0)

	rl.InitWindow(gWindowWidth, gWindowHeight, "Ludum Dare #45 Game")
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		// ECS System Updates -- organized by impact
		physicsSystem(gPositionCompMap, gPhysicsCompMap)
		renderSystem(gPositionCompMap, gRenderCompMap)
		collisionRenderSystem(gPositionCompMap, gCollisionCompMap)

		rl.EndDrawing()
	}
	rl.CloseWindow()
}
