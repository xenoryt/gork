package game

// Generice interfaces that may prove useful

/* Object is the basic type that anything that
can be physically "sensed" should implement.
 - View() should return a description of the of object. */
type Object interface {
	View() string
}

/* Usable is any object that can be used or activated. */
type Useable interface {
	Use() (UseData, error)
}

/* Surface is anything that can have something
physically placed on top of it. */
type Surface interface {
	Receive(*Placeable) error
	Remove(*Placeable) error
	Has() []Placeable
}

/* Placeable is any object that can also be placed (or embedded)
onto a surface. Placeable objects can be taken as well. */
type Placeable interface {
	Place(*Surface) error
	Take() (string, error)
}

/* Wieldable objects are things that can be held in a
person's hands */
type Wieldable interface {
	Item
	Wield() (string, error)
}

// Some embedded interfaces

/* Item is an object that can be placed and used */
type Item interface {
	Placeable
	Useable
	Object
}

/* Weapons are wieldable objects that can also be
used to attack others
Note: Weapon.Use() can be the same thing as Weapon.Attack()
*/
type Weapon interface {
	Wieldable
	Attack(*Object) (string, error)
}
