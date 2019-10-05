package engine

var RenderCoMap map[uint16]RenderCo
var PositionCoMap map[uint16]PositionCo
var OnClickCoMap map[uint16]OnClickCo
var TextCoMap map[uint16]TextCo
var TimerCoMap map[uint16]TimerCo

var PlayerCoSingleton PlayerCo
var CalendarCoSingleton CalendarCo

func init() {
	// init any maps here
	RenderCoMap = make(map[uint16]RenderCo)
	PositionCoMap = make(map[uint16]PositionCo)
	OnClickCoMap = make(map[uint16]OnClickCo)
	TextCoMap = make(map[uint16]TextCo)
	TimerCoMap = make(map[uint16]TimerCo)
}

const (
	GotoWorkButtonID uint16 = 1
	StartDayButtonID uint16 = 2
	MaxReservedID    uint16 = 3
)

func Tick(t float32) {
	timerSystemTick(t)
	onClickSystemTick(t)
	renderSystemTick(t)
	renderTextSystemTick(t)
	renderUITick(t)
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
