package abilities

import "testing"

func Test_outOfRangeError_Error(t *testing.T) {
	errorMessage := "The provided value is out of bounds."

	type fields struct {
		value int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"Too low", fields{0}, errorMessage},
		{"Too high", fields{31}, errorMessage},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := outOfRangeError{
				value: tt.fields.value,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("outOfRangeError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAbilities_SetStrength(t *testing.T) {
	type fields struct {
		Strength             ability
		Dexterity            ability
		Constitution         ability
		Intelligence         ability
		Wisdom               ability
		Charisma             ability
		StrengthModifier     int
		DexterityModifier    int
		ConstitutionModifier int
		IntelligenceModifier int
		WisdomModifier       int
		CharismaModifier     int
	}
	type args struct {
		v int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"Too low", fields{}, args{0}, true},
		{"Too high", fields{}, args{31}, true},
		{"Lower bound", fields{}, args{1}, false},
		{"Higher bound", fields{}, args{30}, false},
		{"Normal value", fields{}, args{5}, false},
		{"Strength pre-set", fields{Strength: ability{value: 3}}, args{5}, false},
		{"Strength pre-set invalid", fields{Strength: ability{value: 3}}, args{0}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Abilities{
				Strength:             tt.fields.Strength,
				Dexterity:            tt.fields.Dexterity,
				Constitution:         tt.fields.Constitution,
				Intelligence:         tt.fields.Intelligence,
				Wisdom:               tt.fields.Wisdom,
				Charisma:             tt.fields.Charisma,
				StrengthModifier:     tt.fields.StrengthModifier,
				DexterityModifier:    tt.fields.DexterityModifier,
				ConstitutionModifier: tt.fields.ConstitutionModifier,
				IntelligenceModifier: tt.fields.IntelligenceModifier,
				WisdomModifier:       tt.fields.WisdomModifier,
				CharismaModifier:     tt.fields.CharismaModifier,
			}
			if err := a.SetStrength(tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("Abilities.SetStrength() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAbilities_SetDexterity(t *testing.T) {
	type fields struct {
		Strength             ability
		Dexterity            ability
		Constitution         ability
		Intelligence         ability
		Wisdom               ability
		Charisma             ability
		StrengthModifier     int
		DexterityModifier    int
		ConstitutionModifier int
		IntelligenceModifier int
		WisdomModifier       int
		CharismaModifier     int
	}
	type args struct {
		v int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"Too low", fields{}, args{0}, true},
		{"Too high", fields{}, args{31}, true},
		{"Lower bound", fields{}, args{1}, false},
		{"Higher bound", fields{}, args{30}, false},
		{"Normal value", fields{}, args{5}, false},
		{"Dexterity pre-set", fields{Dexterity: ability{value: 3}}, args{5}, false},
		{"Dexterity pre-set invalid", fields{Dexterity: ability{value: 3}}, args{0}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Abilities{
				Strength:             tt.fields.Strength,
				Dexterity:            tt.fields.Dexterity,
				Constitution:         tt.fields.Constitution,
				Intelligence:         tt.fields.Intelligence,
				Wisdom:               tt.fields.Wisdom,
				Charisma:             tt.fields.Charisma,
				StrengthModifier:     tt.fields.StrengthModifier,
				DexterityModifier:    tt.fields.DexterityModifier,
				ConstitutionModifier: tt.fields.ConstitutionModifier,
				IntelligenceModifier: tt.fields.IntelligenceModifier,
				WisdomModifier:       tt.fields.WisdomModifier,
				CharismaModifier:     tt.fields.CharismaModifier,
			}
			if err := a.SetDexterity(tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("Abilities.SetDexterity() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAbilities_SetConstitution(t *testing.T) {
	type fields struct {
		Strength             ability
		Dexterity            ability
		Constitution         ability
		Intelligence         ability
		Wisdom               ability
		Charisma             ability
		StrengthModifier     int
		DexterityModifier    int
		ConstitutionModifier int
		IntelligenceModifier int
		WisdomModifier       int
		CharismaModifier     int
	}
	type args struct {
		v int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"Too low", fields{}, args{0}, true},
		{"Too high", fields{}, args{31}, true},
		{"Lower bound", fields{}, args{1}, false},
		{"Higher bound", fields{}, args{30}, false},
		{"Normal value", fields{}, args{5}, false},
		{"Constitution pre-set", fields{Constitution: ability{value: 3}}, args{5}, false},
		{"Constitution pre-set invalid", fields{Constitution: ability{value: 3}}, args{0}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Abilities{
				Strength:             tt.fields.Strength,
				Dexterity:            tt.fields.Dexterity,
				Constitution:         tt.fields.Constitution,
				Intelligence:         tt.fields.Intelligence,
				Wisdom:               tt.fields.Wisdom,
				Charisma:             tt.fields.Charisma,
				StrengthModifier:     tt.fields.StrengthModifier,
				DexterityModifier:    tt.fields.DexterityModifier,
				ConstitutionModifier: tt.fields.ConstitutionModifier,
				IntelligenceModifier: tt.fields.IntelligenceModifier,
				WisdomModifier:       tt.fields.WisdomModifier,
				CharismaModifier:     tt.fields.CharismaModifier,
			}
			if err := a.SetConstitution(tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("Abilities.SetConstitution() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAbilities_SetIntelligence(t *testing.T) {
	type fields struct {
		Strength             ability
		Dexterity            ability
		Constitution         ability
		Intelligence         ability
		Wisdom               ability
		Charisma             ability
		StrengthModifier     int
		DexterityModifier    int
		ConstitutionModifier int
		IntelligenceModifier int
		WisdomModifier       int
		CharismaModifier     int
	}
	type args struct {
		v int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"Too low", fields{}, args{0}, true},
		{"Too high", fields{}, args{31}, true},
		{"Lower bound", fields{}, args{1}, false},
		{"Higher bound", fields{}, args{30}, false},
		{"Normal value", fields{}, args{5}, false},
		{"Intelligence pre-set", fields{Intelligence: ability{value: 3}}, args{5}, false},
		{"Intelligence pre-set invalid", fields{Intelligence: ability{value: 3}}, args{0}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Abilities{
				Strength:             tt.fields.Strength,
				Dexterity:            tt.fields.Dexterity,
				Constitution:         tt.fields.Constitution,
				Intelligence:         tt.fields.Intelligence,
				Wisdom:               tt.fields.Wisdom,
				Charisma:             tt.fields.Charisma,
				StrengthModifier:     tt.fields.StrengthModifier,
				DexterityModifier:    tt.fields.DexterityModifier,
				ConstitutionModifier: tt.fields.ConstitutionModifier,
				IntelligenceModifier: tt.fields.IntelligenceModifier,
				WisdomModifier:       tt.fields.WisdomModifier,
				CharismaModifier:     tt.fields.CharismaModifier,
			}
			if err := a.SetIntelligence(tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("Abilities.SetIntelligence() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAbilities_SetWisdom(t *testing.T) {
	type fields struct {
		Strength             ability
		Dexterity            ability
		Constitution         ability
		Intelligence         ability
		Wisdom               ability
		Charisma             ability
		StrengthModifier     int
		DexterityModifier    int
		ConstitutionModifier int
		IntelligenceModifier int
		WisdomModifier       int
		CharismaModifier     int
	}
	type args struct {
		v int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"Too low", fields{}, args{0}, true},
		{"Too high", fields{}, args{31}, true},
		{"Lower bound", fields{}, args{1}, false},
		{"Higher bound", fields{}, args{30}, false},
		{"Normal value", fields{}, args{5}, false},
		{"Wisdom pre-set", fields{Wisdom: ability{value: 3}}, args{5}, false},
		{"Wisdom pre-set invalid", fields{Wisdom: ability{value: 3}}, args{0}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Abilities{
				Strength:             tt.fields.Strength,
				Dexterity:            tt.fields.Dexterity,
				Constitution:         tt.fields.Constitution,
				Intelligence:         tt.fields.Intelligence,
				Wisdom:               tt.fields.Wisdom,
				Charisma:             tt.fields.Charisma,
				StrengthModifier:     tt.fields.StrengthModifier,
				DexterityModifier:    tt.fields.DexterityModifier,
				ConstitutionModifier: tt.fields.ConstitutionModifier,
				IntelligenceModifier: tt.fields.IntelligenceModifier,
				WisdomModifier:       tt.fields.WisdomModifier,
				CharismaModifier:     tt.fields.CharismaModifier,
			}
			if err := a.SetWisdom(tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("Abilities.SetWisdom() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAbilities_SetCharisma(t *testing.T) {
	type fields struct {
		Strength             ability
		Dexterity            ability
		Constitution         ability
		Intelligence         ability
		Wisdom               ability
		Charisma             ability
		StrengthModifier     int
		DexterityModifier    int
		ConstitutionModifier int
		IntelligenceModifier int
		WisdomModifier       int
		CharismaModifier     int
	}
	type args struct {
		v int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"Too low", fields{}, args{0}, true},
		{"Too high", fields{}, args{31}, true},
		{"Lower bound", fields{}, args{1}, false},
		{"Higher bound", fields{}, args{30}, false},
		{"Normal value", fields{}, args{5}, false},
		{"Charisma pre-set", fields{Charisma: ability{value: 3}}, args{5}, false},
		{"Charisma pre-set invalid", fields{Charisma: ability{value: 3}}, args{0}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Abilities{
				Strength:             tt.fields.Strength,
				Dexterity:            tt.fields.Dexterity,
				Constitution:         tt.fields.Constitution,
				Intelligence:         tt.fields.Intelligence,
				Wisdom:               tt.fields.Wisdom,
				Charisma:             tt.fields.Charisma,
				StrengthModifier:     tt.fields.StrengthModifier,
				DexterityModifier:    tt.fields.DexterityModifier,
				ConstitutionModifier: tt.fields.ConstitutionModifier,
				IntelligenceModifier: tt.fields.IntelligenceModifier,
				WisdomModifier:       tt.fields.WisdomModifier,
				CharismaModifier:     tt.fields.CharismaModifier,
			}
			if err := a.SetCharisma(tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("Abilities.SetCharisma() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
