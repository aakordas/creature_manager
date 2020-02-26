package creature

import (
	"github.com/aakordas/creature_manager/pkg/abilities"
	"github.com/aakordas/creature_manager/pkg/saves"
	"github.com/aakordas/creature_manager/pkg/skills"
)

// Creature models a creature of the game, let that be a player, a monster or
// some non playable character.
type Creature struct {
	abilities.Abilities `json:"abilities"`
	skills.Skills       `json:"skills"`
	saves.SavingThrows  `json:"saving_throws"`
}
