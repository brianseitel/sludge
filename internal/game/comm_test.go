package game_test

import (
	"fmt"
	"net"
	"testing"

	"github.com/brianseitel/sludge/internal/constants"
	"github.com/brianseitel/sludge/internal/game"
	"github.com/stretchr/testify/assert"
)

// TODO: fix tests
func TestNotifyToVictim(t *testing.T) {
	char := game.NewCharacter("Bob", &net.TCPConn{})
	jack := game.NewCharacter("Jack", &net.TCPConn{})
	jim := game.NewCharacter("Jim", &net.TCPConn{})
	opts := game.ActOptions{
		Victim: jack,
	}
	jack.InRoom.People = append(char.InRoom.People, jack)
	jack.InRoom.People = append(char.InRoom.People, jim)

	err := game.Notify("sup y'all", char, constants.ActToVictim, opts)
	assert.Nil(t, err)
}

// TODO: fix tests
func TestNotifyToRoom(t *testing.T) {
	char := game.NewCharacter("Bob", &net.TCPConn{})
	opts := game.ActOptions{
		Victim: game.NewCharacter("Jack", &net.TCPConn{}),
	}

	err := game.Notify("sup y'all", char, constants.ActToRoom, opts)
	assert.Nil(t, err)
}

// TODO: fix tests
func TestNotifyToCharacter(t *testing.T) {
	char := game.NewCharacter("Bob", &net.TCPConn{})
	jack := game.NewCharacter("Jack", &net.TCPConn{})
	opts := game.ActOptions{
		Victim: jack,
	}
	char.InRoom.People = append(char.InRoom.People, jack)

	err := game.Notify("sup y'all", char, constants.ActToCharacter, opts)
	assert.Nil(t, err)
}

// TODO: fix tests
func TestNotifyToNotVictim(t *testing.T) {
	char := game.NewCharacter("Bob", &net.TCPConn{})
	opts := game.ActOptions{
		Victim: game.NewCharacter("Jack", &net.TCPConn{}),
	}

	err := game.Notify("sup y'all", char, constants.ActToNotVictim, opts)
	assert.Nil(t, err)
}

// TODO: fix tests
func TestNotifyNoDescriptor(t *testing.T) {
	char := game.NewCharacter("Bob", &net.TCPConn{})
	opts := game.ActOptions{
		Victim: game.NewCharacter("Jack", nil),
	}

	err := game.Notify("sup y'all", char, constants.ActToVictim, opts)
	assert.Nil(t, err)
}

func TestNotifyWithNoPattern(t *testing.T) {
	err := game.Notify("", game.NewCharacter("Bob", nil), constants.ActToVictim, game.ActOptions{})
	assert.NotNil(t, err)
}

func TestNotifyWithoutVictim(t *testing.T) {
	opts := game.ActOptions{}

	err := game.Notify("sup y'all", game.NewCharacter("Bob", nil), constants.ActToVictim, opts)
	assert.NotNil(t, err)
}

func TestBuildMessage(t *testing.T) {
	character := game.NewCharacter("Bob", nil)
	victim := game.NewCharacter("Jack", nil)

	options := game.ActOptions{
		Object1:    &game.Object{Name: "thingy"},
		Object2:    &game.Object{Name: "whatchamacallit"},
		Character1: game.NewCharacter("Jill", nil),
		Character2: game.NewCharacter("Lori", nil),
	}

	pattern := "$t $T $n $N $e $E $m $M $s $S $p $P"

	message := game.BuildMessage(pattern, character, character, victim, options)
	fmt.Println(message)
	assert.Equal(t, "Jill Lori Bob Jack he he him him his his thingy whatchamacallit", message)
}
