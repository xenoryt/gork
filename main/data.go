package main

type Being interface {
	Move(Direction)
	Attack(Being)
	Visible() bool
}

/*Player struct is used to store information about the player.
Implements the following interfaces:
- Being
- MapObject
*/
type Player struct {
	x, y  int
	sight int
}

func (player *Player) Move(dir Direction) {
	dx, dy := DirectionDelta(dir)
	player.x += dx
	player.y += dy
}
func (player Player) Attack(other Being) {
}
func (player Player) Visible() bool {
	return true
}
func (player Player) GetLoc() (x, y int) {
	return player.x, player.y
}
func (player Player) GetSymbol() byte {
	return '@'
}
