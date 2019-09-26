package main

var gPositionCompMap map[uint32]positionComp
var gRenderCompMap map[uint32]renderComp

func init() {
	gPositionCompMap = make(map[uint32]positionComp)
	gRenderCompMap = make(map[uint32]renderComp)
}
