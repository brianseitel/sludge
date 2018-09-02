package game

var Rooms map[int]*Room

// Room ...
type Room struct {
	People   []*Character
	Contents []*Object
	Light    int
}

// NewRoom ...
func NewRoom() *Room {
	return &Room{
		People: make([]*Character, 0),
	}
}
