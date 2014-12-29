/*Package ui handles rendering all the objects in on the screen.
As well as getting user input and updating the screen. It is also
designed to be easy to switch between text based interface to
a graphical one*/
package ui

import (
	"github.com/xenoryt/gork/errors"
	"github.com/xenoryt/gork/rect"
	"github.com/xenoryt/gork/ui/textdisplay"
	"github.com/xenoryt/gork/world"
)

var display Display
var gui bool
var cam rect.Rect

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

//Update updates the UI
func Update(cam rect.Rect) {
	display.Update(cam)
}

//Init initializes the UI
func Init() error {
	display = TextDisplay.GetDisplay()
	gui = false
	if display == nil {
		return errors.New("Failed to initialize window")
	}
	return nil
}

//Close terminates the UI
func Close() {
	display.Close()
}
