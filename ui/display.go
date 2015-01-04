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

	//Close closes the display
	Close()

	GetInput() int

	//IsGUI returns true iff the ui is graphical
	IsGUI() bool

	//TrackDrawable converts the Drawable into
	//a textObject and will continue to render it
	//on the world
	TrackDrawable(Drawable) error

	//RemoveDrawable stops the tracking of the given
	//Drawable.
	RemoveDrawable(Drawable) error

	//Update updates the display and outputs it
	//to the screen. The Rect passed in is the
	//region of the world to display.
	Update(Rect)

	//Timeout sets the amount of time to wait for a keypress
	Timeout(int)

	//Sleep(int)

	DisplayStats(string)
	DisplayDesc(string)

	//Print will display messages.
	//Can be useful for debugging.
	Print(string)

	Width() int
	Height() int
}
