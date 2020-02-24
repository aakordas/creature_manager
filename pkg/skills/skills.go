package skills

import "github.com/aakordas/creature_manager/pkg/abilities"

// skill holds the value of a skill, whether the creature is proficient on it or
// not, the respective modifier that will be used if proficient is true and a
// description.
type skill struct {
	Value       int               `json:"value"`
	Modifier    abilities.Ability `json:"modifier"` // The ability of which the modifier will be used, if the creature is proficient.
	Description string            `json:"description"`
}

// Skill indicates a skill in an enum.
type Skill string

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

// Skills is the collection of skills the creature is proficient in.
type Skills map[Skill]*skill

func (s *skill) addSkill(v int, mod string, d string) {
	s.Value = v
	s.Modifier = abilities.Ability(mod)
	s.Description = d
}

// AddAcrobatics adds an acrobatics proficiency in the proficiency list.
func (s Skills) AddAcrobatics(v int) {
	s[Acrobatics] = new(skill)
	s[Acrobatics].addSkill(
		v,
		abilities.Dexterity,
		"The dexterity skill covers attempts to stay on one's feet in a tricky situation.",
	)
}

// AddAnimalHandling adds an animal handling proficiency in the proficiency list.
func (s Skills) AddAnimalHandling(v int) {
	s[AnimalHandling] = new(skill)
	s[AnimalHandling].addSkill(
		v,
		abilities.Wisdom,
		"The animal handling skill measures one's ability to clam down a domesticated animal, keep a mount from getting spooked or intuit an animal's intentions.",
	)

}

// AddArcana adds an arcana proficiency in the proficiency list.
func (s Skills) AddArcana(v int) {
	s[Arcana] = new(skill)
	s[Arcana].addSkill(
		v,
		abilities.Intelligence,
		"The arcana skill measures one's ability to recall lore about spells, magic items, eldritch symbols, magical traditions, the planes of existence and the inhabitants of those planes.",
	)

}

// AddAthletics adds an athletics proficiency in the proficiency list.
func (s Skills) AddAthletics(v int) {
	s[Athletics] = new(skill)
	s[Athletics].addSkill(
		v,
		abilities.Strength,
		"The athletics skill covers situations one might encounter while climbing, jumbing or swimming.",
	)
}

// AddDeception adds a deception proficiency in the proficiency list.
func (s Skills) AddDeception(v int) {
	s[Deception] = new(skill)
	s[Deception].addSkill(
		v,
		abilities.Charisma,
		"The deception skill determines whether one can convincingly hide the truth, either verbally or through one's actions.",
	)

}

// AddHistory adds an history proficiency in the proficiency list.
func (s Skills) AddHistory(v int) {
	s[History] = new(skill)
	s[History].addSkill(
		v,
		abilities.Intelligence,
		"The history skill measures one's ability to recall lore about historical events, legendary people, ancient kingdoms, past disputes, recent wars and lost civilizations.",
	)

}

// AddInsight adds an insight proficiency in the proficiency list.
func (s Skills) AddInsight(v int) {
	s[Insight] = new(skill)
	s[Insight].addSkill(
		v,
		abilities.Wisdom,
		"The insight skill decides whether one can determine the true intentions of another creature.",
	)

}

// AddIntimidation adds an intimidation proficiency in the proficiency list.
func (s Skills) AddIntimidation(v int) {
	s[Intimidation] = new(skill)
	s[Intimidation].addSkill(
		v,
		abilities.Charisma,
		"The intimidation skill covers one's attempt to influence someone through overt threats, hostile actions and physical violence.",
	)

}

// AddInvestigation adds an Investigation proficiency in the proficiency list.
func (s Skills) AddInvestigation(v int) {
	s[Investigation] = new(skill)
	s[Investigation].addSkill(
		v,
		abilities.Intelligence,
		"The investigation skill covers one's deductive abilities.",
	)

}

// AddMedicine adds a medicine proficiency in the proficiency list.
func (s Skills) AddMedicine(v int) {
	s[Medicine] = new(skill)
	s[Medicine].addSkill(
		v,
		abilities.Wisdom,
		"The medicine skill lets one try to stabilize a dying companion or diagnose an illness.",
	)

}

// AddNature adds a nature proficiency in the proficiency list.
func (s Skills) AddNature(v int) {
	s[Nature] = new(skill)
	s[Nature].addSkill(
		v,
		abilities.Intelligence,
		"The nature skill measures one's ability to recall lore about terrain, plants and animals, the weather and nature cycles.",
	)

}

// AddPerception adds a perception proficiency in the proficiency list.
func (s Skills) AddPerception(v int) {
	s[Perception] = new(skill)
	s[Perception].addSkill(
		v,
		abilities.Wisdom,
		"The perception skill lets one spot, hear or otherwise detect the presence of something.",
	)

}

// AddPerformance adds a Performance proficiency in the proficiency list.
func (s Skills) AddPerformance(v int) {
	s[Performance] = new(skill)
	s[Performance].addSkill(
		v,
		abilities.Charisma,
		"The performance skill determines how well one can delight an audience with music, dance, acting, storytelling or some other form of entertainment.",
	)

}

// AddPersuasion adds a persuasion proficiency in the proficiency list.
func (s Skills) AddPersuasion(v int) {
	s[Persuasion] = new(skill)
	s[Persuasion].addSkill(
		v,
		abilities.Charisma,
		"The persuasion skill covers one's attempt to influence someone or a group of people with tact, social graces or good nature.",
	)

}

// AddReligion adds a religion proficiency in the proficiency list.
func (s Skills) AddReligion(v int) {
	s[Religion] = new(skill)
	s[Religion].addSkill(
		v,
		abilities.Intelligence,
		"The religion skill measures one's ability to recall lore about deities, rites and prayers, religious hierarchies, holy symbols and the practices of secret cults.",
	)

}

// AddSleightOfHand adds a sleight of hand proficiency in the proficiency list.
func (s Skills) AddSleightOfHand(v int) {
	s[SleightOfHand] = new(skill)
	s[SleightOfHand].addSkill(
		v,
		abilities.Dexterity,
		"The sleight of hand skill covers acts of legerdemain or manual trickery.",
	)

}

// AddStealth adds a stealth proficiency in the proficiency list.
func (s Skills) AddStealth(v int) {
	s[Stealth] = new(skill)
	s[Stealth].addSkill(
		v,
		abilities.Dexterity,
		"The stealth skill covers attempts to conceal yourself from enemies, slink past guards, slip away without being noticed or sneak up on someone without being seen or heard.",
	)

}

// AddSurvival adds a survival proficiency in the proficiency list.
func (s Skills) AddSurvival(v int) {
	s[Survival] = new(skill)
	s[Survival].addSkill(
		v,
		abilities.Wisdom,
		"The survival skill covers one's attempts to follow tracks, hunt wild game, guid one's group through frozen wastelands, identify signs that owlbears live nearby, predict the weather or avoid quicksand and other natural hazards.",
	)

}
