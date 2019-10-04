package engine

var renderCoMap map[uint16]renderCo
var positionCoMap map[uint16]positionCo
var playerCoSingleton playerCo

func init() {
	// init any maps here
	renderCoMap = make(map[uint16]renderCo)
	positionCoMap = make(map[uint16]positionCo)
}

func tick(t float32) {
	renderSystemTick(t)
}
