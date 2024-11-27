package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Boss struct {
	Hp  int
	Dmg int
}

func getBoss() Boss {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	re := regexp.MustCompile(`Hit Points: (\d+)
Damage: (\d+)
`)
	x := re.FindSubmatch(input)

	hp, _ := strconv.Atoi(string(x[1]))
	dmg, _ := strconv.Atoi(string(x[2]))

	boss := Boss{
		Hp:  hp,
		Dmg: dmg,
	}

	return boss
}

func (boss Boss) String() string {
	return fmt.Sprintf("Boss has %d hit points", boss.Hp)
}
