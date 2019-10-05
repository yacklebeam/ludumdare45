package engine

// MarketUIFrameID ...
var MarketUIFrameID uint16

// MarketUIList ...
var MarketUIList []StockUIElem

// PortfolioUIFrameID ...
var PortfolioUIFrameID uint16

// PortfolioUIList ...
var PortfolioUIList []StockUIElem

// StockUIElem ...
type StockUIElem struct {
	TopElemID uint16
	MidElemID uint16
	BotElemID uint16
}

func uiTreeTick(t float32) {
	co, exists := RenderCoMap[MarketUIFrameID]
	if exists {
		if PlayerCoSingleton.ShowMarket {
			co.Visible = true
			for index, stock := range MarketUIList {
				e := TextCoMap[stock.TopElemID]
				e.Visible = true
				TextCoMap[stock.TopElemID] = e
				p := PositionCoMap[stock.TopElemID]
				p.Y = float32(105 + index*60)
				p.X = float32(105)
				PositionCoMap[stock.TopElemID] = p

				e = TextCoMap[stock.MidElemID]
				e.Visible = true
				TextCoMap[stock.MidElemID] = e
				p = PositionCoMap[stock.MidElemID]
				p.Y = float32(120 + index*60)
				p.X = float32(105)
				PositionCoMap[stock.MidElemID] = p

				e = TextCoMap[stock.BotElemID]
				e.Visible = true
				TextCoMap[stock.BotElemID] = e
				p = PositionCoMap[stock.BotElemID]
				p.Y = float32(135 + index*60)
				p.X = float32(105)
				PositionCoMap[stock.BotElemID] = p
			}
		} else {
			co.Visible = false
			for _, stock := range MarketUIList {
				e := TextCoMap[stock.TopElemID]
				e.Visible = false
				TextCoMap[stock.TopElemID] = e
				e = TextCoMap[stock.MidElemID]
				e.Visible = false
				TextCoMap[stock.MidElemID] = e
				e = TextCoMap[stock.BotElemID]
				e.Visible = false
				TextCoMap[stock.BotElemID] = e
			}
		}
		RenderCoMap[MarketUIFrameID] = co
	}

	co, exists = RenderCoMap[PortfolioUIFrameID]
	if exists {
		if PlayerCoSingleton.ShowPortfolio {
			co.Visible = true
		} else {
			co.Visible = false
		}
		RenderCoMap[PortfolioUIFrameID] = co
	}
}
