package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	eng "github.com/yacklebeam/ludumdare45/engine"
	sys "github.com/yacklebeam/ludumdare45/system"
)

func clickMarketStock(id uint16) {

}

func clickPortfolioStock(id uint16) {

}

func clickStartDay(id uint16) {
	// audio
	// we shouldn't use the singleton for one-off sounds, just for background music that is always playing
	if eng.MusicCoSingleton.IsPlaying {
		rl.StopMusicStream(eng.MusicCoSingleton.Music)
		eng.MusicCoSingleton.IsPlaying = false
	}
	//rl.SetMusicVolume(eng.MusicCoSingleton.Music, eng.MusicCoSingleton.Volume)
	rl.PlayMusicStream(eng.MusicCoSingleton.Music)
	eng.MusicCoSingleton.IsPlaying = true
	// player
	eng.PlayerCoSingleton.GamePaused = false
	eng.PlayerCoSingleton.CurrentAccountValue -= 100
	eng.SetDisableOnClick(id, true)
	eng.SetDisableOnClick(eng.GotoWorkButtonID, false)
}

func clickGotoWork(id uint16) {
	// audio
	sfx := sys.GetSound("mccuck.ogg")
	rl.PlaySound(sfx)
	//player
	eng.PlayerCoSingleton.CurrentAccountValue += 500
	eng.EndDay()
}

func clickToggleMarket(id uint16) {
	// audio
	// player
	eng.PlayerCoSingleton.ShowMarket = !eng.PlayerCoSingleton.ShowMarket
	if eng.PlayerCoSingleton.ShowMarket {
		eng.PlayerCoSingleton.ShowPortfolio = false
	}
}

func clickTogglePortfolio(id uint16) {
	// player
	eng.PlayerCoSingleton.ShowPortfolio = !eng.PlayerCoSingleton.ShowPortfolio
	if eng.PlayerCoSingleton.ShowPortfolio {
		eng.PlayerCoSingleton.ShowMarket = false
	}
}
