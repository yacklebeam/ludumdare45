package engine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type PlayerCo struct {
	CurrentAccountValue int64
	GamePaused          bool
	ShowMarket          bool
	ShowPortfolio       bool
}

type CalendarCo struct {
	ElapsedDayCount int64
	AccumulatedSec  float32
}

type RenderCo struct {
	Texture    string
	Tint       rl.Color
	SourceRect rl.Rectangle
	Visible    bool
}

type PositionCo struct {
	X      float32
	Y      float32
	Z      float32
	Width  float32
	Height float32
}

type TextCo struct {
	RawText  string
	Text     string
	Size     int32
	Color    rl.Color
	OffsetX  float32
	OffsetY  float32
	OnUpdate func(uint16)
}

type OnClickCo struct {
	OnClick  func(uint16)
	Disabled bool
}

type PlayerStockCo struct {
	id        uint32
	numShares int32
}

// PortfolioStockCo is used to tie the in game entities back to the stock data
type PortfolioStockCo struct {
	ID uint16
}

// MarketStockCo is used to tie the in game entities back to the stock data
type MarketStockCo struct {
	ID uint16
}

// StockDataLookupCo is a component used to store information about a particular stock type
type StockDataLookupCo struct {
	Name             string
	CurrentPrice     float32
	SharesOut        int32
	PlayerShareCount int32
	Available        bool
}

type TimerCo struct {
	OnTick         func(uint16)
	TickLength     float32
	AccumulatedSec float32
}

type MusicCo struct {
	IsPlaying bool
	Music     rl.Music
}
