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

// Mob ACT Constants
const (
	ActIsNPC = iota
)

// Item Extra Flags

const (
	ItemNoDrop = iota
)
