package world

import (
	"fmt"
	perlin "github.com/xenoryt/gork/pnoise"
	. "github.com/xenoryt/gork/rect"
	_ "io/ioutil"
	_ "strings"
)

type NoSuchTerrainErr rune

func (err NoSuchTerrainErr) Error() string {
	return "Error: No such terrain of type '" + string(err) + "'"
}

func generateScene(alt float64, indx, indy int) (scene Scene) {
	scene.Terrain = newTerrain(int(alt))
	scene.setLoc(indx, indy)
	return
}

/*Create creates a new world using the given parameters.
view is the current slice of the world to look at.
maxalt is the maximum altitude and also the minimum altitude
altoffset is the amount to add to the altitude. (i.e. offset of
	-400 will create a world with a lot of water.)
zoom is how zoomed in the map is. A zoom level of 2 will create
	a grid that is twice the size of view's width and height.
*/
func Create(view Rect, maxalt, altoffset int, zoom ZoomLevel) {
	perlin.SetPersistence(maxalt)
	width := float64(view.Width) / (float64(zoom) * 2)
	height := float64(view.Height) / (float64(zoom) * 2)
	startx := float64(view.X) + float64(view.Width)/2 - width
	starty := float64(view.Y) + float64(view.Height)/2 - height
	fmt.Println(startx, starty)

	Grid = make([][]Scene, view.Height)
	for y := view.Y; y < view.Height; y++ {
		r := int(y - view.Y)
		Grid[r] = make([]Scene, view.Width)
		for x := view.X; x < view.Width; x++ {
			c := int(x - view.X)
			worldx := float64(startx) + float64(x)/float64(zoom)
			worldy := float64(starty) + float64(y)/float64(zoom)
			altitude := perlin.Noise(worldx, worldy) + float64(altoffset)
			Grid[r][c] = generateScene(altitude, c, r)
		}
	}
	Width = view.Width
	Height = view.Height
}

/*
func loadTerrain(c byte) (Terrain, error) {
	var terr Terrain
	var err error
	switch c {
	case '^':
		terr = MountainTerrain()
	case '&':
		terr = ForestTerrain()
	case '~':
		terr = WaterTerrain()
	case '.':
		terr = GrassTerrain()
	default:
		err = NoSuchTerrainErr(c)
	}
	return terr, err
}

func LoadWorld(file string) (World, error) {
	dat, err := ioutil.ReadFile(file)
	var world World
	if err == nil {
		lines := strings.Split(string(dat), "\n")
		numrows := len(lines) - 1
		numcols := len(lines[0])

		world.Init(numrows, numcols)

		for row := 0; row < numrows; row++ {
			for col := 0; col < numcols; col++ {
				terr, err := loadTerrain(lines[row][col])
				if err == nil {
					world.Grid[row][col] = terr.Develop()
					world.Grid[row][col].SetLit(false)
				}
			}
		}
	}

	return world, err
}
*/
