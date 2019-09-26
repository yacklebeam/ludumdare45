package main

import rl "github.com/gen2brain/raylib-go/raylib"

type positionComp struct {
	x      int32
	y      int32
	width  int32
	height int32
}

type renderComp struct {
	color rl.Color
}
