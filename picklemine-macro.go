package main

import (
	"github.com/go-vgo/robotgo"
	api "github.com/probablynotkai/api"
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

func (p *PickleMines) registerActions(c chan error) {
	createConditionalMacro("b", func() {
		pickleMineMacros.buyStage1()
	})

	createConditionalMacro("n", func() {
		pickleMineMacros.buyStage2()
	})

	createConditionalMacro("u", func() {
		pickleMineMacros.doUpgrades(1)
	})

	createConditionalMacro("i", func() {
		pickleMineMacros.doUpgrades(2)
	})

	createMacro("/", func() {
		Typing = true
	})

	createMacro("enter", func() {
		if Typing {
			Typing = false
		}
	})

	createMacro("esc", func() {
		if Typing {
			Typing = false
		}
	})

	c <- nil
}

func (p *PickleMines) buyStage1() {
	p.buyItem(api.COBBLE, 160)
	p.buyItem(api.COAL, 160)
	p.buyItem(api.QUARTZ, 160)
	p.buyItem(api.IRON, 160)
	p.buyItem(api.GOLD, 160)
	p.buyItem(api.LAPIS, 160)
	p.buyItem(api.REDSTONE, 160)
	p.buyItem(api.EMERALD, 160)
	robotgo.KeyPress("esc")
}

func (p *PickleMines) buyStage2() {
	p.buyItem(api.DIAMOND, 160)
	p.buyItem(api.SCRAPS, 160)
	p.buyItem(api.OBSIDIAN, 64)
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
