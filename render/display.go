package render

import (
	"github.com/xenoryt/gork/world"
)

const (
	_MinWidth  = 80
	_MinHeight = 24
)

/*Display is an object that is able to render things in a buffer
then finally display it to the user*/
type Display interface {
	//For initializing the display
	Init() error

	//IsGUI is true iff this display can draw pictures
	IsGUI() bool
	//Renders an object onto the display
	Render(Renderable) error

	//Updates the display and outputs it to the screen
	Update()
	DrawWorld(world.World, camera)
	DrawStats()
	DrawDesc()

	Width() int
	Height() int
}
