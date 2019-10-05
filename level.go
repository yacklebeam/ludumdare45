package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	eng "github.com/yacklebeam/ludumdare45/engine"
	sys "github.com/yacklebeam/ludumdare45/system"
)

func loadStocks() {
	// these IDs correspond to stock IDs, not
	eng.StockDataLookupCoMap[0] = eng.StockDataLookupCo{Name: "ABC", CurrentPrice: 100.0, SharesOut: 100, PlayerShareCount: 0, Available: true}
	eng.StockDataLookupCoMap[1] = eng.StockDataLookupCo{Name: "GOOG", CurrentPrice: 60.0, SharesOut: 100, PlayerShareCount: 0, Available: true}
	eng.StockDataLookupCoMap[2] = eng.StockDataLookupCo{Name: "MSFT", CurrentPrice: 200.0, SharesOut: 100, PlayerShareCount: 0, Available: true}
	eng.StockDataLookupCoMap[3] = eng.StockDataLookupCo{Name: "ABC", CurrentPrice: 100.0, SharesOut: 100, PlayerShareCount: 0, Available: false}
	eng.StockDataLookupCoMap[4] = eng.StockDataLookupCo{Name: "ABC", CurrentPrice: 100.0, SharesOut: 100, PlayerShareCount: 0, Available: false}
	eng.StockDataLookupCoMap[5] = eng.StockDataLookupCo{Name: "ABC", CurrentPrice: 100.0, SharesOut: 100, PlayerShareCount: 0, Available: false}
	eng.StockDataLookupCoMap[6] = eng.StockDataLookupCo{Name: "ABC", CurrentPrice: 100.0, SharesOut: 100, PlayerShareCount: 0, Available: false}
}

func loadLevel() {
	sys.LoadTextureFromFile("example.png")
	sys.LoadTextureFromFile("ui_frame.png")
	sys.LoadSoundFromFile("mccuck.ogg")
	sys.LoadMusicFromFile("typewriter.ogg")

	loadStocks()

	// player singleton
	eng.PlayerCoSingleton = eng.PlayerCo{CurrentAccountValue: 0.0, GamePaused: true, ShowMarket: false, ShowPortfolio: false}
	eng.CalendarCoSingleton = eng.CalendarCo{ElapsedDayCount: 0, AccumulatedSec: 0}

	// music singleton
	eng.MusicCoSingleton = eng.MusicCo{Music: sys.GetMusic("typewriter.ogg"), Volume: 5.0}

	// click to work button
	eng.CoID = eng.GotoWorkButtonID
	eng.RenderCoMap[eng.CoID] = eng.RenderCo{Texture: "example.png", SourceRect: rl.NewRectangle(0, 0, 30, 30), Tint: rl.White, Visible: true}
	eng.PositionCoMap[eng.CoID] = eng.PositionCo{X: 10, Y: 200, Width: 200, Height: 30}
	eng.TextCoMap[eng.CoID] = eng.TextCo{Text: "Go to work...", Color: rl.Black, Size: 20, OffsetX: 10, OffsetY: 5}
	eng.OnClickCoMap[eng.CoID] = eng.OnClickCo{Disabled: true, OnClick: clickGotoWork}
	eng.CoID = eng.StartDayButtonID

	// click to start day button
	eng.CoID = eng.StartDayButtonID
	eng.RenderCoMap[eng.CoID] = eng.RenderCo{Texture: "example.png", SourceRect: rl.NewRectangle(0, 0, 30, 30), Tint: rl.White, Visible: true}
	eng.PositionCoMap[eng.CoID] = eng.PositionCo{X: 10, Y: 300, Width: 200, Height: 30}
	eng.TextCoMap[eng.CoID] = eng.TextCo{Text: "Start day...", Color: rl.Black, Size: 20, OffsetX: 10, OffsetY: 5, Visible: true, OnUpdate: func(id uint16) {

	}}
	eng.OnClickCoMap[eng.CoID] = eng.OnClickCo{Disabled: false, OnClick: clickStartDay}

	// toggle market view button
	eng.CoID = eng.ToggleMarketViewID
	eng.RenderCoMap[eng.CoID] = eng.RenderCo{Texture: "example.png", SourceRect: rl.NewRectangle(0, 0, 30, 30), Tint: rl.White, Visible: true}
	eng.PositionCoMap[eng.CoID] = eng.PositionCo{X: 10, Y: 400, Width: 200, Height: 30}
	eng.TextCoMap[eng.CoID] = eng.TextCo{Text: "Show Market", Color: rl.Black, Size: 20, OffsetX: 10, OffsetY: 5}
	eng.OnClickCoMap[eng.CoID] = eng.OnClickCo{Disabled: false, OnClick: clickToggleMarket}

	// toggle portfolio view button
	eng.CoID = eng.TogglePortfolioViewID
	eng.RenderCoMap[eng.CoID] = eng.RenderCo{Texture: "example.png", SourceRect: rl.NewRectangle(0, 0, 30, 30), Tint: rl.White, Visible: true}
	eng.PositionCoMap[eng.CoID] = eng.PositionCo{X: 10, Y: 450, Width: 200, Height: 30}
	eng.TextCoMap[eng.CoID] = eng.TextCo{Text: "Show Portfolio", Color: rl.Black, Size: 20, OffsetX: 10, OffsetY: 5}
	eng.OnClickCoMap[eng.CoID] = eng.OnClickCo{Disabled: false, OnClick: clickTogglePortfolio}

	eng.CoID = eng.MaxReservedID

	// HUD elements
	eng.PositionCoMap[eng.CoID] = eng.PositionCo{X: 10, Y: 10}
	eng.TextCoMap[eng.CoID] = eng.TextCo{RawText: "Account Balance: $%v", Color: rl.Black, Size: 20, OffsetX: 0, OffsetY: 0, OnUpdate: func(id uint16) {
		tmp := eng.TextCoMap[id]
		tmp.Text = fmt.Sprintf(tmp.RawText, eng.PlayerCoSingleton.CurrentAccountValue)
		eng.TextCoMap[id] = tmp
	}}
	eng.CoID++

	eng.PositionCoMap[eng.CoID] = eng.PositionCo{X: 10, Y: 40}
	eng.TextCoMap[eng.CoID] = eng.TextCo{RawText: "Day Ends In: %.fs", Color: rl.Black, Size: 20, OffsetX: 0, OffsetY: 0, OnUpdate: func(id uint16) {
		tmp := eng.TextCoMap[id]
		tmp.Text = fmt.Sprintf(tmp.RawText, 60.0-eng.CalendarCoSingleton.AccumulatedSec)
		eng.TextCoMap[id] = tmp
	}}
	eng.CoID++

	eng.PositionCoMap[eng.CoID] = eng.PositionCo{X: 10, Y: 70}
	eng.TextCoMap[eng.CoID] = eng.TextCo{RawText: "Day #%v", Color: rl.Black, Size: 20, OffsetX: 0, OffsetY: 0, OnUpdate: func(id uint16) {
		tmp := eng.TextCoMap[id]
		tmp.Text = fmt.Sprintf(tmp.RawText, eng.CalendarCoSingleton.ElapsedDayCount+1)
		eng.TextCoMap[id] = tmp
	}}
	eng.CoID++

	// Market Frame
	eng.PositionCoMap[eng.CoID] = eng.PositionCo{X: 100, Y: 100, Width: 200, Height: 200}
	eng.RenderCoMap[eng.CoID] = eng.RenderCo{Texture: "ui_frame.png", SourceRect: rl.NewRectangle(0, 0, 100, 100), Tint: rl.White, Visible: false}
	eng.MarketUIFrameID = eng.CoID
	eng.CoID++

	// add entities for rendering stock data
	/*for stockID, stockData := range eng.StockDataLookupCoMap {
		eng.MarketStockCoMap[eng.CoID] = eng.MarketStockCo{ID: stockID}
		eng.PositionCoMap[eng.CoID] = eng.PositionCo{X: 0, Y: 0}
		eng.TextCoMap[eng.CoID] = eng.TextCo{Text: stockData.Name, Color: rl.Black, Size: 15, OffsetX: 0, OffsetY: 0}
		eng.CoID++
		eng.MarketStockCoMap[eng.CoID] = eng.MarketStockCo{ID: stockID}
		eng.PositionCoMap[eng.CoID] = eng.PositionCo{X: 0, Y: 15}
		eng.TextCoMap[eng.CoID] = eng.TextCo{Text: "Price", Color: rl.Black, Size: 15, OffsetX: 0, OffsetY: 0}
		eng.CoID++
		eng.MarketStockCoMap[eng.CoID] = eng.MarketStockCo{ID: stockID}
		eng.PositionCoMap[eng.CoID] = eng.PositionCo{X: 0, Y: 30}
		eng.TextCoMap[eng.CoID] = eng.TextCo{Text: "Shares Out", Color: rl.Black, Size: 15, OffsetX: 0, OffsetY: 0}
		eng.CoID++
	}*/
}
