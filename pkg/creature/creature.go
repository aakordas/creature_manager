package creature

import (
	"github.com/aakordas/creature_manager/pkg/abilities"
	"github.com/aakordas/creature_manager/pkg/saves"
	"github.com/aakordas/creature_manager/pkg/skills"
)

// Creature models a creature of the game, let that be a player, a monster or
// some non playable character.
type Creature struct {
	Name string `json:"name" bson:"name"`

	CurrentHitPoints int `json:"hit_points" bson:"hit_points"`
	Level            int `json:"level" bson:"level"`

	abilities.Abilities `json:"abilities" bson:"abilities"`
	skills.Skills       `json:"skills,omitempty" bson:"skills,omitempty"`
	saves.SavingThrows  `json:"saving_throws,omitempty" bson:"saving_throws,omitempty"`

	ProficiencyBonus int `json:"proficiency_bonus" bson:"proficiency_bonus"`
	ArmorClass       int `json:"armor_class" bson:"armor_class"`

	PassivePerception int `json:"passive_perception" bson:"passive_perception"` // FIXME: Include any other bonuses.
}

// ProficiencyBonusPerLevel maps a character level to the proficiency bonus.
var ProficiencyBonusPerLevel = map[int]int{
	1:  2,
	2:  2,
	3:  2,
	4:  2,
	5:  3,
	6:  3,
	7:  3,
	8:  3,
	9:  4,
	10: 4,
	11: 4,
	12: 4,
	13: 5,
	14: 5,
	15: 5,
	16: 5,
	17: 6,
	18: 6,
	19: 6,
	20: 6,
}
