package game

// Room ...
type Room struct {
	Area *Area

	Vnum     int
	People   []*Character
	Mobs     []*Mob
	Contents []*Object
	Light    int

	Name             string
	Description      string
	ExtraDescription string
	Flags            int
	SectorType       int
	Exits            [6]*Exit
}

// Exit ...
type Exit struct {
	Vnum int

	Description string
	Keywords    string
	Info        int
	Key         int
	Locks       int
}

// Exit flags
const (
	ExitIsDoor    = 1
	ExitPickProof = 2
	ExitIsClosed  = 3
)

// NewRoom ...
func NewRoom() *Room {
	return &Room{
		People: make([]*Character, 0),
	}
}
