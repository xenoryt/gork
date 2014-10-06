package game

//const (
//	EFFECT_HEAL = 1 << iota
//	EFFECT_DAMAGE
//	EFFECT_BOOST

type Effect interface {
	Effect() error //This is for delayed effects
	Duration() int
	Name() string
	Status() string
}

type boost struct {
	maxhp int
	hp    int
	str   int
	def   int
}

type basicEffect struct {
	name     string
	status   string
	duration int // in number of turns
	damage   int
	heal     int
	boosts   boost
}

func (eff basicEffect) Effect() error {
	return nil
}
func (eff basicEffect) Duration() int {
	return eff.duration
}
func (eff basicEffect) Name() string {
	return eff.name
}
func (eff basicEffect) Status() string {
	if eff.status == "" {
		return eff.name + "ed"
	}
	return eff.status
}

func GenerateEffects() map[string]Effect {
	alleffects := make(map[string]Effect)

	// Lets create a simple Poison effect
	effect := basicEffect{
		name:     "poison",
		status:   "poisoned", //Can leave this blank. default is name+"ed"
		duration: 10,
		heal:     0,
		boosts:   boost{},
	}
	alleffects[effect.name] = effect

	return alleffects
}
