package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
)

type box struct {
	l int
	w int
	h int
}

func (b *box) PaperNeeded() int {
	amts := []int{b.l * b.w, b.l * b.h, b.w * b.h}
	sort.Ints(amts)
	amount := 0
	for i, a := range amts {
		if i == 0 {
			amount += 3 * a
		} else {
			amount += 2 * a
		}
	}
	return amount
}

func (b *box) RibbonNeeded() int {
	amts := []int{b.l, b.h, b.w}
	sort.Ints(amts)
	return 2*amts[0] + 2*amts[1] + b.l*b.h*b.w
}

func main() {
	fmt.Println("Day 2")

	file, err := os.Open(`input\input2.txt`)
	if err != nil {
		log.Fatal(err)
	}

	re := regexp.MustCompile(`(\d+)x(\d+)x(\d+)`)

	boxes := []box{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		r := re.FindStringSubmatch(scanner.Text())
		l, err := strconv.Atoi(r[1])
		if err != nil {
			log.Fatal(err)
		}
		w, err := strconv.Atoi(r[2])
		if err != nil {
			log.Fatal(err)
		}
		h, err := strconv.Atoi(r[3])
		if err != nil {
			log.Fatal(err)
		}
		boxes = append(boxes, box{l: l, w: w, h: h})
	}

	paper := 0
	ribbon := 0
	for _, b := range boxes {
		paper += b.PaperNeeded()
		ribbon += b.RibbonNeeded()
	}

	fmt.Println("paper:", paper)   // 1588178
	fmt.Println("ribbon:", ribbon) // 3783758

}
