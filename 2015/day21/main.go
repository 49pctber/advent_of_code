package main

import (
	"fmt"
	"math"
)

func search(boss Player) (int, int) {

	min := math.MaxInt
	max := math.MinInt

	for _, weapon := range weapons {
		for _, armor := range armors {
			for i := 0; i < len(rings); i++ {
				for j := i + 1; j < len(rings); j++ {
					var player Player
					player.hp = 100
					player.AddItem(weapon)
					player.AddItem(armor)
					player.AddItem(rings[i])
					player.AddItem(rings[j])
					if playerWin(player, boss) {
						if min > player.cost {
							min = player.cost
						}
					} else {
						if max < player.cost {
							max = player.cost
						}
					}
				}
			}
		}
	}

	return min, max

}

func main() {
	fmt.Println("Day 21")

	boss := getBoss()
	part1, part2 := search(boss)

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
