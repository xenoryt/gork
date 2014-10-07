package game

// This file mainly deals with the way data is
// handled and passed between objects

/*UseData is a data structure that should contain
all the information the usage of an item should return */
type UseData struct {
	effect  Effect
	ability Ability
}
