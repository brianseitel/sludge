package game

import "github.com/brianseitel/sludge/internal/constants"

// Class ...
type Class struct {
	// The name of the class
	Name string
	/*
		The three letter name is used for the 'who' listing. it's also
		used for the list of classes shown to new characters for selecting a class,
		and for matching the character's input in selecting a class.
	*/
	WhoName string

	/*
		This attribute is initialized to 16 for new characters of this class. It costs
		only three practices to train one's prime attribute, versus five practices for
		other attributes. In addition, characters may increase their prime attribute
		(only) over 18 by using magic items.
	*/
	PrimeAttribute int

	// Weapon (vnum) is given to new characters of this class for their first weapon
	Weapon VNUM

	// This room (vnum) is off limits to characters of other classes
	Guild VNUM

	// The maximum level to which a character may train a skill or spell
	SkillAdept int

	Thac0_00 int // Thac0 for level 0
	Thac0_32 int // Thac0 for level 32

	HPmin     int
	HPmax     int
	GainsMana bool
}

// Classes ...
var Classes = map[string]Class{
	"war": Class{
		Name:           "Warrior",
		WhoName:        "war",
		PrimeAttribute: constants.Strength,
		Weapon:         constants.VnumObjectSchoolSword,
		Guild:          constants.VnumRoomGuildWarrior,
		SkillAdept:     85,
		Thac0_00:       18,
		Thac0_32:       6,
		HPmin:          11,
		HPmax:          15,
		GainsMana:      false,
	},
	"mag": Class{
		Name:           "Mage",
		WhoName:        "mag",
		PrimeAttribute: constants.Wisdom,
		Weapon:         constants.VnumObjectSchoolDagger,
		Guild:          constants.VnumRoomGuildMage,
		SkillAdept:     95,
		Thac0_00:       18,
		Thac0_32:       10,
		HPmin:          6,
		HPmax:          8,
		GainsMana:      true,
	},
	"cle": Class{
		Name:           "Cleric",
		WhoName:        "cle",
		PrimeAttribute: constants.Wisdom,
		Weapon:         constants.VnumObjectSchoolMace,
		Guild:          constants.VnumRoomGuildCleric,
		SkillAdept:     95,
		Thac0_00:       18,
		Thac0_32:       12,
		HPmin:          7,
		HPmax:          10,
		GainsMana:      true,
	},
	"thi": Class{
		Name:           "Thief",
		WhoName:        "thi",
		PrimeAttribute: constants.Wisdom,
		Weapon:         constants.VnumObjectSchoolDagger,
		Guild:          constants.VnumRoomGuildThief,
		SkillAdept:     85,
		Thac0_00:       18,
		Thac0_32:       8,
		HPmin:          8,
		HPmax:          13,
		GainsMana:      false,
	},
}

// Titles Map
var Titles = map[string][][]string{
	"mag": [][]string{
		[]string{"Apprentice of Magic", "Apprentice of Magic"},
		[]string{"Spell Student", "Spell Student"},
		[]string{"Scholar of Magic", "Scholar of Magic"},
		[]string{"Delver in Spells", "Delveress in Spells"},
		[]string{"Medium of Magic", "Medium of Magic"},

		[]string{"Scribe of Magic", "Scribess of Magic"},
		[]string{"Seer", "Seeress"},
		[]string{"Sage", "Sage"},
		[]string{"Illusionist", "Illusionist"},
		[]string{"Abjurer", "Abjuress"},

		[]string{"Invoker", "Invoker"},
		[]string{"Enchanter", "Enchantress"},
		[]string{"Conjurer", "Conjuress"},
		[]string{"Magician", "Witch"},
		[]string{"Creator", "Creator"},

		[]string{"Savant", "Savant"},
		[]string{"Magus", "Craftess"},
		[]string{"Wizard", "Wizard"},
		[]string{"Warlock", "War Witch"},
		[]string{"Sorcerer", "Sorceress"},

		[]string{"Elder Sorcerer", "Elder Sorceress"},
		[]string{"Grand Sorcerer", "Grand Sorceress"},
		[]string{"Great Sorcerer", "Great Sorceress"},
		[]string{"Golem Maker", "Golem Maker"},
		[]string{"Greater Golem Maker", "Greater Golem Maker"},

		[]string{"Maker of Stones", "Maker of Stones"},
		[]string{"Maker of Potions", "Maker of Potions"},
		[]string{"Maker of Scrolls", "Maker of Scrolls"},
		[]string{"Maker of Wands", "Maker of Wands"},
		[]string{"Maker of Staves", "Maker of Staves"},

		[]string{"Demon Summoner", "Demon Summoner"},
		[]string{"Greater Demon Summoner", "Greater Demon Summoner"},
		[]string{"Dragon Charmer", "Dragon Charmer"},
		[]string{"Greater Dragon Charmer", "Greater Dragon Charmer"},
		[]string{"Master of all Magic", "Master of all Magic"},

		[]string{"Mage Hero", "Mage Heroine"},
		[]string{"Angel of Magic", "Angel of Magic"},
		[]string{"Deity of Magic", "Deity of Magic"},
		[]string{"Supremity of Magic", "Supremity of Magic"},
		[]string{"Implementor", "Implementress"},
	},
	"cle": [][]string{
		[]string{"Believer", "Believer"},
		[]string{"Attendant", "Attendant"},
		[]string{"Acolyte", "Acolyte"},
		[]string{"Novice", "Novice"},
		[]string{"Missionary", "Missionary"},

		[]string{"Adept", "Adept"},
		[]string{"Deacon", "Deaconess"},
		[]string{"Vicar", "Vicaress"},
		[]string{"Priest", "Priestess"},
		[]string{"Minister", "Lady Minister"},

		[]string{"Canon", "Canon"},
		[]string{"Levite", "Levitess"},
		[]string{"Curate", "Curess"},
		[]string{"Monk", "Nun"},
		[]string{"Healer", "Healess"},

		[]string{"Chaplain", "Chaplain"},
		[]string{"Expositor", "Expositress"},
		[]string{"Bishop", "Bishop"},
		[]string{"Arch Bishop", "Arch Lady of the Church"},
		[]string{"Patriarch", "Matriarch"},

		[]string{"Elder Patriarch", "Elder Matriarch"},
		[]string{"Grand Patriarch", "Grand Matriarch"},
		[]string{"Great Patriarch", "Great Matriarch"},
		[]string{"Demon Killer", "Demon Killer"},
		[]string{"Greater Demon Killer", "Greater Demon Killer"},

		[]string{"Cardinal of the Sea", "Cardinal of the Sea"},
		[]string{"Cardinal of the Earth", "Cardinal of the Earth"},
		[]string{"Cardinal of the Air", "Cardinal of the Air"},
		[]string{"Cardinal of the Ether", "Cardinal of the Ether"},
		[]string{"Cardinal of the Heavens", "Cardinal of the Heavens"},

		[]string{"Avatar of an Immortal", "Avatar of an Immortal"},
		[]string{"Avatar of a Deity", "Avatar of a Deity"},
		[]string{"Avatar of a Supremity", "Avatar of a Supremity"},
		[]string{"Avatar of an Implementor", "Avatar of an Implementor"},
		[]string{"Master of all Divinity", "Mistress of all Divinity"},

		[]string{"Holy Hero", "Holy Heroine"},
		[]string{"Angel", "Angel"},
		[]string{"Deity", "Deity"},
		[]string{"Supreme Master", "Supreme Mistress"},
		[]string{"Implementor", "Implementress"},
	},
	"thi": [][]string{
		[]string{"Pilferer", "Pilferess"},
		[]string{"Footpad", "Footpad"},
		[]string{"Filcher", "Filcheress"},
		[]string{"Pick-Pocket", "Pick-Pocket"},
		[]string{"Sneak", "Sneak"},

		[]string{"Pincher", "Pincheress"},
		[]string{"Cut-Purse", "Cut-Purse"},
		[]string{"Snatcher", "Snatcheress"},
		[]string{"Sharper", "Sharpress"},
		[]string{"Rogue", "Rogue"},

		[]string{"Robber", "Robber"},
		[]string{"Magsman", "Magswoman"},
		[]string{"Highwayman", "Highwaywoman"},
		[]string{"Burglar", "Burglaress"},
		[]string{"Thief", "Thief"},

		[]string{"Knifer", "Knifer"},
		[]string{"Quick-Blade", "Quick-Blade"},
		[]string{"Killer", "Murderess"},
		[]string{"Brigand", "Brigand"},
		[]string{"Cut-Throat", "Cut-Throat"},

		[]string{"Spy", "Spy"},
		[]string{"Grand Spy", "Grand Spy"},
		[]string{"Master Spy", "Master Spy"},
		[]string{"Assassin", "Assassin"},
		[]string{"Greater Assassin", "Greater Assassin"},

		[]string{"Master of Vision", "Mistress of Vision"},
		[]string{"Master of Hearing", "Mistress of Hearing"},
		[]string{"Master of Smell", "Mistress of Smell"},
		[]string{"Master of Taste", "Mistress of Taste"},
		[]string{"Master of Touch", "Mistress of Touch"},

		[]string{"Crime Lord", "Crime Mistress"},
		[]string{"Infamous Crime Lord", "Infamous Crime Mistress"},
		[]string{"Greater Crime Lord", "Greater Crime Mistress"},
		[]string{"Master Crime Lord", "Master Crime Mistress"},
		[]string{"Godfather", "Godmother"},

		[]string{"Assassin Hero", "Assassin Heroine"},
		[]string{"Angel of Death", "Angel of Death"},
		[]string{"Deity of Assassins", "Deity of Assassins"},
		[]string{"Supreme Master", "Supreme Mistress"},
		[]string{"Implementor", "Implementress"},
	},
	"war": [][]string{
		[]string{"Swordpupil", "Swordpupil"},
		[]string{"Recruit", "Recruit"},
		[]string{"Sentry", "Sentress"},
		[]string{"Fighter", "Fighter"},
		[]string{"Soldier", "Soldier"},

		[]string{"Warrior", "Warrior"},
		[]string{"Veteran", "Veteran"},
		[]string{"Swordsman", "Swordswoman"},
		[]string{"Fencer", "Fenceress"},
		[]string{"Combatant", "Combatess"},

		[]string{"Hero", "Heroine"},
		[]string{"Myrmidon", "Myrmidon"},
		[]string{"Swashbuckler", "Swashbuckleress"},
		[]string{"Mercenary", "Mercenaress"},
		[]string{"Swordmaster", "Swordmistress"},

		[]string{"Lieutenant", "Lieutenant"},
		[]string{"Champion", "Lady Champion"},
		[]string{"Dragoon", "Lady Dragoon"},
		[]string{"Cavalier", "Lady Cavalier"},
		[]string{"Knight", "Lady Knight"},

		[]string{"Grand Knight", "Grand Knight"},
		[]string{"Master Knight", "Master Knight"},
		[]string{"Paladin", "Paladin"},
		[]string{"Grand Paladin", "Grand Paladin"},
		[]string{"Demon Slayer", "Demon Slayer"},

		[]string{"Greater Demon Slayer", "Greater Demon Slayer"},
		[]string{"Dragon Slayer", "Dragon Slayer"},
		[]string{"Greater Dragon Slayer", "Greater Dragon Slayer"},
		[]string{"Underlord", "Underlord"},
		[]string{"Overlord", "Overlord"},

		[]string{"Baron of Thunder", "Baroness of Thunder"},
		[]string{"Baron of Storms", "Baroness of Storms"},
		[]string{"Baron of Tornadoes", "Baroness of Tornadoes"},
		[]string{"Baron of Hurricanes", "Baroness of Hurricanes"},
		[]string{"Baron of Meteors", "Baroness of Meteors"},

		[]string{"Knight Hero", "Knight Heroine"},
		[]string{"Angel of War", "Angel of War"},
		[]string{"Deity of War", "Deity of War"},
		[]string{"Supreme Master of War", "Supreme Mistress of War"},
		[]string{"Implementor", "Implementress"},
	},
}
