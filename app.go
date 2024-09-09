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
	pmChannel := make(chan error)

	go pickleMineMacros.registerActions(pmChannel)
	<-pmChannel

	s := hook.Start()
	<-hook.Process(s)
}
