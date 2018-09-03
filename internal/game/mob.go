package game

// Mob ...
type Mob struct {
	Vnum int

	Name             string
	ShortDescription string
	LongDescription  string
	Description      string

	Act        int
	AffectedBy int
	Shop       *Shop
	Alignment  int
	Level      int

	Sex int

	SpecFun string
}
