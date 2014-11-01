package world

import (
	"github.com/xenoryt/gork/ui/color"
)

type Terrain struct {
	TerrType string
	Symbol   byte
	Altitude int
	Color    int16
	//How much further/less can the player see
	//if they are standing on this terrain.
	SightBonus int

	//How dense the terrain is. This number will be subtracted from
	//the player's sight
	//Difference between this and SightBonus is that this affects player
	//even if they are not standing on this tile.
	Density float32

	//If the terrain is tall, players cannot see past it unless
	//they themselves are at same altitude or higher up
	Tall bool
}

//Returns the terrain given the altitude.
func newTerrain(alt int) (t Terrain) {
	t.Altitude = alt
	switch {
	case alt < -200:
		t.TerrType = "Deep Water"
		t.Color = color.DarkBlue
		t.Symbol = '~'
		t.SightBonus = -2
	case alt < 0:
		t.TerrType = "Water"
		t.Color = color.LightBlue
		t.Symbol = '~'
	case alt < 50:
		t.TerrType = "Sand"
		t.Color = color.Yellow
		t.Symbol = ','
	case alt < 350:
		t.TerrType = "Grass"
		t.Color = color.Green
		t.Symbol = '.'
	case alt < 650:
		t.TerrType = "Forest"
		t.Color = color.Green
		t.Symbol = '&'
		t.Density = 1
	case alt < 900:
		t.TerrType = "Mountain"
		t.Color = color.Gray
		t.Symbol = '^'
		t.SightBonus = 2
		t.Tall = true
	default:
		t.TerrType = "Snowy Peak"
		t.Color = color.White
		t.Symbol = '*'
		t.SightBonus = 4
		t.Tall = true
	}
	return
}
