package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/webdevwilson/goskifree/core"
	"github.com/webdevwilson/goskifree/html/graphics"
	"github.com/webdevwilson/goskifree/html/input"
)

func main() {
	document := js.Global.Get("document")

	document.Call("addEventListener", "DOMContentLoaded", func() {
		canvas := graphics.NewWebGLCanvas()
		controller := input.NewHTMLController()
		core.Start(controller, canvas)
	})
}
