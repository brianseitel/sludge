package game_test

import (
	"testing"

	"github.com/brianseitel/sludge/internal/game"
	"github.com/stretchr/testify/assert"
)

func TestNewCharacter(t *testing.T) {
	char := game.NewCharacter("Bob", nil)
	assert.Equal(t, "Bob", char.Name)
}

func TestCharIsAwake(t *testing.T) {
	char := game.NewCharacter("Bob", nil)
	assert.True(t, char.IsAwake())
}

func TestCharIsImmortal(t *testing.T) {
	char := game.NewCharacter("Bob", nil)
	assert.False(t, char.IsImmortal())
}

// TODO: Finish this test
func TestCharDo(t *testing.T) {
	char := game.NewCharacter("Bob", nil)
	char.Do("something")
}

// TODO: Finish this test
func TestCharSeenBy(t *testing.T) {
	char := game.NewCharacter("Bob", nil)
	other := game.NewCharacter("Jack", nil)
	pers := char.SeenBy(other)

	assert.Equal(t, char.Name, pers)
}
