package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type Card struct {
	index           int
	winning_numbers [10]int
	our_numbers     [25]int
	n_cards         uint64
}

func (c *Card) Print() {
	fmt.Printf("Card %d: ", c.index)
	for _, n := range c.winning_numbers {
		fmt.Printf("%d ", n)
	}
	fmt.Printf("| ")
	for _, n := range c.our_numbers {
		fmt.Printf("%d ", n)
	}
	fmt.Printf("\n")
}

func (c *Card) Sort() {
	slices.Sort(c.winning_numbers[:])
	slices.Sort(c.our_numbers[:])
}

func (c *Card) NWinningNumbers() int {
	c.Sort()

	var i, j, n int = 0, 0, 0

	for i < 10 && j < 25 {
		if c.winning_numbers[i] < c.our_numbers[j] {
			i++
		} else if c.winning_numbers[i] > c.our_numbers[j] {
			j++
		} else {
			i++
			j++
			n++
		}
	}

	return n
}

func (c *Card) Score() int {

	n := c.NWinningNumbers()

	if n > 0 {
		return 1 << (n - 1)
	} else {
		return 0
	}
}

func main() {
	fmt.Println("Day 4")

	var cards []Card
	sum := 0

	re := regexp.MustCompile(`Card (\d+): ((?:\d+ *)+) \| ((?: *\d+)+)`)
	re_remove_space := regexp.MustCompile(` +`)

	file, err := os.Open(filepath.Join("input", "input4.txt"))
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := re_remove_space.ReplaceAllString(scanner.Text(), " ")
		r := re.FindStringSubmatch(line)

		var card Card
		card.n_cards = 1
		card.index, err = strconv.Atoi(r[1])
		if err != nil {
			log.Fatal(err)
		}

		for i, n := range strings.Split(r[2], " ") {
			card.winning_numbers[i], err = strconv.Atoi(n)
			if err != nil {
				log.Fatal(err)
			}
		}

		for i, n := range strings.Split(r[3], " ") {
			card.our_numbers[i], err = strconv.Atoi(n)
			if err != nil {
				log.Fatal(err)
			}
		}

		sum += card.Score()
		cards = append(cards, card)
	}

	fmt.Println("Sum:", sum) // 22897

	var n_cards uint64 = 0

	for i, card := range cards {

		n_cards += cards[i].n_cards

		high := i + card.NWinningNumbers()

		for j := i + 1; j <= high && j < 209; j++ {
			cards[j].n_cards += cards[i].n_cards
		}
	}

	fmt.Println("Number of Scratchcards:", n_cards) // 5095824
}
