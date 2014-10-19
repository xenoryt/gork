package render

import (
	"fmt"
	"github.com/xenoryt/gork/world"
)

type pane struct {
	buffer []byte
	width  int
	height int
}

func newPane(buffer []byte, bwidth, x, y, w, h int) pane {
	start := y*bwidth + x
	end := (y+h)*bwidth + x + w
	return pane{buffer[start:end], w, h}
}

func (p pane) inBounds(x, y int) bool {
	if x < 0 || x >= p.width {
		return false
	}
	if y < 0 || y >= p.height {
		return false
	}
	return true
}
func (p *pane) print(text []byte, x, y int) error {
	if p.inBounds(x, y) {
		for i := range text {
			p.buffer[i+x] = text[i]
		}
		return nil
	}
	return GenericError("Out of bounds: (" + string(x) + ", " + string(y) + ")")
}

/*TextDisplay is a text-based implementation of Display*/
type TextDisplay struct {
	buffer    []byte
	worldpane pane
	statpane  pane
	descpane  pane

	width, height uint16
	initialized   bool
	padding       int
}

//Init initializes the display to a specific width and height
func (display TextDisplay) Init(width, height uint16) error {
	if !display.initialized {
		if width < _MinWidth || height < _MinHeight {
			return InitError("Window does not meet minimum width and height requirements")
		}
		display.width = width
		display.height = height
		display.padding = 1

		display.buffer = make([]byte, height*width)

		//We want to split up the screen into 3 sections like so:
		//
		//	-------------------------
		//	|            |          |
		//	| (1) world  | (2) desc |
		//	|            |          |
		//	|            |          |
		//	-------------------------
		//	| (3) status            |
		//	-------------------------
		//
		//	(1) is where the world is drawn
		//	(2) is where the description of the world is displayed
		//	(3) is where the status of the player is displayed
		//
		//	The world pane will be 24 chars wide and have 15 rows.
		//  (2) width will the the remaining space, height same as (1)
		//  (3)

		display.initialized = true
		return nil
	}
	return InitError("Display already initialized!")
}

func (display TextDisplay) IsGUI() bool {
	return false
}

/*
func (display *TextDisplay) Render(obj Renderable) error {
	//Check if obj is a TextObject, if so add it to the buffer
	if txtobj, ok := obj.(TextObject); ok {
		return display.put(txtobj)
	}
	return GenericError("Invalid object recieved: not a TextObject")
}*/

func (display *TextDisplay) DrawWorld(wmap world.World, cam camera) {
	//Get the current view of the world and put it into the world pane
}

func (display TextDisplay) Update() {
	for i := 0; i < display.Height(); i++ {
		fmt.Println(string(display.buffer[i : display.Width()*(i+1)]))
	}
}

func (display TextDisplay) Width() int {
	return int(display.width)
}
func (display TextDisplay) Height() int {
	return int(display.height)
}
