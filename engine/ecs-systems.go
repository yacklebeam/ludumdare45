package engine

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	sys "github.com/yacklebeam/ludumdare45/system"
)

func renderSystemTick(t float32) { // RenderCo, PositionCo
	for id, r := range RenderCoMap {
		p, hasPosition := PositionCoMap[id]
		o, hasOnClick := OnClickCoMap[id]
		if hasPosition {
			if hasOnClick && o.Disabled {
				rl.DrawTexturePro(sys.GetTexture(r.Texture), r.SourceRect, rl.NewRectangle(p.X, p.Y, p.Width, p.Height), rl.NewVector2(0, 0), 0, rl.Black)

			} else {
				rl.DrawTexturePro(sys.GetTexture(r.Texture), r.SourceRect, rl.NewRectangle(p.X, p.Y, p.Width, p.Height), rl.NewVector2(0, 0), 0, r.Tint)
			}
		}
	}
}

func renderUITick(t float32) {
	valueStr := fmt.Sprintf("Account Value: $%v", PlayerCoSingleton.CurrentAccountValue)
	rl.DrawText(valueStr, 10, 10, 20, rl.Black)

	valueStr = fmt.Sprintf("Day Ends In: %.fs", 60.0-CalendarCoSingleton.AccumulatedSec)
	rl.DrawText(valueStr, 10, 40, 20, rl.Black)

	valueStr = fmt.Sprintf("Day #%v", CalendarCoSingleton.ElapsedDayCount)
	rl.DrawText(valueStr, 10, 70, 20, rl.Black)
}

func renderTextSystemTick(t float32) {
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
			if !o.Disabled && rl.IsMouseButtonPressed(rl.MouseLeftButton) && rl.CheckCollisionPointRec(mPos, rl.NewRectangle(p.X, p.Y, p.Width, p.Height)) {
				o.OnClick(id)
			}
		}
	}
}

func timerSystemTick(t float32) {
	if PlayerCoSingleton.GamePaused {
		return
	}
	// see here for example of updating a component within a map
	for id, tc := range TimerCoMap {
		tc.AccumulatedSec += t
		if tc.AccumulatedSec > tc.TickLength {
			tc.AccumulatedSec -= tc.TickLength
			tc.OnTick(id)
		}
		TimerCoMap[id] = tc
	}

	CalendarCoSingleton.AccumulatedSec += t
	if CalendarCoSingleton.AccumulatedSec > 60 {
		CalendarCoSingleton.AccumulatedSec -= 60
		CalendarCoSingleton.ElapsedDayCount++
		PlayerCoSingleton.GamePaused = true
	}
}
