package game

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/brianseitel/sludge/internal/constants"
	"github.com/brianseitel/sludge/internal/helpers"
)

// Parser for reading files
type Parser struct {
	*bytes.Buffer
}

// Area ...
type Area struct {
	Name       string
	Age        int
	NumPlayers int

	Resets []*Reset
}

// Reset ...
type Reset struct {
	Command string
	Arg1    int
	Arg2    int
	Arg3    int
}

// Help ...
type Help struct {
	Level    int
	Keywords string
	Text     string
}

// LoadAreas ...
func LoadAreas() {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b) + "/../../"

	areaFiles, _ := filepath.Glob(basepath + "areas/*.are")

	for _, file := range areaFiles {
		f, err := ioutil.ReadFile(file)
		if err != nil {
			panic(err)
		}

		fmt.Println(file)
		load(f)
	}
}

func load(f []byte) {
	buf := bytes.NewBuffer(f)
	p := &Parser{buf}

	l := p.letter()
	if l != "#" {
		panic("# not found")
	}
	p.UnreadByte()

	var area *Area
	for p.Len() > 0 {
		w := p.word()

		switch w {
		case "#AREA":
			fmt.Println("\t", w)
			area = LoadArea(p)
		case "#HELPS":
			fmt.Println("\t", w)
			LoadHelps(p)
		case "#MOBILES":
			fmt.Println("\t", w)
			LoadMobiles(p)
		case "#OBJECTS":
			fmt.Println("\t", w)
			LoadObjects(p)
		case "#ROOMS":
			fmt.Println("\t", w)
			LoadRooms(p, area)
		case "#RESETS":
			fmt.Println("\t", w)
			LoadResets(p, area)
		case "#SPECIALS":
			fmt.Println("\t", w)
			LoadSpecials(p)
		case "#SHOPS":
			fmt.Println("\t", w)
			LoadShops(p)
		case "#$":
			fmt.Println("done")
			break
		}
	}
	// os.Exit(1)
	// fmt.Println(p.letter())
	// fmt.Println(p.word())
	// fmt.Println(p.line())
	// fmt.Println(p.eol())
}

func LoadArea(p *Parser) *Area {
	area := &Area{}
	area.Name = p.line()
	area.Age = 15
	area.NumPlayers = 0
	fmt.Println(area.Name)
	return area
}

func LoadHelps(p *Parser) {
	for {
		help := &Help{}
		p.trim()

		help.Level = p.number()
		help.Keywords = p.line()
		if help.Keywords == "$" {
			return
		}
		help.Text = p.line()
		world.Helps[help.Keywords] = help
	}
}

func LoadMobiles(p *Parser) {
	p.trim()
	for {
		mob := &Mob{}

		l := p.letter()
		if l != "#" {
			p.UnreadByte()
			fmt.Println("line", p.line())
			log.Fatalf("load mobiles: expected #, got %s", l)
		}

		vnum := p.number()
		if vnum == 0 {
			break
		}

		if _, ok := world.Mobs[vnum]; ok {
			log.Fatalf("load mobiles: vnum %d duplicated", vnum)
		}

		mob.Vnum = vnum
		mob.Name = p.line()
		mob.ShortDescription = p.line()
		mob.LongDescription = p.line()
		mob.Description = p.line()

		mob.Act = p.number() | constants.ActIsNPC
		mob.AffectedBy = p.number()
		mob.Shop = nil
		mob.Alignment = p.number()
		p.letter() // throwaway i think?
		mob.Level = helpers.Fuzzy(p.number())

		p.word()
		p.word()
		p.word()
		p.word()
		p.word()
		p.word()
		p.word()
		p.word()
		mob.Sex = p.number()

		world.Mobs[vnum] = mob
	}
}

func LoadShops(p *Parser) {
	p.trim()
	for {
		shop := &Shop{}

		shop.Keeper = p.number()
		if shop.Keeper == 0 {
			break
		}

		maxTrade := 5
		for t := 0; t < maxTrade; t++ {
			shop.BuyType[t] = p.number()
		}
		shop.ProfitBuy = p.number()
		shop.ProfitSell = p.number()
		shop.OpenHour = p.number()
		shop.CloseHour = p.number()

		mob := world.Mobs[shop.Keeper]
		mob.Shop = shop
		world.Shops = append(world.Shops, shop)
	}
}

func LoadObjects(p *Parser) {
	p.trim()
	for {
		obj := &Object{}

		l := p.letter()
		if l != "#" {
			log.Fatalf("load objects: expected #, got %s", l)
		}

		vnum := p.number()
		if vnum == 0 {
			return
		}

		if _, ok := world.Objects[vnum]; ok {
			log.Fatalf("load object: vnum %d duplicated", vnum)
		}

		obj.Vnum = vnum
		obj.Name = p.line()
		obj.ShortDescription = p.line()
		obj.Description = p.line()
		p.line() // unused: action description

		obj.ShortDescription = strings.ToLower(string(obj.ShortDescription[0])) + obj.ShortDescription[1:]

		if len(obj.Description) > 1 {
			obj.Description = strings.ToUpper(string(obj.Description[0])) + obj.Description[1:]
		}

		obj.ItemType = p.number()
		obj.ExtraFlags = p.number()
		obj.WearFlags = p.number()

		obj.Values = make([]int, 4)
		obj.Values[0] = p.number()
		obj.Values[1] = p.number()
		obj.Values[2] = p.number()
		obj.Values[3] = p.number()
		obj.Weight = p.number()
		obj.Cost = p.number()
		p.number() // unused: cost per day

		if obj.ItemType == ItemPotion {
			obj.ExtraFlags = constants.ItemNoDrop
		}

		for {
			l := p.letter()
			if l == "~" {
				l = p.letter()
			}
			if l == "A" {
				affect := &Affect{}

				affect.Duration = -1
				affect.Location = p.number()
				affect.Modifier = p.number()
				affect.Bitvector = 0

				obj.Affected = append(obj.Affected, affect)
			} else if l == "E" {
				// TODO: extra descriptions

				p.line()
				p.line()

			} else {
				p.UnreadByte()
				break
			}
		}

		// TODO: slot lookups
		world.Objects[vnum] = obj
	}
}

func LoadRooms(p *Parser, area *Area) {
	for {
		p.trim()
		room := &Room{}

		l := p.letter()
		if l != "#" {
			p.UnreadByte()
			log.Fatalf("load rooms: expected #, got %v", l)
		}

		vnum := p.number()
		if vnum == 0 {
			break
		}

		if _, ok := world.Rooms[vnum]; ok {
			log.Fatalf("load room: vnum %d duplicated", vnum)
		}

		room.Vnum = vnum
		room.Area = area

		room.People = make([]*Character, 0)
		room.Contents = make([]*Object, 0)
		room.ExtraDescription = ""

		room.Name = p.line()
		room.Description = strings.Trim(p.line(), " ")
		p.number() // unused: area number
		room.Flags = p.number()
		room.SectorType = p.number()
		room.Light = 0
		room.Exits = make([]*Exit, 6)
		for d := 0; d <= 5; d++ {
			room.Exits[d] = nil
		}

		for {
			l := p.letter()
			if l == "S" {
				break
			}

			// Handle Door
			if l == "D" {
				exit := &Exit{}
				door := p.number()

				if door < 0 || door > 5 {
					log.Fatalf("load room: vnum %d has bad number", vnum)
				}

				exit.Description = p.line()
				exit.Keywords = p.line()
				exit.Info = 0
				locks := p.number()
				exit.Key = p.number()
				exit.Vnum = p.number()

				switch locks {
				case 1:
					exit.Info = ExitIsDoor
				case 2:
					exit.Info = ExitIsDoor | ExitPickProof
				}

				room.Exits[door] = exit
			} else if l == "E" {
				// TODO: handle extra descriptions
				p.line()
				p.line()
			} else {
				log.Printf("Load rooms: vnum %d has flag not DES\n", vnum)
			}
		}

		world.Rooms[vnum] = room
	}
}

func LoadResets(p *Parser, area *Area) {
	p.trim()
	for {

		if area == nil {
			log.Fatalf("load resets: no #AREA seen yet")
		}

		l := p.letter()

		if l == "S" {
			break
		}

		if l == "*" {
			p.eol()
			continue
		}

		reset := &Reset{}
		reset.Command = l
		// p.number() // unused: if flag
		reset.Arg1 = p.number()
		reset.Arg2 = p.number()
		reset.Arg3 = p.number()

		p.eol()

		area.Resets = append(area.Resets, reset)
	}
}

func LoadSpecials(p *Parser) {
	for {
		p.trim()
		l := p.letter()
		switch l {
		case "S":
			return

		case "*":
			break

		case "M":
			vnum := p.number()
			if _, ok := world.Mobs[vnum]; ok {
				// TODO: lookup spec fun
				world.Mobs[vnum].SpecFun = p.word()
			}
		}
	}
}

func (p *Parser) trim() {
	for {
		b, _ := p.ReadByte()
		if b == ' ' || b == '\t' || b == '\n' {
			continue
		} else {
			p.UnreadByte() // put the byte back!
			break
		}
	}
}

func (p *Parser) letter() string {
	var r rune
	for {
		r, _, _ = p.ReadRune()
		if !(r == ' ' || r == '\t' || r == '\n' || r == '\v' || r == '\f' || r == '\r') {
			break
		}
	}

	return string(r)
}

func (p *Parser) word() string {
	var word string
	for {
		b, _ := p.ReadByte()
		if b == ' ' || b == '\t' || b == '\n' {
			return strings.Trim(word, " ")
		}
		word += string(b)
	}
}

func (p *Parser) eol() string {
	var word string
	for {
		b, _ := p.ReadByte()
		word += string(b)
		if b == '\n' || b == '\r' {
			return strings.Trim(word, " ")
		}
	}
}

func (p *Parser) line() string {
	l, _ := p.ReadString(byte('~'))

	l = strings.TrimLeft(l, "\n\r")
	l = strings.TrimRight(l, "~\n\r")
	return l
}

func (p *Parser) number() int {
	var word string
	for {
		word = p.word()
		if word != "" {
			break
		}
	}

	if strings.Contains(word, "|") {
		parts := strings.Split(word, "|")
		num := 0
		for _, p := range parts {
			n, _ := strconv.Atoi(p)
			num += n
		}
		return num
	}

	if strings.Contains(word, "+") {
		parts := strings.SplitN(word, "+", 2)
		diceParts := strings.Split(parts[0], "d")
		numDice, _ := strconv.Atoi(diceParts[0])
		sides, _ := strconv.Atoi(diceParts[1])
		plus, _ := strconv.Atoi(parts[2])

		fmt.Println(word)
		return numDice*sides + plus
	}

	num, err := strconv.Atoi(word)
	if err != nil {
		return 0
		// panic(err)
	}
	return num
}
