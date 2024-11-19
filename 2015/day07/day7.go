package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type wire_t uint16

var vars map[string]wire_t
var instructions map[string]string

func compute(v string) wire_t {
	n, ok := vars[v]
	if ok {
		return n
	}

	nn, err := strconv.Atoi(v)
	if err == nil {
		return wire_t(nn)
	}

	if regexp.MustCompile(`^(\w+)$`).MatchString(instructions[v]) {
		r := compute(instructions[v])
		vars[v] = r
		return r
	}

	var m []string

	m = regexp.MustCompile(`NOT (.*)`).FindStringSubmatch(instructions[v])
	if m != nil {
		r := ^compute(m[1])
		vars[v] = r
		return r
	}

	m = regexp.MustCompile(`(.*) AND (.*)`).FindStringSubmatch(instructions[v])
	if m != nil {
		r := compute(m[1]) & compute(m[2])
		vars[v] = r
		return r
	}

	m = regexp.MustCompile(`(.*) OR (.*)`).FindStringSubmatch(instructions[v])
	if m != nil {
		r := compute(m[1]) | compute(m[2])
		vars[v] = r
		return r
	}

	m = regexp.MustCompile(`(.*) RSHIFT (.*)`).FindStringSubmatch(instructions[v])
	if m != nil {
		r := compute(m[1]) >> compute(m[2])
		vars[v] = r
		return r
	}

	m = regexp.MustCompile(`(.*) LSHIFT (.*)`).FindStringSubmatch(instructions[v])
	if m != nil {
		r := compute(m[1]) << compute(m[2])
		vars[v] = r
		return r
	}

	log.Fatalf("You shouldn't get here. (%s)", v)
	return 0
}

func load() {

	vars = make(map[string]wire_t)
	instructions = make(map[string]string)

	file, err := os.Open(`input.txt`)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	re := regexp.MustCompile(`^(.*) -> (.*)`)

	// load instructions
	for scanner.Scan() {
		r := re.FindStringSubmatch(scanner.Text())
		instructions[r[2]] = r[1]
		n, err := strconv.Atoi(r[1])
		if err == nil {
			vars[r[2]] = wire_t(n)
		}
	}
}

func main() {
	fmt.Println("Day 7")

	// Part 1
	load()
	aval := compute("a")
	fmt.Printf("aval: %v\n", aval) // 16076

	// fmt.Println(vars)
	load()
	vars["b"] = aval
	newaval := compute("a")
	fmt.Printf("newaval: %v\n", newaval) //

}
