package main

import (
	"fmt"
	"github.com/xenoryt/gork/fov"
	_ "github.com/xenoryt/gork/game"
	"github.com/xenoryt/gork/world"
	"strings"
)

const mapfile = "data/map.txt"

var player Player

func scene2cell(w *world.World) [][]fov.Cell {
	newgrid := make([][]fov.Cell, len(w.Grid))
	for i := 0; i < len(w.Grid); i++ {
		newgrid[i] = make([]fov.Cell, len(w.Grid[i]))
		for n := 0; n < len(w.Grid[i]); n++ {
			newgrid[i][n] = fov.Cell(&w.Grid[i][n])
			//w.Grid[i][n].SetLit(true)
		}
	}
	return newgrid
}

func main() {
	worldmap, err := world.LoadWorld(mapfile)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(worldmap.Width, worldmap.Height)

	//Init player
	player.x = worldmap.Width / 2
	player.y = worldmap.Height / 2
	player.sight = 6

	//Add the player to the world.
	worldmap.AddObject(&player)

	var inp string
	for {
		fov.CastShadows(scene2cell(&worldmap), player.x, player.y, player.sight)
		//Print the map
		fmt.Println(worldmap)
		fmt.Println(player.x, player.y)

		//Get user input
		fmt.Scanf("%s", &inp)
		inp = strings.ToLower(inp)
		if inp == "q" {
			break
		}

		//Handle movement
		switch strings.ToLower(inp) {
		case "w":
			player.Move(NORTH)
		case "a":
			player.Move(WEST)
		case "s":
			player.Move(SOUTH)
		case "d":
			player.Move(EAST)
		}

		fov.Blackout()
	}
}
