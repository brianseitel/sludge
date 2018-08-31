package game

// World ...
type World struct {
	Wizlocked  bool
	Characters []*Player
}

type Player struct {
	Name   string
	Banned bool
	Exists bool

	Password string

	Sex   int
	Class *Class

	Level int
	XP    int

	HP      int
	MaxHP   int
	Mana    int
	MaxMana int
	Move    int
	MaxMove int
}

// NewPlayer ...
func NewPlayer(name string) *Player {
	return &Player{
		Name:    name,
		Level:   0,
		MaxHP:   100,
		MaxMana: 100,
		MaxMove: 100,
	}
}

// Class ...
type Class struct {
	Name    string
	WhoName string
}

// Classes ...
var Classes = []Class{
	Class{
		Name:    "Warrior",
		WhoName: "war",
	},
	Class{
		Name:    "Mage",
		WhoName: "mag",
	},
}

var MOTD = "Welcome!!!\n"
