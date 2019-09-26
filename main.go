package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	gWindowWidth  = 800
	gWindowHeight = 600
)

type menu struct {
	buttons []button
	x       int32
	y       int32
}

type button struct {
	text string
}

func intersect(a, b rl.RectangleInt32) bool {
	if a.X+a.Width < b.X || a.Y+a.Height < b.Y || a.X > b.X+b.Width || a.Y > b.Y+b.Height {
		return false
	}
	return true
}

func main() {
	rl.InitWindow(gWindowWidth, gWindowHeight, "Ludum Dare #45 Game")

	m := menu{x: (gWindowWidth / 2) - 100, y: 200}
	m.buttons = append(m.buttons, button{text: "Option 1"})
	m.buttons = append(m.buttons, button{text: "Option 2"})
	m.buttons = append(m.buttons, button{text: "Option 3"})

	//rl.SetTargetFPS(200)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		mouseX := rl.GetMouseX()
		mouseY := rl.GetMouseY()

		for i, b := range m.buttons {
			buttonColor := rl.Red
			margin := 10
			buttonWidth := 200
			buttonHeight := 50

			if intersect(rl.RectangleInt32{X: m.x, Y: m.y + int32(60*i), Width: 200, Height: 50}, rl.RectangleInt32{X: mouseX, Y: mouseY, Width: 0, Height: 0}) {
				buttonColor = rl.Blue
				buttonWidth = 210
				buttonHeight = 50
			}

			rl.DrawRectangle(m.x, m.y+int32((margin+buttonHeight)*i), int32(buttonWidth), int32(buttonHeight), buttonColor)
			rl.DrawText(b.text, m.x+10, m.y+int32(60*i)+10, 30, rl.White)
		}

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
