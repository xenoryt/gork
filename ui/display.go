package ui

import (
	. "github.com/xenoryt/gork/rect"
	. "github.com/xenoryt/gork/ui/drawable"
)

/*Display is an object that is able to render things in a buffer
then finally display it to the user*/
type Display interface {
	//For initializing the display
	Init() error

	//IsGUI is true iff this display can draw pictures
	IsGUI() bool

	//TrackDrawable converts the Drawable into a textObject and will continue
	//to render it on the world
	TrackDrawable(Drawable)

	//RemoveDrawable stops the tracking of the given Drawable.
	RemoveDrawable(Drawable)

	//Update updates the display and outputs it to the screen
	Update(Rect)

	DisplayStats(string)
	DisplayDesc(string)

	//PrintMessage will display messages. Can be useful for debugging.
	PrintMessage(string)

	Width() int
	Height() int
}
