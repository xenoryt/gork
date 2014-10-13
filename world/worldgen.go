package world

/*This file covers the world generation related functions.
Mainly the Gen() function which generates everything*/

/*pangea is an undeveloped world with only the basic landscape features.
By developing story and history, pangea can evolve into a world. */
type pangea struct {
	grid   [][]Terrain
	width  int
	height int
}

func (p pangea) Init(numRows, numCols int) {
	p.grid = make([][]Terrain, numRows, numCols)
	p.width = numCols
	p.height = numRows
}

func (p pangea) String() string {
	str := ""
	for _, row := range p.grid {
		for _, item := range row {
			str += string(item.Symbol())
		}
	}
	return str
}

/*Gen generates and returns a world*/
func Gen(row, col int) World {
	/* Some algorithms I'm thinking of:
	First set the entire map to some basic terrain such as grass or forest
	depending on the biome.
	Then generate high altitude landmarks like mountains followed by
	decreasingly lower landmarks until you start generating lakes/oceans.
	Each time you generate a new type of landmark you overwrite the existing
	one if it happens to generate on the same place. Of course, the amount
	of each landmark generated should depend on the biome.

	Each landmark should also have its own generating algorithm.
	For example, mountains should probably generate in a slightly jagged
	line. Lakes should generate as batches of ovular shapes. Rivers should
	generate as thin lines that branch out.
	*/

	var world pangea

	//TEMPORARY
	return World{}
}
