package main

import "fmt"

/* Object is the basic type that anything that
can be physically "sensed" should implement.
 - View() should return a description of the of object. */
type Object interface {
	View() string
}

/* Surface is anything that can have something
physically placed on top of it. */
type Surface interface {
	Recieve(*Placeable) error
	Remove(*Placeable) error
}

/* Placeable is any object that can also be placed (or embedded)
onto a surface. Placeable objects can be taken as well. */
type Placeable interface {
	Place(*Surface) (string, error)
	Take() (string, error)
}

/* Usable is any object that can be used or activated. */
type Useable interface {
	Use() error
}

/* Item is an object that can be placed and used */
type Item interface {
	Object
	Placeable
	Useable
}

type Player interface {
	Stats() string
}

type Being struct {
	name  string
	hp    int
	maxhp int
	def   int
	str   int
}

func (b Being) Stats() string {
	return fmt.Sprintf(
		"- %s -\n"+
			"hp: %d/%d\n"+
			"def: %d\n"+
			"str: %d", b.name, b.hp, b.maxhp, b.def, b.str)
}

type Human struct {
	Being
}

func (h Human) Greet() string {
	return "Hi! I am " + h.name + " and my stats are:"
}

func main() {
	sprite := Being{"Spirit", 10, 10, 0, 1}
	player := Human{Being{"Player", 15, 15, 3, 4}}
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
