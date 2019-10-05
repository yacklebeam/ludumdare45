package engine

// Market Pane
var MarketUIFrameID uint16
var MarketUIList []StockUIElem

// Portfolio Pane
var PortfolioUIFrameID uint16
var PortfolioUIList []StockUIElem

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
			for _, stock := range MarketUIList {
				e := TextCoMap[stock.TopElemID]
				e.Visible = true
				TextCoMap[stock.TopElemID] = e
				e = TextCoMap[stock.MidElemID]
				e.Visible = true
				TextCoMap[stock.TopElemID] = e
				e = TextCoMap[stock.BotElemID]
				e.Visible = true
				TextCoMap[stock.TopElemID] = e
			}
		} else {
			co.Visible = false
			for _, stock := range MarketUIList {
				e := TextCoMap[stock.TopElemID]
				e.Visible = false
				TextCoMap[stock.TopElemID] = e
				e = TextCoMap[stock.MidElemID]
				e.Visible = false
				TextCoMap[stock.TopElemID] = e
				e = TextCoMap[stock.BotElemID]
				e.Visible = false
				TextCoMap[stock.TopElemID] = e
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
