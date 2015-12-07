package ui

import (
	"github.com/xenoryt/gork/shape"
)

//Type specifies the type of display
type Type int

const (
	TextDisplay Type = iota
	GraphicalDisplay
)

/*Display is the interface that wraps all the functionality for displaying
things to the user.
*/
type Display interface {
	Init() error
	//Close closes the display
	Close()

	//GetInput returns a channel that keypresses will be sent through
	GetInput() int

	//Type returns the display type
	Type() Type

	//DrawChar creates a Drawable at the location (x, y)
	DrawRune(char rune, x, y int) Drawable

	//DrawBG stores the background
	DrawBG(bg []string)

	//Update updates the display and outputs it
	//to the screen. The Rect passed in is the
	//region of the world to display.
	Update(shape.Rect)

	// DisplayStats(string)
	// DisplayDesc(string)

	Width() int
	Height() int
}
