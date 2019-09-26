package main

import rl "github.com/gen2brain/raylib-go/raylib"

type positionComp struct { // remove width from this
	x      float32
	y      float32
	width  float32
	height float32
}

type renderComp struct {
	// this is standin, eventually this will be a texture component
	color rl.Color
}

type physicsComp struct {
	x float32
	y float32
}

type collisionComp struct {
	xOffset       float32
	yOffset       float32
	width         float32
	height        float32
	collisionType uint8
}

type textComponent struct {
	text    string
	size    int32
	color   rl.Color
	xOffset float32
	yOffset float32
}
