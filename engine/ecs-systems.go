package engine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	sys "github.com/yacklebeam/ludumdare45/system"
)

func renderSystemTick(t float32) {
	for id, r := range renderCoMap {
		p, hasPosition := positionCoMap[id]
		if hasPosition {
			rl.DrawTexturePro(sys.GetTexture(r.texture), r.sourceRect, rl.NewRectangle(p.x, p.y, r.width, r.height), rl.NewVector2(0, 0), 0, r.tint)
		}
	}
}
