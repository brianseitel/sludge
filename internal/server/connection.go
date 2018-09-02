package server

import (
	"bufio"
	"fmt"
	"net"
	"strings"

	"github.com/brianseitel/sludge/internal/constants"
	"github.com/brianseitel/sludge/internal/game"
)

// Connection ...
type Connection struct {
	Conn      net.Conn
	Connected int
	Character *game.Character
	Host      string
}

// Interpret ...
func (c Connection) Interpret(input string) {

}

// Greet the user
func (c Connection) Greet() {
	c.Write("Welcome to Sludge MUD\n")
	c.Write(`What's your name? `)
}

// LoadChar ...
func (c *Connection) LoadChar(name string) {
	var err error
	c.Character, err = game.LoadPlayerFile(name)
	if err == nil {
		c.Character.Exists = true
	}
	c.Character.Descriptor = c.Conn

	if c.Character.InRoom == nil {
		c.Character.InRoom = &game.Room{}
	}
}

func (c Connection) Write(message string, args ...interface{}) string {
	msg := fmt.Sprintf(message, args...)
	c.Conn.Write(constants.EchoOn)
	c.Conn.Write([]byte(msg))
	return msg
}

// Read
func (c Connection) Read() string {
	message, _ := bufio.NewReader(c.Conn).ReadString('\n')
	message = strings.Replace(message, "\r\n", "", -1)
	message = strings.Replace(message, "\xff\xfd\x01", "", -1)

	return message
}
