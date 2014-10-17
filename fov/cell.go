package fov

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
