package main

type Direction uint8

const (
	NORTH Direction = iota
	NORTHEAST
	EAST
	SOUTHEAST
	SOUTH
	SOUTHWEST
	WEST
	NORTHWEST
)

func DirectionDelta(dir Direction) (x, y int) {
	switch dir {
	case NORTH:
		return 0, -1
	case EAST:
		return 1, 0
	case SOUTH:
		return 0, 1
	case WEST:
		return -1, 0
	case NORTHEAST:
		return 1, -1
	case NORTHWEST:
		return -1, -1
	case SOUTHEAST:
		return 1, 1
	case SOUTHWEST:
		return -1, 1
	}
	return 0, 0
}
