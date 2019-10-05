package engine

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	sys "github.com/yacklebeam/ludumdare45/system"
)

func renderSystemTick(t float32) { // RenderCo, PositionCo
	for id, r := range RenderCoMap {
		p, hasPosition := PositionCoMap[id]
		if hasPosition {
			rl.DrawTexturePro(sys.GetTexture(r.Texture), r.SourceRect, rl.NewRectangle(p.X, p.Y, p.Width, p.Height), rl.NewVector2(0, 0), 0, r.Tint)
		}
	}
}

func renderPlayerSystemTick(t float32) {
	valueStr := fmt.Sprintf("Account Value: $%v", PlayerCoSingleton.CurrentAccountValue)
	rl.DrawText(valueStr, 10, 10, 20, rl.Black)
}

func renderTextSystem(t float32) {
	for id, t := range TextCoMap {
		p, hasPosition := PositionCoMap[id]
		if hasPosition {
			rl.DrawText(t.Text, int32(p.X+t.OffsetX), int32(p.Y+t.OffsetY), t.Size, t.Color)
		}
	}
}

func onClickSystemTick(t float32) {
	for id, o := range OnClickCoMap {
		p, hasPosition := PositionCoMap[id]
		if hasPosition {
			mPos := rl.GetMousePosition()
			if rl.IsMouseButtonPressed(rl.MouseLeftButton) && rl.CheckCollisionPointRec(mPos, rl.NewRectangle(p.X, p.Y, p.Width, p.Height)) {
				o.OnClick()
			}
		}
	}
}
