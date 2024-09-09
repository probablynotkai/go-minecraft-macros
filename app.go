package main

import (
	"strings"

	"github.com/go-vgo/robotgo"
	helpers "github.com/probablynotkai/helpers"
	"github.com/probablynotkai/shop"
	hook "github.com/robotn/gohook"
)

// delay sub 310 won't work well
const (
	Delay        int = 310
	MinimalDelay int = 280
)

var (
	Typing     bool = false
	AutoTyping bool = false
)

func main() {
	helpers.RegisterSmallInventoryLocations()
	helpers.RegisterLargeInventoryLocations()

	startListener()
}

func startListener() {
	hook.Register(hook.KeyDown, []string{"b"}, func(e hook.Event) {
		if canExecute() {
			buyStage1()
		}
	})

	hook.Register(hook.KeyDown, []string{"n"}, func(e hook.Event) {
		if canExecute() {
			buyStage2()
		}
	})

	hook.Register(hook.KeyDown, []string{"u"}, func(e hook.Event) {
		if canExecute() {
			doUpgrades(1)
		}
	})

	hook.Register(hook.KeyDown, []string{"i"}, func(e hook.Event) {
		if canExecute() {
			doUpgrades(2)
		}
	})

	hook.Register(hook.KeyDown, []string{"/"}, func(e hook.Event) {
		Typing = true
	})

	hook.Register(hook.KeyDown, []string{"enter"}, func(e hook.Event) {
		if Typing {
			Typing = false
		}
	})

	hook.Register(hook.KeyDown, []string{"esc"}, func(e hook.Event) {
		if Typing {
			Typing = false
		}
	})

	s := hook.Start()
	<-hook.Process(s)
}

func goToSlot(slot [2]int) {
	robotgo.Move(slot[0], slot[1])
}

func leftClickSlot() {
	robotgo.Click()
}

func buyItem(itemArr [2]int, targetAmount int) {
	itemReturn, itemSlot := itemArr[0], itemArr[1]

	goToSlot(helpers.LargeLocations[itemSlot])

	for i := targetAmount; i > 0; i -= itemReturn {
		robotgo.MouseSleep = Delay
		leftClickSlot()
		robotgo.MouseSleep = Delay
	}
}

// TOTAL PRESTIGE: 98 coins
func buyStage1() {
	buyItem(shop.COBBLE, 160)
	buyItem(shop.COAL, 160)
	buyItem(shop.QUARTZ, 160)
	buyItem(shop.IRON, 160)
	buyItem(shop.GOLD, 160)
	buyItem(shop.LAPIS, 160)
	buyItem(shop.REDSTONE, 160)
	buyItem(shop.EMERALD, 160)
	robotgo.KeyPress("esc")
}

func buyStage2() {
	buyItem(shop.DIAMOND, 160)
	buyItem(shop.SCRAPS, 160)
	buyItem(shop.OBSIDIAN, 64)
	robotgo.KeyPress("esc")
}

func doUpgrades(stage int) {
	if stage == 1 {
		for i := 0; i < 31; i++ {
			goToSlot(helpers.SmallLocations[10])
			robotgo.MouseSleep = MinimalDelay
			leftClickSlot()
			robotgo.MouseSleep = Delay
		}
	} else if stage == 2 {
		for i := 0; i < 9; i++ {
			goToSlot(helpers.SmallLocations[10])
			robotgo.MouseSleep = MinimalDelay
			leftClickSlot()
			robotgo.MouseSleep = Delay
		}
	}

	robotgo.KeyPress("esc")

	if stage == 2 {
		robotgo.MilliSleep(500)
		rebirth()
	}
}

func rebirth() {
	autoType("/rebirth")
	robotgo.KeyTap("enter")
	autoType("/rebirth")
	robotgo.KeyTap("enter")
}

func autoType(str string) {
	AutoTyping = true

	args := strings.Split(str, "")

	for i := 0; i < len(args); i++ {
		robotgo.MilliSleep(250)
		robotgo.KeyTap(args[i])
	}

	AutoTyping = false
}

func canExecute() bool {
	return !AutoTyping && !Typing
}
