package main

var gPositionCompMap map[uint32]positionComp
var gRenderCompMap map[uint32]renderComp
var gPhysicsCompMap map[uint32]physicsComp

func init() {
	gPositionCompMap = make(map[uint32]positionComp)
	gRenderCompMap = make(map[uint32]renderComp)
	gPhysicsCompMap = make(map[uint32]physicsComp)
}
