package render

type camera struct {
	x, y          int
	width, height int
}

func (cam camera) InBounds(x, y int) bool {
	if x < cam.x || x >= cam.x {
		return false
	}
	if y < cam.y || y >= cam.y {
		return false
	}
	return true
}

func (cam camera) Center(x, y int) {
	cam.x = x - cam.width/2
	cam.y = y - cam.height/2
}
