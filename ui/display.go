package ui

import (
	. "github.com/xenoryt/gork/rect"
	. "github.com/xenoryt/gork/ui/uiconstants"
)

/*Display is an object that is able to render things in a buffer
then finally display it to the user*/
type Display interface {
	//For initializing the display
	Init() error

	//IsGUI is true iff this display can draw pictures
	IsGUI() bool

	//Updates the display and outputs it to the screen
	Update()
	DisplayWorld(Rect)
	DisplayStats()
	DisplayDesc()
	PrintMessage(string)

	Width() int
	Height() int
}
