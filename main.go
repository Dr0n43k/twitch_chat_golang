package main

import (
	"runtime"

	"github.com/rajveermalviya/gamen/display"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	d, err := display.NewDisplay()
	if err != nil {
		panic(err)
	}
	defer d.Destroy()

	w, err := display.NewWindow(d)
	if err != nil {
		panic(err)
	}
	defer w.Destroy()

	w.SetCloseRequestedCallback(func() { d.Destroy() })
	w.SetTitle("twitch chat by Dr0n43k")
	for {
		if !d.Poll() {
			break
		}

	}
}
