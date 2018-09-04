package game

import (
	"strings"

	"github.com/brianseitel/sludge/internal/constants"
	"github.com/sanity-io/litter"
)

// Command ...
type Command struct {
	Keywords string
	Func     CommandFunc
	Level    int
	Position int
	// TODO: add log levels
}

// CommandFunc ...
type CommandFunc func(ch *Character, args string)

// Commands ...
var Commands map[string]Command

func init() {
	cmds := []Command{
		Command{Keywords: "north", Func: moveChar(constants.DirNorth), Position: constants.PositionStanding, Level: 0},
		Command{Keywords: "south", Func: moveChar(constants.DirSouth), Position: constants.PositionStanding, Level: 0},
		Command{Keywords: "east", Func: moveChar(constants.DirEast), Position: constants.PositionStanding, Level: 0},
		Command{Keywords: "west", Func: moveChar(constants.DirWest), Position: constants.PositionStanding, Level: 0},
		Command{Keywords: "up", Func: moveChar(constants.DirUp), Position: constants.PositionStanding, Level: 0},
		Command{Keywords: "down", Func: moveChar(constants.DirDown), Position: constants.PositionStanding, Level: 0},

		Command{Keywords: "look", Func: look(), Position: constants.PositionStanding, Level: 0},
	}

	Commands = make(map[string]Command, len(cmds))
	for _, c := range cmds {
		Commands[c.Keywords] = c
	}
}

// Do command
func (c Command) Do(ch *Character, args string) {
	c.Func(ch, args)
}

func moveChar(direction int) CommandFunc {
	return func(ch *Character, args string) {
		ch.Move(direction)
	}
}

func look() CommandFunc {
	return func(ch *Character, args string) {
		if ch.Descriptor == nil {
			return
		}

		if ch.Position < constants.PositionSleeping {
			ch.Send("You can't see anything but stars!\n")
			return
		}

		if ch.Position == constants.PositionSleeping {
			ch.Send("You can't see anything, dumbass. You're sleeping!\n")
			return
		}

		// TODO: check blind

		// TODO: check darkness
		arg1, args := one(args)
		arg2, args := one(args)

		litter.Dump(ch.InRoom)

		if arg1 == "" || arg1 == "auto" {
			// 'look' or 'look auto'
			ch.Send(ch.InRoom.Name + "\n")
			ch.Send(ch.InRoom.Description)
			ch.Send("\n")
			// TODO: send contents
			// TODO: send people
			return
		}

		if arg1 == "i" || arg1 == "in" {
			// look in
			if arg2 == "" {
				ch.Send("Look in what?\n")
				return
			}

			// TODO: look in object
			return
		}

		var victim *Character
		for _, p := range ch.InRoom.People {
			if strings.HasPrefix(p.Name, arg1) {
				victim = p
				break
			}
		}

		if victim != nil {
			// TODO: show victim
		}

		for _, obj := range ch.Carrying {
			if strings.HasPrefix(obj.Name, arg1) {
				// TODO: add can see
				ch.Send(obj.Description + "\n")
				return
			}
		}

		for _, obj := range ch.InRoom.Contents {
			if strings.HasPrefix(obj.Name, arg1) {
				ch.Send(obj.Description + "\n")
				return
			}
		}

		// check doors
		var door int
		switch arg1 {
		case "n", "north":
			door = constants.DirNorth
		case "e", "east":
			door = constants.DirEast
		case "s", "south":
			door = constants.DirSouth
		case "w", "west":
			door = constants.DirWest
		case "u", "up":
			door = constants.DirUp
		case "d", "down":
			door = constants.DirDown
		default:
			ch.Send("You do not see that here.\n")
			return
		}

		if ch.InRoom.Exits[door] == nil {
			ch.Send("Nothing special there.\n")
			return
		}

		ex := ch.InRoom.Exits[door]
		if ex.Description == "" {
			ch.Send("Nothing special there.\n")
		} else {
			ch.Send(ex.Description + "\n")
		}

		if ex.Keywords != "" {
			if ex.Info == ExitIsClosed {
				ch.Send("The %s is closed.\n", ex.Keywords)
			} else if ex.Info == ExitIsDoor {
				ch.Send("The %s is open.\n", ex.Keywords)
			}
		}
	}
}
