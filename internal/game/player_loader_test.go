package game_test

import (
	"testing"

	"github.com/brianseitel/sludge/internal/game"
	"github.com/stretchr/testify/assert"
)

func init() {
	game.NewWorld()
}

func TestLoadPlayerFile(t *testing.T) {
	ch, err := game.LoadPlayerFile("locke")
	if err != nil {
		panic(err)
	}
	assert.NotNil(t, ch)
}

func TestPlayerFileToCharacter(t *testing.T) {
	data := `#PLAYER
Name locke
Password lamora
Level 1
Act 1
AffectedBy 1
Alignment 1000
Armor -12
AttrMod 12 12 12 12 12
AttrPerm 10 10 10 10 10
Affect 1 60 0 0 0 0
Bamfin holla!
Bamfout what!
Class war
Damroll 10
Deaf false
Description I am Locke Lamora
Exp 1000
Gold 1000
Hitroll 10
HpManaMove 100 100 100 100 100 100
Position standing
Practice 32
LongDescr I am Locke Lamora
Race elf
Room 1
SavingThrow 3
Sex male
ShortDescr Me llamo Locke Lamora
Skill 0 0 0 0
Trust 1
Wimpy 100
Unrecognized Key Goes Here

#OBJECT
Affect 1 60 0 0 0
AffectedBy [stuff]
Cost 1
Description A thingery for thinging
ExtraFlags 1
ItemType 1
Level 1
Name Thinger
Nest 1
ShortDescr 1
Spell 1 0 0 0 0
Timer 0
Values 1 2 3 4
VNUM 1
WearFlags 1
WearLoc 1
Weight 0
Unrecognized Key Goes Here
#END
`

	ch := game.ConvertPlayerFileToCharacter("locke", data)

	assert.NotNil(t, ch)
	assert.Equal(t, "locke", ch.Name)
	assert.Equal(t, "lamora", ch.PCData.Password)
	assert.Equal(t, 1, ch.Level)
}
