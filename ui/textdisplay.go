package ui

import (
	"fmt"
	"log"
	"sync"

	gc "github.com/rthornton128/goncurses"
	"github.com/xenoryt/gork/shape"
	uiconst "github.com/xenoryt/gork/ui/uiconstants"
	//"github.com/xenoryt/gork/world"
)

var (
	gcMu sync.Mutex
)

/*textDisplay is a text-based implementation of Display*/
type textDisplay struct {
	stdscr *gc.Window

	background  *gc.Pad
	worldBounds shape.Rect
	objs        []*textDrawable

	width, height int
	initialized   bool
	padding       int

	timeout int
}

//Init initializes the display to a specific width and height
func (display *textDisplay) Init() error {
	if display.initialized {
		return fmt.Errorf("Error: TextDisplay already initialized!")
	}
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

	stdscr, err := gc.Init()
	if err != nil {
		log.Fatal("init:", err)
		return err
	}
	defer gc.End()
	display.stdscr = stdscr

	gc.StartColor()
	gc.Raw(true)
	gc.Echo(false)
	gc.Cursor(0)
	stdscr.Clear()
	stdscr.Keypad(true)

	height, width := stdscr.MaxYX()

	if width < uiconst.MinWidth || height < uiconst.MinHeight {
		err := fmt.Errorf("Window does not meet minimum width and height requirements\n"+
			"currently: %dx%d, requires: %dx%d\n", width, height, uiconst.MinWidth, uiconst.MinHeight)
		log.Println(err)
		return err
	}
	display.width = width
	display.height = height
	display.padding = 1
	display.initialized = true
	return nil
}

func (display *textDisplay) Close() {
	display.stdscr.Delete()
}

//GetInput returns an integer representing the key the user pressed.
//Returns 0 if no key has been pressed
func (display *textDisplay) GetInput() int {
	gcMu.Lock()
	defer gcMu.Unlock()
	return int(display.stdscr.GetChar())
}

func (display textDisplay) Type() Type {
	return TextDisplay
}

func (display *textDisplay) DrawRune(symbol rune, x, y int) Drawable {
	// Create textDrawable and add it to objs
	drawable := newTextDrawable(symbol, x, y, gc.A_NORMAL)
	display.objs = append(display.objs, drawable)
	return drawable
}
func (display *textDisplay) DrawBG(field []string) {
}

//Update updates the UI and draws everything onto the screen
func (display *textDisplay) Update(r shape.Rect) {
	//TODO: render the background
	// Make sure we only draw the objects that are in r
	for _, obj := range display.objs {
		if r.InBounds(obj.Loc()) {
			display.stdscr.Overlay(obj.Window)
		}
	}
	display.stdscr.Refresh()
}

func (display textDisplay) Width() int {
	return int(display.width)
}
func (display textDisplay) Height() int {
	return int(display.height)
}
