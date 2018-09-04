package game

import (
	"log"
)

// Mob ...
type Mob struct {
	Vnum int

	Name             string
	ShortDescription string
	LongDescription  string
	Description      string

	Act        int
	AffectedBy int
	Shop       *Shop
	Alignment  int
	Level      int

	Sex int

	InRoom *Room

	SpecFun string

	Count int // total number in the world, I guess
}

// FromRoom ...
func (m *Mob) FromRoom() {
	if m.InRoom == nil {
		log.Println("char from room: null room")
		return
	}

	for i, person := range m.InRoom.Mobs {
		if person == m {
			m.InRoom.People = append(m.InRoom.People[:i], m.InRoom.People[i+1:]...)
			break
		}
	}

	// if obj := m.ItemFromLocation(constants.WearLight); obj != nil && obj.ItemType == ItemLight && obj.Values[2] != 0 {
	// 	m.InRoom.Light--
	// }

	m.InRoom = nil
}

// ToRoom ...
func (m *Mob) ToRoom(room *Room) {
	if room == nil {
		log.Println("char to room: null room")
		return
	}

	m.InRoom = room
	room.Mobs = append(room.Mobs, m)

	// TODO: detect PCs, add to area list

	// if obj := m.ItemFromLocation(constants.WearLight); obj != nil && obj.ItemType == ItemLight && obj.Values[2] != 0 {
	// 	m.InRoom.Light++
	// }
}
