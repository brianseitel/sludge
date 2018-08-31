package main

import (
	"fmt"
	"log"
	"strings"
	"time"
	"unicode"

	"github.com/brianseitel/sludge/internal/constants"
	"github.com/brianseitel/sludge/internal/game"
	"github.com/brianseitel/sludge/internal/server"
)

var world *game.World
var gameServer *server.Server

func main() {
	world = &game.World{
		Wizlocked: false,
	}
	gameServer = server.NewServer()
	gameServer.Start()

	go ticker()

	for {
		conn, err := gameServer.Listener.Accept()
		if err != nil {
			log.Printf("Failed to accept connection %v", err)
			continue
		}

		// Init connection
		// Send the greeting
		c := server.Connection{
			Conn:      conn,
			Connected: constants.GetName,
			Host:      conn.RemoteAddr().String(),
		}
		c.Greet()
		go loop(c)
	}
}

func loop(conn server.Connection) {
	for !gameServer.Down {

		input := conn.Read()

		if conn.Connected == constants.Playing {
			conn.Interpret(input)
		} else {
			nanny(&conn, input)
		}

		_, err := conn.Conn.Write([]byte{})
		if err != nil {
			return
		}
		// wait
		// read input
		// show output
	}
}

func ticker() {
	t := time.NewTicker(time.Second * 5)
	go func() {
		for t := range t.C {
			fmt.Println("Tick at", t)
		}
	}()
}

func nanny(conn *server.Connection, input string) {
	ch := conn.Character

	switch conn.Connected {
	case constants.GetName:
		if len(input) <= 0 {
			conn.Conn.Close()
			return
		}

		if !checkName(input) {
			conn.Write("Illegal name, try another.%s\rName: ", constants.EOL)
			return
		}

		conn.LoadChar(input)
		ch = conn.Character

		if ch.Banned {
			conn.Write("Access denied.%s", constants.EOL)
			conn.Conn.Close()
			return
		}

		if world.Wizlocked {
			conn.Write("The game is wizlocked.%s", constants.EOL)
			conn.Conn.Close()
			return
		}

		if ch.Exists {
			// ask for password
			conn.Write("Password: ", constants.EchoOff)
			conn.Connected = constants.GetOldPassword
		} else {
			// new player
			conn.Write("Did I get that right, %s? (Y/N) ", input)
			conn.Connected = constants.ConfirmNewName
		}

	case constants.GetOldPassword:
		if input != ch.Password {
			conn.Write("Wrong password.%s", constants.EOL)
			conn.Conn.Close()
			return
		}

		// check to see if currently playing
		if 1 == 2 {
			conn.Write("Already playing!%s", constants.EOL)
			conn.Conn.Close()
			return
		}

		log.Printf("%s has connected.%s", ch.Name, constants.EOL)
		// show MOTD
		conn.Connected = constants.ReadMOTD

	case constants.ConfirmNewName:
		switch strings.ToLower(input) {
		case "y":
			conn.Write("New character.%s", constants.EOL)
			conn.Write("Give me a password for %s: %s", ch.Name, constants.EchoOff)
			conn.Connected = constants.GetNewPassword
		case "n":
			conn.Write("Ok, what IS it then? ")
			conn.Character = nil
			conn.Connected = constants.GetName
		default:
			conn.Write("Please write Y or N.%s", constants.EOL)
		}

	case constants.GetNewPassword:
		if len(input) < 5 {
			conn.Write("Password must be at least five characters long.%s", constants.EOL)
			conn.Write("Password: %s", constants.EchoOff)
			return
		}

		pwd, err := crypt(input)
		if err != nil {
			conn.Write("New password not acceptable. Try again.%s", constants.EOL)
			conn.Write("Password: %s", constants.EchoOff)
		}

		ch.Password = pwd
		conn.Write(constants.EOL)
		conn.Write("Please retype password: %s", constants.EchoOff)
		conn.Connected = constants.ConfirmNewPassword

	case constants.ConfirmNewPassword:
		newPwd, err := crypt(input)
		if err != nil || newPwd != ch.Password {
			conn.Write(constants.EOL)
			conn.Write("Passwords don't match.%s", constants.EOL)
			conn.Write("Password: %s", constants.EchoOff)
			conn.Connected = constants.GetNewPassword
			return
		}

		conn.Write(constants.EOL)
		conn.Write("What is your sex? (M/F/N) ")
		conn.Connected = constants.GetNewSex

	case constants.GetNewSex:
		switch strings.ToLower(input) {
		case "m":
			ch.Sex = constants.SexMale
		case "f":
			ch.Sex = constants.SexFemale
		case "n":
			ch.Sex = constants.SexNeutral
		default:
			conn.Write("%sThat's not a valid sex.%s", constants.EOL)
			conn.Write("What IS your sex? (M/F/N) ")
			return
		}

		conn.Write(constants.EOL)
		conn.Write("Select a class: [")
		for i, class := range game.Classes {
			if i > 0 {
				conn.Write(" ")
			}
			conn.Write(class.WhoName)
		}
		conn.Write("]: ")
		conn.Connected = constants.GetNewClass

	case constants.GetNewClass:
		for _, c := range game.Classes {
			if strings.ToLower(input) == c.WhoName {
				ch.Class = &c
				break
			}
		}

		if ch.Class == nil {
			conn.Write("That's not a class.%s", constants.EOL)
			conn.Write("What IS your class?")
			return
		}

		log.Printf("%s@%s new player.%s", ch.Name, conn.Host, constants.EOL)
		conn.Write("%s", constants.EOL)

		conn.Connected = constants.ReadMOTD

	case constants.ReadMOTD:
		conn.Write(game.MOTD)

		world.Characters = append(world.Characters, conn.Character)
		conn.Connected = constants.Playing

		if ch.Level == 0 {
			ch.Level = 1
			ch.XP = 1000

			ch.Mana = ch.MaxMana
			ch.HP = ch.MaxHP
			ch.Move = ch.MaxMove
		}
	}
}

func crypt(pwd string) (string, error) {
	return pwd, nil
}

func checkName(name string) bool {
	if name == "all" || name == "auto" || name == "immortal" || name == "self" || name == "someone" {
		fmt.Println("1")
		return false
	}

	if len(name) < 4 {
		fmt.Println("4")
		return false
	}

	// Only allow letters, no numbers
	for _, r := range name {
		if !unicode.IsLetter(r) {
			return false
		}
	}

	// TODO: prevent users from naming themselves after mobs

	return true
}
