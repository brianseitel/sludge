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
	Character *game.Player
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
	c.Character = game.NewPlayer(name)
}

func (c Connection) Write(message string, args ...interface{}) {
	c.Conn.Write(constants.EchoOn)
	c.Conn.Write([]byte(fmt.Sprintf(message, args...)))
}

// Read
func (c Connection) Read() string {
	message, _ := bufio.NewReader(c.Conn).ReadString('\n')
	message = strings.Replace(message, "\r\n", "", -1)

	return message
}
