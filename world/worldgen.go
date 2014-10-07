package world

/*This file covers the world generation related functions.
Mainly the Gen() function which generates everything*/

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
}
