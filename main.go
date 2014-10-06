package main

import "fmt"

type Object interface {
	View() string
	Embed(*Object) (string, error)
}
type Surface interface {
	Recieve(*Placeable) error
	Remove(*Placeable) error
}
type Placeable interface {
	Place(*Surface) (string, error)
}
type Useable interface {
	Use() error
}
type Item interface {
	Placeable
	Useable
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

func main() {
	sprite := Being{"Spirit", 10, 10, 0, 1}
	fmt.Println(sprite.Stats())
	player := Human{Being{"Player", 15, 15, 3, 4}}
	player.maxhp += 10 //level up!
	fmt.Println(player.Stats())
}
