/*Package fov implements a way to get the player's current field of view
using technniques such as shadowcasting.*/
package fov

import (
	//This is needed as we will be changing light values in the world
	//"github.com/xenoryt/gork/world"
	"fmt"
)

//Keeps track of all the lit cells.
var litCells []Cell

//A struct to store data
type worldData struct {
	grid     [][]Cell
	maxdepth int
	xcenter  int
	ycenter  int
}

/*Blackout blacks out all the cells that were previously lit using CastShadows*/
func Blackout() {
	fmt.Println("blacking out", len(litCells))
	for _, cell := range litCells {
		cell.SetLit(false)
	}
}

//http://www.roguebasin.com/index.php?title=Computing_LOS_for_Large_Areas
//http://www.roguebasin.com/index.php?title=FOV_using_recursive_shadowcasting

/*CastShadows lights up cells in the given grid assuming that the object
located at (x,y) is luminous*/
func CastShadows(grid [][]Cell, x, y int, maxdepth int) {
	//First create the worldData object
	data := worldData{
		grid:     grid,
		maxdepth: 6,
		xcenter:  x,
		ycenter:  y,
	}
	chans := make([]chan Cell, 8)
	for i := 0; i < 8; i++ {
		chans[i] = make(chan Cell)
	}

	/* We scan 8 different sectors
	\1|2/
	8\|/3
	-----
	7/|\4
	/6|5\
	*/

	//scan sector 2
	go data.scan(true, -maxdepth, x, y, -1, 0, chans[0])
	//scan sector 1
	go data.scan(true, -maxdepth, x, y, 1, 0, chans[1])
	//scan sector 6
	go data.scan(true, maxdepth, x, y, 1, 0, chans[2])
	//scan sector 5
	go data.scan(true, maxdepth, x, y, -1, 0, chans[3])

	//scan sector 3
	go data.scan(false, maxdepth, y, x, 1, 0, chans[4])
	//scan sector 3
	go data.scan(false, maxdepth, y, x, -1, 0, chans[5])
	//scan sector 8
	go data.scan(false, -maxdepth, y, x, -1, 0, chans[6])
	//scan sector 7
	go data.scan(false, -maxdepth, y, x, 1, 0, chans[7])

	//The (x,y) location itself should be lit.
	grid[y][x].SetLit(true)

	litCells = []Cell{grid[y][x]}

	//for i := 2; i < 8; i++ {
	//	close(chans[i])
	//}

	//Now we wait until it has finished scanning
	for i := 0; i < 8; i++ {
		fmt.Println("waiting", i)
		for cell := range chans[i] {
			litCells = append(litCells, cell)
		}
	}
}

func (data *worldData) getCell(scanrow bool, variable, static int) Cell {
	if !scanrow {
		return data.grid[variable][static]
	}
	return data.grid[static][variable]
}

/*scans the grid given. If scanrow is true, vcoordinate should be the x coordinate
and scoordinate should be the y coordinate. If scanrow is false, then vcoordinate
and scoordinate are reversed.*/
func (data *worldData) scan(scanrow bool, maxdepth, vcoordinate, scoordinate int, startslope, endslope slope, ch chan Cell) {
	defer func() {
		close(ch)
	}()

	var curdepth int = -1
	var depthdir int = -1
	if maxdepth >= 0 {
		curdepth = 1
		depthdir = 1
	}
	var step int = 1 * depthdir
	if startslope > endslope {
		step = -1 * depthdir
	}

	var chans []chan Cell

	//We want to make sure we recieved all the data from all the channels before
	//closing ours.
	defer func() {
		for i := range chans {
			for cell := range chans[i] {
				ch <- cell
			}
		}
	}()

	newchan := func() chan Cell {
		ch := make(chan Cell)
		chans = append(chans, ch)
		return ch
	}

	for ; abs(curdepth) < abs(maxdepth); curdepth += 1 * depthdir {
		start := int(startslope*slope(curdepth)) + vcoordinate
		end := int(endslope*slope(curdepth)) + vcoordinate
		var static int = scoordinate + curdepth

		//Double check to make sure we aren't going out of bounds
		if static < 0 || (scanrow && static >= len(data.grid)) || (!scanrow && static >= len(data.grid[0])) {
			break
		}

		//Make sure these stay within bounds!
		if start < 0 {
			start = 0
		} else if scanrow && start >= len(data.grid[0]) {
			start = len(data.grid[0]) - 1
		} else if !scanrow && start >= len(data.grid) {
			start = len(data.grid) - 1
		}
		if end < 0 {
			end = 0
		} else if scanrow && end >= len(data.grid[0]) {
			end = len(data.grid[0]) - 1
		} else if !scanrow && end >= len(data.grid) {
			end = len(data.grid) - 1
		}

		//Returns true if a <= b (depending on which way step is incrementing a)
		compare := func(a, b int) bool {
			if step > 0 {
				return a <= b
			}
			return a >= b
		}

		//start the loop to check for shadows.
		isBlocked := false
		for variable := start; compare(variable, end); variable += step {
			cell := data.getCell(scanrow, variable, static)
			cell.SetLit(true)
			ch <- cell

			//if this is the first cell in the "wall"
			if !isBlocked {
				if cell.Opacity() == 0 {
					newslope := calcSlope(variable-step, static, vcoordinate, scoordinate)
					go data.scan(scanrow, maxdepth-curdepth, variable-step, static, startslope, newslope, newchan())
					isBlocked = true
				}
			} else if cell.Opacity() != 0 {
				//This is the last cell in the wall
				startslope = calcSlope(variable+step, static, vcoordinate, scoordinate)
				fmt.Println("new slope", startslope)
				//go data.scan(scanrow, maxdepth-curdepth, variable+step, static, newslope, endslope, newchan())
			}
		}

		//if the last cell is opaque then we don't need to continue
		if isBlocked {
			break
		}
	}

}
