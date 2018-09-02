package constants

// Sexes
const (
	SexFemale = iota
	SexMale
	SexNeutral
)

// Prime Attributes
const (
	Strength = iota
	Intelligence
	Wisdom
	Dexterity
	Constitution
)

// WearLocation ...
type WearLocation int

// Wear Locations
const (
	WearNone  WearLocation = -1
	WearLight WearLocation = iota
	WearFingerLeft
	WearFingerRight
	WearNeck1
	WearNeck2
	WearBody
	WearHead
	WearLegs
	WearFeet
	WearHands
	WearArms
	WearShield
	WearAbout
	WearWaist
	WearWristLeft
	WearWristRight
	WearWield
	WearHold
)
