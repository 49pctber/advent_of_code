package main

import (
	"fmt"
)

type Player struct {
	Hp     int
	Mana   int
	Armor  int
	Cost   int
	Depth  int
	Spells map[string]*Spell
}

func getPlayer() Player {
	return Player{
		Hp:     50,
		Mana:   500,
		Armor:  0,
		Cost:   0,
		Depth:  0,
		Spells: make(map[string]*Spell),
	}
}

func (player Player) String() string {
	return fmt.Sprintf("Player has %d hit points, %d armor, %d mana", player.Hp, player.Armor, player.Mana)
}

func (player *Player) LogGameState(boss Boss) {
	player.Depth++
	player.Log(fmt.Sprint(player, "/", boss))
}

func (player Player) Log(msg string) {
	// spaces := strings.Repeat("-", player.Depth)
	// fmt.Println(spaces, msg)
}

func (player Player) Duplicate() Player {
	new_player := Player{
		Hp:     player.Hp,
		Mana:   player.Mana,
		Armor:  player.Armor,
		Cost:   player.Cost,
		Depth:  player.Depth,
		Spells: make(map[string]*Spell),
	}

	for k, v := range player.Spells {
		new_player.Spells[k] = &Spell{
			Cost:        v.Cost,
			DamageBonus: v.DamageBonus,
			ArmorBonus:  v.ArmorBonus,
			ManaBonus:   v.ManaBonus,
			HpBonus:     v.HpBonus,
			Duration:    v.Duration,
		}
	}

	return new_player
}
