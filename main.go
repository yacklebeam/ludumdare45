package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	maxColumns          = 20
	maxSamples          = 22050
	maxSamplesPerUpdate = 4096
)

func main() {
	rl.InitWindow(800, 450, "Trent's Epic Adventure")
	rl.InitAudioDevice()

	stream := rl.InitAudioStream(maxSamples, 32, 1)
	data := make([]float32, maxSamples)

	for i := 0; i < maxSamples; i++ {
		data[i] = float32(math.Sin(float64((2*rl.Pi*float32(i))/2) * rl.Deg2rad))
	}
	rl.PlayAudioStream(stream)

	totalSamples := int32(maxSamples)
	samplesLeft := int32(totalSamples)

	camera := rl.Camera3D{}
	camera.Position = rl.NewVector3(4.0, 2.0, 4.0)
	camera.Target = rl.NewVector3(0.0, 1.93, 0.0)
	camera.Up = rl.NewVector3(0.0, 1.0, 0.0)
	camera.Fovy = 60.0
	camera.Type = rl.CameraPerspective

	// Generates some random columns
	heights := make([]float32, maxColumns)
	positions := make([]rl.Vector3, maxColumns)
	colors := make([]rl.Color, maxColumns)

	// size the player
	playerSize := rl.NewVector3(1.0, 2.0, 1.0)

	// make an enemy box
	enemyBoxPos := rl.NewVector3(-4.0, 1.0, 0.0)
	enemyBoxSize := rl.NewVector3(2.0, 2.0, 2.0)
	enemyBoxColor := rl.Green

	for i := 0; i < maxColumns; i++ {
		heights[i] = float32(rl.GetRandomValue(1, 12))
		positions[i] = rl.NewVector3(float32(rl.GetRandomValue(-30, 30)), heights[i]/2, float32(rl.GetRandomValue(-30, 30)))
		colors[i] = rl.NewColor(uint8(rl.GetRandomValue(20, 255)), uint8(rl.GetRandomValue(10, 55)), 30, 255)
	}

	rl.SetCameraMode(camera, rl.CameraFirstPerson) // Set a first person camera mode
	rl.SetTargetFPS(60)

	collision := false

	for !rl.WindowShouldClose() {
		rl.UpdateCamera(&camera) // Update camera
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.BeginMode3D(camera)

		// draw enemy-box
		rl.DrawCube(enemyBoxPos, enemyBoxSize.X, enemyBoxSize.Y, enemyBoxSize.Z, enemyBoxColor)
		rl.DrawCubeWires(enemyBoxPos, enemyBoxSize.X, enemyBoxSize.Y, enemyBoxSize.Z, enemyBoxColor)

		rl.DrawPlane(rl.NewVector3(0.0, 0.0, 0.0), rl.NewVector2(64.0, 64.0), rl.LightGray) // Draw ground
		rl.DrawCube(rl.NewVector3(-32.0, 2.5, 0.0), 1.0, 5.0, 64.0, rl.Blue)                // Draw a blue wall
		rl.DrawCube(rl.NewVector3(32.0, 2.5, 0.0), 1.0, 5.0, 64.0, rl.Lime)                 // Draw a green wall
		rl.DrawCube(rl.NewVector3(0.0, 2.5, 32.0), 64.0, 5.0, 1.0, rl.Gold)                 // Draw a yellow wall

		// Draw some cubes around
		for i := 0; i < maxColumns; i++ {
			rl.DrawCube(positions[i], 2.0, heights[i], 2.0, colors[i])
			rl.DrawCubeWires(positions[i], 2.0, heights[i], 2.0, rl.Maroon)
		}

		collision = false

		// check player collision
		if rl.CheckCollisionBoxes(
			rl.NewBoundingBox(
				rl.NewVector3(camera.Position.X-playerSize.X/2,
					camera.Position.Y-playerSize.Y/2,
					camera.Position.Z-playerSize.Z/2),
				rl.NewVector3(camera.Position.X+playerSize.X/2,
					camera.Position.Y+playerSize.Y/2,
					camera.Position.Z+playerSize.Z/2)),
			rl.NewBoundingBox(
				rl.NewVector3(enemyBoxPos.X-enemyBoxSize.X/2,
					enemyBoxPos.Y-enemyBoxSize.Y/2,
					enemyBoxPos.Z-enemyBoxSize.Z/2),
				rl.NewVector3(enemyBoxPos.X+enemyBoxSize.X/2,
					enemyBoxPos.Y+enemyBoxSize.Y/2,
					enemyBoxPos.Z+enemyBoxSize.Z/2)),
		) {
			collision = true
		}

		if collision {
			enemyBoxColor = rl.Red
			// play a sine wave
			if rl.IsAudioBufferProcessed(stream) {
				numSamples := int32(0)
				if samplesLeft >= maxSamplesPerUpdate {
					numSamples = maxSamplesPerUpdate
				} else {
					numSamples = samplesLeft
				}

				rl.UpdateAudioStream(stream, data[totalSamples-samplesLeft:], numSamples)
				samplesLeft -= numSamples

				// loop audio
				if samplesLeft <= 0 {
					samplesLeft = totalSamples
				}
			}
		} else {
			enemyBoxColor = rl.Green
		}

		rl.EndMode3D()
		rl.EndDrawing()
	}

	rl.CloseAudioStream(stream)
	rl.CloseAudioDevice()
	rl.CloseWindow()
}
