package sprites

import "github.com/webdevwilson/goskifree/core/graphics"

var Skier = graphics.Sprite{
	file: "characters.png",
	frames: map[string][4]int{
		"blank":       graphics.Rect{0, 0, 0, 0},
		"east":        graphics.Rect{0, 0, 24, 34},
		"ese":         graphics.Rect{24, 0, 24, 34},
		"se":          graphics.Rect{49, 0, 17, 34},
		"south":       graphics.Rect{65, 0, 17, 34},
		"sw":          graphics.Rect{49, 37, 17, 34},
		"wsw":         graphics.Rect{24, 37, 24, 34},
		"west":        graphics.Rect{0, 37, 24, 34},
		"hit":         graphics.Rect{0, 78, 31, 31},
		"jumping":     graphics.Rect{84, 0, 32, 34},
		"somersault1": graphics.Rect{116, 0, 32, 34},
		"somersault2": graphics.Rect{148, 0, 32, 34},
	},
}
