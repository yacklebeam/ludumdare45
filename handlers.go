package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	eng "github.com/yacklebeam/ludumdare45/engine"
	sys "github.com/yacklebeam/ludumdare45/system"
)

func clickMarketStock(id uint16) {
	// add 1 of this stock to player portfolio
	eng.PortfolioStockCoList = append(eng.PortfolioStockCoList, id)
	s, exists := eng.PortfolioStockCoMap[id]
	if exists {
		s.CurrentCount++
		eng.PortfolioStockCoMap[id] = s
	} else {
		eng.PortfolioStockCoMap[id] = eng.PortfolioStockCo{CurrentCount: 1}
	}
}

func clickStartDay(id uint16) {
	eng.PlayerCoSingleton.GamePaused = false
	eng.PlayerCoSingleton.CurrentAccountValue -= 100
	eng.SetDisableOnClick(id, true)
	eng.SetDisableOnClick(eng.GotoWorkButtonID, false)
}

func clickGotoWork(id uint16) {
	gotoWorkSoundFile := "mccuck.ogg"
	eng.PlayerCoSingleton.CurrentAccountValue += 500
	sys.LoadSoundFromFile(gotoWorkSoundFile)
	eng.SoundCoSingleton.Sound = sys.GetSound(gotoWorkSoundFile)
	rl.PlaySound(eng.SoundCoSingleton.Sound)
	eng.EndDay()
}

func clickToggleMarket(id uint16) {
	eng.PlayerCoSingleton.ShowMarket = !eng.PlayerCoSingleton.ShowMarket
	if eng.PlayerCoSingleton.ShowMarket {
		eng.PlayerCoSingleton.ShowPortfolio = false
	}
}

func clickTogglePortfolio(id uint16) {
	eng.PlayerCoSingleton.ShowPortfolio = !eng.PlayerCoSingleton.ShowPortfolio
	if eng.PlayerCoSingleton.ShowPortfolio {
		eng.PlayerCoSingleton.ShowMarket = false
	}
}
