package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	eng "github.com/yacklebeam/ludumdare45/engine"
	sys "github.com/yacklebeam/ludumdare45/system"
)

func loadLevel() {
	var coID uint16 = 0
	sys.LoadTextureFromFile("example.png")

	// load the audio assets
	sys.LoadAudioFromFile("McCuckolds_Jingle_(Min).wav")
	workMusic := sys.GetAudio("McCuckolds_Jingle_(Min).wav")

	// player singleton
	eng.PlayerCoSingleton = eng.PlayerCo{CurrentAccountValue: 0.0, GamePaused: true}
	eng.CalendarCoSingleton = eng.CalendarCo{ElapsedDayCount: 0, AccumulatedSec: 0}

	// click to work button
	coID = eng.GotoWorkButtonID

	eng.RenderCoMap[coID] = eng.RenderCo{Texture: "example.png", SourceRect: rl.NewRectangle(0, 0, 30, 30), Tint: rl.White}
	eng.PositionCoMap[coID] = eng.PositionCo{X: 10, Y: 200, Width: 200, Height: 30}
	eng.TextCoMap[coID] = eng.TextCo{Text: "Go to work...", Color: rl.Black, Size: 20, OffsetX: 10, OffsetY: 5}
	eng.OnClickCoMap[coID] = eng.OnClickCo{Disabled: true, OnClick: func(id uint16) {
		eng.PlayerCoSingleton.CurrentAccountValue += 500
		rl.PlaySound(workMusic)
		eng.EndDay()
	}}

	rl.UnloadSound(workMusic)

	coID = eng.StartDayButtonID

	// click to start day button
	eng.RenderCoMap[coID] = eng.RenderCo{Texture: "example.png", SourceRect: rl.NewRectangle(0, 0, 30, 30), Tint: rl.White}
	eng.PositionCoMap[coID] = eng.PositionCo{X: 10, Y: 300, Width: 200, Height: 30}
	eng.TextCoMap[coID] = eng.TextCo{Text: "Start day...", Color: rl.Black, Size: 20, OffsetX: 10, OffsetY: 5}
	eng.OnClickCoMap[coID] = eng.OnClickCo{Disabled: false, OnClick: func(id uint16) {
		eng.PlayerCoSingleton.GamePaused = false
		eng.PlayerCoSingleton.CurrentAccountValue -= 100
		eng.SetDisableOnClick(id, true)
		eng.SetDisableOnClick(eng.GotoWorkButtonID, false)
	}}

	coID = eng.MaxReservedID
}
