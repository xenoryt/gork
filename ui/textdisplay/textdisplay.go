package TextDisplay

import (
	gc "github.com/rthornton128/goncurses"
	"github.com/xenoryt/gork/rect"
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

//command is a struct that contains all the necessary information to give
//an order to the main thread
type command struct {
	cmdType uiconst.CommandType
	data    interface{}
}

func (cmd command) String() string {
	if str, ok := cmd.data.(string); ok {
		return str
	}
	return ""
}

func (cmd command) TextObject() textObject {
	if obj, ok := cmd.data.(textObject); ok {
		return obj
	}
	return textObject{}
}

//textObject is to store information of objects that need to be rendered
type textObject struct {
	*gc.Window
	Drawable
}

////dialog creates a new dialog window
//type dialog struct {
//*gc.Window
//parent  *gc.Window
//visible bool
//}

//func (dlg *dialog) Draw() {
//dlg.parent.Overlay(dlg.Window)
//}
//func (dlg *dialog) Show(parent *gc.Window, msg string) error {
//dlg.parent = parent
//y, x := parent.YX()
//w, err := gc.NewWindow(3, len(msg)+2, y/2, x/2)
//if err != nil {
//return err
//}
//dlg.Window = w
//dlg.MovePrint(1, 1, msg)
//dlg.visible = true
//return nil
//}

//draw Draws the object onto the window
func draw(obj textObject, w *gc.Window) {
	x, y := obj.Loc()
	obj.MoveWindow(y, x)
	w.Overlay(obj.Window)
}

/*textDisplay is a text-based implementation of Display*/
type textDisplay struct {
	world       [][]textObject
	worldBounds rect.Rect
	objs        []textObject

	width, height int
	initialized   bool
	padding       int

	updateChan chan command
	errChan    chan error
	inputChan  chan int

	closing bool
	timeout int
}

//Init initializes the display to a specific width and height
func (display *textDisplay) Init() error {
	if !display.initialized {
		display.closing = false

		//Allow the channel to store 5 values
		display.updateChan = make(chan command)
		display.errChan = make(chan error)
		display.inputChan = make(chan int, 5)

		go display.mainLoop()

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
	}
	return nil
}

//mainLoop is a go routine that will run in the background
func (display *textDisplay) mainLoop() {
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

	height, width := stdscr.MaxYX()

	if width < uiconst.MinWidth || height < uiconst.MinHeight {
		fmt.Printf("Window does not meet minimum width and height requirements\n"+
			"currently: %dx%d, requires: %dx%d\n", width, height, uiconst.MinWidth, uiconst.MinHeight)
	}
	display.width = width
	display.height = height
	display.padding = 1

	//start main loop here
mainloop:
	for !display.closing {
		select {
		case cmd := <-display.updateChan:
			switch cmd.cmdType {
			case uiconst.CMD_TRACK:
				display.objs = append(display.objs)
			case uiconst.CMD_REMOVE:
			case uiconst.CMD_UPDATE:
			case uiconst.CMD_EXIT:
				break mainloop
			}
		default:
		}
	}

}

func (display *textDisplay) Close() {
	display.closing = true
	display.updateChan <- command{uiconst.CMD_EXIT, nil}
}

//GetInputChan returns a channel that user input will be passed into
func (display *textDisplay) GetInputChan() chan int {
	return display.inputChan
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
	w.Print(drawable.Symbol())
	display.updateChan <- command{uiconst.CMD_TRACK, textObject{w, drawable}}
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
func (display *textDisplay) LoadWorld(rect rect.Rect) {

}

//Update updates the UI. Makes changes to all the different panes accordingly.
func (display *textDisplay) Update(rect rect.Rect) {
	//display.stdscr.Erase()
	//for _, obj := range display.objs {
	//draw(obj, display.stdscr)
	//}
	//display.stdscr.Refresh()
	//gc.Nap(100)
}

func (display *textDisplay) Timeout(delay int) {
	//display.stdscr.Timeout(delay)
	display.timeout = delay
}

func (display textDisplay) Sleep(ms int) {
	gc.Nap(ms)
}

func (display *textDisplay) DisplayStats(stats string) {
}
func (display *textDisplay) DisplayDesc(desc string) {
}

func (display *textDisplay) Error(message string) {
	//w, err := gc.NewWindow(3, len(message)+2, display.height/2, display.width/2)
	//if err != nil {
	//log.Fatal("Popup error:", err)
	//log.Println("Message:", message)
	//}
	//w.Box(gc.ACS_VLINE, gc.ACS_HLINE)
	//w.MovePrint(1, 1, "Error: "+message)
	//display.stdscr.Overlay(w)
	//display.stdscr.Refresh()
	//w.Timeout(-1)
	//w.GetChar()
	//w.Delete()
	//display.stdscr.Erase()
}

func (display textDisplay) Width() int {
	return int(display.width)
}
func (display textDisplay) Height() int {
	return int(display.height)
}
