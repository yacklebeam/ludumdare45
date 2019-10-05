package engine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	sys "github.com/yacklebeam/ludumdare45/system"
)

func renderSystemTick(t float32) { // RenderCo, PositionCo
	for id, r := range RenderCoMap {
		p, hasPosition := PositionCoMap[id]
		o, hasOnClick := OnClickCoMap[id]
		if hasPosition && r.Visible {
			if hasOnClick && o.Disabled {
				//rl.DrawTexturePro(sys.GetTexture(r.Texture), r.SourceRect, rl.NewRectangle(p.X, p.Y, p.Width, p.Height), rl.NewVector2(0, 0), 0, rl.Black)
				addToRenderPipeline(renderable{renderType: 0, texture: r.Texture, sourceRec: r.SourceRect, position: rl.NewRectangle(p.X, p.Y, p.Width, p.Height), color: rl.Black, zIndex: p.Z})
			} else {
				//rl.DrawTexturePro(sys.GetTexture(r.Texture), r.SourceRect, rl.NewRectangle(p.X, p.Y, p.Width, p.Height), rl.NewVector2(0, 0), 0, r.Tint)
				addToRenderPipeline(renderable{renderType: 0, texture: r.Texture, sourceRec: r.SourceRect, position: rl.NewRectangle(p.X, p.Y, p.Width, p.Height), color: r.Tint, zIndex: p.Z})
			}
		}
	}
}

func renderTextSystemTick(t float32) {
	for id, t := range TextCoMap {
		if t.OnUpdate != nil {
			t.OnUpdate(id)
		}
		p, hasPosition := PositionCoMap[id]
		if hasPosition && t.Visible {
			//rl.DrawText(t.Text, int32(p.X+t.OffsetX), int32(p.Y+t.OffsetY), t.Size, t.Color)
			addToRenderPipeline(renderable{renderType: 1, text: t.Text, position: rl.NewRectangle(p.X+t.OffsetX, p.Y+t.OffsetY, 0, 0), textSize: t.Size, color: t.Color, zIndex: p.Z + 1})
		}
	}
}

func onClickSystemTick(t float32) {
	for id, o := range OnClickCoMap {
		p, hasPosition := PositionCoMap[id]
		if hasPosition {
			mPos := rl.GetMousePosition()
			if !o.Disabled && rl.IsMouseButtonPressed(rl.MouseLeftButton) && rl.CheckCollisionPointRec(mPos, rl.NewRectangle(p.X, p.Y, p.Width, p.Height)) {
				o.OnClick(id)
			}
		}
	}
}

func timerSystemTick(t float32) {
	if PlayerCoSingleton.GamePaused {
		return
	}
	// see here for example of updating a component within a map
	for id, tc := range TimerCoMap {
		tc.AccumulatedSec += t
		if tc.AccumulatedSec > tc.TickLength {
			tc.AccumulatedSec -= tc.TickLength
			tc.OnTick(id)
		}
		TimerCoMap[id] = tc
	}

	CalendarCoSingleton.AccumulatedSec += t
	if CalendarCoSingleton.AccumulatedSec > 60 {
		EndDay()
	}
}

func musicStreamingTick(t float32) {
	if MusicCoSingleton.IsPlaying {
		rl.UpdateMusicStream(MusicCoSingleton.Music)
	}
}

func renderPipelineTick(t float32) {
	for _, r := range renderPipeline {
		if r.renderType == 0 { // texture
			rl.DrawTexturePro(sys.GetTexture(r.texture), r.sourceRec, rl.NewRectangle(r.position.X, r.position.Y, r.position.Width, r.position.Height), rl.NewVector2(0, 0), 0, r.color)
		} else if r.renderType == 1 { // text
			rl.DrawText(r.text, int32(r.position.X), int32(r.position.Y), r.textSize, r.color)
		}
	}
	renderPipeline = nil
}
