package skills

import "github.com/aakordas/creature_manager/pkg/abilities"

// skill holds the value of a skill, whether the creature is proficient on it or
// not and the respective modifier that will be used if proficient is true.
type skill struct {
	Value    int    `json:"value" bson:"value"`
	Modifier string `json:"modifier" bson:"modifier"` // The ability of which the modifier will be used, if the creature is proficient.
}

const (
	// Acrobatics means the Acrobatics skill will be added in the proficiency list.
	Acrobatics = "acrobatics"
	// AnimalHandling means the AnimalHandling skill will be added in the proficiency list.
	AnimalHandling = "animal_handling"
	// Arcana means the Arcana skill will be added in the proficiency list.
	Arcana = "arcana"
	// Athletics means the Athletics skill will be added in the proficiency list.
	Athletics = "athletics"
	// Deception means the Deception skill will be added in the proficiency list.
	Deception = "deception"
	// History means the History skill will be added in the proficiency list.
	History = "history"
	// Insight means the Insight skill will be added in the proficiency list.
	Insight = "insight"
	// Intimidation means the Intimidation skill will be added in the proficiency list.
	Intimidation = "intimidation"
	// Investigation means the Investigation skill will be added in the proficiency list.
	Investigation = "investigation"
	// Medicine means the Medicine skill will be added in the proficiency list.
	Medicine = "medicine"
	// Nature means the Nature skill will be added in the proficiency list.
	Nature = "nature"
	// Perception means the Perception skill will be added in the proficiency list.
	Perception = "perception"
	// Performance means the Performance skill will be added in the proficiency list.
	Performance = "performance"
	// Persuasion means the Persuasion skill will be added in the proficiency list.
	Persuasion = "persuasion"
	// Religion means the Religion skill will be added in the proficiency list.
	Religion = "religion"
	// SleightOfHand means the SleightOfHand skill will be added in the proficiency list.
	SleightOfHand = "sleight_of_hand"
	// Stealth means the Stealth skill will be added in the proficiency list.
	Stealth = "stealth"
	// Survival means the Survival skill will be added in the proficiency list.
	Survival = "survival"
)

// SkillToAbility maps a skill to the ability modifier it needs, when the
// creature is proficient in it.
var SkillToAbility = map[string]string{
	Acrobatics:     abilities.Dexterity,
	AnimalHandling: abilities.Wisdom,
	Arcana:         abilities.Intelligence,
	Athletics:      abilities.Strength,
	Deception:      abilities.Charisma,
	History:        abilities.Intelligence,
	Insight:        abilities.Wisdom,
	Intimidation:   abilities.Charisma,
	Investigation:  abilities.Intelligence,
	Medicine:       abilities.Wisdom,
	Nature:         abilities.Intelligence,
	Perception:     abilities.Wisdom,
	Performance:    abilities.Charisma,
	Persuasion:     abilities.Charisma,
	Religion:       abilities.Intelligence,
	SleightOfHand:  abilities.Dexterity,
	Stealth:        abilities.Dexterity,
	Survival:       abilities.Wisdom,
}

// Skills is the collection of skills the creature is proficient in.
type Skills map[string]*skill

func (s *skill) addSkill(v int, mod string) {
	s.Value = v
	s.Modifier = mod
}

// AddAcrobatics adds an acrobatics proficiency in the proficiency list.
func (s Skills) AddAcrobatics(v int) {
	s[Acrobatics] = new(skill)
	s[Acrobatics].addSkill(v, abilities.Dexterity)
}

// AddAnimalHandling adds an animal handling proficiency in the proficiency list.
func (s Skills) AddAnimalHandling(v int) {
	s[AnimalHandling] = new(skill)
	s[AnimalHandling].addSkill(v, abilities.Wisdom)

}

// AddArcana adds an arcana proficiency in the proficiency list.
func (s Skills) AddArcana(v int) {
	s[Arcana] = new(skill)
	s[Arcana].addSkill(v, abilities.Intelligence)

}

// AddAthletics adds an athletics proficiency in the proficiency list.
func (s Skills) AddAthletics(v int) {
	s[Athletics] = new(skill)
	s[Athletics].addSkill(v, abilities.Strength)
}

// AddDeception adds a deception proficiency in the proficiency list.
func (s Skills) AddDeception(v int) {
	s[Deception] = new(skill)
	s[Deception].addSkill(v, abilities.Charisma)

}

// AddHistory adds an history proficiency in the proficiency list.
func (s Skills) AddHistory(v int) {
	s[History] = new(skill)
	s[History].addSkill(v, abilities.Intelligence)

}

// AddInsight adds an insight proficiency in the proficiency list.
func (s Skills) AddInsight(v int) {
	s[Insight] = new(skill)
	s[Insight].addSkill(v, abilities.Wisdom)

}

// AddIntimidation adds an intimidation proficiency in the proficiency list.
func (s Skills) AddIntimidation(v int) {
	s[Intimidation] = new(skill)
	s[Intimidation].addSkill(v, abilities.Charisma)

}

// AddInvestigation adds an Investigation proficiency in the proficiency list.
func (s Skills) AddInvestigation(v int) {
	s[Investigation] = new(skill)
	s[Investigation].addSkill(v, abilities.Intelligence)

}

// AddMedicine adds a medicine proficiency in the proficiency list.
func (s Skills) AddMedicine(v int) {
	s[Medicine] = new(skill)
	s[Medicine].addSkill(v, abilities.Wisdom)

}

// AddNature adds a nature proficiency in the proficiency list.
func (s Skills) AddNature(v int) {
	s[Nature] = new(skill)
	s[Nature].addSkill(v, abilities.Intelligence)

}

// AddPerception adds a perception proficiency in the proficiency list.
func (s Skills) AddPerception(v int) {
	s[Perception] = new(skill)
	s[Perception].addSkill(v, abilities.Wisdom)

}

// AddPerformance adds a Performance proficiency in the proficiency list.
func (s Skills) AddPerformance(v int) {
	s[Performance] = new(skill)
	s[Performance].addSkill(v, abilities.Charisma)

}

// AddPersuasion adds a persuasion proficiency in the proficiency list.
func (s Skills) AddPersuasion(v int) {
	s[Persuasion] = new(skill)
	s[Persuasion].addSkill(v, abilities.Charisma)

}

// AddReligion adds a religion proficiency in the proficiency list.
func (s Skills) AddReligion(v int) {
	s[Religion] = new(skill)
	s[Religion].addSkill(v, abilities.Intelligence)

}

// AddSleightOfHand adds a sleight of hand proficiency in the proficiency list.
func (s Skills) AddSleightOfHand(v int) {
	s[SleightOfHand] = new(skill)
	s[SleightOfHand].addSkill(v, abilities.Dexterity)

}

// AddStealth adds a stealth proficiency in the proficiency list.
func (s Skills) AddStealth(v int) {
	s[Stealth] = new(skill)
	s[Stealth].addSkill(v, abilities.Dexterity)

}

// AddSurvival adds a survival proficiency in the proficiency list.
func (s Skills) AddSurvival(v int) {
	s[Survival] = new(skill)
	s[Survival].addSkill(v, abilities.Wisdom)

}
