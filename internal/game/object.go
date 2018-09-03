package game

import (
	"log"

	"github.com/brianseitel/sludge/internal/constants"
	"github.com/brianseitel/sludge/internal/helpers"
)

// Object ...
type Object struct {
	Name             string
	Cost             int
	Description      string
	ExtraFlags       int
	ItemType         int
	Level            int
	ShortDescription string
	Timer            int
	Values           []int
	WearFlags        int
	WearLocation     constants.WearLocation
	Weight           int
	Vnum             int
	indexData        *Object

	CarriedBy *Character
	InRoom    *Room
	InObject  *Object

	Affected []*Affect
}

// Item types
const (
	ItemLight = iota
	ItemTreasure
	ItemFurniture
	ItemTrash
	ItemContainer
	ItemDrinkContainer
	ItemKey
	ItemFood
	ItemBoat
	ItemCorpseNPC
	ItemCorpsePC
	ItemFountain
	ItemScroll
	ItemWand
	ItemStaff
	ItemArmor
	ItemWeapon
	ItemPotion
	ItemPill
	ItemMoney
)

// FromChar ...
func (obj *Object) FromChar() {
	if obj.CarriedBy == nil {
		log.Println("obj from char: null char")
		return
	}
	ch := obj.CarriedBy

	// TODO: check for wear location

	for i, o := range ch.Carrying {
		if o == obj {
			ch.Carrying = append(ch.Carrying[:i], ch.Carrying[i+1:]...)
		}
	}

	obj.CarriedBy = nil
	ch.CarryNumber--
	ch.CarryWeight -= obj.Weight
}

// ToChar ...
func (obj *Object) ToChar(char *Character) {
	char.Carrying = append(char.Carrying, obj)
	obj.CarriedBy = char
	obj.InRoom = nil
	obj.InObject = nil
	char.CarryNumber++
	char.CarryWeight += obj.Weight
}

// CreateObject ...
func CreateObject(index *Object, level int) *Object {
	if index == nil {
		log.Println("create object: nil index")
		return nil
	}

	// Set basic object data
	obj := &Object{}
	obj.indexData = index
	obj.InRoom = nil
	obj.Level = level
	obj.WearLocation = -1

	obj.Name = index.Name
	obj.ShortDescription = index.ShortDescription
	obj.Description = index.Description
	obj.ItemType = index.ItemType
	obj.ExtraFlags = index.ExtraFlags
	obj.WearFlags = index.WearFlags
	obj.Values = index.Values
	obj.Weight = index.Weight
	obj.Cost = helpers.Fuzzy(10) * helpers.Fuzzy(level) * helpers.Fuzzy(level)

	// Mess with properties

	switch obj.ItemType {
	case ItemLight, ItemTreasure, ItemFurniture, ItemTrash, ItemContainer, ItemDrinkContainer, ItemKey, ItemFood, ItemBoat, ItemCorpseNPC, ItemCorpsePC, ItemFountain:
		break
	case ItemScroll:
		obj.Values[0] = helpers.Fuzzy(obj.Values[0])

	case ItemWand, ItemStaff:
		obj.Values[0] = helpers.Fuzzy(obj.Values[0])
		obj.Values[1] = helpers.Fuzzy(obj.Values[1])
		obj.Values[2] = obj.Values[1]

	case ItemWeapon:
		obj.Values[1] = helpers.Fuzzy(helpers.Fuzzy(1*level/4 + 2))
		obj.Values[2] = helpers.Fuzzy(helpers.Fuzzy(3*level/4 + 6))

	case ItemArmor:
		obj.Values[0] = helpers.Fuzzy(level/4 + 2)

	case ItemPotion, ItemPill:
		obj.Values[0] = helpers.Fuzzy(helpers.Fuzzy(obj.Values[0]))

	case ItemMoney:
		obj.Values[0] = obj.Cost
	}

	return obj
}

// IsAntiGood ... TODO: finish
func (obj *Object) IsAntiGood() bool {
	return false
}

// IsAntiEvil ... TODO: finish
func (obj *Object) IsAntiEvil() bool {
	return false
}

// IsAntiNeutral ... TODO: finish
func (obj *Object) IsAntiNeutral() bool {
	return false
}

// ToRoom ...
func (obj *Object) ToRoom(room *Room) {
	if room == nil {
		log.Println("obj to room: room does not exist")
		return
	}
	obj.InRoom = room
	obj.CarriedBy = nil
	obj.InObject = nil

	room.Contents = append(room.Contents, obj)
}
