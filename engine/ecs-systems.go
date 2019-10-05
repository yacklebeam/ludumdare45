package engine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	sys "github.com/yacklebeam/ludumdare45/system"
)

func renderSystemTick(t float32) { // RenderCo, PositionCo
	for id, r := range RenderCoMap {
		p, hasPosition := PositionCoMap[id]
		o, hasOnClick := OnClickCoMap[id]
		if hasPosition && r.Visible {
			if hasOnClick && o.Disabled {
				rl.DrawTexturePro(sys.GetTexture(r.Texture), r.SourceRect, rl.NewRectangle(p.X, p.Y, p.Width, p.Height), rl.NewVector2(0, 0), 0, rl.Black)
			} else {
				rl.DrawTexturePro(sys.GetTexture(r.Texture), r.SourceRect, rl.NewRectangle(p.X, p.Y, p.Width, p.Height), rl.NewVector2(0, 0), 0, r.Tint)
			}
		}
	}
}

func renderTextSystemTick(t float32) {
	for id, t := range TextCoMap {
		if t.OnUpdate != nil {
			t.OnUpdate(id)
		}
		p, hasPosition := PositionCoMap[id]
		if hasPosition {
			rl.DrawText(t.Text, int32(p.X+t.OffsetX), int32(p.Y+t.OffsetY), t.Size, t.Color)
		}
	}
}

/*func renderMarketStockTick(t float32) {
	if PlayerCoSingleton.ShowMarket {
		rl.DrawTexturePro(sys.GetTexture("ui_frame.png"), rl.NewRectangle(0, 0, 100, 100), rl.NewRectangle(100, 100, 400, 400), rl.NewVector2(0, 0), 0, rl.White)
		for index, id := range MarketStockCoList {
			ms := MarketStockCoMap[id]
			rl.DrawRectangleLines(105, int32(105+index*60), 100, 55, rl.Black)
			rl.DrawText(fmt.Sprintf("%v", ms.Name), 110, int32(110+index*60), 15, rl.Black)
			rl.DrawText(fmt.Sprintf("Val: $%v", ms.CurrentValue), 110, int32(125+index*60), 15, rl.Black)
			rl.DrawText(fmt.Sprintf("Avail: %v", ms.SharesOut), 110, int32(140+index*60), 15, rl.Black)

			o, hasOnClick := OnClickCoMap[id]
			p, hasPosition := PositionCoMap[id]
			if hasOnClick && hasPosition {
				o.Disabled = false
				OnClickCoMap[id] = o

				p.X = 105
				p.Y = float32(105 + index*60)
				p.Width = 100
				p.Height = 55
				PositionCoMap[id] = p
			}
		}
	} else {
		for _, id := range MarketStockCoList {

			o, hasOnClick := OnClickCoMap[id]
			if hasOnClick {
				o.Disabled = true
				OnClickCoMap[id] = o
			}
		}
	}
}*/

/*func renderPortfolioStockTick(t float32) {
	if PlayerCoSingleton.ShowPortfolio {
		rl.DrawTexturePro(sys.GetTexture("ui_frame.png"), rl.NewRectangle(0, 0, 100, 100), rl.NewRectangle(100, 100, 400, 400), rl.NewVector2(0, 0), 0, rl.White)
		for index, id := range PortfolioStockCoList {
			ps := PortfolioStockCoMap[id]
			ms, _ := MarketStockCoMap[id]
			rl.DrawRectangleLines(105, int32(105+index*60), 100, 55, rl.Black)
			rl.DrawText(fmt.Sprintf("%v", ms.Name), 110, int32(110+index*60), 15, rl.Black)
			rl.DrawText(fmt.Sprintf("Val: $%v", ms.CurrentValue), 110, int32(125+index*60), 15, rl.Black)
			rl.DrawText(fmt.Sprintf("Current: %v", ps.CurrentCount), 110, int32(140+index*60), 15, rl.Black)

			o, hasOnClick := OnClickCoMap[id]
			p, hasPosition := PositionCoMap[id]
			if hasOnClick && hasPosition {
				o.Disabled = false
				OnClickCoMap[id] = o

				p.X = 105
				p.Y = float32(105 + index*60)
				p.Width = 100
				p.Height = 55
				PositionCoMap[id] = p
			}
		}
	} else {
		for _, id := range PortfolioStockCoList {

			o, hasOnClick := OnClickCoMap[id]
			if hasOnClick {
				o.Disabled = true
				OnClickCoMap[id] = o
			}
		}
	}
}*/

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
		EndDay()
	}
}

func audioStreamingTick(t float32) {
}
