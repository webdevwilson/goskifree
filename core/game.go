package core

import (
	"github.com/webdevwilson/goskifree/core/graphics"
	"github.com/webdevwilson/goskifree/core/input"
)

func Start(controller input.Controller, g graphics.Canvas) {
	controller.Poll()
}
