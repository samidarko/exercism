package dndcharacter

import (
	"math"
	"math/rand"
	"time"
)

type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
	value := float64(score-10) / 2
	if value < 0 {
		return int(math.Round(value))
	}
	return int(math.Floor(value*100) / 100)
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
	ability := 0
	minValue := 7
	for i := 0; i < 4; i++ {
		value := rollSixSideDice()
		ability += value
		if value < minValue {
			minValue = value
		}
	}

	return ability - minValue
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	character := Character{
		Strength:     Ability(),
		Dexterity:    Ability(),
		Constitution: Ability(),
		Intelligence: Ability(),
		Wisdom:       Ability(),
		Charisma:     Ability(),
	}

	hitPoints := 10 + Modifier(character.Constitution)
	character.Hitpoints = hitPoints

	return character
}

func rollSixSideDice() int {
	rand.NewSource(time.Now().UnixNano())
	return rand.Intn(6) + 1
}
