package engine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type PlayerCo struct {
	CurrentAccountValue int64
}

type RenderCo struct {
	Texture    string
	Tint       rl.Color
	SourceRect rl.Rectangle
}

type PositionCo struct {
	X      float32
	Y      float32
	Width  float32
	Height float32
}

type TextCo struct {
	Text    string
	Size    int32
	Color   rl.Color
	OffsetX float32
	OffsetY float32
}

type OnClickCo struct {
	OnClick func()
}
