/*Package world provides necessary structures and functions for rendering,
describing and generating worlds.
This package is designed to be independent of any other non-standard packages.*/
package world

import (
	"bytes"
)

type MapObject interface {
	GetLoc() (x, y int)
	GetSymbol() byte
	Visible() bool
}

/*World contains a grid of Scenes to visit*/
type World struct {
	Inited   bool
	Grid     [][]Scene
	Width    int
	Height   int
	Objects  []MapObject
	litCells []Scene
}

func (world *World) Init(rows, cols int) {
	world.Grid = make([][]Scene, rows)
	for r := 0; r < rows; r++ {
		world.Grid[r] = make([]Scene, cols)
	}
	world.Width = cols
	world.Height = rows
}

func (world World) String() string {
	var buffer bytes.Buffer
	for row := 0; row < world.Height; row++ {
		for col := 0; col < world.Width; col++ {
			if world.Grid[row][col].GetLit() {
				buffer.WriteString(world.Grid[row][col].String())
			} else {
				buffer.WriteString(" ")
			}
		}
		buffer.WriteString("\n")
	}
	bytes := buffer.Bytes()
	//Add each object in Objects into our new temporary grid
	for _, obj := range world.Objects {
		x, y := obj.GetLoc()
		if world.Grid[y][x].GetLit() && obj.Visible() {
			bytes[(world.Width+1)*y+x] = obj.GetSymbol()
		}
	}
	return string(bytes)
}

func (world *World) AddObject(obj MapObject) {
	world.Objects = append(world.Objects, obj)
}

func (world World) Size() (row, col int) {
	return len(world.Grid), len(world.Grid[0])
}

func (world World) GetScene(x, y int) Scene {
	return world.Grid[y][x]
}

/*Update will update the state of the world. It will check for luminous
objects and cast shadows accordingly*/
func (world World) Update() {
}
