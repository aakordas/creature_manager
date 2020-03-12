package saves

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
type SavingThrows map[string]bool

