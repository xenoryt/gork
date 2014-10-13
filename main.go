package main

import (
	"fmt"
	_ "github.com/xenoryt/gork/fov"
	_ "github.com/xenoryt/gork/game"
	"github.com/xenoryt/gork/world"
)

const mapfile = "/home/xenoryt/programming/gocode/src/github.com/xenoryt/gork/map1.txt"

var player Player

func main() {
	worldmap, err := world.LoadWorld(mapfile)
	if err != nil {
		fmt.Println(err)
		return
	}

	//Init player
	player.x = worldmap.Width / 2
	player.y = worldmap.Height / 2

	//Something basic
	fmt.Println(worldmap)
}
