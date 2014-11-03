/*Package world provides necessary structures and functions for rendering,
describing and generating worlds.
This package is designed to be independent of the game*/
package world

import (
	"bytes"
	. "github.com/xenoryt/gork/rect"
)

//MapObject are objects that can be seen on the map.
type MapObject interface {
	GetLoc() (x, y int)
	SetLoc(x, y int)
	GetSymbol() byte
	Visible() bool
	Update()
}

type ZoomLevel uint

const (
	ZoomInstance ZoomLevel = 32
	ZoomNormal   ZoomLevel = 8
	ZoomWorld    ZoomLevel = 2
)

//The current zoom level of the world
var currZoom ZoomLevel = ZoomNormal
var prevZoom ZoomLevel = ZoomNormal

//Global Variables
var (
	Inited   bool
	Grid     [][]Scene
	Width    int
	Height   int
	Objects  []MapObject
	litCells []Scene
)

//this is used to temporarily store the Grid when zooming in/out
var tmpGrid *[][]Scene

var altOffset float64

//CurrentView represents which part of the world the current Grid represents.
var CurrentView Rect

func ZoomIn(view Rect) {

	switch currZoom {
	case ZoomWorld:
		currZoom = ZoomNormal
	case ZoomNormal:
		currZoom = ZoomInstance
	default:
		return
	}

	tmpGrid = &Grid

}

func Move(obj MapObject, x, y int) {
	if Grid[y][x].traversable {
		obj.SetLoc(x, y)
	}
}

//String returns string representation of the world.
//Mainly for debugging purposes.
func String() string {
	var buffer bytes.Buffer
	for row := 0; row < Height; row++ {
		for col := 0; col < Width; col++ {
			//if Grid[row][col].GetLit() {
			//	buffer.WriteString(Grid[row][col].String())
			//} else {
			//	buffer.WriteString(" ")
			//}
			buffer.WriteString(Grid[row][col].String())
		}
		buffer.WriteString("\n")
	}
	bytes := buffer.Bytes()
	//Add each object in Objects into our new temporary grid
	for _, obj := range Objects {
		x, y := obj.GetLoc()
		if Grid[y][x].GetLit() && obj.Visible() {
			bytes[(Width+1)*y+x] = obj.GetSymbol()
		}
	}
	return string(bytes)
}

//Adds a mapobject to the world
func AddObject(obj MapObject) {
	Objects = append(Objects, obj)
}
