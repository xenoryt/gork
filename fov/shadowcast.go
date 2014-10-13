/*Implements a way to get the player's current field of view
using technniques such as shadowcasting.*/
package fov

import (
//This is needed as we will be changing light values in the world
//"github.com/xenoryt/gork/world"
)

type slope float64

//Define an arbitrarily large number to represent "NaN"
const _Nan slope = 99991

func (s slope) isNan() bool {
	return s == _Nan
}

type Cell interface {
	GetLit() bool
	SetLit(bool) //This should be a pointer method

	//Opacity returns how see through the cell is.
	//-1 - can see through it
	//0  - not see through
	//1  - can see one cell past it
	//n  - can see n cells past it
	Opacity() int
}

//Some helper functions
func isOpaque(cell Cell) bool {
	return cell.Opacity() == 0
}
func calcSlope(startx, starty int, endx, endy int) slope {
	dx := float64(starty - endy)
	dy := float64(startx - endx)
	if dx == 0 {
		return _Nan
	}
	return slope(dy / dx)
}
func calcInvSlope(startx, starty int, endx, endy int) slope {
	dx := float64(starty - endy)
	dy := float64(startx - endx)
	if dy == 0 {
		return _Nan
	}
	return slope(dx / dy)
}

//A struct to store data
type worldData struct {
	grid     [][]Cell
	maxdepth int
	xcenter  int
	ycenter  int
}

//http://www.roguebasin.com/index.php?title=Computing_LOS_for_Large_Areas
//http://www.roguebasin.com/index.php?title=FOV_using_recursive_shadowcasting
func CastShadows(grid [][]Cell, x, y int) {
	//Create worldData object and scan the world
	//using goroutines.
}

func (data *worldData) scanRow(depth int, startslope, endslope slope, ch chan bool) {
	if depth >= data.maxdepth {
		ch <- true //Notify other goroutines we're done.
		return
	}

	//The direction we will be recursing
	//If the initial depth is less than 0, then we want to recurse
	//downwards.
	dir := 1
	if depth < 0 {
		dir = -dir
	}

	//The row we are currently on.
	row := data.ycenter - depth

	startx := int(startslope*slope(depth) + slope(data.xcenter) + 0.5)
	endx := int(endslope*slope(depth) + slope(data.xcenter) + 0.5)

	//how much to increment counter by
	step := 1
	chsize := endx - startx
	if endx < startx {
		step = -1
		chsize = -chsize //make sure chsize >= 0
	}

	newChan := make(chan bool, chsize)

	//This is true iff all the previous cells before current is a blocker
	blockerFound := false

	//using != in the condition since endx can be greater or smaller than col
	for col := startx; col != endx; col += step {
		var cell Cell = data.grid[row][col]
		if blockerFound {
			if !isOpaque(cell) {
				//This is the last cell in the "wall"
				//Calculate new start slope
				newslope := calcInvSlope(col+step, row, data.xcenter, data.ycenter)
				go data.scanRow(depth+1*dir, newslope, endslope, newChan)
			} else { // this cell is the first cell in the "wall"
				newslope := calcInvSlope(col-step, row, data.xcenter, data.ycenter)
				go data.scanRow(depth+1*dir, startslope, newslope, newChan)
			}
		} else {
			cell.SetLit(true)
			newChan <- true
		}
	}

	//Wait for the recursive goroutines we called to finish
	for i := 0; i < chsize; i++ {
		<-newChan
	}

	//Notify that we're done
	ch <- true
}
