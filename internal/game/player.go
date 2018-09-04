package game

import (
	"fmt"
	"log"
	"net"

	"github.com/brianseitel/sludge/internal/constants"
)

// Character ...
type Character struct {
	Descriptor net.Conn

	Name   string
	Title  string
	Banned bool
	Exists bool

	Played bool

	Position int
	Practice int

	Sex   int
	Class *Class
	Race  *Race

	Level int
	XP    int
	Gold  int
	Trust int

	VNUM int

	DamageRoll  int
	HitRoll     int
	SavingThrow int

	Deaf             bool
	Description      string
	ShortDescription string
	LongDescrition   string

	PCData      *PCData
	HP          int
	MaxHP       int
	Mana        int
	MaxMana     int
	Movement    int
	MaxMovement int
	Wimpy       int

	InRoom *Room

	Act        int
	AffectedBy int
	Alignment  int
	Armor      int
	Affected   *Affect
	Skills     []Skill

	Carrying       []*Object
	CarryNumber    int
	CarryNumberMax int
	CarryWeight    int
	CarryWeightMax int
}

// PCData represents PC data
type PCData struct {
	Password string
	BamfIn   string
	BamfOut  string
	Title    string

	PermanentStrength     int
	PermanentIntelligence int
	PermanentWisdom       int
	PermanentDexterity    int
	PermanentConstitution int

	ModifiedStrength     int
	ModifiedIntelligence int
	ModifiedWisdom       int
	ModifiedDexterity    int
	ModifiedConstitution int
}

// NewCharacter ...
func NewCharacter(name string, desc net.Conn) *Character {
	char := &Character{
		Descriptor: desc,

		Sex: constants.SexMale,

		Name:        name,
		Level:       0,
		MaxHP:       100,
		MaxMana:     100,
		MaxMovement: 100,

		PCData: &PCData{},
		InRoom: NewRoom(),
	}

	char.InRoom.People = append(char.InRoom.People, char)
	return char
}

// Equip an object to a character
func (c *Character) Equip(obj *Object, location constants.WearLocation) {
	// TODO: detect if already equipped

	if (obj.IsAntiEvil() && c.IsEvil()) ||
		(obj.IsAntiGood() && c.IsGood()) ||
		(obj.IsAntiNeutral() && c.IsNeutral()) {
		Notify("You are zapped by $p and drop it.", c, constants.ActToCharacter, ActOptions{
			Object1: obj,
		})
		Notify("$n is zapped by $p and drops it.", c, constants.ActToRoom, ActOptions{
			Object1: obj,
		})

		obj.FromChar()
		obj.ToRoom(c.InRoom)
		return
	}

	// TODO: apply ac

	obj.WearLocation = location

	for _, _ = range obj.indexData.Affected {
		// TODO: modify character with affect
	}

	for _, _ = range obj.Affected {
		// TODO: modify character with affect
	}

	if obj.ItemType == ItemLight && obj.Values[2] != 0 && c.InRoom != nil {
		c.InRoom.Light++
	}
}

// Interpret ...
func (c *Character) Interpret(args string) {
	interpreter := NewInterpreter(c, args)
	interpreter.Do()
}

// ItemFromLocation ...
func (c *Character) ItemFromLocation(loc constants.WearLocation) *Object {
	for _, obj := range c.Carrying {
		if obj.WearLocation == constants.WearLight {
			return obj
		}
	}
	return nil
}

// IsEvil ...
func (c *Character) IsEvil() bool {
	return c.Alignment < -250
}

// IsGood ...
func (c *Character) IsGood() bool {
	return c.Alignment > 250
}

// IsNeutral ...
func (c *Character) IsNeutral() bool {
	return c.Alignment <= 250 && c.Alignment >= -250
}

// IsAwake ... TODO: finish
func (c *Character) IsAwake() bool {
	return true
}

// IsImmortal ... TODO: finish
func (c *Character) IsImmortal() bool {
	return false
}

// Send message to char
func (c *Character) Send(message string, args ...interface{}) {
	if c.Descriptor != nil {
		c.Descriptor.Write([]byte(fmt.Sprintf(message, args...)))
	}
}

// FromRoom ...
func (c *Character) FromRoom() {
	if c.InRoom == nil {
		log.Println("char from room: null room")
		return
	}

	for i, person := range c.InRoom.People {
		if person == c {
			c.InRoom.People = append(c.InRoom.People[:i], c.InRoom.People[i+1:]...)
			break
		}
	}

	if obj := c.ItemFromLocation(constants.WearLight); obj != nil && obj.ItemType == ItemLight && obj.Values[2] != 0 {
		c.InRoom.Light--
	}

	c.InRoom = nil
}

// ToRoom ...
func (c *Character) ToRoom(room *Room) {
	if room == nil {
		log.Println("char to room: null room")
		return
	}

	c.InRoom = room
	room.People = append(room.People, c)

	// TODO: detect PCs, add to area list

	if obj := c.ItemFromLocation(constants.WearLight); obj != nil && obj.ItemType == ItemLight && obj.Values[2] != 0 {
		c.InRoom.Light++
	}
}

// SeenBy ... TODO: fix this
func (c *Character) SeenBy(other *Character) string {
	return c.Name
}
