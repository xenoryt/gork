package world

import (
	"math/rand"
	"time"
)

//generator takes a size and returns what a basic "unit" of a terrain
//should look like
type generator func(self basicTerrain, world *pangea, x, y, sizex, sizey int)

//paver takes the current direction it is heading in and returns
type paver func(curx, cury int) (int, int)

/* Scratch space:
My idea on the different terrain:
Use an equation to represent where the terrain is.
e.g. Mountain ranges can be linear/quadratic and lakes are elliptical
To introduce randomness, if it a cell on the world grid is close to
or on the equation (i.e. f(x,y) = 0), then possibly do not set the
terrain on that cell (random chance). This should cause jaggedness.
We when it is close as the "width" of the line.
Note: x and y is 0 when it is at the centre of where the terrain should
be placed.
If two terrain overlap, the terrain type with higher priority should
appear on map.

Some equations:
Mountain ranges
- f(x) = (+-)x

Lakes
- x^2/a^2 + y^2/b^2 = 1
- f(x) = (+-)sqrt(b^2 - x^2/a^2)

Forests
Same as lakes but larger scale (a and b is generally larger) and much lower
priority. Or spawn in smaller scale but spawns many instances of it.

Rivers
- f(x) = (x-a)^3
- f(x) = (x-a)(x-b)(x-c)



End scratch spaces*/

//Terrain just contains basic information about the land
type Terrain interface {
	//Symbol should return a character representation of the terrain
	Symbol() byte

	Name() string
	Altitude() int

	priority() int

	stampgen() generator
	pathgen() paver
}

type basicTerrain struct {
	name     string
	symbol   byte
	altitude int
	priority int

	stamp generator
	path  paver
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
func (t basicTerrain) stampgen() generator {
	return t.stamp
}
func (t basicTerrain) pathgen() paver {
	return t.path
}

func mountainstamp(self basicTerrain, world *pangea, x, y, sizex, sizey int) {
	/*The stamp should look something like
	 ^
	^^^
	 ^

	  ^
	 ^^^
	^^^^^
	 ^^^
	  ^
	*/

	if size == 0 {
		return
	}

	rand.Seed(time.Now().UnixNano())
	// Note: To create randomness in the shape, have a 10% chance of "failure"
	if rand.Intn(10) < 3 {
		newtrn := self
		world.grid[y][x] = newtrn
	}

	// Now recursively create more terrain
	mountainstamp(self, world, x+1, y, sizex-1, sizey)
	mountainstamp(self, world, x-1, y, sizex-1, sizey)
	mountainstamp(self, world, x, y+1, sizex, sizey-1)
	mountainstamp(self, world, x, y-1, sizex, sizey-1)
}
func mountainpath(x, y int) (int, int) {
	if rand.Intn(2) == 0 {
		return y, x
	}
	return x, y
}

func MountainTerrain() basicTerrain {
	terr := basicTerrain{
		name:     "Mountain",
		symbol:   '^',
		altitude: 1000 + rand.Intn(500) - 250,
	}
	terr.stamp = mountainstamp
	terr.path = mountainpath
	return terr
}

func WaterStamp(self basicTerrain, world *pangea, x, y, sizex, sizey int) {
	//The equation of an eclipse is x^2/a^2 + y^2/b^2 = 1
	//where 2*a is the width and 2*b is the height

	for row := 0; row < sizey; row++ {
		for col := 0; col < sizex; col++ {
			//We want the to use the values relative to (x,y) which is the
			//centre
			curx := col - x
			cury := row - y

			//We want to introduce randomness but we don't want to have holes
			//in middle of ocean!
		}
	}
}

func WaterTerrain() basicTerrain {
	terr := basicTerrain{
		name:     "Water",
		symbol:   '~',
		altitude: rand.Intn(300) - 150,
	}
	terr.stamp = mountainstamp
	terr.path = mountainpath
	return terr
}
