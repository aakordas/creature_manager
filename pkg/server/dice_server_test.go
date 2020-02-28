package server

import (
	"testing"

	"github.com/aakordas/creature_manager/pkg/dice"
)

func Test_getSides(t *testing.T) {
	type args struct {
		sides string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"No sides", args{""}, 20},
		{"Valid sides", args{"4"}, 4},
		{"Valid sides, invalid dice", args{"1"}, 1},
		{"Invalid sides", args{"a"}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSides(tt.args.sides); got != tt.want {
				t.Errorf("getSides() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getCount(t *testing.T) {
	type args struct {
		count string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"No count", args{""}, 1},
		{"Valid count", args{"2"}, 2},
		{"Invalid count", args{"a"}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCount(tt.args.count); got != tt.want {
				t.Errorf("getCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_chooseDice(t *testing.T) {
	type args struct {
		sides int
	}
	tests := []struct {
		name string
		args args
		want dice.Dice
	}{
		{"Valid sides", args{4}, dice.D4},
		{"Invalid sides", args{5}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := chooseDice(tt.args.sides); &got == &tt.want {
				t.Errorf("chooseDice() = %v, want %v", got, tt.want)
			}
		})
	}
}
