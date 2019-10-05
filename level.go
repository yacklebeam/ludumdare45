package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	eng "github.com/yacklebeam/ludumdare45/engine"
	sys "github.com/yacklebeam/ludumdare45/system"
)

func loadLevel() {
	var coID uint16 = 0
	sys.LoadTextureFromFile("example.png")
	sys.LoadTextureFromFile("ui_frame.png")

	// player singleton
	eng.PlayerCoSingleton = eng.PlayerCo{CurrentAccountValue: 0.0, GamePaused: true, ShowMarket: false, ShowPortfolio: false}
	eng.CalendarCoSingleton = eng.CalendarCo{ElapsedDayCount: 0, AccumulatedSec: 0}

	// click to work button
	coID = eng.GotoWorkButtonID
	eng.RenderCoMap[coID] = eng.RenderCo{Texture: "example.png", SourceRect: rl.NewRectangle(0, 0, 30, 30), Tint: rl.White}
	eng.PositionCoMap[coID] = eng.PositionCo{X: 10, Y: 200, Width: 200, Height: 30}
	eng.TextCoMap[coID] = eng.TextCo{Text: "Go to work...", Color: rl.Black, Size: 20, OffsetX: 10, OffsetY: 5}
	eng.OnClickCoMap[coID] = eng.OnClickCo{Disabled: true, OnClick: clickGotoWork}
	coID = eng.StartDayButtonID

	// click to start day button
	coID = eng.StartDayButtonID
	eng.RenderCoMap[coID] = eng.RenderCo{Texture: "example.png", SourceRect: rl.NewRectangle(0, 0, 30, 30), Tint: rl.White}
	eng.PositionCoMap[coID] = eng.PositionCo{X: 10, Y: 300, Width: 200, Height: 30}
	eng.TextCoMap[coID] = eng.TextCo{Text: "Start day...", Color: rl.Black, Size: 20, OffsetX: 10, OffsetY: 5}
	eng.OnClickCoMap[coID] = eng.OnClickCo{Disabled: false, OnClick: clickStartDay}

	// toggle market view button
	coID = eng.ToggleMarketViewID
	eng.RenderCoMap[coID] = eng.RenderCo{Texture: "example.png", SourceRect: rl.NewRectangle(0, 0, 30, 30), Tint: rl.White}
	eng.PositionCoMap[coID] = eng.PositionCo{X: 10, Y: 400, Width: 200, Height: 30}
	eng.TextCoMap[coID] = eng.TextCo{Text: "Show Market", Color: rl.Black, Size: 20, OffsetX: 10, OffsetY: 5}
	eng.OnClickCoMap[coID] = eng.OnClickCo{Disabled: false, OnClick: clickToggleMarket}

	// toggle portfolio view button
	coID = eng.TogglePortfolioViewID
	eng.RenderCoMap[coID] = eng.RenderCo{Texture: "example.png", SourceRect: rl.NewRectangle(0, 0, 30, 30), Tint: rl.White}
	eng.PositionCoMap[coID] = eng.PositionCo{X: 10, Y: 450, Width: 200, Height: 30}
	eng.TextCoMap[coID] = eng.TextCo{Text: "Show Portfolio", Color: rl.Black, Size: 20, OffsetX: 10, OffsetY: 5}
	eng.OnClickCoMap[coID] = eng.OnClickCo{Disabled: false, OnClick: clickTogglePortfolio}

	coID = eng.MaxReservedID
	// load starting stocks
	eng.MarketStockCoMap[coID] = eng.MarketStockCo{Name: "BANANA", CurrentValue: 100.0, SharesOut: 100}
	eng.PositionCoMap[coID] = eng.PositionCo{X: 10, Y: 300, Width: 200, Height: 30}
	eng.OnClickCoMap[coID] = eng.OnClickCo{Disabled: false, OnClick: clickMarketStock}
	eng.MarketStockCoList = append(eng.MarketStockCoList, coID)
	coID++
	eng.MarketStockCoMap[coID] = eng.MarketStockCo{Name: "APPLE", CurrentValue: 200.0, SharesOut: 55}
	eng.PositionCoMap[coID] = eng.PositionCo{X: 10, Y: 300, Width: 200, Height: 30}
	eng.OnClickCoMap[coID] = eng.OnClickCo{Disabled: false, OnClick: clickMarketStock}
	eng.MarketStockCoList = append(eng.MarketStockCoList, coID)
	coID++
	eng.MarketStockCoMap[coID] = eng.MarketStockCo{Name: "ORANGE", CurrentValue: 300.0, SharesOut: 20}
	eng.PositionCoMap[coID] = eng.PositionCo{X: 10, Y: 300, Width: 200, Height: 30}
	eng.OnClickCoMap[coID] = eng.OnClickCo{Disabled: false, OnClick: clickMarketStock}
	eng.MarketStockCoList = append(eng.MarketStockCoList, coID)
	coID++
}
