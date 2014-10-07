package main

import (
	"fmt"
	_ "github.com/xenoryt/gork/game"
)

type Player interface {
	Stats() string
}

type Being struct {
	name  string
	hp    int
	maxhp int
	def   int
	str   int
	//bag   []game.Item
	//held  game.Wieldable
}

func (b Being) Stats() string {
	return fmt.Sprintf("- %s -\n"+
		"hp: %d/%d\n"+
		"def: %d\n"+
		"str: %d", b.name, b.hp, b.maxhp, b.def, b.str)
}

type Human struct {
	Being
	maxhp int
}

func (h Human) Greet() string {
	return "Hi! I am " + h.name + " and my stats are:"
}

func main() {
	sprite := Being{"Spirit", 10, 10, 0, 1}
	player := Human{Being{"Player", 15, 15, 3, 4}, 40}
	player.maxhp += 10 //level up!

	beings := make([]Player, 2)
	beings[0] = sprite
	beings[1] = player
	for _, being := range beings {
		if h, ok := being.(Human); ok {
			fmt.Println(h.Greet())
		}
		fmt.Println(being.Stats())
	}
}
