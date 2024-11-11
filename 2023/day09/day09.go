package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

func Predict(row []int) int {

	done := true
	for _, x := range row {
		if x != 0 {
			done = false
			break
		}
	}

	if done {
		return 0
	} else {
		var y []int
		for i := 1; i < len(row); i++ {
			y = append(y, row[i]-row[i-1])
		}
		return Predict(y) + row[len(row)-1]
	}

}

func Postdict(row []int) int {

	done := true
	for _, x := range row {
		if x != 0 {
			done = false
			break
		}
	}

	if done {
		return 0
	} else {
		var y []int
		for i := 1; i < len(row); i++ {
			y = append(y, row[i]-row[i-1])
		}
		// fmt.Println(Postdict, row[0])
		return row[0] - Postdict(y)
	}

}

func main() {
	fmt.Println("Day 9")

	file, err := os.Open(filepath.Join("input", "input9.txt"))
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	regexp.MustCompile(`()`)
	sum := 0
	sum2 := 0
	for scanner.Scan() {
		tokens := strings.Fields(scanner.Text())
		row := make([]int, len(tokens))
		for i, t := range tokens {
			n, err := strconv.Atoi(t)
			if err != nil {
				log.Fatal(err)
			}
			row[i] = n
		}
		sum += Predict(row)
		sum2 += Postdict(row)
	}

	fmt.Println(sum) // 1898776583
	fmt.Println(sum2)
}
