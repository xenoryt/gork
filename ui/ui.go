/*Package ui handles rendering all the objects in on the screen.
As well as getting user input and updating the screen. It is also
designed to be easy to switch between text based interface to
a graphical one*/
package ui

import "errors"

var display Display

//Init initializes the UI and returns the display
//Only one display can be open at a time.
func Init(dtype Type) (Display, error) {
	if display != nil {
		if display.Type() == dtype {
			return display, nil
		}
		display.Close()
	}

	if dtype == TextDisplay {
		display = new(textDisplay)
	} else {
		return nil, errors.New("GUI is not yet supported")
	}

	return display, display.Init()
}

//Exit closes any open display
func Exit() {
	display.Close()
}
