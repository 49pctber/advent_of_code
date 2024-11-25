package main

type Item struct {
	Cost   int
	Damage int
	Armor  int
}

var weapons []Item
var armors []Item
var rings []Item

func init() {
	weapons = []Item{
		{Cost: 8, Damage: 4, Armor: 0},
		{Cost: 10, Damage: 5, Armor: 0},
		{Cost: 25, Damage: 6, Armor: 0},
		{Cost: 40, Damage: 7, Armor: 0},
		{Cost: 74, Damage: 8, Armor: 0},
	}

	armors = []Item{
		{Cost: 0, Damage: 0, Armor: 0},
		{Cost: 13, Damage: 0, Armor: 1},
		{Cost: 31, Damage: 0, Armor: 2},
		{Cost: 53, Damage: 0, Armor: 3},
		{Cost: 75, Damage: 0, Armor: 4},
		{Cost: 102, Damage: 0, Armor: 5},
	}

	rings = []Item{
		{Cost: 0, Damage: 0, Armor: 0},
		{Cost: 0, Damage: 0, Armor: 0},
		{Cost: 25, Damage: 1, Armor: 0},
		{Cost: 50, Damage: 2, Armor: 0},
		{Cost: 100, Damage: 3, Armor: 0},
		{Cost: 20, Damage: 0, Armor: 1},
		{Cost: 40, Damage: 0, Armor: 2},
		{Cost: 80, Damage: 0, Armor: 3},
	}
}
