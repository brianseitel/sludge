package game

func init() {
	Skills = make(map[int]Skill, 0)
}

// Skill ...
type Skill struct {
	Name string
}

// Skills ...
var Skills = map[int]Skill{}
