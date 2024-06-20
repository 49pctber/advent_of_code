package main

import (
	"fmt"
	"os"
	"slices"
	"strings"

	_ "embed"
)

type BotId_t int
type MicrochipId_t int
type OutputId_t int

type Bot struct {
	chips    []MicrochipId_t
	id       BotId_t
	passlow  BotId_t
	passhigh BotId_t
}

func GiveChip(bots []*Bot, to BotId_t, chip MicrochipId_t) {
	bot := bots[to]
	bot.chips = append(bot.chips, chip)

	// check if has two chips
	if len(bot.chips) == 2 {
		lc := slices.Min(bot.chips)
		hc := slices.Max(bot.chips)
		bot.chips = make([]MicrochipId_t, 0)
		GiveChip(bots, bot.passlow, lc)
		GiveChip(bots, bot.passhigh, hc)
	}
}

func GetBot(input string) int {

	nbots := 256
	bots := make([]*Bot, nbots)
	for i := 0; i < nbots; i++ {
		bots[i] = &Bot{id: BotId_t(i), chips: make([]MicrochipId_t, 0)}
	}

	instructions := strings.Split(input, "\n")
	for i := range instructions {
		instructions[i] = strings.TrimSpace(instructions[i])
	}

	var val MicrochipId_t
	var bot, bot1, bot2 BotId_t
	var output1, output2 OutputId_t
	var v1, v2 string

	// TODO load all passing rules first, then give chips.

	for _, i := range instructions {
		i = strings.TrimSpace(i)
		if n, err := fmt.Sscanf(i, "value %d goes to bot %d", &val, &bot); n == 2 && err == nil {
			GiveChip(bots, bot, val)
		} else if n, err := fmt.Sscanf(i, "bot %d gives %s to bot %d and %s to bot %d", &bot, &v1, &bot1, &v2, &bot2); n == 5 && err == nil {
			//
		} else if n, err := fmt.Sscanf(i, "bot %d gives %s to output %d and %s to bot %d", &bot, &v1, &output1, &v2, &bot2); n == 5 && err == nil {
			//
		} else if n, err := fmt.Sscanf(i, "bot %d gives %s to output %d and %s to output %d", &bot, &v1, &output1, &v2, &output2); n == 5 && err == nil {
			//
		} else {
			panic(fmt.Sprintf("error processing instruction: %s", i))
		}
	}

	for _, bot := range bots {
		if len(bot.chips) > 0 {
			fmt.Println(*bot)
		}
	}
	return 0
}

func Part1() {
	fmt.Printf("Part 1: %v\n", "?")
}

func Part2() {
	fmt.Printf("Part 2: %v\n", "?")
}

func InputString(fname string) string {
	b, err := os.ReadFile(fname)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func main() {
	Part1()
	Part2()
}
