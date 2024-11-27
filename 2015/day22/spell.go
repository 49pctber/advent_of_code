package main

type Spell struct {
	Cost        int
	DamageBonus int
	ArmorBonus  int
	ManaBonus   int
	HpBonus     int
	Duration    int
}

var spells map[string]Spell

func init() {
	spells = map[string]Spell{
		"Magic Missile": {
			Cost:        53,
			DamageBonus: 4,
			Duration:    0,
		},
		"Drain": {
			Cost:        73,
			DamageBonus: 2,
			HpBonus:     2,
			Duration:    0,
		},
		"Shield": {
			Cost:       113,
			Duration:   6,
			ArmorBonus: 7,
		},
		"Poison": {
			Cost:        173,
			Duration:    6,
			DamageBonus: 3,
		},
		"Recharge": {
			Cost:      229,
			Duration:  5,
			ManaBonus: 101,
		},
	}
}
