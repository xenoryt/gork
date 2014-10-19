package render

//TODO: create GUIObject -- once done game

/*Renderable is an object that can be drawn onto a Display*/
type Renderable interface {
	GetLoc() (int, int)
	GetSize() (int, int)
	Visible() bool
	Text() []byte
}

/*TextObject is a text-based Renderable object*/
type TextObject struct {
	x, y          int
	width, height int
	text          string
	visible       bool
}

func (txtobj *TextObject) New(text string, w, h int) {
	txtobj.text = text
	txtobj.visible = true
	txtobj.x, txtobj.y = 0, 0
	txtobj.width = w
	txtobj.height = h
}

func (txtobj TextObject) GetLoc() (int, int) {
	return txtobj.x, txtobj.y
}
func (txtobj TextObject) GetSize() (int, int) {
	return txtobj.width, txtobj.height
}

func (txtobj TextObject) Visible() bool {
	return txtobj.visible
}

func (txtobj TextObject) Text() []byte {
	return []byte(txtobj.text)
}
