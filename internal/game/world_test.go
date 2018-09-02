package game_test

import (
	"testing"

	"github.com/brianseitel/sludge/internal/game"
	"github.com/stretchr/testify/assert"
)

func TestNewWorld(t *testing.T) {
	world := game.NewWorld()
	assert.NotNil(t, world)
	assert.False(t, world.Wizlocked)
}

func TestGetWorld(t *testing.T) {
	game.NewWorld()

	assert.NotNil(t, game.GetWorld())
}
