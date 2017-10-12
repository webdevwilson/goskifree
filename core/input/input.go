package input

type Controller interface {
	Poll()
	Up() bool
	Down() bool
	Left() bool
	Right() bool
	Jump() bool
}
