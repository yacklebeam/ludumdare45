package engine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type playerCo struct {
	currentAccountValue int64
}

type renderCo struct {
	texture    string
	tint       rl.Color
	sourceRect rl.Rectangle
	width      float32
	height     float32
}

type positionCo struct {
	x float32
	y float32
}
