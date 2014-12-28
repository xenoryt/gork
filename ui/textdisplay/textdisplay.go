package TextDisplay

import (
	gc "code.google.com/p/goncurses"
	"github.com/xenoryt/gork/errors"
	. "github.com/xenoryt/gork/rect"
	. "github.com/xenoryt/gork/ui/drawable"
	uiconst "github.com/xenoryt/gork/ui/uiconstants"
	//"github.com/xenoryt/gork/world"
	"log"
)

var (
	tdInstance *textDisplay = new(textDisplay)
	objs       []textObject
)

//GetDisplay gets an instance of textDisplay
func GetDisplay() *textDisplay {
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

	stdscr *gc.Window
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
			return errors.New("Window does not meet minimum width and height requirements")
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

func (display *textDisplay) TrackDrawable(drawable Drawable) {
	w, err := gc.NewWindow(1, 1, 0, 0) //h,w, y,x
	w.Print(drawable.GetSymbol())
	objs = append(objs, textObject{w, drawable})
}

func (display *textDisplay) RemoveDrawable(drawable Drawable) {
	//Loop through to find the element
	for i, obj := range objs {
		if obj.Drawable == drawable {
			objs = append(objs[:i], objs[i+1:])
			break
		}
	}
}

//LoadWorld converts the current world into bunch of textObjects so it
//will  be easier to render using goncurses
func (display *textDisplay) LoadWorld(rect Rect) {

}

//Update updates the UI. Makes changes to all the different panes accordingly.
func (display *textDisplay) Update(rect Rect) {
	for i, obj := range objs {
		draw(obj, display.stdscr)
	}
}

func (display *textDisplay) DisplayStats(stats string) {
}
func (display *textDisplay) DisplayDesc(desc string) {
}

//PrintMessage will display messages. Can be useful for debugging.
func (display *textDisplay) PrintMessage(message string) {
}

func (display textDisplay) Width() int {
	return int(display.width)
}
func (display textDisplay) Height() int {
	return int(display.height)
}
