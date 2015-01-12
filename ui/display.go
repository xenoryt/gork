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

	GetInputChan() chan int

	//IsGUI returns true iff the ui is graphical
	IsGUI() bool

	//LoadWorld loads the world into memory. This should
	//be called before doing anything.
	//TODO: Possibly merge this into Init()?
	LoadWorld(worldmap [][]Drawable) error

	//TrackDrawable converts the Drawable into
	//a textObject and will continue to render it
	//on the world.
	//NOTE: Please pass in a POINTER to the object
	TrackDrawable(Drawable) error

	//RemoveDrawable stops the tracking of the given
	//Drawable.
	RemoveDrawable(Drawable) error

	//Update updates the display and outputs it
	//to the screen. The Rect passed in is the
	//region of the world to display.
	Update(Rect)

	//SetView sets where the camera is looking at
	SetView(Rect)

	//Timeout sets the amount of time to wait for a keypress
	Timeout(int)

	//Sleep(int)

	DisplayStats(string)
	DisplayDesc(string)

	//Error will display a message in a popup window.
	//Can be useful for debugging.
	Error(string)

	Width() int
	Height() int
}
