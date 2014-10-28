package world

import (
	"math/rand"
	_ "time"
)

//Terrain just contains basic information about the land
type Terrain interface {
	//Symbol should return a character representation of the terrain
	Symbol() byte

	Name() string
	Altitude() int
	Develop() Scene

	priority() int
}

type basicTerrain struct {
	name     string
	symbol   byte
	altitude int
	zlevel   int
	opacity  int
	path     paver
}

func (t basicTerrain) Name() string {
	return t.name
}
func (t basicTerrain) Symbol() byte {
	return t.symbol
}
func (t basicTerrain) Altitude() int {
	return t.altitude
}
func (t basicTerrain) priority() int {
	return t.zlevel
}
