package render

//TODO: create GUIObject -- once done game

/*Renderable is an object that can be drawn onto a Display*/
type Renderable interface {
	GetLoc() (x, y int)
	Render(Display)
	Visible() bool
}

/*TextObject is a text-based Renderable object*/
type TextObject struct {
	x, y          int
	width, height uint
	text          string
}
