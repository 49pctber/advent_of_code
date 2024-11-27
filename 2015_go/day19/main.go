package main

import (
	"fmt"
	"math"
	"os"
	"strings"

	_ "embed"
)

var medicine_molecule string
var rules map[string][]string
var brules map[string][]string
var min int
var searched map[string]any

func init() {
	// read file
	inputbytes, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := string(inputbytes)

	out := strings.Split(input, "\n\n")
	rulestrs := strings.Split(out[0], "\n")
	medicine_molecule = out[1][:len(out[1])-1]

	// parse rules
	rules = make(map[string][]string)
	brules = make(map[string][]string) // rules, but backwards
	for _, rulestr := range rulestrs {
		out := strings.Split(rulestr, " => ")

		if _, ok := rules[out[0]]; !ok {
			rules[out[0]] = make([]string, 0)
		}
		if _, ok := brules[out[1]]; !ok {
			rules[out[1]] = make([]string, 0)
		}

		rules[out[0]] = append(rules[out[0]], out[1])
		brules[out[1]] = append(brules[out[1]], out[0])
	}

	min = math.MaxInt

	searched = make(map[string]any)
}

func search(remaining string, depth int) {
	if remaining[0] == 'e' {
		if len(remaining) == 1 && depth < min {
			min = depth
			fmt.Printf("New min found: %v\n", min)
		}
		return
	}

	if depth >= min {
		return
	}

	if _, already_searched := searched[remaining]; already_searched {
		return
	} else {
		searched[remaining] = nil
	}

	for k, v := range brules {
		for i := 0; ; i++ {
			pos := strings.Index(remaining[i:], k)
			if pos == -1 {
				break
			}
			i += pos
			for _, rep := range v {
				// fmt.Printf("(%d) %s (%s->%s) %s %s %s\n", depth, remaining, k, rep, remaining[:i], rep, remaining[i+len(k):])
				search(remaining[:i]+rep+remaining[i+len(k):], depth+1)
			}
		}
	}
}

func main() {
	fmt.Println("Day 19")

	// part 1
	molecules := make(map[string]any, 0)
	for i := range medicine_molecule {
		for _, o := range []int{1, 2} {
			if i+o >= len(medicine_molecule) {
				continue
			}
			key := medicine_molecule[i : i+o]
			if replacements, ok := rules[key]; ok {
				for _, r := range replacements {
					new_molecule := medicine_molecule[:i] + r + medicine_molecule[i+o:]
					molecules[new_molecule] = nil
				}
			}
		}
	}
	fmt.Println("Part 1:", len(molecules))

	// part 2
	fmt.Println("Currently, this program will find the molecule almost immediately or (practically) never. If you don't immediately get a result, just rerun it a few times until you get it.")
	search(medicine_molecule, 0)
	fmt.Println("Part 2:", min)
}
