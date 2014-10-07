//Package world provides necessary structures and functions for rendering,
//describing and generating worlds
package world

/*Scene is a location on the world map. It should
include brief narrative description and paths to other Scenes*/
type Scene struct {
	description string
	paths       []*Scene
}

/*World contains a grid of Scenes to visit*/
type World struct {
	inited bool
	grid   [][]Scene
}

func (world World) World(rows, cols int) {
	world.grid = make([][]Scene, rows)
	for r := 0; r < rows; r++ {
		world.grid[r] = make([]Scene, cols)
	}
}

func (world World) Size() (row, col int) {
	return len(world.grid), len(world.grid[0])
}

func (world World) GetScene() (x, y int) {
	return world.grid[y][x]
}
