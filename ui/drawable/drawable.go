package drawable

//Drawable is anything that you can draw onto the screen.
type Drawable interface {
	Loc() (x, y int)
	Symbol() byte
}
