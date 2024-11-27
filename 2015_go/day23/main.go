package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var instructions []string
var registers map[byte]uint

func init() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		instructions = append(instructions, scanner.Text())
	}

	registers = make(map[byte]uint)
}

func execute() {
	n := 0

instructionLoop:
	for {
		tokens := strings.Split(instructions[n], " ")
		instruction := tokens[0]

		// fmt.Println(n, instructions[n], registers)

		switch instruction {
		case "hlf":
			reg := tokens[1][0]
			registers[reg] /= 2
			n++
		case "tpl":
			reg := tokens[1][0]
			registers[reg] *= 3
			n++
		case "inc":
			reg := tokens[1][0]
			registers[reg] += 1
			n++
		case "jmp":
			offset, err := strconv.Atoi(tokens[1])
			if err != nil {
				panic(err)
			}
			n += offset
		case "jie":
			reg := tokens[1][0]
			if registers[reg]%2 == 0 {
				offset, err := strconv.Atoi(tokens[2])
				if err != nil {
					panic(err)
				}
				n += offset
			} else {
				n++
			}
		case "jio":
			reg := tokens[1][0]
			if registers[reg] == 1 {
				offset, err := strconv.Atoi(tokens[2])
				if err != nil {
					panic(err)
				}
				n += offset
			} else {
				n++
			}
		default:
			panic("invalid instruction")
		}

		if n >= len(instructions) {
			break instructionLoop
		}
	}
}

func main() {

	registers['a'] = 0
	registers['b'] = 0
	execute()
	fmt.Println("Part 1:", registers['b'])

	registers['a'] = 1
	registers['b'] = 0
	execute()
	fmt.Println("Part 2:", registers['b'])
}
