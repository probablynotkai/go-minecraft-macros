package main

import (
	helpers "github.com/probablynotkai/helpers"
	hook "github.com/robotn/gohook"
)

var (
	pickleMineMacros PickleMines = PickleMines{}
)

func main() {
	ThreeRow = helpers.GetLocationsForSizedInventory(3)
	FiveRow = helpers.GetLocationsForSizedInventory(5)

	startListener()
}

func startListener() {
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

	s := hook.Start()
	<-hook.Process(s)
}
