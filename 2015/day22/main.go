package main

import (
	"fmt"
	"math"
	"strconv"
)

var min_mana int

func applyEffects(player *Player, boss *Boss) {
	for name, spell := range player.Spells {
		if spell.Duration > 0 {
			switch name {
			case "Poison":
				player.Log("Poison deals " + strconv.Itoa(spell.DamageBonus) + " damage.")
				boss.Hp -= spell.DamageBonus
			case "Recharge":
				player.Log("Recharge provides " + strconv.Itoa(spell.ManaBonus) + " mana.")
				player.Mana += spell.ManaBonus
			}
		}

		s := player.Spells[name]
		s.Duration--

		if s.Duration <= 0 {
			switch name {
			case "Shield":
				player.Armor -= spell.ArmorBonus
				player.Log("Shield wore off.")
			case "Poison":
				player.Log("Poison wore off.")
			case "Recharge":
				player.Log("Recharge wore off.")
			}
			delete(player.Spells, name)
		} else {
			player.Log(name + " timer at " + strconv.Itoa(s.Duration))
		}
	}
}

func battle(player_state Player, boss Boss) (bool, int) {

	if player_state.Cost > min_mana {
		return false, math.MaxInt
	}

	player_win := false
	min_cost := math.MaxInt

	// beginning of player turn
	player_state.LogGameState(boss)

	// apply effects
	applyEffects(&player_state, &boss)

	// check for win
	if boss.Hp <= 0 {
		player_state.Log("Player wins with " + strconv.Itoa(player_state.Cost) + " mana")
		min_mana = min(min_mana, player_state.Cost)
		return true, player_state.Cost
	}

	// player casts spell
	for name, spell := range spells {
		player := player_state.Duplicate()
		boss := boss

		if _, active := player.Spells[name]; active {
			player.Log("Player cannot cast " + name + " [Already Active]")

			continue
		}

		if player.Mana < spell.Cost {
			player.Log("Player cannot cast " + name + " [Not Enough Mana]")
			continue
		}

		// add to active spells
		player.Spells[name] = &Spell{
			Cost:        spell.Cost,
			DamageBonus: spell.DamageBonus,
			ArmorBonus:  spell.ArmorBonus,
			ManaBonus:   spell.ManaBonus,
			HpBonus:     spell.HpBonus,
			Duration:    spell.Duration,
		}
		player.Log("Player casts " + name)
		player.Cost += spell.Cost
		player.Mana -= spell.Cost

		// check for immediate spells
		switch name {
		case "Magic Missile":
			player.Log("Magic missile does 4 damage.")
			boss.Hp -= spell.DamageBonus
		case "Drain":
			player.Log("Drain deals 2 damage and heals 2 hit points.")
			boss.Hp -= spell.DamageBonus
			player.Hp += spell.HpBonus
		case "Shield":
			player.Armor += spell.ArmorBonus
		}

		// check for win
		if boss.Hp <= 0 {
			player_win = true
			min_cost = min(min_cost, player.Cost)
			player.Log("Player wins with " + strconv.Itoa(player.Cost) + " mana")
			continue
		}

		// beginning of boss turn
		player.LogGameState(boss)

		// apply effects
		applyEffects(&player, &boss)

		// check for win
		if boss.Hp <= 0 {
			player_win = true
			min_cost = min(min_cost, player.Cost)
			player.Log("Player wins with " + strconv.Itoa(player.Cost) + " mana")
			continue
		}

		// boss attacks
		dmg := boss.Dmg - player.Armor
		if dmg < 1 {
			dmg = 1
		}
		player.Log("Boss does " + strconv.Itoa(dmg) + " damage")
		player.Hp -= dmg

		// check for loss
		if player.Hp <= 0 {
			player.Log("Player loses")
			continue
		}

		win, cost := battle(player, boss)
		if win {
			player_win = true
			min_cost = min(min_cost, cost)
		}

	}
	min_mana = min(min_mana, min_cost)
	return player_win, min_cost
}

func battleHard(player_state Player, boss Boss) (bool, int) {

	if player_state.Cost > min_mana {
		return false, math.MaxInt
	}

	player_win := false
	min_cost := math.MaxInt

	// beginning of player turn
	player_state.LogGameState(boss)

	// make it hard
	player_state.Hp -= 1
	if player_state.Hp <= 0 {
		player_state.Log("Player loses.")
		return false, math.MaxInt
	}

	// apply effects
	applyEffects(&player_state, &boss)

	// check for win
	if boss.Hp <= 0 {
		player_state.Log("Player wins with " + strconv.Itoa(player_state.Cost) + " mana")
		min_mana = min(min_mana, player_state.Cost)
		return true, player_state.Cost
	}

	// player casts spell
	for name, spell := range spells {
		player := player_state.Duplicate()
		boss := boss

		if _, active := player.Spells[name]; active {
			player.Log("Player cannot cast " + name + " [Already Active]")

			continue
		}

		if player.Mana < spell.Cost {
			player.Log("Player cannot cast " + name + " [Not Enough Mana]")
			continue
		}

		// add to active spells
		player.Spells[name] = &Spell{
			Cost:        spell.Cost,
			DamageBonus: spell.DamageBonus,
			ArmorBonus:  spell.ArmorBonus,
			ManaBonus:   spell.ManaBonus,
			HpBonus:     spell.HpBonus,
			Duration:    spell.Duration,
		}
		player.Log("Player casts " + name)
		player.Cost += spell.Cost
		player.Mana -= spell.Cost

		// check for immediate spells
		switch name {
		case "Magic Missile":
			player.Log("Magic missile does 4 damage.")
			boss.Hp -= spell.DamageBonus
		case "Drain":
			player.Log("Drain deals 2 damage and heals 2 hit points.")
			boss.Hp -= spell.DamageBonus
			player.Hp += spell.HpBonus
		case "Shield":
			player.Armor += spell.ArmorBonus
		}

		// check for win
		if boss.Hp <= 0 {
			player_win = true
			min_cost = min(min_cost, player.Cost)
			player.Log("Player wins with " + strconv.Itoa(player.Cost) + " mana")
			continue
		}

		// beginning of boss turn
		player.LogGameState(boss)

		// apply effects
		applyEffects(&player, &boss)

		// check for win
		if boss.Hp <= 0 {
			player_win = true
			min_cost = min(min_cost, player.Cost)
			player.Log("Player wins with " + strconv.Itoa(player.Cost) + " mana")
			continue
		}

		// boss attacks
		dmg := boss.Dmg - player.Armor
		if dmg < 1 {
			dmg = 1
		}
		player.Log("Boss does " + strconv.Itoa(dmg) + " damage")
		player.Hp -= dmg

		// check for loss
		if player.Hp <= 0 {
			player.Log("Player loses")
			continue
		}

		win, cost := battleHard(player, boss)
		if win {
			player_win = true
			min_cost = min(min_cost, cost)
		}

	}
	min_mana = min(min_mana, min_cost)
	return player_win, min_cost
}

func main() {
	fmt.Println("Day 22")

	// player := Player{
	// 	Hp:     10,
	// 	Armor:  0,
	// 	Mana:   250,
	// 	Cost:   0,
	// 	Spells: make(map[string]*Spell),
	// }

	// boss := Boss{
	// 	Hp:  14,
	// 	Dmg: 8,
	// }

	min_mana = math.MaxInt
	player := getPlayer()
	boss := getBoss()
	_, part1 := battle(player, boss)
	fmt.Println("Part 1:", part1)

	min_mana = math.MaxInt
	player2 := getPlayer()
	boss2 := getBoss()
	_, part2 := battleHard(player2, boss2)
	fmt.Println("Part 2:", part2)

}
