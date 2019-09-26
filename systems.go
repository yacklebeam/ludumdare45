package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// because these are in the main package right now, technically the maps are global...I don't want to access them like that
// 	since down the road we won't be using one package probably maybe
func renderSystem(ps map[uint32]positionComp, rs map[uint32]renderComp) {
	for id, r := range rs {
		p, hasPosition := ps[id]
		if hasPosition {
			rl.DrawRectangle(p.x, p.y, p.width, p.height, r.color)
		}
	}
}
