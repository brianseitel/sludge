package game

import (
	"errors"
	"log"
	"strings"

	"github.com/brianseitel/sludge/internal/constants"
)

// Pronouns for notifications
var (
	HeShe  = []string{"it", "he", "she"}
	HimHer = []string{"it", "him", "her"}
	HisHer = []string{"its", "his", "her"}
)

// ActOptions ...
type ActOptions struct {
	Object1    *Object
	Object2    *Object
	Character1 *Character
	Character2 *Character
	Victim     *Character
}

// Notify ...
// TODO: Fix notify
func Notify(pattern string, character *Character, messageType constants.ActType, options ActOptions) error {
	if len(pattern) == 0 {
		return errors.New("no pattern specified")
	}

	targets := character.InRoom.People

	if messageType == constants.ActToVictim {
		if options.Victim == nil {
			log.Printf("Act: null victim with ToVictim.")
			return errors.New("null victim")
		}
		targets = options.Victim.InRoom.People
	}

	for _, to := range targets {
		if to.Descriptor == nil || !to.IsAwake() {
			continue
		}

		if messageType == constants.ActToCharacter && to != character {
			continue
		}
		if messageType == constants.ActToVictim && (to != options.Victim || to == character) {
			continue
		}
		if messageType == constants.ActToRoom && to == character {
			continue
		}
		if messageType == constants.ActToNotVictim && (to == character || to == options.Victim) {
			continue
		}

		message := BuildMessage(pattern, character, to, options.Victim, options)

		to.Descriptor.Write([]byte(message))
	}

	return nil
}

// BuildMessage from pattern
func BuildMessage(pattern string, character *Character, to *Character, victim *Character, options ActOptions) string {
	if options.Character1 != nil {
		pattern = strings.Replace(pattern, "$t", options.Character1.Name, -1)
	}
	if options.Character2 != nil {
		pattern = strings.Replace(pattern, "$T", options.Character2.Name, -1)
	}

	pattern = strings.Replace(pattern, "$n", character.SeenBy(to), -1)
	pattern = strings.Replace(pattern, "$N", victim.SeenBy(to), -1)
	pattern = strings.Replace(pattern, "$e", HeShe[character.Sex], -1)
	pattern = strings.Replace(pattern, "$E", HeShe[victim.Sex], -1)
	pattern = strings.Replace(pattern, "$m", HimHer[character.Sex], -1)
	pattern = strings.Replace(pattern, "$M", HimHer[victim.Sex], -1)
	pattern = strings.Replace(pattern, "$s", HisHer[character.Sex], -1)
	pattern = strings.Replace(pattern, "$S", HisHer[character.Sex], -1)
	if options.Object1 != nil {
		pattern = strings.Replace(pattern, "$p", options.Object1.Name, -1)
	}

	if options.Object2 != nil {
		pattern = strings.Replace(pattern, "$P", options.Object2.Name, -1)
	}

	return pattern

}
