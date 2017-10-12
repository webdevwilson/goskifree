package input

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/webdevwilson/goskifree/core/input"
)

const (
	UpArrow    = 38
	DownArrow  = 40
	LeftArrow  = 37
	RightArrow = 39
	Spacebar   = 32
)

type htmlController struct {
	jump  bool
	up    bool
	down  bool
	left  bool
	right bool
}

var _ input.Controller = &htmlController{}

func NewHTMLController() input.Controller {
	c := &htmlController{}
	window := js.Global.Get("window")
	window.Call("addEventListener", "keydown", toggle(c, true))
	window.Call("addEventListener", "keyup", toggle(c, false))
	return c
}

func toggle(c *htmlController, mode bool) func(e *js.Object) {
	return func(e *js.Object) {
		keyCode := e.Get("keyCode").Int()
		switch keyCode {
		case Spacebar:
			c.jump = mode
		case UpArrow:
			c.up = mode
		case DownArrow:
			c.down = mode
		case LeftArrow:
			c.left = mode
		case RightArrow:
			c.right = mode
		}
	}
}

func (c *htmlController) Up() bool {
	return c.up
}

func (c *htmlController) Down() bool {
	return c.down
}

func (c *htmlController) Left() bool {
	return c.left
}

func (c *htmlController) Right() bool {
	return c.right
}

func (c *htmlController) Jump() bool {
	return c.down
}

func (c *htmlController) Poll() {
}
