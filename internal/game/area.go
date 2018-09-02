package game

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/sanity-io/litter"
)

// Area ...
type Area struct {
	Name       string
	Age        int
	NumPlayers int
}

// Help ...
type Help struct {
	Level    int
	Keywords string
	Text     string
}

// LoadAreas ...
func LoadAreas() {
	world := GetWorld()
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b) + "/../../"

	areaFiles, _ := filepath.Glob(basepath + "areas/*.are")

	for _, file := range areaFiles {
		f, err := ioutil.ReadFile(file)
		if err != nil {
			panic(err)
		}

		area := Area{}

		lines := strings.Split(string(f), "\n")

		var help *Help
		var mob *Mob

		state := "parsing"
		for _, line := range lines {
			if len(line) == 0 {
				continue
			}

			if state == "parsing" && line[0] == '#' {
				parts := strings.SplitN(line, "\t", 2)
				word := parts[0]
				var rest string
				if len(parts) > 1 {
					rest = parts[1]
				}
				state = word

				if state == "#AREA" {
					area.Name = rest
					area.Age = 15
					area.NumPlayers = 0
				}
				continue
			}

			switch state {
			case "#HELPS":
				if help == nil {
					help = &Help{}
					parts := strings.SplitN(line, " ", 2)
					help.Level = readNumber(parts[0])
					help.Keywords = strings.Trim(parts[1], "~")
					continue
				}

				if line != "~" {
					help.Text += line + "\n"
				} else {
					world.Helps[help.Keywords] = help
					help = nil
				}
			case "#MOBILES":
				if mob == nil {
					mob = &Mob{}
					continue
				}

				if line[0] == '#' {
					// VNUM
					vnum := readNumber(line[1:])
					if vnum == 0 {
						break
					}
					if _, ok := world.Mobs[vnum]; ok {
						log.Println("load mobs: vnum ", vnum, " duplicated")
						os.Exit(1)
					}
					mob.Vnum = vnum
				}
			}
		}

		litter.Dump(world.Helps)
	}
}
