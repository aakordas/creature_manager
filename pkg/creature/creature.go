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

	abilities.Abilities `json:"abilities" bson:"abilities"`
	skills.Skills       `json:"skills,omitempty" bson:"skills,omitempty"`
	saves.SavingThrows  `json:"saving_throws,omitempty" bson:"saving_throws,omitempty"`

	ProficiencyBonus int `json:"proficiency_bonus" bson:"proficiency_bonus"`
	ArmorClass       int `json:"armor_class" bson:"armor_class"`

	PassivePerception int `json:"passive_perception" bson:"passive_perception"` // FIXME: Include any other bonuses.
}
