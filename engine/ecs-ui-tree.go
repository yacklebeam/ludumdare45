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
			//for index, stock := range MarketUIList {

			//}
		} else {
			co.Visible = false
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
