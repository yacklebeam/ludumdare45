package engine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type PlayerCo struct {
	CurrentAccountValue int64
}

type PlayerPortfolioCo struct {
	CurrentStocks []StockCo
}

type MarketCo struct {
	AvailableStocks []StockCo
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
	OnClick func()
}

type PlayerStockCo struct {
	id        uint32
	numShares int32
}

type StockCo struct {
	id           uint32
	name         string
	currentValue float32
	sharesOut    int32
}

type TimerCo struct {
	OnTick        func()
	TickLength    float32
	AccumulatedMS float32
}
