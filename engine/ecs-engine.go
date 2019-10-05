package engine

var RenderCoMap map[uint16]RenderCo
var PositionCoMap map[uint16]PositionCo
var OnClickCoMap map[uint16]OnClickCo
var TextCoMap map[uint16]TextCo
var TimerCoMap map[uint16]TimerCo

var PlayerCoSingleton PlayerCo

func init() {
	// init any maps here
	RenderCoMap = make(map[uint16]RenderCo)
	PositionCoMap = make(map[uint16]PositionCo)
	OnClickCoMap = make(map[uint16]OnClickCo)
	TextCoMap = make(map[uint16]TextCo)
	TimerCoMap = make(map[uint16]TimerCo)
}

func Tick(t float32) {
	timerSystemTick(t)
	onClickSystemTick(t)
	renderSystemTick(t)
	renderPlayerSystemTick(t)
	renderTextSystemTick(t)
}
