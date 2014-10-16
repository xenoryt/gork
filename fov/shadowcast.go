/*Implements a way to get the player's current field of view
using technniques such as shadowcasting.*/
package fov

import (
	//This is needed as we will be changing light values in the world
	//"github.com/xenoryt/gork/world"
	"fmt"
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
func calcSlope(inversed bool, startx, starty int, endx, endy int) slope {
	dx := float64(starty - endy)
	dy := float64(startx - endx)
	if inversed {
		if dy != 0 {
			return slope(dx / dy)
		} else {
			return 0
		}
	} else {
		if dx == 0 {
			return _Nan
		}
		return slope(dy / dx)
	}
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
/*CastShadows lights up cells in the given grid assuming that the object
located at (x,y) is luminous*/
func CastShadows(grid [][]Cell, x, y int) {
	//First create the worldData object
	data := worldData{
		grid:     grid,
		maxdepth: 6,
		xcenter:  x,
		ycenter:  y,
	}
	ch := make(chan bool)

	/* We scan 8 different sectors
	\1|2/
	8\|/3
	-----
	7/|\4
	/6|5\
	*/

	//scan sector 1
	go data.scanRow(1, -1, 0, ch)
	//scan sector 2
	go data.scanRow(1, 1, 0, ch)
	//scan sector 6
	go data.scanRow(-1, 1, 0, ch)
	//scan sector 5
	go data.scanRow(-1, -1, 0, ch)

	//scan sector 3
	go data.scanCol(1, 1, 0, ch)
	//scan sector 3
	go data.scanCol(1, -1, 0, ch)
	//scan sector 8
	go data.scanCol(-1, -1, 0, ch)
	//scan sector 7
	go data.scanCol(-1, 1, 0, ch)

	//The (x,y) location itself should be lit.
	grid[y][x].SetLit(true)

	//Now we wait until it has finished scanning
	for i := 0; i < 8; i++ {
		fmt.Println("waiting", i)
		<-ch
	}
}

/*
func (data *worldData) scan(srow bool, start, end, ind int, f func(int, slope, slope, chan bool)) {
	var iter, row, col int
	if srow == -1 {
		iter = &row
		col = ind
	} else {
		iter = &col
		row = ind
	}
	for *iter = start; *iter != endx; *iter += step {
		var cell Cell = data.grid[row][col]
		if blockerFound {
			if !isOpaque(cell) {
				//This is the last cell in the "wall"
				//Calculate new start slope
				newslope := calcSlope(srow, col+step, row, data.xcenter, data.ycenter)
				go f(depth+1*dir, newslope, endslope, newChan)
			} else { // this cell is the first cell in the "wall"
				newslope := calcSlope(srow, col-step, row, data.xcenter, data.ycenter)
				go f(depth+1*dir, startslope, newslope, newChan)
			}
		} else {
			cell.SetLit(true)
			newChan <- true
		}
	}
}
*/

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

	//Double check we are still within bounds.
	if row < 0 || row >= len(data.grid) || startx < 0 ||
		startx >= len(data.grid[0]) {

		ch <- true
		return
	}

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
				newslope := calcSlope(true, col+step, row, data.xcenter, data.ycenter)
				go data.scanRow(depth+1*dir, newslope, endslope, newChan)
				cell.SetLit(true)
			}
		} else if isOpaque(cell) {
			// this cell is the first cell in the "wall"
			newslope := calcSlope(true, col-step, row, data.xcenter, data.ycenter)
			go data.scanRow(depth+1*dir, startslope, newslope, newChan)
			blockerFound = true
		} else {
			cell.SetLit(true)
			chsize -= 1
		}
	}

	//Wait for the recursive goroutines we called to finish
	for i := 0; i < chsize; i++ {
		<-newChan
	}

	//Notify that we're done
	ch <- true
}
func (data *worldData) scanCol(depth int, startslope, endslope slope, ch chan bool) {
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
	col := data.xcenter + depth

	starty := int(startslope*slope(depth) + slope(data.ycenter) + 0.5)
	endy := int(endslope*slope(depth) + slope(data.ycenter) + 0.5)

	//Double check we are still within bounds.
	if col < 0 || col >= len(data.grid) || starty < 0 ||
		starty >= len(data.grid[0]) {

		ch <- true
		return
	}

	//how much to increment counter by
	step := 1
	chsize := endy - starty
	if endy < starty {
		step = -1
		chsize = -chsize //make sure chsize >= 0
	}

	newChan := make(chan bool, chsize)

	//This is true iff all the previous cells before current is a blocker
	blockerFound := false

	//using != in the condition since endx can be greater or smaller than col
	for row := starty; row != endy; row += step {
		var cell Cell = data.grid[row][col]
		if blockerFound {
			if !isOpaque(cell) {
				//This is the last cell in the "wall"
				//Calculate new start slope
				newslope := calcSlope(false, col, row+step, data.xcenter, data.ycenter)
				go data.scanRow(depth+1*dir, newslope, endslope, newChan)
			}
		} else if isOpaque(cell) {
			// this cell is the first cell in the "wall"
			newslope := calcSlope(false, col, row-step, data.xcenter, data.ycenter)
			go data.scanRow(depth+1*dir, startslope, newslope, newChan)
			blockerFound = true
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
