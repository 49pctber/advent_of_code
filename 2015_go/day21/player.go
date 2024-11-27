package main

import (
	"os"
	"regexp"
	"strconv"
)

type Player struct {
	hp     int
	damage int
	armor  int
	cost   int
}

func (player *Player) AddItem(item Item) {
	player.cost += item.Cost
	player.damage += item.Damage
	player.armor += item.Armor
}

func playerWin(player, boss Player) bool {
	for {
		// player turn
		dmg := player.damage - boss.armor
		if dmg < 1 {
			dmg = 1
		}
		boss.hp -= dmg
		if boss.hp <= 0 {
			return true
		}

		// boss turn
		dmg = boss.damage - player.armor
		if dmg < 1 {
			dmg = 1
		}
		player.hp -= dmg
		if player.hp <= 0 {
			return false
		}
	}
}

func getBoss() Player {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	re := regexp.MustCompile(`Hit Points: (\d+)
Damage: (\d+)
Armor: (\d+)
`)
	x := re.FindSubmatch(input)

	hp, _ := strconv.Atoi(string(x[1]))
	dmg, _ := strconv.Atoi(string(x[2]))
	armor, _ := strconv.Atoi(string(x[3]))

	boss := Player{
		hp:     hp,
		damage: dmg,
		armor:  armor,
	}

	return boss
}
