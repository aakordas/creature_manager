package dice

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

// Dice models a dice. A function that returns some number.
type Dice func() int

// D4 models a 4-sided die.
func D4() int {
	return rand.Intn(4) + 1
}

// D6 models a 6-sided die.
func D6() int {
	return rand.Intn(6) + 1
}

// D8 models an 8-sided die.
func D8() int {
	return rand.Intn(8) + 1
}

// D10 models a 10-sided die.
func D10() int {
	return rand.Intn(10) + 1
}

// D12 models a 12-sided die.
func D12() int {
	return rand.Intn(12) + 1
}

// D20 models a 20-sided die.
func D20() int {
	return rand.Intn(20) + 1
}

// D100 models a 100-sided die.
func D100() int {
	return rand.Intn(100) + 1
}

// Advantage returns the result of a D20 roll with advantage.
func Advantage() int {
	roll1 := D20()
	roll2 := D20()

	if roll1 > roll2 {
		return roll1
	}

	return roll2
}

// Disadvantage returns the result of a D20 roll with disadvantage.
func Disadvantage() int {
	roll1 := D20()
	roll2 := D20()

	if roll1 < roll2 {
		return roll1
	}

	return roll2
}

// Inspiration returns the result of a D20 roll with inspiration.
func Inspiration() int {
	return Advantage()
}
