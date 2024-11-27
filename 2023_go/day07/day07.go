package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

const (
	_        int = iota
	HighCard     = iota
	OnePair
	TwoPair
	ThreeKind
	FullHouse
	FourKind
	FiveKind
)

const (
	_ int = iota
	CardJoker
	CardTwo
	CardThree
	CardFour
	CardFive
	CardSix
	CardSeven
	CardEight
	CardNine
	CardTen
	CardJack
	CardQueen
	CardKing
	CardAce
)

var CardLut map[byte]int

func init() {
	CardLut = make(map[byte]int)
	CardLut['*'] = CardJoker
	CardLut['2'] = CardTwo
	CardLut['3'] = CardThree
	CardLut['4'] = CardFour
	CardLut['5'] = CardFive
	CardLut['6'] = CardSix
	CardLut['7'] = CardSeven
	CardLut['8'] = CardEight
	CardLut['9'] = CardNine
	CardLut['T'] = CardTen
	CardLut['J'] = CardJack
	CardLut['Q'] = CardQueen
	CardLut['K'] = CardKing
	CardLut['A'] = CardAce
}

type Hand struct {
	cards     string
	bet       int
	hand_type int
}

type Hands struct {
	hands []Hand
}

func (h Hands) String() string {
	r := ""
	for _, h := range h.hands {
		r += h.String() + "\n"
	}
	return r
}

func (h Hands) Len() int {
	return len(h.hands)
}

func (h Hands) Swap(i, j int) {
	h.hands[i], h.hands[j] = h.hands[j], h.hands[i]
}

func (h Hands) Less(i, j int) bool {
	if h.hands[i].hand_type == 0 {
		h.hands[i].Classify()
	}

	if h.hands[j].hand_type == 0 {
		h.hands[j].Classify()
	}

	if h.hands[i].hand_type < h.hands[j].hand_type {
		return true
	} else if h.hands[i].hand_type > h.hands[j].hand_type {
		return false
	} else {
		for k := 0; k < len(h.hands[i].cards); k++ {
			if CardLut[h.hands[i].cards[k]] < CardLut[h.hands[j].cards[k]] {
				return true
			} else if CardLut[h.hands[i].cards[k]] > CardLut[h.hands[j].cards[k]] {
				return false
			}
		}
	}

	return false // they're the same
}

func (h *Hand) Classify() int {
	counts := make(map[byte]int)
	jokers := 0

	for i := 0; i < len(h.cards); i++ {
		if h.cards[i] == '*' {
			jokers++
		} else {
			counts[h.cards[i]] += 1
		}
	}

	var x []int
	for _, j := range counts {
		x = append(x, j)
	}

	if len(x) == 0 {
		// covers case of all jokers
		x = append(x, 0)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(x)))

	h.hand_type = UpgradeHand(x, jokers)

	return h.hand_type
}

func ClassifyHand(x []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(x)))

	switch {
	case x[0] == 5:
		return FiveKind
	case x[0] == 4:
		return FourKind
	case x[0] == 3 && x[1] == 2:
		return FullHouse
	case x[0] == 3 && x[1] == 1:
		return ThreeKind
	case x[0] == 2 && x[1] == 2:
		return TwoPair
	case x[0] == 2 && x[1] == 1:
		return OnePair
	default:
		return HighCard
	}
}

func UpgradeHand(x []int, n_jokers int) int {
	if n_jokers == 0 {
		return ClassifyHand(x)
	}

	y := make([]int, len(x))
	copy(y, x)
	y[0] += 1
	return UpgradeHand(y, n_jokers-1)
}

func (h Hand) String() string {
	return fmt.Sprintf("%v (Bet: %d) (Hand: %v)", h.cards, h.bet, h.Classify())
}

func (h *Hands) Score() int {
	sort.Sort(h)
	score := 0
	for i, hand := range h.hands {
		score += (i + 1) * hand.bet
	}
	return score
}

func (h *Hands) ConvertJacks() {
	for i := range h.hands {
		h.hands[i].cards = strings.ReplaceAll(h.hands[i].cards, "J", "*")
		h.hands[i].Classify()
	}
}

func main() {
	fmt.Println("Day7")

	file, err := os.Open(filepath.Join("input", "input7.txt"))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var hands Hands
	re := regexp.MustCompile(`(.{5}) (\d+)`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var hand Hand
		r := re.FindStringSubmatch(scanner.Text())
		hand.cards = r[1]
		hand.bet, err = strconv.Atoi(r[2])
		if err != nil {
			log.Fatal(err)
		}
		hand.Classify()
		hands.hands = append(hands.hands, hand)
	}
	sort.Sort(hands)
	fmt.Printf("Part 1: %d\n", hands.Score()) // 249390788

	hands.ConvertJacks()
	sort.Sort(hands)
	fmt.Printf("Part 2: %d\n", hands.Score()) // 248750248

}
