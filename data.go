package main

type Being interface {
	Move(Direction)
	Attack(Being)
}

type Player struct {
	x, y int
}

func (player Player) Move(dir Direction) {
	dx, dy := DirectionDelta(dir)
	player.x += dx
	player.y += dy
}
