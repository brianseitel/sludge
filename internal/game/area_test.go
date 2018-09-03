package game_test

import (
	"bytes"
	"testing"

	"github.com/brianseitel/sludge/internal/game"
	"github.com/stretchr/testify/assert"
)

func TestLoadArea(t *testing.T) {
	f := []byte(`{15 25} Copper  Chapel Catacombs~`)
	buf := bytes.NewBuffer(f)

	parser := game.Parser{buf}

	area := game.LoadArea(&parser)

	assert.Equal(t, "{15 25} Copper  Chapel Catacombs", area.Name)
}

func TestLoadHelps(t *testing.T) {
	f := []byte(`
-1 DIKU~
.                    Original game idea, concept, and design:

          Katja Nyboe               [Superwoman] (katz@freja.diku.dk)
          Tom Madsen              [Stormbringer] (noop@freja.diku.dk)
          Hans Henrik Staerfeldt           [God] (bombman@freja.diku.dk)
          Michael Seifert                 [Papi] (seifert@freja.diku.dk)
          Sebastian Hammer               [Quinn] (quinn@freja.diku.dk)

                     Additional contributions from:

Michael Curran  - the player title collection and additional locations.
Ragnar Loenn    - the bulletin board.
Bill Wisner     - for being the first to successfully port the game,
                  uncovering several old bugs, uh, inconsistencies,
                  in the process.

And: Mads Haar and Stephan Dahl for additional locations.

Developed at: DIKU -- The Department of Computer Science
                      at the University of Copenhagen.

~

`)
	buf := bytes.NewBuffer(f)

	parser := game.Parser{buf}

	game.LoadHelps(&parser)

	help := game.GetWorld().Helps["DIKU"]

	assert.Equal(t, -1, help.Level)
	assert.Equal(t, "DIKU", help.Keywords)
}

func TestLoadMobiles(t *testing.T) {
	f := []byte(`
#9301
poor mudder~
a poor mudder~
A poor mudder, lost in this place.
~
You will become him if you are not careful.
~
68 0 0 S
4 16 8 10d5+20 1d6+3
100 600
8 8 1
#9302
star~
a beautiful star~
A beautiful white star smiles at you.
~
You can't tell what the star looks like, she is too bright!
~
66 0 500 S
8 15 4 4d580 2d3+3
500 2000
8 8 2
#9303
nebula young~
a young nebula~
A young nebula, waiting to become a star.
~
It hasn't got a definite shape, it is just a cloud of mist.
~
64 0 200 S
10 10 2 3d6+240 2d4+4
1000 5000
8 8 2
#0

`)
	buf := bytes.NewBuffer(f)

	parser := game.Parser{buf}

	game.LoadMobiles(&parser)

	mob := game.GetWorld().Mobs[9302]

	assert.Equal(t, "star", mob.Name)

}

func TestLoadShops(t *testing.T) {
	game.GetWorld().Mobs[9238] = &game.Mob{}

	f := []byte(`9238	 0  0  0  0  0	 105  15	 0 23	; Elixir Vendor
0
`)
	buf := bytes.NewBuffer(f)

	parser := game.Parser{buf}

	game.LoadShops(&parser)

	shop := game.GetWorld().Shops[0]

	assert.Equal(t, 9238, shop.Keeper)
}

func TestLoadObjects(t *testing.T) {
	f := []byte(`#1001
rose~
a red rose~
A red rose is lying on the ground.~
~
9 0 17
6 0 0 0
8 0 200
E
rose~
This is no ordinary rose, as it is a gift from Dionysus to his
dearest, Wendella.
~
#0
`)

	buf := bytes.NewBuffer(f)

	parser := game.Parser{buf}

	game.LoadObjects(&parser)

	obj := game.GetWorld().Objects[1001]

	assert.Equal(t, 1001, obj.Vnum)

}

func TestLoadRooms(t *testing.T) {
	f := []byte(`#1001
In the air ...~
You are flying!
Currently you are about 40 meters above the ground north of the city.
It seems like you can see forever!
~
30 0 9
D1
More of the same.
~
~
0 0 1006
D2
More of the same.
~
~
0 0 1002
D4
More of the same.
~
~
0 0 1021
S
#0
`)
	buf := bytes.NewBuffer(f)

	parser := game.Parser{buf}

	game.LoadRooms(&parser, &game.Area{})

	room := game.GetWorld().Rooms[1001]
	assert.Equal(t, 1001, room.Vnum)
	assert.Equal(t, "In the air ...", room.Name)

}

func TestLoadResets(t *testing.T) {
	f := []byte(`R 0 9144 6
R 0 9145 6
R 0 9146 6
R 0 9147 6
R 0 9148 6
R 0 9149 6
R 0 9150 6
R 0 9151 6
R 0 9152 6
*
M 0 9101 40 9102        Lots of Harpies!!
M 0 9101 40 9102
M 0 9101 40 9105
M 0 9101 40 9105
M 0 9101 40 9127
M 0 9101 40 9127
M 0 9101 40 9127
M 0 9101 40 9108
M 0 9101 40 9108
M 0 9101 40 9108
M 0 9101 40 9112
M 0 9101 40 9112
M 0 9101 40 9117
M 0 9101 40 9119
M 0 9101 40 9119
M 0 9101 40 9119
M 0 9101 40 9121
M 0 9101 40 9122
M 0 9101 40 9122
M 0 9101 40 9129
M 0 9101 40 9131
M 0 9101 40 9133
M 0 9101 40 9135
M 0 9101 40 9144
M 0 9101 40 9145
M 0 9101 40 9146
M 0 9101 40 9147
M 0 9101 40 9149
M 0 9101 40 9150
M 0 9101 40 9151
M 0 9101 40 9152
M 0 9101 40 9158
M 0 9101 40 9158
M 0 9101 40 9162
M 0 9101 40 9162
M 0 9102 2 9123         Harpy Leader %1
G 1 9104 50                     magenta potion
M 0 9102 2 9124         Harpy Leader %2
G 1 9105 1                      ancient parchment
M 0 9103 5 9148
M 0 9103 5 9109
M 0 9103 5 9137
M 0 9103 5 9163
M 0 9103 5 9177
M 0 9104 2 9182         Efreet %1
G 1 9102 9                      bronze bracers
M 0 9104 2 9187         Efreet %2
G 1 9103 20                     ornate brooch
M 0 9105 1 9156         Female roc
M 0 9106 1 9153         Male roc
M 0 9107 1 9141         Wicked Witch
D 0 9141 1 1
S
`)
	buf := bytes.NewBuffer(f)

	parser := game.Parser{buf}

	area := &game.Area{}
	game.LoadResets(&parser, area)

	assert.Equal(t, "R", area.Resets[0].Command)
}

func TestLoadSpecials(t *testing.T) {
	f := []byte(`M  1000 spec_cast_mage
S
`)
	buf := bytes.NewBuffer(f)

	parser := game.Parser{buf}

	world := game.GetWorld()
	world.Mobs[1000] = &game.Mob{}

	game.LoadSpecials(&parser)

	mob := world.Mobs[1000]

	assert.Equal(t, "spec_cast_mage", mob.SpecFun)
}
