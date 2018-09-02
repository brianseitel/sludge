package server_test

import (
	"net"
	"testing"

	"github.com/brianseitel/sludge/internal/server"
	"github.com/stretchr/testify/assert"
)

func TestLoadChar(t *testing.T) {
	c := server.Connection{}

	c.LoadChar("Bob")

	assert.Equal(t, "Bob", c.Character.Name)
}

func TestConnectionWrite(t *testing.T) {
	c := server.Connection{
		Conn: &net.TCPConn{},
	}

	out := c.Write("Sup %d", 1)
	assert.Equal(t, "Sup 1", out)
}

func TestConnectionRead(t *testing.T) {
	srv, client := net.Pipe()
	defer client.Close()
	go func() {
		srv.Write([]byte("this is a test"))
		srv.Close()
	}()

	c := server.Connection{
		Conn: client,
	}

	message := c.Read()

	assert.Equal(t, "this is a test", message)
}

func TestConnectionGreet(t *testing.T) {
	c := server.Connection{
		Conn: &net.TCPConn{},
	}

	c.Greet()
}
