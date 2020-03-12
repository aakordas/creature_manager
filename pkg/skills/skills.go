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
