package main

import (
	"log"

	"golang.design/x/hotkey"
	"golang.design/x/hotkey/mainthread"
	"github.com/go-vgo/robotgo"
)

const pageUp = 0xff55
var hk *hotkey.Hotkey

func main() {
	println("Page up hotkey started!")

	robotgo.KeySleep = 10

	hk = hotkey.New([]hotkey.Modifier{}, pageUp)
	err := hk.Register()
	if err != nil {
		log.Fatal(err)
	}

	mainthread.Init(fn)
}

func fn() {
	for {
		<-hk.Keyup()
		robotgo.KeyPress("down")
		robotgo.KeyTap("up")
		robotgo.KeyTap("right")
		robotgo.KeyTap("left")
		robotgo.KeyTap("up")
		robotgo.KeyTap("up")
		robotgo.KeyTap("down")
		robotgo.KeyTap("down")
	}
}
