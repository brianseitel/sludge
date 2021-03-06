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
#2200
hatchling dragon baby~
the dragon hatchling~
A dragon hatchling is here, chewing on a bone.
~
This small dragon hisses at you as you enter the room.  Although it is only a
few feet long, its sharp teeth make you think twice about petting it.
~
2|64 524288 -900 S
7 12 4 4d8+100 1d8+2
2000 2250
8 8 0
#2201
draconian~
the Draconian~
The Draconian is standing here.
~
This horrible creature is a bizarre cross between a man and a dragon.  He
has black scales and a seven foot wingspan.  He scowls at you and hefts his
spear as you enter the room.
~
2|4|32 8 -900 S
13 8 2 13d13+130 2d6+3
2000 15000
8 8 1
#2202
master~
the dragon master~
The lord of this crypt is here.
~
He doesn't look happy to see you.
~
2|4|32 524288 -1000 S
27 0 -2 27d27+270 3d6+8
35000 60000
8 8 1
#2203
man mage wizard~
a powerful mage~
A man is here, studying some books.
~
A powerful looking mage is standing in this room studying his spells.  You are
surprised at the fact that he is human, and realize that he must be very
powerful to have been accepted by the creatures which live here.
~
2 2|8|128 -900 S
22 0 3 22d22+220 1d8+10
600 45000
8 8 1
#2204
cleric draconian~
the holy Draconian~
A Draconian is here, deep in thought.
~
A Draconian stands here, dressed in simple robes.  Around his neck you
notice a golden medallion in the shape of a five headed dragon.
~
2|4|32 8|128 -900 S
22 1 0 22d22+220 3d5+6
400 45000
8 8 1
#2205
king draconian~
the Draconian King~
A Draconian wearing fine clothes is here, pondering his greatness.
~
The king of the draconians sits here on a golden throne.  He looks as
though he could crush your head with a single blow.  Yet you sense that he
is controlled by a greater power.
~
2|4 8|128 -900 S
26 0 1 26d26+260 2d8+12
800 100000
8 8 1
#2206
concubine draconian~
a concubine~
A concubine is resting here.
~
This is a concubine of the king.  She is lounging here, wearing nothing at
all.  You find yourself strangely repulsed as she draws a knife and
prepares to defend herself.
~
2 0 -500 S
6 16 5 6d10+50 2d3+1
100 1000
8 8 2
#2207
bodyguard draconian~
A bodyguard~
A bodyguard is here, staring at you menacingly.
~
This creature has devoted its life to defending the king.  Your intrusion has
not pleased it.
~
2|32 4|8|32 -800 S
18 3 2 18d18+180 2d6+10
5000 30000
8 8 1
#2220
tiamat dragon~
Tiamat~
A five headed dragon hisses at you as you enter this room.
~
You see before you the master of this Tower, Tiamat.  She frowns at you as she
prepares to make you pay for your insolence ... with your lives!
~
2|4|32 8|32 -1000 S
35 0 -10 35d35+1000 4d9+15
50000 3500000
8 8 2
#2221
dragon red~
the Great Red Dragon~
A red dragon is here, contemplating your existence.
~
This huge red dragon dominates the chamber.  As you turn to flee, he grins at
you and invites you to stay ... forever!
~
2|4|32 8|32 -1000 S
27 0 -5 27d27+270 2d8+10
10000 300000
8 8 1
#2222
dragon black~
the Great Black Dragon~
A black dragon is here, laughing at your insolence.
~
This huge black wyrm laughs at your puny weapons.  You realize that he is not
going to roll over and die for you.
~
2|4|32 8|32 -900 S
26 1 -3 26d26+260 3d8+4
20000 250000
8 8 1
#2223
dragon white~
the Great White Dragon~
A white dragon is here, waiting for you.
~
This dragon towers over you.  Recalling your previous experiences with white
dragons, you are not afraid.  She smiles at you and says 'I think you will find
me a greater challenge than my offspring.'
~
2|4|32 8|32 -1000 S
27 0 -3 27d27+270 3d8+4
35000 200000
8 8 2
#2225
dragon green~
the Ancient Green Dragon~
A green dragon is here, looking distraught.
~
This huge beast appears to be the most directly related to the draconians.
She looks at you with sorrow in her eyes and says 'You have slaughtered my
children.  Prepare to die.'
~
2|4|32 8|32 -1000 S
31 0 -4 31d31+310 3d8+15
20000 500000
8 8 2
#2226
hydra~
a large hydra~
A hydra is here, blocking the doorway.
~
A hydra is here, guarding the entrance to a sealed vault.  You have the feeling
it isn't happy to see you.
~
2|4|32 8|32 -800 S
18 3 1 18d18+180 2d8+7
5000 35000
8 8 0
#2227
slave human~
A human slave~
A Human slave is here, hard at work.
~
He looks like he's under a lot of stress.  He would probably like it if you
left him alone.
~
2 0 0 S
5 16 6 1d10+50 1d6+4
100 900
8 8 1
#2240
zombie draconian~
a Draconian zombie~
A Draconian zombie is here, staring at nothing in particular.
~
This used to be one of the warriors of this tribe.  He was denied his final
rest, however, and now guards this corridor tirelessly as a zombie.
~
2|4|32 8|32 -400 S
9 9 4 1d100+175 1d8+6
1500 3000
8 8 0
#2241
dragon phase~
the phase dragon~
A Phase Dragon is darting around the room.
~
This small dragon looks like he's up to no good.
~
4|32|64|128 8|32 -300 S
10 11 3 10d10+150 2d6+6
5000 6000
8 8 1
#2242
fool draconian~
the draconian fool~
A Fool is here, making fun of you.
~
This dragon man looks VERY foolish, dressed in a green and blue striped
suit.  As you enter the room, he points at you and laughs.
~
64|128 0 0 S
5 16 5 1d10+50 3d3+4
500 800
8 8 1
#2243
queen draconian~
the Draconian Queen~
A Female Draconian sits, here, looking very important.
~
This draconian looks VERY busy as she orders her slaves around.  You think
it might be a good idea to leave her alone.
~
2|4 8|32|128 -600 S
20 1 2 20d20+200 3d4+8
10000 45000
8 8 2
#0
`)
	buf := bytes.NewBuffer(f)

	parser := game.Parser{buf}

	game.LoadMobiles(&parser)

	mob := game.GetWorld().Mobs[2240]

	assert.Equal(t, "zombie draconian", mob.Name)

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
