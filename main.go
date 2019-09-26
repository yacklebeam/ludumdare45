package main

import (
	"github.com/gen2brain/raylib-go/raylib"
)

const (
	maxColumns = 20
)

func main() {
	rl.InitWindow(800, 450, "Trent's Epic Adventure")

	camera := rl.Camera3D{}
	camera.Position = rl.NewVector3(4.0, 2.0, 4.0)
	camera.Target = rl.NewVector3(0.0, 1.8, 0.0)
	camera.Up = rl.NewVector3(0.0, 1.0, 0.0)
	camera.Fovy = 60.0
	camera.Type = rl.CameraPerspective

	// Generates some random columns
	heights := make([]float32, maxColumns)
	positions := make([]rl.Vector3, maxColumns)
	colors := make([]rl.Color, maxColumns)

	for i := 0; i < maxColumns; i++ {
		heights[i] = float32(rl.GetRandomValue(1, 12))
		positions[i] = rl.NewVector3(float32(rl.GetRandomValue(-15, 15)), heights[i]/2, float32(rl.GetRandomValue(-15, 15)))
		colors[i] = rl.NewColor(uint8(rl.GetRandomValue(20, 255)), uint8(rl.GetRandomValue(10, 55)), 30, 255)
	}

	rl.SetCameraMode(camera, rl.CameraFirstPerson) // Set a first person camera mode

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.UpdateCamera(&camera) // Update camera

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode3D(camera)

		rl.DrawPlane(rl.NewVector3(0.0, 0.0, 0.0), rl.NewVector2(32.0, 32.0), rl.LightGray) // Draw ground
		rl.DrawCube(rl.NewVector3(-16.0, 2.5, 0.0), 1.0, 5.0, 32.0, rl.Blue)                // Draw a blue wall
		rl.DrawCube(rl.NewVector3(16.0, 2.5, 0.0), 1.0, 5.0, 32.0, rl.Lime)                 // Draw a green wall
		rl.DrawCube(rl.NewVector3(0.0, 2.5, 16.0), 32.0, 5.0, 1.0, rl.Gold)                 // Draw a yellow wall

		// Draw some cubes around
		for i := 0; i < maxColumns; i++ {
			rl.DrawCube(positions[i], 2.0, heights[i], 2.0, colors[i])
			rl.DrawCubeWires(positions[i], 2.0, heights[i], 2.0, rl.Maroon)
		}

		rl.EndMode3D()

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
