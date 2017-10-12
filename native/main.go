package main

import (
	"github.com/webdevwilson/goskifree/core"
	"github.com/webdevwilson/goskifree/native/graphics"
	"github.com/webdevwilson/goskifree/native/input"
	"runtime"
)

func main() {

	runtime.LockOSThread()

	canvas := graphics.NewSDLCanvas()
	controller := input.SDLController()
	core.Start(controller, canvas)
}
