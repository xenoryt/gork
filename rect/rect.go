package rect

func NewRect(x, y, width, height int) Rect {
	return Rect{x, y, width, height}
}

//Camera is used to indicate which the area we want to see
type Rect struct {
	X, Y          int
	Width, Height int
}

func (r Rect) InBounds(x, y int) bool {
	if x < r.X || x >= r.X {
		return false
	}
	if y < r.Y || y >= r.Y {
		return false
	}
	return true
}

func (r *Rect) Center(x, y int) {
	r.X = x - r.Width/2
	r.Y = y - r.Height/2
}
