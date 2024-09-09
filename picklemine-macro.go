package main

import (
	"github.com/go-vgo/robotgo"
	shop "github.com/probablynotkai/api"
)

// TOTAL PRESTIGE: 98 coins
type PickleMines struct{}

func (p *PickleMines) buyItem(itemArr [2]int, targetAmount int) {
	itemReturn, itemSlot := itemArr[0], itemArr[1]

	goToSlot(FiveRow.MappedLocations[itemSlot])

	for i := targetAmount; i > 0; i -= itemReturn {
		robotgo.MouseSleep = Delay
		leftClick()
		robotgo.MouseSleep = Delay
	}
}

func (p *PickleMines) buyStage1() {
	p.buyItem(shop.COBBLE, 160)
	p.buyItem(shop.COAL, 160)
	p.buyItem(shop.QUARTZ, 160)
	p.buyItem(shop.IRON, 160)
	p.buyItem(shop.GOLD, 160)
	p.buyItem(shop.LAPIS, 160)
	p.buyItem(shop.REDSTONE, 160)
	p.buyItem(shop.EMERALD, 160)
	robotgo.KeyPress("esc")
}

func (p *PickleMines) buyStage2() {
	p.buyItem(shop.DIAMOND, 160)
	p.buyItem(shop.SCRAPS, 160)
	p.buyItem(shop.OBSIDIAN, 64)
	robotgo.KeyPress("esc")
}

func (p *PickleMines) doUpgrades(stage int) {
	if stage == 1 {
		for i := 0; i < 31; i++ {
			goToSlot(ThreeRow.MappedLocations[10])
			robotgo.MouseSleep = MinimalDelay
			leftClick()
			robotgo.MouseSleep = Delay
		}
	} else if stage == 2 {
		for i := 0; i < 9; i++ {
			goToSlot(ThreeRow.MappedLocations[10])
			robotgo.MouseSleep = MinimalDelay
			leftClick()
			robotgo.MouseSleep = Delay
		}
	}

	robotgo.KeyPress("esc")

	if stage == 2 {
		robotgo.MilliSleep(500)
		p.rebirth()
	}
}

func (p *PickleMines) rebirth() {
	autoType("/rebirth")
	robotgo.KeyTap("enter")
	autoType("/rebirth")
	robotgo.KeyTap("enter")
}
