package game

import (
	"fmt"
	"log"
	"strings"

	"github.com/brianseitel/sludge/internal/constants"
)

// Logging Levels
const (
	LogNormal = iota
	LogAlways
	LogNever
)

// Interpreter ...
type Interpreter struct {
	Character   *Character
	OriginalArg string
	Args        string
}

// NewInterpreter ...
func NewInterpreter(ch *Character, args string) *Interpreter {
	return &Interpreter{
		Character:   ch,
		OriginalArg: strings.Trim(args, " "),
		Args:        strings.Trim(args, " "),
	}
}

// Do ...
func (i *Interpreter) Do() {
	if len(i.Args) == 0 {
		return
	}

	// TODO: remove hide

	// TODO: check for frozen status

	// look up command
	argument := i.one()
	trust := i.Character.Trust

	found := false
	var command Command
	for k, cmd := range Commands {
		fmt.Println(argument)
		if strings.HasPrefix(k, argument) && cmd.Level <= trust {
			found = true
			command = cmd
			break
		}
	}

	// log and snoop
	if i.Character.Descriptor != nil {
		log.Printf("[LOG: %s] %s", i.Character.Name, i.OriginalArg)
	}

	if !found {
		// TODO: check socials

		for key, social := range Socials {
			if strings.HasPrefix(key, argument) {
				social.Do(i)
				return
			}
		}
		i.Character.Send("Huh?\n")
		return
	}

	// Check to see if the user can do the command
	if i.Character.Position < command.Position {
		switch i.Character.Position {
		case constants.PositionDead:
		case constants.PositionMortal:
		case constants.PositionIncapactitated:
		case constants.PositionStunned:
		case constants.PositionSleeping:
		case constants.PositionResting:
		case constants.PositionFighting:

		}
	}

	command.Do(i.Character, i.Args)
}

func (i *Interpreter) one() string {
	o, args := one(i.Args)
	i.Args = strings.Trim(args, " ")
	return o
}

func one(args string) (string, string) {
	if len(args) == 0 {
		return "", ""
	}

	fields := strings.Fields(args)
	if len(fields) > 1 {
		return fields[0], strings.Join(fields[1:], " ")
	}

	return fields[0], ""
}
