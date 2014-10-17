/*Package fov implements a way to get the player's current field of view
using technniques such as shadowcasting.*/
package fov

import (
	//This is needed as we will be changing light values in the world
	//"github.com/xenoryt/gork/world"
	"fmt"
)

type slope float64

//Keeps track of all the lit cells.
var litCells []Cell

func calcSlope(newvar, newstatic int, oldvar, oldstatic int) slope {
	ds := float64(newstatic - oldstatic)
	dv := float64(newvar - oldvar)
	if ds != 0 {
		return slope(dv / ds)
	}
	return 0
}

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

	//scan sector 1
	go data.scan(true, maxdepth, y, -1, 0, chans[0])
	//scan sector 2
	go data.scan(true, maxdepth, y, 1, 0, chans[1])
	//scan sector 6
	go data.scan(true, maxdepth, y, 1, 0, chans[2])
	//scan sector 5
	go data.scan(true, maxdepth, y, -1, 0, chans[3])

	//scan sector 3
	go data.scan(false, maxdepth, x, 1, 0, chans[4])
	//scan sector 3
	go data.scan(false, maxdepth, x, -1, 0, chans[5])
	//scan sector 8
	go data.scan(false, maxdepth, x, -1, 0, chans[6])
	//scan sector 7
	go data.scan(false, maxdepth, x, 1, 0, chans[7])

	//The (x,y) location itself should be lit.
	grid[y][x].SetLit(true)

	litCells = []Cell{grid[y][x]}

	//Now we wait until it has finished scanning
	for i := 0; i < 8; i++ {
		fmt.Println("waiting", i)
		for cell := range chans[i] {
			litCells = append(litCells, cell)
		}
	}
}

func (data *worldData) getData(scanrow bool, variable, static int) {
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
	var step int = 1
	if startslope > 0 {
		step = -1
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

	for ; math.abs(curdepth) < math.abs(maxdepth); curdepth += 1 * depthdir {
		start := int(startslope*slope(curdepth) + 0.5)
		end := int(endslope*slope(curdepth) + 0.5)
		var static int = scoordinate + curdepth

		//Double check to make sure we aren't going out of bounds
		if static < 0 || static > len(data.grid) {
			break
		}

		//Make sure these stay within bounds!
		if start < 0 {
			start = 0
		} else if start >= len(data.grid[0]) {
			start = len(data.grid[0]) - 1
		}
		if end < 0 {
			end = 0
		} else if end >= len(data.grid[0]) {
			end = len(data.grid[0]) - 1
		}

		//start the loop to check for shadows.
		isBlocked := false
		for variable := start; math.abs(variable) <= math.abs(end); variable += step {
			cell := data.getCell(scanrow, variable, static)
			cell.SetLit(true)
			ch <- cell

			//if this is the first cell in the "wall"
			if !isBlocked {
				if cell.Opacity() == 0 {
					newslope := calcSlope(scanrow, variable-step, static, vcoordinate, scoordinate)
					go data.scan(maxdepth-curdepth, variable-step, static, startslope, newslope, newchan())
				}
			} else if cell.Opacity() != 0 {
				//This is the last cell in the wall
				newslope := calcSlope(scanrow, variable+step, static, vcoordinate, scoordinate)
				go data.scan(maxdepth-curdepth, variable+step, static, newslope, endslope, newchan())
			}
		}

		//if the last cell is opaque then we don't need to continue
		if isBlocked {
			break
		}
	}

}
