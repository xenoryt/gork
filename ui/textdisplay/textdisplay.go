package TextDisplay

import (
	gc "code.google.com/p/goncurses"
	"github.com/xenoryt/gork/errors"
	. "github.com/xenoryt/gork/rect"
	. "github.com/xenoryt/gork/ui/drawable"
	uiconst "github.com/xenoryt/gork/ui/uiconstants"
	//"github.com/xenoryt/gork/world"
	"fmt"
	"log"
)

var (
	tdInstance *textDisplay
)

//GetDisplay gets an instance of textDisplay
func GetDisplay() *textDisplay {
	if tdInstance == nil {
		tdInstance = new(textDisplay)
	}
	return tdInstance
}

//textObject is to store information of objects that need to be rendered
type textObject struct {
	*gc.Window
	Drawable
}

//draw Draws the object onto the window
func draw(obj textObject, w *gc.Window) {
	x, y := obj.GetLoc()
	obj.MoveWindow(y, x)
	w.Overlay(obj.Window)
}

/*textDisplay is a text-based implementation of Display*/
type textDisplay struct {
	world       [][]textObject
	worldBounds Rect
	objs        []textObject

	width, height int
	initialized   bool
	padding       int

	stdscr  *gc.Window
	closing bool
}

//Init initializes the display to a specific width and height
func (display *textDisplay) Init() error {
	if !display.initialized {
		display.closing = false

		stdscr, err := gc.Init()
		if err != nil {
			log.Fatal("init:", err)
		}

		gc.StartColor()
		gc.Raw(true)
		gc.Echo(false)
		gc.Cursor(0)
		stdscr.Clear()
		stdscr.Keypad(true)
		stdscr.Timeout(0)

		height, width := stdscr.MaxYX()

		if width < uiconst.MinWidth || height < uiconst.MinHeight {
			display.Close()
			return fmt.Errorf("Window does not meet minimum width and height requirements\n"+
				"currently: %dx%d, requires: %dx%d", width, height, uiconst.MinWidth, uiconst.MinHeight)
		}
		display.stdscr = stdscr
		display.width = width
		display.height = height
		display.padding = 1

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
	return errors.New("Display already initialized!")
}

func (display textDisplay) Close() {
	display.closing = true
	gc.End()
}

//GetInputChan returns a channel that user input will be passed into
func (display *textDisplay) GetInput() int {
	return int(display.stdscr.GetChar())
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

func (display *textDisplay) TrackDrawable(drawable Drawable) error {
	w, err := gc.NewWindow(1, 1, 0, 0) //h,w, y,x
	if err != nil {
		return err
	}
	w.Print(drawable.GetSymbol())
	display.objs = append(display.objs, textObject{w, drawable})
	return nil
}

func (display *textDisplay) RemoveDrawable(drawable Drawable) error {
	//Loop through to find the element
	for i, obj := range display.objs {
		if obj.Drawable == drawable {
			display.objs = append(display.objs[:i], display.objs[i+1:]...)
			break
		}
	}
	return nil
}

//LoadWorld converts the current world into bunch of textObjects so it
//will  be easier to render using goncurses
func (display *textDisplay) LoadWorld(rect Rect) {

}

//Update updates the UI. Makes changes to all the different panes accordingly.
func (display *textDisplay) Update(rect Rect) {
	for _, obj := range display.objs {
		draw(obj, display.stdscr)
	}
	display.stdscr.Refresh()
}

func (display textDisplay) Sleep(ms int) {
	gc.Nap(ms)
}

func (display *textDisplay) DisplayStats(stats string) {
}
func (display *textDisplay) DisplayDesc(desc string) {
}

//Printwill display messages. Can be useful for debugging.
func (display *textDisplay) Print(message string) {
	display.stdscr.MovePrint(0, 0, message)
	display.stdscr.Refresh()
}

func (display textDisplay) Width() int {
	return int(display.width)
}
func (display textDisplay) Height() int {
	return int(display.height)
}
