package graphics

type Canvas interface {
	Write(text string)
	Draw(sprite Sprite, point Point)
	Destroy()
}

type Point struct {
	X int
	Y int
}

type Rect struct {
	X1 int
	Y1 int
	X2 int
	Y2 int
}
