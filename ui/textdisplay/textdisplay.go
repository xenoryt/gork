package TextDisplay

import (
	gc "code.google.com/p/goncurses"
	"fmt"
	. "github.com/xenoryt/gork/error"
	. "github.com/xenoryt/gork/rect"
	uiconst "github.com/xenoryt/gork/ui/uiconstants"
	"github.com/xenoryt/gork/world"
	"log"
)

var (
	tdInstance textDisplay
)

//Get gets an instance of textDisplay
func Get() textDisplay {
	return tdInstance
}

type pane struct {
	buffer   [][]byte
	width    int
	height   int
	caretPos int
}

func newPane(buffer []byte, bwidth, x, y, w, h int) pane {
	panelBuffer := make([][]byte, h)
	for i := 0; i < h; i++ {
		start := (y+i)*bwidth + x
		end := (y+i)*bwidth + x + w
		panelBuffer[i] = buffer[start:end]
	}
	return pane{panelBuffer, w, h, 0}
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

func (p *pane) put(index int, char byte) {
	p.buffer[index/p.width][index%p.width] = char
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
		caret := i + bpos
		//Make sure we're not out of bounds
		if caret >= p.width*p.height {
			return GenericError("Error: couldn't print everything. Pane overflowed!")
		}

		//Add the character to the buffer
		switch text[i] {
		case '\n':
			bpos += p.width % i
		case '\t':
			bpos += 4 - (i % 4)
			//if (i+bpos) % p.width < 4 {
			//	//We went onto another line.
			//}
		default:
			p.put(caret, text[i])
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
		p.put(*i, ' ')
	}
}

/*textDisplay is a text-based implementation of Display*/
type textDisplay struct {
	buffer    []byte
	worldpane pane
	statpane  pane
	descpane  pane

	width, height int
	initialized   bool
	padding       int
}

//Init initializes the display to a specific width and height
func (display textDisplay) Init() error {
	if !display.initialized {

		stdscr, err := gc.Init()
		if err != nil {
			log.Fatal("init:", err)
		}
		defer gc.End()

		gc.StartColor()
		gc.Raw(true)
		gc.Echo(false)
		gc.Cursor(0)
		stdscr.Clear()
		stdscr.Keypad(true)

		width, height := stdscr.MaxYX()

		if width < uiconst.MinWidth || height < uiconst.MinHeight {
			return InitError("Window does not meet minimum width and height requirements")
		}
		display.width = width
		display.height = height
		display.padding = 1

		display.buffer = make([]byte, height*width)

		//We want to split up the screen into 3 sections like so:
		//
		//	+-------------+----------+
		//	|             | (3) stat |
		//	| (1) world   +----------+
		//	|             |          |
		//	|             | (2) desc |
		//	|             |          |
		//	|             |          |
		//	+-------------+----------+
		//
		//	(1) is where the world is drawn
		//	(2) is where the description of the world is displayed
		//	(3) is where the status of the player is displayed
		//
		//	The world pane and status will be 25 chars wide while the
		//  terrain and scene descriptions will take the remaining space
		//  The status and terrain will take up 8 rows
		//

		//Set the the panes

		display.initialized = true
		return nil
	}
	return InitError("Display already initialized!")
}

func (display textDisplay) IsGUI() bool {
	return false
}

/*
func (display *textDisplay) Render(obj Renderable) error {
	//Check if obj is a TextObject, if so add it to the buffer
	if txtobj, ok := obj.(TextObject); ok {
		return display.put(txtobj)
	}
	return GenericError("Invalid object recieved: not a TextObject")
}*/

func (display *textDisplay) DrawWorld(wmap [][]world.Scene, cam Rect) {
	//Get the current view of the world and put it into the world pane
}

func (display textDisplay) Update() {
	for i := 0; i < display.Height(); i++ {
		fmt.Println(string(display.buffer[i : display.Width()*(i+1)]))
	}
}

func (display textDisplay) Width() int {
	return int(display.width)
}
func (display textDisplay) Height() int {
	return int(display.height)
}
