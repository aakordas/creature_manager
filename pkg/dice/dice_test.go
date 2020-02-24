package dice

import (
	"flag"
	"testing"
)

var iterations = flag.Int("iters", 100, "Specify the number of iterations for each test function.")

func TestD4(t *testing.T) {
	sides := 4

	for i := 0; i <= *iterations; i++ {
		got := D4()

		if got < 1 || got > sides {
			t.Errorf("D4 rolled out of range.")
		}
	}
}

func TestD6(t *testing.T) {
	sides := 6

	for i := 0; i <= *iterations; i++ {
		got := D6()

		if got < 1 || got > sides {
			t.Errorf("D6 rolled out of range.")
		}
	}
}

func TestD8(t *testing.T) {
	sides := 8

	for i := 0; i <= *iterations; i++ {
		got := D8()

		if got < 1 || got > sides {
			t.Errorf("D8 rolled out of range.")
		}
	}
}

func TestD10(t *testing.T) {
	sides := 10

	for i := 0; i <= *iterations; i++ {
		got := D10()

		if got < 1 || got > sides {
			t.Errorf("D10 rolled out of range.")
		}
	}
}

func TestD12(t *testing.T) {
	sides := 12

	for i := 0; i <= *iterations; i++ {
		got := D12()

		if got < 1 || got > sides {
			t.Errorf("D12 rolled out of range.")
		}
	}
}

func TestD20(t *testing.T) {
	sides := 20

	for i := 0; i <= *iterations; i++ {
		got := D20()

		if got < 1 || got > sides {
			t.Errorf("D20 rolled out of range.")
		}
	}
}

func TestD100(t *testing.T) {
	sides := 100

	for i := 0; i <= *iterations; i++ {
		got := D100()

		if got < 1 || got > sides {
			t.Errorf("D100 rolled out of range.")
		}
	}
}
