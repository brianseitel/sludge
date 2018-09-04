package game

import (
	"math/rand"
	"sync"
)

var world *World
var worldOnce sync.Once

// World ...
type World struct {
	Wizlocked  bool
	Characters []*Character
	Areas      []*Area
	Shops      []*Shop

	Rooms   map[int]*Room
	Mobs    map[int]*Mob
	Helps   map[string]*Help
	Objects map[int]*Object
}

// NewWorld ...
func NewWorld() *World {
	worldOnce.Do(func() {
		world = &World{
			Wizlocked: false,
			Rooms:     make(map[int]*Room, 0),
			Mobs:      make(map[int]*Mob, 0),
			Helps:     make(map[string]*Help, 0),
			Objects:   make(map[int]*Object, 0),
		}

		// TODO: set weather
	})

	return world
}

// GetWorld ...
func GetWorld() *World {
	return world
}

// UpdateAreas ...
func (w *World) UpdateAreas() {
	for _, a := range w.Areas {
		a.Age++
		if a.Age < 3 {
			continue
		}

		if a.NumPlayers == 0 && a.Age >= 15 {
			a.Reset()
			a.Age = rand.Intn(3)
		}
	}
}
