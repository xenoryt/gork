package render

import (
	"fmt"
	"github.com/xenoryt/gork/world"
)

type pane struct {
	buffer   []byte
	width    int
	height   int
	caretPos int
}

func newPane(buffer []byte, bwidth, x, y, w, h int) pane {
	start := y*bwidth + x
	end := (y+h)*bwidth + x + w
	return pane{buffer[start:end], w, h, 0}
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

//Prints text to the pane
func (p *pane) print(text []byte) error {
	//We want to start from the previous print's position
	bpos := p.caretPos

	//We want to save the new position of the caret after printing
	var i int = 0
	defer func() {
		p.caretPos = bpos + i
	}()

	for i = range text {
		//Make sure we're not out of bounds
		if i+bpos >= len(p.buffer) {
			return GenericError("Error: couldn't print everything. Pane overflowed!")
		}

		//Add the character to the buffer
		switch text[i] {
		case '\n':
			bpos += p.width % i
		case '\t':
			bpos += 4 - (i % 4)
		default:
			p.buffer[i+bpos] = text[i]
		}
	}
	return nil
}

//Clears the pane of any text.
func (p *pane) clear() {
	//We shall start from the current caret position and keep
	//rolling back until we hit the start again.
	var i *int = &p.caretPos
	for ; *i >= 0; *i-- {
		p.buffer[*i] = ' '
	}
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
