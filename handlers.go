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
	// player
	eng.PlayerCoSingleton.GamePaused = false
	eng.PlayerCoSingleton.CurrentAccountValue -= 100
	eng.SetDisableOnClick(id, true)
	eng.SetDisableOnClick(eng.GotoWorkButtonID, false)
}

func clickGotoWork(id uint16) {
	// audio
	if eng.SoundCoSingleton.IsPlaying {
		rl.StopSound(eng.SoundCoSingleton.Sound)
		eng.SoundCoSingleton.IsPlaying = false
	}
	// we shouldn't use the singleton for one-off sounds, just for background music that is always playing
	eng.SoundCoSingleton.Sound = sys.GetSound("mccuck.ogg")
	eng.SoundCoSingleton.IsPlaying = true
	rl.PlaySound(eng.SoundCoSingleton.Sound)
	//player
	eng.PlayerCoSingleton.CurrentAccountValue += 500
	eng.EndDay()
}

func clickToggleMarket(id uint16) {
	// audio
	/*if eng.SoundCoSingleton.IsPlaying {
		rl.StopSound(eng.SoundCoSingleton.Sound)
		eng.SoundCoSingleton.IsPlaying = false
	}
	eng.SoundCoSingleton.Sound = sys.GetSound("typewriter.ogg")
	eng.SoundCoSingleton.IsPlaying = true
	rl.PlaySound(eng.SoundCoSingleton.Sound)*/
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
