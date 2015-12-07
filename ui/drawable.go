package ui

//Drawable is an object that can be rendered on the screen.
//This interface allows manipulating objects that are drawn on the display
type Drawable interface {
	Loc() (x, y int)
	Symbol() rune

	//Moves the object by x,y
	Move(x, y int)
	//Moves the Drawable to location (x,y)
	MoveTo(x, y int)

	Delete()
}
