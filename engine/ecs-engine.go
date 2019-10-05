package engine

import (
	"sort"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var RenderCoMap map[uint16]RenderCo
var PositionCoMap map[uint16]PositionCo
var OnClickCoMap map[uint16]OnClickCo
var TextCoMap map[uint16]TextCo
var TimerCoMap map[uint16]TimerCo
var StockDataLookupCoMap map[uint16]StockDataLookupCo

var MarketStockCoMap map[uint16]MarketStockCo
var PortfolioStockCoMap map[uint16]PortfolioStockCo

var MusicCoMap map[uint16]MusicCo

var PlayerCoSingleton PlayerCo
var CalendarCoSingleton CalendarCo
var MusicCoSingleton MusicCo

var MarketStockCoList []uint16
var PortfolioStockCoList []uint16

var renderPipeline []renderable

type renderable struct {
	renderType uint8
	text       string
	texture    string
	sourceRec  rl.Rectangle
	position   rl.Rectangle
	textSize   int32
	color      rl.Color
	zIndex     float32
}

func init() {
	// init any maps here
	RenderCoMap = make(map[uint16]RenderCo)
	PositionCoMap = make(map[uint16]PositionCo)
	OnClickCoMap = make(map[uint16]OnClickCo)
	TextCoMap = make(map[uint16]TextCo)
	TimerCoMap = make(map[uint16]TimerCo)
	StockDataLookupCoMap = make(map[uint16]StockDataLookupCo)
	MusicCoMap = make(map[uint16]MusicCo)

	MarketStockCoMap = make(map[uint16]MarketStockCo)
	PortfolioStockCoMap = make(map[uint16]PortfolioStockCo)
}

// TODO (JT) remove this
const (
	GotoWorkButtonID      uint16 = 1
	StartDayButtonID      uint16 = 2
	ToggleMarketViewID    uint16 = 3
	TogglePortfolioViewID uint16 = 4
	MaxReservedID         uint16 = 5
)

func Tick(t float32) {
	uiTreeTick(t)

	timerSystemTick(t)
	onClickSystemTick(t)
	renderSystemTick(t)
	renderTextSystemTick(t)
	musicStreamingTick(t)
	//renderMarketStockTick(t)
	//renderPortfolioStockTick(t)

	// last!
	renderPipelineTick(t)
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

// renderPipeline sorting stuff
type byZIndex []renderable

func (a byZIndex) Len() int           { return len(a) }
func (a byZIndex) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byZIndex) Less(i, j int) bool { return a[i].zIndex < a[j].zIndex }

func addToRenderPipeline(r renderable) {
	// this is supposed to be performant
	renderPipeline = append(renderPipeline, r)
	sort.Sort(byZIndex(renderPipeline))
}
