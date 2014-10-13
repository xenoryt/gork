package world

import (
	"io/ioutil"
	"strings"
)

type NoSuchTerrainErr rune

func (err NoSuchTerrainErr) Error() string {
	return "Error: No such terrain of type '" + string(err) + "'"
}

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
				}
			}
		}
	}

	return world, err
}
