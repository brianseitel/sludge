package game

// Room ...
type Room struct {
	Area *Area

	Vnum     int
	People   []*Character
	Contents []*Object
	Light    int

	Name             string
	Description      string
	ExtraDescription string
	Flags            int
	SectorType       int
	Exits            []*Exit
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
)

// NewRoom ...
func NewRoom() *Room {
	return &Room{
		People: make([]*Character, 0),
	}
}
