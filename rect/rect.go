package rect

//New creates a new Rect object
func New(x, y, width, height int) Rect {
	return Rect{x, y, width, height}
}

//Rect is used to indicate a rectangular region
type Rect struct {
	X, Y          int
	Width, Height int
}

//InBounds checks if the (x,y) coordinate is within the bounds of the rectangle.
func (r Rect) InBounds(x, y int) bool {
	if x < r.X || x >= r.X {
		return false
	}
	if y < r.Y || y >= r.Y {
		return false
	}
	return true
}

//Center centers the rectangle at the (x,y) position
func (r *Rect) Center(x, y int) {
	r.X = x - r.Width/2
	r.Y = y - r.Height/2
}
