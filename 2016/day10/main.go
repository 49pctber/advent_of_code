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

var low_chip_value, high_chip_value int

func SetSearchParamters(lc, hc int) {
	low_chip_value = lc
	high_chip_value = hc
}

func (bot Bot) String() string {
	return fmt.Sprintf("Bot %d has chips %v and will pass low to %d and high to %d", bot.id, bot.chips, bot.passlow, bot.passhigh)
}

func GiveChip(bots []*Bot, to BotId_t, chip MicrochipId_t) {
	bot := bots[to]
	bot.chips = append(bot.chips, chip)

	// check if has two chips
	if len(bot.chips) == 2 {
		lc := slices.Min(bot.chips)
		hc := slices.Max(bot.chips)
		bot.chips = make([]MicrochipId_t, 0)

		if lc == -1 || hc == -1 {
			fmt.Println(bot)
			panic("uninitialized values")
		}

		if lc == MicrochipId_t(low_chip_value) && hc == MicrochipId_t(high_chip_value) {
			fmt.Printf("Bot %d\n", bot.id)
		}

		GiveChip(bots, bot.passlow, lc)
		GiveChip(bots, bot.passhigh, hc)
	}
}

func GetBot(input string) {

	var output_offset OutputId_t = 210
	nbots := 256
	bots := make([]*Bot, nbots)
	for i := 0; i < nbots; i++ {
		bots[i] = &Bot{id: BotId_t(i), chips: make([]MicrochipId_t, 0), passlow: -1, passhigh: -1}
	}

	instructions := strings.Split(input, "\n")
	for i := range instructions {
		instructions[i] = strings.TrimSpace(instructions[i])
	}

	var val MicrochipId_t
	var bot, bot1, bot2 BotId_t
	var output1, output2 OutputId_t
	var v1, v2 string

	// load all passing rules first
	for _, i := range instructions {
		if n, err := fmt.Sscanf(i, "bot %d gives %s to bot %d and %s to bot %d", &bot, &v1, &bot1, &v2, &bot2); n == 5 && err == nil {
			b := bots[bot]
			if v1 == "high" {
				b.passhigh = bot1
				b.passlow = bot2
			} else {
				b.passhigh = bot2
				b.passlow = bot1
			}
		} else if n, err := fmt.Sscanf(i, "bot %d gives %s to output %d and %s to bot %d", &bot, &v1, &output1, &v2, &bot2); n == 5 && err == nil {
			b := bots[bot]
			if v2 == "high" {
				b.passhigh = bot2
				b.passlow = BotId_t(output1 + output_offset)
			} else {
				b.passlow = bot2
				b.passhigh = BotId_t(output1 + output_offset)
			}
		} else if n, err := fmt.Sscanf(i, "bot %d gives %s to output %d and %s to output %d", &bot, &v1, &output1, &v2, &output2); n == 5 && err == nil {
			b := bots[bot]
			if v1 == "high" {
				b.passhigh = BotId_t(output1 + output_offset)
				b.passlow = BotId_t(output2 + output_offset)
			} else {
				b.passlow = BotId_t(output1 + output_offset)
				b.passhigh = BotId_t(output2 + output_offset)
			}
		}
	}

	// give chips
	for _, i := range instructions {
		if n, err := fmt.Sscanf(i, "value %d goes to bot %d", &val, &bot); n == 2 && err == nil {
			GiveChip(bots, bot, val)
		}
	}

	product := bots[output_offset+0].chips[0] * bots[output_offset+1].chips[0] * bots[output_offset+2].chips[0]

	fmt.Printf("Part 2: %d\n", product)

}

func InputString(fname string) string {
	b, err := os.ReadFile(fname)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func main() {
	SetSearchParamters(17, 61)
	fmt.Printf("Part 1: ")
	GetBot(InputString(`input.txt`))
}
