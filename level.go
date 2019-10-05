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
	eng.StockDataLookupCoMap[2] = eng.StockDataLookupCo{Name: "MSFT", CurrentPrice: 200.0, SharesOut: 25, PlayerShareCount: 0, Available: true}
}

func loadLevel() {
	var coID uint16 = 0
	sys.LoadTextureFromFile("example.png")
	sys.LoadTextureFromFile("ui_frame.png")
	//sys.LoadSoundFromFile("mccuck.ogg")
	//sys.LoadMusicFromFile("wind.ogg")

	loadStocks()

	// player singleton
	eng.PlayerCoSingleton = eng.PlayerCo{CurrentAccountValue: 0.0, GamePaused: true, ShowMarket: false, ShowPortfolio: false}
	eng.CalendarCoSingleton = eng.CalendarCo{ElapsedDayCount: 0, AccumulatedSec: 0}

	// click to work button
	coID = eng.GotoWorkButtonID
	eng.RenderCoMap[coID] = eng.RenderCo{Texture: "example.png", SourceRect: rl.NewRectangle(0, 0, 30, 30), Tint: rl.White, Visible: true}
	eng.PositionCoMap[coID] = eng.PositionCo{X: 10, Y: 200, Width: 200, Height: 30}
	eng.TextCoMap[coID] = eng.TextCo{Text: "Go to work...", Color: rl.Black, Size: 20, OffsetX: 10, OffsetY: 5, Visible: true}
	eng.OnClickCoMap[coID] = eng.OnClickCo{Disabled: true, OnClick: clickGotoWork}
	coID = eng.StartDayButtonID

	// click to start day button
	coID = eng.StartDayButtonID
	eng.RenderCoMap[coID] = eng.RenderCo{Texture: "example.png", SourceRect: rl.NewRectangle(0, 0, 30, 30), Tint: rl.White, Visible: true}
	eng.PositionCoMap[coID] = eng.PositionCo{X: 10, Y: 300, Width: 200, Height: 30}
	eng.TextCoMap[coID] = eng.TextCo{Text: "Start day...", Color: rl.Black, Size: 20, OffsetX: 10, OffsetY: 5, Visible: true}
	eng.OnClickCoMap[coID] = eng.OnClickCo{Disabled: false, OnClick: clickStartDay}

	// toggle market view button
	coID = eng.ToggleMarketViewID
	eng.RenderCoMap[coID] = eng.RenderCo{Texture: "example.png", SourceRect: rl.NewRectangle(0, 0, 30, 30), Tint: rl.White, Visible: true}
	eng.PositionCoMap[coID] = eng.PositionCo{X: 10, Y: 400, Width: 200, Height: 30}
	eng.TextCoMap[coID] = eng.TextCo{Text: "Show Market", Color: rl.Black, Size: 20, OffsetX: 10, OffsetY: 5, Visible: true}
	eng.OnClickCoMap[coID] = eng.OnClickCo{Disabled: false, OnClick: clickToggleMarket}

	// toggle portfolio view button
	coID = eng.TogglePortfolioViewID
	eng.RenderCoMap[coID] = eng.RenderCo{Texture: "example.png", SourceRect: rl.NewRectangle(0, 0, 30, 30), Tint: rl.White, Visible: true}
	eng.PositionCoMap[coID] = eng.PositionCo{X: 10, Y: 450, Width: 200, Height: 30}
	eng.TextCoMap[coID] = eng.TextCo{Text: "Show Portfolio", Color: rl.Black, Size: 20, OffsetX: 10, OffsetY: 5, Visible: true}
	eng.OnClickCoMap[coID] = eng.OnClickCo{Disabled: false, OnClick: clickTogglePortfolio}

	coID = eng.MaxReservedID

	// HUD elements
	eng.PositionCoMap[coID] = eng.PositionCo{X: 10, Y: 10}
	eng.TextCoMap[coID] = eng.TextCo{RawText: "Account Balance: $%v", Color: rl.Black, Size: 20, OffsetX: 0, OffsetY: 0, Visible: true, OnUpdate: func(id uint16) {
		tmp := eng.TextCoMap[id]
		tmp.Text = fmt.Sprintf(tmp.RawText, eng.PlayerCoSingleton.CurrentAccountValue)
		eng.TextCoMap[id] = tmp
	}}
	coID++

	eng.PositionCoMap[coID] = eng.PositionCo{X: 10, Y: 40}
	eng.TextCoMap[coID] = eng.TextCo{RawText: "Day Ends In: %.fs", Color: rl.Black, Size: 20, OffsetX: 0, OffsetY: 0, Visible: true, OnUpdate: func(id uint16) {
		tmp := eng.TextCoMap[id]
		tmp.Text = fmt.Sprintf(tmp.RawText, 60.0-eng.CalendarCoSingleton.AccumulatedSec)
		eng.TextCoMap[id] = tmp
	}}
	coID++

	eng.PositionCoMap[coID] = eng.PositionCo{X: 10, Y: 70}
	eng.TextCoMap[coID] = eng.TextCo{RawText: "Day #%v", Color: rl.Black, Size: 20, OffsetX: 0, OffsetY: 0, Visible: true, OnUpdate: func(id uint16) {
		tmp := eng.TextCoMap[id]
		tmp.Text = fmt.Sprintf(tmp.RawText, eng.CalendarCoSingleton.ElapsedDayCount+1)
		eng.TextCoMap[id] = tmp
	}}
	coID++

	// Market Frame
	eng.PositionCoMap[coID] = eng.PositionCo{X: 100, Y: 100, Width: 200, Height: 200, Z: 100}
	eng.RenderCoMap[coID] = eng.RenderCo{Texture: "ui_frame.png", SourceRect: rl.NewRectangle(0, 0, 100, 100), Tint: rl.White, Visible: false}
	eng.MarketUIFrameID = coID
	coID++

	// Portfolio Frame
	eng.PositionCoMap[coID] = eng.PositionCo{X: 200, Y: 100, Width: 200, Height: 200, Z: 100}
	eng.RenderCoMap[coID] = eng.RenderCo{Texture: "ui_frame.png", SourceRect: rl.NewRectangle(0, 0, 100, 100), Tint: rl.White, Visible: false}
	eng.PortfolioUIFrameID = coID
	coID++

	// add entities for rendering stock data
	for stockID, stockData := range eng.StockDataLookupCoMap {
		stockElem := eng.StockUIElem{}
		eng.MarketStockCoMap[coID] = eng.MarketStockCo{ID: stockID}
		eng.PositionCoMap[coID] = eng.PositionCo{X: 0, Y: 0, Z: 101}
		eng.TextCoMap[coID] = eng.TextCo{Text: stockData.Name, Color: rl.Black, Size: 15, OffsetX: 0, OffsetY: 0}
		stockElem.TopElemID = coID
		coID++
		eng.MarketStockCoMap[coID] = eng.MarketStockCo{ID: stockID}
		eng.PositionCoMap[coID] = eng.PositionCo{X: 0, Y: 15, Z: 101}
		eng.TextCoMap[coID] = eng.TextCo{RawText: "Price: $%v", Color: rl.Black, Size: 15, OffsetX: 0, OffsetY: 0, Visible: false, OnUpdate: func(id uint16) {
			tmp := eng.TextCoMap[id]
			stock, exists := eng.MarketStockCoMap[id]
			if exists {
				stockID := stock.ID
				tmp.Text = fmt.Sprintf(tmp.RawText, eng.StockDataLookupCoMap[stockID].CurrentPrice)
				eng.TextCoMap[id] = tmp
			}
		}}
		stockElem.MidElemID = coID
		coID++
		eng.MarketStockCoMap[coID] = eng.MarketStockCo{ID: stockID}
		eng.PositionCoMap[coID] = eng.PositionCo{X: 0, Y: 30, Z: 101}
		eng.TextCoMap[coID] = eng.TextCo{RawText: "Shares Out: %v", Color: rl.Black, Size: 15, OffsetX: 0, OffsetY: 0, Visible: false, OnUpdate: func(id uint16) {
			tmp := eng.TextCoMap[id]
			stock, exists := eng.MarketStockCoMap[id]
			if exists {
				stockID := stock.ID
				tmp.Text = fmt.Sprintf(tmp.RawText, eng.StockDataLookupCoMap[stockID].SharesOut)
				eng.TextCoMap[id] = tmp
			}
		}}
		stockElem.BotElemID = coID
		coID++

		eng.PositionCoMap[coID] = eng.PositionCo{X: 0, Y: 30, Z: 101, Width: 100, Height: 50}
		eng.OnClickCoMap[coID] = eng.OnClickCo{Disabled: true, OnClick: func(id uint16) {
			fmt.Printf("Clicked stock at id %v\n", id)
		}}
		stockElem.SelfElemID = coID
		coID++
		eng.MarketUIList = append(eng.MarketUIList, stockElem)
	}
}
