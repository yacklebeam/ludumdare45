package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func collision(camera, positions[]) {

	collision := false

	// Check collisions player vs enemy-box
	for i := 0; i < maxColumns; i++ {
		if rl.CheckCollisionBoxes(
			rl.NewBoundingBox(
				rl.NewVector3(camera.Position.X-camera.Target.X/2, camera.Position.Y-camera.Target.Y/2, camera.Position.Z-camera.Target.Z/2),
				rl.NewVector3(camera.Position.X+camera.Target.X/2, camera.Position.Y+camera.Target.Y/2, camera.Position.Z+camera.Target.Z/2)),
			rl.NewBoundingBox(
				rl.NewVector3(positions[i].X-2.0/2, positions[i].Y-heights[i]/2, positions[i].Z-2.0/2),
				rl.NewVector3(positions[i].X+2.0/2, positions[i].Y+heights[i]/2, positions[i].Z+2.0/2)),
		) {
			collision = true
		}
		if collision {
			colors[i] = rl.Red
		}
	}
}
