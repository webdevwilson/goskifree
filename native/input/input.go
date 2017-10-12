package input

import (
	"github.com/veandco/go-sdl2/sdl"
	coreInput "github.com/webdevwilson/goskifree/core/input"
)

const (
	UpArrow    = 82
	DownArrow  = 81
	LeftArrow  = 80
	RightArrow = 79
	Spacebar   = 44
)

type sdlController struct {
	jump  bool
	up    bool
	down  bool
	left  bool
	right bool
}

var _ coreInput.Controller = &sdlController{}

func SDLController() coreInput.Controller {
	return &sdlController{}
}

func (c *sdlController) Up() bool {
	return c.up
}

func (c *sdlController) Down() bool {
	return c.down
}

func (c *sdlController) Left() bool {
	return c.left
}

func (c *sdlController) Right() bool {
	return c.right
}

func (c *sdlController) Jump() bool {
	return c.jump
}

func (c *sdlController) Poll() {
	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.KeyDownEvent:
				code := event.(*sdl.KeyDownEvent).Keysym.Scancode
				switch code {
				case UpArrow:
					c.up = true
				case DownArrow:
					c.down = true
				case LeftArrow:
					c.left = true
				case RightArrow:
					c.right = true
				case Spacebar:
					c.jump = true
				}
			case *sdl.KeyUpEvent:
				code := event.(*sdl.KeyUpEvent).Keysym.Scancode
				switch code {
				case UpArrow:
					c.up = false
				case DownArrow:
					c.down = false
				case LeftArrow:
					c.left = false
				case RightArrow:
					c.right = false
				case Spacebar:
					c.jump = false
				}
			}
		}
		sdl.Delay(16)
	}
}
