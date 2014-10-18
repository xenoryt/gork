package render

/*Display is an object that is able to render things in a buffer
then finally display it to the user*/
type Display interface {
	//IsGUI is true iff this display can draw pictures
	IsGUI() bool
	//Renders an object onto the display
	Render(interface{}, int, int)
	//Updates the display
	Refresh()

	Width() uint16
	Height() uint16
}

/*TextDisplay is a text-based implementation of Display*/
type TextDisplay struct {
	buffer        [][]TextObject
	width, height uint16
}

func (display TextDisplay) IsGUI() bool {
	return false
}

func (display TextDisplay) Render(obj interface{}, x, y int) {
	//Check if obj is a TextObject, if so add it to the buffer
}
