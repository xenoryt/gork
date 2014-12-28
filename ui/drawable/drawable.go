package drawable

//Drawable is anything that you can get the location of
type Drawable interface {
	GetLoc() (x, y int)
	GetSymbol() byte
}
