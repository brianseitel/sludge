package game

import (
	"log"

	"github.com/brianseitel/sludge/internal/constants"
)

// Move ...
func (c *Character) Move(direction int) {
	if direction < 0 || direction > 5 {
		log.Printf("do move: bad direction %d\n", direction)
		return
	}

	room := c.InRoom
	if room == nil {
		log.Printf("do move: character not in a room")
		return
	}

	exit := room.Exits[direction]
	if exit == nil {
		c.Send("Alas, you cannot go that way.\n")
		return
	}

	nextRoom, ok := world.Rooms[exit.Vnum]
	if !ok {
		c.Send("Alas, you cannot go that way.\n")
		return
	}

	// TODO: check for closed door

	// TODO: check for charm

	// TODO: check for private

	// TODO: check for guild

	// TODO: check for water / boat

	// TODO: check for movement loss based on sector type

	c.Movement--

	// TODO: check for sneak
	Notify("$n leaves $T.", c, constants.ActToRoom, ActOptions{})

	c.FromRoom()
	c.ToRoom(nextRoom)

	// TODO: check for sneak
	Notify("$n has arrived.", c, constants.ActToRoom, ActOptions{})

	c.Interpret("look")

	// TODO: implement following
}
