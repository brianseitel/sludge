package game

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/brianseitel/sludge/internal/constants"
)

// PlayerPath ...
const PlayerPath = "players/"

// LoadPlayerFile ...
func LoadPlayerFile(name string) (*Character, error) {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b) + "/../../"

	path := basepath + PlayerPath
	f, err := ioutil.ReadFile(path + strings.ToLower(name) + ".pfile")
	if err != nil {
		return nil, err
	}

	ch := ConvertPlayerFileToCharacter(name, string(f))

	return ch, nil
}

func readNumber(num string) int {
	n, _ := strconv.Atoi(num)
	return n
}

func readBoolean(val string) bool {
	val = strings.ToLower(val)
	if val == "true" || val == "t" || val == "1" {
		return true
	}
	return false
}

// ConvertPlayerFileToCharacter ...
func ConvertPlayerFileToCharacter(name string, data string) *Character {
	ch := &Character{
		Name: name,
		PCData: &PCData{
			Password: "",
			BamfIn:   "",
			BamfOut:  "",
			Title:    "",

			PermanentStrength:     12,
			PermanentIntelligence: 12,
			PermanentWisdom:       12,
			PermanentDexterity:    12,
			PermanentConstitution: 12,
		},
	}

	state := "none"
	for _, line := range strings.Split(data, "\n") {
		if len(line) == 0 {
			continue
		}

		if line[0] == '#' {
			switch line {
			case "#PLAYER":
				state = "player"
			case "#OBJECT":
				state = "object"
			case "#END":
				state = "end"
			}
			continue
		}

		if state == "end" {
			break
		}

		// key is first word, value is the rest
		parts := strings.SplitN(line, " ", 2)
		key, value := parts[0], parts[1]
		if state == "player" {
			switch key {
			case "Act":
				ch.Act = readNumber(value)
			case "AffectedBy":
				ch.AffectedBy = readNumber(value)
			case "Alignment":
				ch.Alignment = readNumber(value)
			case "Armor":
				ch.Armor = readNumber(value)
			case "Affect":
				parts := strings.Split(value, " ")
				affect := Affect{}
				skill, ok := Skills[readNumber(parts[0])]
				if !ok {
					log.Println("player load: unknown skill")
				} else {
					affect.Type = skill
				}

				affect.Duration = readNumber(parts[1])
				affect.Modifier = readNumber(parts[2])
				affect.Location = readNumber(parts[3])
				affect.Bitvector = readNumber(parts[4])
				ch.Affected = &affect

			case "AttrMod":
				attrs := strings.Split(value, " ")
				ch.PCData.ModifiedStrength = readNumber(attrs[0])
				ch.PCData.ModifiedIntelligence = readNumber(attrs[1])
				ch.PCData.ModifiedWisdom = readNumber(attrs[1])
				ch.PCData.ModifiedDexterity = readNumber(attrs[2])
				ch.PCData.ModifiedConstitution = readNumber(attrs[3])
			case "AttrPerm":
				attrs := strings.Split(value, " ")
				ch.PCData.PermanentStrength = readNumber(attrs[0])
				ch.PCData.PermanentIntelligence = readNumber(attrs[1])
				ch.PCData.PermanentWisdom = readNumber(attrs[1])
				ch.PCData.PermanentDexterity = readNumber(attrs[2])
				ch.PCData.PermanentConstitution = readNumber(attrs[3])
			case "Bamfin":
				ch.PCData.BamfIn = value
			case "Bamfout":
				ch.PCData.BamfOut = value
			case "Class":
				class, ok := Classes[value]
				if !ok {
					log.Println("load player: class not found", value)
					break
				}
				ch.Class = &class
			case "Damroll":
				ch.DamageRoll = readNumber(value)
			case "Deaf":
				ch.Deaf = readBoolean(value)
			case "Description":
				ch.Description = value
			case "Exp":
				ch.XP = readNumber(value)
			case "Gold":
				ch.Gold = readNumber(value)
			case "Hitroll":
				ch.HitRoll = readNumber(value)
			case "HpManaMove":
				attrs := strings.Split(value, " ")
				ch.HP = readNumber(attrs[0])
				ch.MaxHP = readNumber(attrs[1])
				ch.Mana = readNumber(attrs[2])
				ch.MaxMana = readNumber(attrs[3])
				ch.Move = readNumber(attrs[4])
				ch.MaxMove = readNumber(attrs[5])
			case "Level":
				ch.Level = readNumber(value)
			case "LongDescr":
				ch.LongDescrition = value
			case "Name":
				// Do nothing. Already set externally
			case "Password":
				ch.PCData.Password = value
			case "Played":
				ch.Played = readBoolean(value)
			case "Position":
				ch.Position = readNumber(value)
			case "Practice":
				ch.Practice = readNumber(value)
			case "Race":
				race, ok := Races[value]
				if !ok {
					log.Println("load player: race not found", value)
					break
				}
				ch.Race = &race
			case "Room":
				room, ok := GetWorld().Rooms[readNumber(value)]
				if !ok {
					log.Println("player load: room not found")
					break
				}
				ch.InRoom = room
			case "SavingThrow":
				ch.SavingThrow = readNumber(value)
			case "Sex":
				ch.Sex = readNumber(value)
			case "ShortDescr":
				ch.ShortDescription = value
			case "Skill":
				skill, ok := Skills[readNumber(value)]
				if !ok {
					log.Println("player load: skill not found")
					break
				}
				ch.Skills = append(ch.Skills, skill)
			case "Trust":
				ch.Trust = readNumber(value)
			case "Vnum":
				_, ok := GetWorld().Mobs[readNumber(value)]
				if !ok {
					log.Println("player load: mob not found")
					break
				}
				ch.VNUM = readNumber(value)
			case "Wimpy":
				ch.Wimpy = readNumber(value)
			default:
				log.Println("player load: no match on ", key)
			}
		} else if state == "object" {
			obj := &Object{}
			switch key {
			case "Affect", "AffectData":
				// TODO: load affects
			case "Cost":
				obj.Cost = readNumber(value)
			case "Description":
				obj.Description = value
			case "ExtraFlags":
				obj.ExtraFlags = readNumber(value)
			case "ExtraDescr":
				// TODO: load extra description
			case "ItemType":
				obj.ItemType = readNumber(value)
			case "Level":
				obj.Level = readNumber(value)
			case "Name":
				obj.Name = value
			case "Nest":
				// TODO: handle nested objects
			case "ShortDescr":
				obj.ShortDescription = value
			case "Spell":
				parts := strings.Split(value, " ")
				idx, sn := readNumber(parts[0]), readNumber(parts[1])
				_, ok := Skills[sn]
				if !ok {
					log.Println("load object: spell not found")
					break
				}
				if idx < 0 || idx > 3 {
					log.Println("load object: bad skill index")
				}
				obj.Values[idx] = sn
			case "Timer":
				obj.Timer = readNumber(value)
			case "Values":
				attrs := strings.Split(value, " ")
				for _, v := range attrs {
					obj.Values = append(obj.Values, readNumber(v))
				}
			case "Vnum":
				_, ok := Objects[readNumber(value)]
				if !ok {
					log.Println("load object: vnum not found")
					break
				}
				obj.Vnum = readNumber(value)
			case "WearFlags":
				obj.WearFlags = readNumber(value)
			case "WearLoc":
				obj.WearLocation = constants.WearLocation(readNumber(value))
			case "Weight":
				obj.Weight = readNumber(value)
			default:
				log.Println("load object: no match")
			}
		}
	}

	return ch
}
