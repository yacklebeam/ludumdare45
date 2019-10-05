package engine

var MarketUIFrameID uint16
var MarketUIList []uint16

func uiTreeTick(t float32) {
	co, exists := RenderCoMap[MarketUIFrameID]
	if exists {
		if PlayerCoSingleton.ShowMarket {
			co.Visible = true
		} else {
			co.Visible = false
		}
		RenderCoMap[MarketUIFrameID] = co
	}
}
