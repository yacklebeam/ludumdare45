package main

var gPositionCompMap map[uint32]positionComp
var gRenderCompMap map[uint32]renderComp
var gPhysicsCompMap map[uint32]physicsComp
var gCollisionCompMap map[uint32]collisionComp

func init() {
	gPositionCompMap = make(map[uint32]positionComp)
	gRenderCompMap = make(map[uint32]renderComp)
	gPhysicsCompMap = make(map[uint32]physicsComp)
	gCollisionCompMap = make(map[uint32]collisionComp)
}
