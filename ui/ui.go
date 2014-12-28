/*Package ui handles rendering all the objects in on the screen.
As well as getting user input and updating the screen. It is also
designed to be easy to switch between text based interface to
a graphical one*/
package ui

import (
	. "github.com/xenoryt/gork/rect"
	"github.com/xenoryt/gork/ui/textdisplay"
	"github.com/xenoryt/gork/world"
)

var display Display
var gui bool
var cam Rect

//This channel is used to send user input
var input chan byte

//This channel is used to recieve update information
var update chan int

//GetInputChan gets a channel that contains user input
func GetInputChan() chan byte {
	return input
}

//LoadWorld loads the world into a buffer
func LoadWorld([][]world.Scene) {
}

//DisplayWorld displays a section of the world on the screen.
func DisplayWorld(cam Rect) {
	display.DisplayWorld(cam)
}

//Init initializes the UI
func Init(textbased bool) error {
	if textbased {
		display = TextDisplay.GetDisplay()
	}
	gui = !true
	return nil
}
