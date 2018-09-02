package constants

// ActType ...
type ActType int

// World constants
const (
	ActToRoom ActType = iota
	ActToVictim
	ActToCharacter
	ActToNotVictim
)
