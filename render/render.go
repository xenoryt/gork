/*Package render handles rendering all the objects in on the screen.
It is designed to be easily modified to handle GUI if that is ever needed.*/
package render

import (
	"github.com/xenoryt/gork/world"
)

var display Display
var gui bool
var cam camera

func Init(x, y, int, width, height, uint16, textbased bool) error {
	if textbased {
		//display = TextDisplay{}
		gui = false
	}
	gui = true
	return nil
}

func CenterView(x, y int) {
	cam.Center(x, y)
}

func DrawWorld(wmap world.World) {
	display.DrawWorld(wmap, cam)
}
