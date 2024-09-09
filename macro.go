package main

import (
	"strings"

	"github.com/go-vgo/robotgo"
	"github.com/probablynotkai/helpers"
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
	ThreeRow   helpers.Inventory
	FiveRow    helpers.Inventory
)

func createMacro(key string, actions func()) {
	hook.Register(hook.KeyDown, []string{key}, func(e hook.Event) {
		actions()
	})
}

func createConditionalMacro(key string, actions func()) {
	hook.Register(hook.KeyDown, []string{key}, func(e hook.Event) {
		if canExecute() {
			actions()
		}
	})
}

func canExecute() bool {
	return !AutoTyping && !Typing
}

func goToSlot(coords [2]int) {
	robotgo.Move(coords[0], coords[1])
}

func leftClick() {
	robotgo.Click()
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
