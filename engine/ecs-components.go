package engine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type playerCoSingleton struct {
	currentAccountValue int64
}

type renderCo struct {
	texture    string
	tint       rl.Color
	sourceRect rl.Rectangle
}

type positionCo struct {
	x float32
	y float32
}
