package saves

// SavingThrow indicates the type of a saving throw.
type SavingThrow string

const (
	// Strength means the Strength saving throw will be used.
	Strength = "strength"
	// Dexterity means the Dexterity saving throw will be used.
	Dexterity = "dexterity"
	// Constitution means the Constitution saving throw will be used.
	Constitution = "constitution"
	// Intelligence means the Intelligence saving throw will be used.
	Intelligence = "intelligence"
	// Wisdom means the Wisdom saving throw will be used.
	Wisdom = "wisdom"
	// Charisma means the Charisma saving throw will be used.
	Charisma = "charisma"
)

// SavingThrows is the collection of saving throws the creature is proficient at.
type SavingThrows map[SavingThrow]bool

// AddStrength adds proficiency to the Strength saving throw.
func (s SavingThrows) AddStrength() {
	s[Strength] = true
}

// AddDexterity adds proficiency to the Dexterity saving throw.
func (s SavingThrows) AddDexterity() {
	s[Dexterity] = true
}

// AddConstitution adds proficiency to the Constitution saving throw.
func (s SavingThrows) AddConstitution() {
	s[Constitution] = true
}

// AddIntelligence adds proficienty to the Intelligence saving throw.
func (s SavingThrows) AddIntelligence() {
	s[Intelligence] = true
}

// AddWisdom adds proficiency to the Wisdom saving throw.
func (s SavingThrows) AddWisdom() {
	s[Wisdom] = true
}

// AddCharisma adds proficiency to the Charisma saving throw.
func (s SavingThrows) AddCharisma() {
	s[Charisma] = true
}
