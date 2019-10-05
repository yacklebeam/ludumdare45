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
}

type PositionCo struct {
	X      float32
	Y      float32
	Width  float32
	Height float32
}

type TextCo struct {
	Text    string
	Size    int32
	Color   rl.Color
	OffsetX float32
	OffsetY float32
}

type OnClickCo struct {
	OnClick  func(uint16)
	Disabled bool
}

type PlayerStockCo struct {
	id        uint32
	numShares int32
}

type MarketStockCo struct {
	Name         string
	CurrentValue float32
	SharesOut    int32
}

type PortfolioStockCo struct {
	CurrentCount int32
}

type TimerCo struct {
	OnTick         func(uint16)
	TickLength     float32
	AccumulatedSec float32
}
