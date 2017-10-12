package graphics

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/webdevwilson/goskifree/core/graphics"
	"os"
)

type sdlCanvas struct {
	window  *sdl.Window
	surface *sdl.Surface
}

var _ graphics.Canvas = &sdlCanvas{}

func NewSDLCanvas() graphics.Canvas {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}

	return &sdlCanvas{
		window:  window,
		surface: surface,
	}
}

func (c *sdlCanvas) Write(text string) {
	os.Stdout.Write([]byte(text))
}

func (c *sdlCanvas) Draw(sprite graphics.Sprite, point graphics.Point) {

}

func (c *sdlCanvas) Destroy() {
	c.window.Destroy()
}
