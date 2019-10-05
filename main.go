package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	eng "github.com/yacklebeam/ludumdare45/engine"
	sys "github.com/yacklebeam/ludumdare45/system"
)

const (
	gWindowWidth  = 800
	gWindowHeight = 600
)

func main() {
	rl.InitWindow(gWindowWidth, gWindowHeight, "Ludum Dare #45 Game")
	sys.LoadDefaults()

	sys.LoadTextureFromFile("example.png")

	// dummy load level for now
	eng.PlayerCoSingleton = eng.PlayerCo{CurrentAccountValue: 0.0}

	eng.RenderCoMap[0] = eng.RenderCo{Texture: "example.png", SourceRect: rl.NewRectangle(0, 0, 30, 30), Tint: rl.White}
	eng.PositionCoMap[0] = eng.PositionCo{X: 10, Y: 200, Width: 200, Height: 30}
	eng.TextCoMap[0] = eng.TextCo{Text: "Go to work...", Color: rl.Black, Size: 20, OffsetX: 10, OffsetY: 5}
	eng.OnClickCoMap[0] = eng.OnClickCo{OnClick: func() {
		eng.PlayerCoSingleton.CurrentAccountValue += 5
	}}

	for !rl.WindowShouldClose() {
		t := rl.GetFrameTime()
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		eng.Tick(t)
		rl.EndDrawing()
	}

	rl.CloseWindow()
}
