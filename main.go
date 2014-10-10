package main

import (
	"fmt"
	_ "github.com/xenoryt/gork/game"
	"github.com/xenoryt/gork/world"
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
	newWorld := world.Gen(20, 20)
	fmt.Println(newWorld)
}
