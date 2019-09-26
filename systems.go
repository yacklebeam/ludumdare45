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
			rl.DrawRectangle(int32(p.x), int32(p.y), int32(p.width), int32(p.height), r.color)
		}
	}
}

func physicsSystem(ps map[uint32]positionComp, vs map[uint32]physicsComp) {
	for id, v := range vs {
		p, hasPosition := ps[id]
		if hasPosition {
			p.x += v.x
			p.y += v.y
			ps[id] = p
		}
	}
}

func collisionRenderSystem(ps map[uint32]positionComp, cs map[uint32]collisionComp) {
	for id, c := range cs {
		p, hasPosition := ps[id]
		if hasPosition {
			rl.DrawRectangleLines(int32(p.x+c.xOffset), int32(p.y+c.yOffset), int32(c.width), int32(c.height), rl.Yellow)
		}
	}
}

func textRenderSystem(ps map[uint32]positionComp, ts map[uint32]textComponent) {
	for id, t := range ts {
		p, hasPosition := ps[id]
		if hasPosition {
			rl.DrawText(t.text, int32(p.x+t.xOffset), int32(p.y+t.yOffset), t.size, t.color)
		}
	}
}
