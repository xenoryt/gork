/*Package ui handles rendering all the objects in on the screen.
As well as getting user input and updating the screen. It is also
designed to be easy to switch between text based interface to
a graphical one*/
package ui

import (
	"github.com/xenoryt/gork/errors"
	"github.com/xenoryt/gork/ui/textdisplay"
)

var display Display

//Init initializes the UI and returns the display
//Only one display can be open at a time.
func Init(gui bool) (Display, error) {
	if display != nil {
		if display.IsGUI() == gui {
			return display, nil
		} else {
			display.Close()
		}
	}

	display = TextDisplay.GetDisplay()

	//Can't think of any reason why this may fail, but just in case...
	if display == nil {
		return nil, errors.New("Failed to create instance of text display")
	}
	return display, display.Init()
}

//Close closes any open display
func Close() {
	display.Close()
}
