package world

import (
	_ "fmt"
	"github.com/xenoryt/gork/path"
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

func LoadWorld(view Rect) {
	width := float64(view.Width) / (float64(currZoom) * 2)
	height := float64(view.Height) / (float64(currZoom) * 2)
	startx := float64(view.X) + float64(view.Width)/2 - width
	starty := float64(view.Y) + float64(view.Height)/2 - height

	Grid = make([][]Scene, view.Height)
	for y := 0; y < view.Height; y++ {
		r := int(y - view.Y)
		Grid[r] = make([]Scene, view.Width)
		for x := 0; x < view.Width; x++ {
			c := int(x - view.X)
			worldx := float64(startx) + float64(x)/float64(currZoom)
			worldy := float64(starty) + float64(y)/float64(currZoom)
			altitude := perlin.Noise(worldx, worldy) + float64(altOffset)
			Grid[r][c] = generateScene(altitude, c, r)
		}
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func genPath(x, y int) {
	var scene *Scene = &Grid[y][x]
	scene.Terrain = pathTerrain(scene.Altitude)
	//TODO: possibly flatten out the area
	return
}
func getPaths(x, y int) []path.Point {
	var paths []path.Point
	if x > 0 {
		paths = append(paths, &Grid[y][x-1])
	}
	if x < Width-1 {
		paths = append(paths, &Grid[y][x+1])
	}
	if y > 0 {
		paths = append(paths, &Grid[y-1][x])
	}
	if y < Height-1 {
		paths = append(paths, &Grid[y+1][x])
	}
	return paths
}
func pathCost(x0, y0, x1, y1 int) int {
	//s0 := &Grid[y0][x0]
	s1 := &Grid[y1][x1]
	//moving from one tile to another is 10 so we don't have to deal
	//with decimals
	penalty := 10
	switch s1.Symbol {
	case '~':
		penalty += abs(s1.Altitude)/50 + 5
		return penalty
	case '&':
		penalty += 2
	case '^':
		penalty += 5
	case '*':
		penalty += 15
	}

	//if s1.Altitude > s0.Altitude {
	//	penalty += (s1.Altitude - s0.Altitude) / 3
	//}
	return penalty
}

//heuristic using manhatten distance
func pathHeur(x0, y0, x1, y1 int) int {
	return abs(x0-x1) + abs(y0-y1)
}
func DrawPath(startx, starty, endx, endy int) {
	//path.Line(startx, starty, endx, endy, genPath)
	start := &Grid[starty][startx]
	end := &Grid[endy][endx]
	path.WeightedAStar(start, end, pathCost, pathHeur, getPaths, genPath, 10)
	/*
		perlin.SetPersistence(2)
		perlin.SetNumOctaves(4)
		slope := float64(starty-endy) / float64(startx-endx)
		//x, y := startx, starty
		for x := startx; x < endx; x++ {
			n := perlin.IntNoise(x, y)
			fmt.Println(n)
			col := int(x + n)
			row := int(y + n)
			Grid[row][col] = genPath(col, row)
		}
	*/
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
