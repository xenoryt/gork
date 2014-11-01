package world

/*Scene is a location on the world map. It should
include brief narrative description and paths to other Scenes*/
type Scene struct {
	Terrain //The location of the scene is also a type of terrain

	//where the scene is on the map
	x, y int
	/*Whether or not the player can see this location.
	  0 - can't see it (at all)
	  1 - Saw it before (greyed out due to fog of war)
	  2 - Can see it fully
	*/
	lit         int8
	opacity     int
	traversable bool
}

func (scene *Scene) setLoc(x, y int) {
	scene.x, scene.y = x, y
}
func (scene *Scene) GetLoc() (x, y int) {
	return scene.x, scene.y
}

func (scene Scene) String() string {
	return string(scene.Symbol)
}

//GetLit returns true iff the player can see this scene
func (scene Scene) GetLit() bool {
	return scene.lit == 2
}

//SetLit sets whether or not the player can see this scene
func (scene *Scene) SetLit(lit bool) {
	if lit {
		scene.lit = 2
	} else {
		scene.lit = 1
	}
}

/*Opacity returns the level of opacity of the scene
-1	- See through.
0	- Can't see through it.
n	- Can see n cells past this scene.
*/
func (scene Scene) Opacity() int {
	return scene.opacity
}
