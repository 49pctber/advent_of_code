package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"

	_ "embed"
)

func DecompressedLength(s string) int {
	l := 0
	r := regexp.MustCompile(`\((\d+)x(\d+)\)`)

	for {
		loc := r.FindStringSubmatch(s)
		if loc == nil {
			return l + len(s)
		} else {

			a, err := strconv.Atoi(loc[1])
			if err != nil {
				panic(err)
			}

			b, err := strconv.Atoi(loc[2])
			if err != nil {
				panic(err)
			}

			s = s[len(loc[0])+a:]

			l += a * b
		}
	}
}

func main() {
	b, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %v\n", DecompressedLength(string(b)))
	fmt.Printf("Part 1: %v\n", DecompressedLength2(string(b)))
}
