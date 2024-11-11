package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

type Node struct {
	left  string
	right string
}

func main() {
	fmt.Println("Day 8")

	var Map map[string]Node = make(map[string]Node)

	file, err := os.Open(filepath.Join("input", "input8.txt"))
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	instructions := scanner.Text()
	scanner.Scan()

	fmt.Println(instructions)

	re := regexp.MustCompile(`(\w{3}) = \((\w{3}), (\w{3})\)`)

	// var locs []string

	for scanner.Scan() {
		r := re.FindStringSubmatch(scanner.Text())
		Map[r[1]] = Node{left: r[2], right: r[3]}

		// if r[1][2] == 'A' {
		// 	locs = append(locs, r[1])
		// }
	}

	count := 0

	for loc := "AAA"; loc != "ZZZ"; count++ {
		dir := instructions[count%len(instructions)]
		if dir == 'L' {
			loc = Map[loc].left
		} else {
			loc = Map[loc].right
		}
	}

	fmt.Println("Steps:", count)

	// count = 0
	// for {
	// 	done := true
	// 	dir := instructions[count%len(instructions)]
	// 	for i := 0; i < len(locs); i++ {
	// 		if locs[i][2] != 'Z' {
	// 			done = false
	// 		} else {
	// 			fmt.Println(count, i, locs[i])
	// 		}
	// 		if dir == 'L' {
	// 			locs[i] = Map[locs[i]].left
	// 		} else {
	// 			locs[i] = Map[locs[i]].right
	// 		}
	// 	}

	// 	count++

	// 	if done {
	// 		break
	// 	}
	// }

	// fmt.Println("Steps:", count)

	fmt.Println(Factor(11309))
	fmt.Println(Factor(13939))
	fmt.Println(Factor(15517))
	fmt.Println(Factor(17621))
	fmt.Println(Factor(19199))
	fmt.Println(Factor(20777))

}

var Primes []int

func init() {
	Primes = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199}
}

func Factor(i int) []int {

	if i == 1 {
		return nil
	}

	for _, p := range Primes {
		if i%p == 0 {
			return append(Factor(i/p), p)
		}
	}

	return []int{i}
}
