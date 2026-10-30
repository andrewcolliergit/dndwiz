// Package character/character.go
package character

import (
	cryptoRand "crypto/rand"
	"encoding/hex"
	"slices"

	"dndwiz/character/ability"
	"dndwiz/character/alignment"
	"dndwiz/character/background"
	"dndwiz/character/class"
	"dndwiz/character/score"
	"dndwiz/character/species"
	draconic "dndwiz/character/species/draconic_ancestry"
	elven "dndwiz/character/species/elven_lineage"
)

type CharacterID string

type Character struct {
	ID                     CharacterID
	Name                   string
	Class                  class.Class
	Species                species.Species
	DraconicAncestry       draconic.Ancestry
	ElvenLineage           elven.Lineage
	Size                   string
	Background             background.Background
	PrimaryAbilities       ability.Abilities
	AbilityScore           score.Ability
	HasRolledAbilityScores bool
	AbilityScoreMethod     score.Method
	Alignment              alignment.Alignment
}

func generateRandomID() CharacterID {
	bytes := make([]byte, 32)
	cryptoRand.Read(bytes)
	return CharacterID(hex.EncodeToString(bytes))
}

func NewCharacter() Character {
	return Character{
		ID:   generateRandomID(),
		Name: "New Adventurer",
	}
}

func (c *Character) HasPrimaryAbility(a ability.Ability) bool {
	return slices.Contains(c.PrimaryAbilities, a)
}

func (c *Character) SetSpecies(s species.Species) {
	c.Species = s
	c.DraconicAncestry = draconic.None
	c.ElvenLineage = elven.None
}
