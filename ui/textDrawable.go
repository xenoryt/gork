package ui

import (
	"log"

	gc "github.com/rthornton128/goncurses"
)

func newTextDrawable(symbol rune, x, y int, attr gc.Char) *textDrawable {
	w, err := gc.NewWindow(1, 1, 1, 1)
	if err != nil {
		log.Println("newTextDrawable:", err)
	}
	w.AttrOn(attr)
	w.Print(symbol)
	return &textDrawable{w, symbol, x, y, false}
}

//textDrawable implements draw.Drawable for text based display using goncurses
type textDrawable struct {
	*gc.Window
	symbol  rune
	x, y    int
	expired bool
}

func (td textDrawable) Loc() (x, y int) {
	return td.x, td.y
}

func (td textDrawable) Symbol() rune {
	return td.symbol
}

func (td *textDrawable) Move(x, y int) {
	td.x += x
	td.y += y
}
func (td *textDrawable) MoveTo(x, y int) {
	td.x = x
	td.y = y
}

func (td *textDrawable) Delete() {
	td.expired = true
}
