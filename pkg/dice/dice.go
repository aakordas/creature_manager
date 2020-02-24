package dice

import "math/rand"

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
