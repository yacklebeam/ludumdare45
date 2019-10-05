package engine

var RenderCoMap map[uint16]RenderCo
var PositionCoMap map[uint16]PositionCo
var OnClickCoMap map[uint16]OnClickCo
var TextCoMap map[uint16]TextCo
var TimerCoMap map[uint16]TimerCo
var MarketStockCoMap map[uint16]MarketStockCo
var PortfolioStockCoMap map[uint16]PortfolioStockCo

var PlayerCoSingleton PlayerCo
var CalendarCoSingleton CalendarCo
var SoundCoSingleton SoundCo

var MarketStockCoList []uint16
var PortfolioStockCoList []uint16

func init() {
	// init any maps here
	RenderCoMap = make(map[uint16]RenderCo)
	PositionCoMap = make(map[uint16]PositionCo)
	OnClickCoMap = make(map[uint16]OnClickCo)
	TextCoMap = make(map[uint16]TextCo)
	TimerCoMap = make(map[uint16]TimerCo)
	MarketStockCoMap = make(map[uint16]MarketStockCo)
	PortfolioStockCoMap = make(map[uint16]PortfolioStockCo)
}

const (
	GotoWorkButtonID      uint16 = 1
	StartDayButtonID      uint16 = 2
	ToggleMarketViewID    uint16 = 3
	TogglePortfolioViewID uint16 = 4
	MaxReservedID         uint16 = 5
)

func Tick(t float32) {
	timerSystemTick(t)
	onClickSystemTick(t)
	renderSystemTick(t)
	renderTextSystemTick(t)
	renderUITick(t)
	renderMarketStockTick(t)
	renderPortfolioStockTick(t)
}

func SetDisableOnClick(id uint16, isDisabled bool) {
	tmp := OnClickCoMap[id]
	tmp.Disabled = isDisabled
	OnClickCoMap[id] = tmp
}

func EndDay() {
	CalendarCoSingleton.ElapsedDayCount++
	PlayerCoSingleton.GamePaused = true
	SetDisableOnClick(StartDayButtonID, false)
	SetDisableOnClick(GotoWorkButtonID, true)
	CalendarCoSingleton.AccumulatedSec = 0
}
