// Package character/character.go
package character

import (
	cryptoRand "crypto/rand"
	"encoding/hex"
	"slices"
)

type CharacterID string

type Character struct {
	ID                     CharacterID
	Name                   string
	Class                  Class
	Species                Species
	DraconicAncestry       DraconicAncestry
	ElvenLineage           ElvenLineage
	Size                   string
	Background             Background
	PrimaryAbilities       Abilities
	AbilityScore           AbilityScores
	HasRolledAbilityScores bool
	AbilityScoreMethod     AbilityScoreSelectionMethod
	Alignment              Alignment
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

func (c *Character) HasPrimaryAbility(a Ability) bool {
	return slices.Contains(c.PrimaryAbilities, a)
}

func (c *Character) SetSpecies(s Species) {
	c.Species = s
	c.DraconicAncestry = DraconicAncestryNone
	c.ElvenLineage = ElvenLineageNone
}
