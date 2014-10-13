//Package world provides necessary structures and functions for rendering,
//describing and generating worlds
package world

/*Scene is a location on the world map. It should
include brief narrative description and paths to other Scenes*/
type Scene struct {
	Terrain     //The location of the scene is also a type of terrain
	description string

	paths []*Scene

	/*Whether or not the player can see this location.
	  0 - can't see it (at all)
	  1 - Saw it before (greyed out due to fog of war)
	  2 - Can see it fully
	*/
	lit int8

	opacity int
}

//Returns true iff the player can see this scene
func (scene Scene) GetLit() bool {
	return scene.lit == 2
}

//Sets whether or not the player can see this scene
func (scene *Scene) SetLit(lit bool) {
	if lit {
		scene.lit = 2
	} else {
		scene.lit = 1
	}
}

/*Returns the level of opacity.
-1	- See through.
0	- Can't see through it.
n	- Can see n cells past this scene.
*/
func (scene Scene) Opacity() int {
	return scene.opacity
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

func (world World) GetScene(x, y int) Scene {
	return world.grid[y][x]
}
