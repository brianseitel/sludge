package server_test

import (
	"testing"

	"github.com/brianseitel/sludge/internal/server"
	"github.com/stretchr/testify/assert"
)

func TestNewServer(t *testing.T) {
	s := server.NewServer()

	assert.NotNil(t, s)
	assert.True(t, s.Down)
}

func TestStart(t *testing.T) {
	s := server.NewServer()

	err := s.Start("0.0.0.0:9932")
	assert.Nil(t, err)
}

func TestStartErr(t *testing.T) {
	s := server.NewServer()

	err := s.Start("your mom")
	assert.NotNil(t, err)
}
