package abilities

// Abilities indicates a creature's basic abilities.
type Abilities struct {
	Strength     int `json:"strength" bson:"strength"`
	Dexterity    int `json:"dexterity" bson:"dexterity"`
	Constitution int `json:"constitution¨ bson:"constitution¨"`
	Intelligence int `json:"intelligence" bson:"intelligence"`
	Wisdom       int `json:"wisdom" bson:"wisdom"`
	Charisma     int `json:"charisma" bson:"charisma"`

	StrengthModifier     int `json:"strength_modifier" bson:"strength_modifier"`
	DexterityModifier    int `json:"dexterity_modifier" bson:"dexterity_modifier"`
	ConstitutionModifier int `json:"constitution_modifier" bson:"constitution_modifier"`
	IntelligenceModifier int `json:"intelligence_modifier" bson:"intelligence_modifier"`
	WisdomModifier       int `json:"wisdom_modifier" bson:"wisdom_modifier"`
	CharismaModifier     int `json:"charisma_modifier" bson:"charisma_modifier"`
}

// minimumAbilityScore indicates the minimum acceptable value for an ability score.
const minimumAbilityScore = 1

// maximumAbilityScore indicates the maximum acceptable value for an ability score.
const maximumAbilityScore = 30

// abilityScoresAndModifiers maps an ability score to an ability modifier.
var abilityScoresAndModifiers = map[int]int{
	1:  -5,
	2:  -4,
	3:  -4,
	4:  -3,
	5:  -3,
	6:  -2,
	7:  -2,
	8:  -1,
	9:  -1,
	10: 0,
	11: 0,
	12: 1,
	13: 1,
	14: 2,
	15: 2,
	16: 3,
	17: 3,
	18: 4,
	19: 4,
	20: 5,
	21: 5,
	22: 6,
	23: 6,
	24: 7,
	25: 7,
	26: 8,
	27: 8,
	28: 9,
	29: 9,
	30: 10,
}

// Ability indicates the type of an ability.
type Ability string

const (
	// Strength means the StrengthModifier has to used.
	Strength = "strength"
	// Dexterity means the DexterityModifier has to be used.
	Dexterity = "dexterity"
	// Constitution means the ConstitutionModifier has to be used.
	Constitution = "constitution"
	// Intelligence means the IntelligenceModifier has to be used.
	Intelligence = "intelligence"
	// Wisdom means the WisdomModifier has to be used.
	Wisdom = "wisdom"
	// Charisma means the CharismaModifier has to be used.
	Charisma = "charisma"
)

// outOfRange checks whether the provided value is withing the acceptable range.
func outOfRange(v int) bool {
	if v >= minimumAbilityScore && v <= maximumAbilityScore {
		return false
	}

	return true
}

type outOfRangeError struct {
	value int // The value that caused the error.
}

func (e outOfRangeError) Error() string {
	return "The provided value is out of bounds."
}

// SetStrength sets the Strength value of a creature.
func (a *Abilities) SetStrength(v int) error {
	if outOfRange(v) {
		return outOfRangeError{v}
	}

	a.Strength = v
	a.StrengthModifier = abilityScoresAndModifiers[v]

	return nil
}

// SetDexterity sets the Dexterity value of a creature.
func (a *Abilities) SetDexterity(v int) error {
	if outOfRange(v) {
		return outOfRangeError{v}
	}

	a.Dexterity = v
	a.DexterityModifier = abilityScoresAndModifiers[v]

	return nil
}

// SetConstitution sets the Constitution value of a creature.
func (a *Abilities) SetConstitution(v int) error {
	if outOfRange(v) {
		return outOfRangeError{v}
	}

	a.Constitution = v
	a.ConstitutionModifier = abilityScoresAndModifiers[v]

	return nil
}

// SetIntelligence sets the Intelligence value of a creature.
func (a *Abilities) SetIntelligence(v int) error {
	if outOfRange(v) {
		return outOfRangeError{v}
	}

	a.Intelligence = v
	a.IntelligenceModifier = abilityScoresAndModifiers[v]

	return nil
}

// SetWisdom sets the Wisdom value of a creature.
func (a *Abilities) SetWisdom(v int) error {
	if outOfRange(v) {
		return outOfRangeError{v}
	}

	a.Wisdom = v
	a.WisdomModifier = abilityScoresAndModifiers[v]

	return nil
}

// SetCharisma sets the Charisma value of a creature.
func (a *Abilities) SetCharisma(v int) error {
	if outOfRange(v) {
		return outOfRangeError{v}
	}

	a.Charisma = v
	a.CharismaModifier = abilityScoresAndModifiers[v]

	return nil
}
