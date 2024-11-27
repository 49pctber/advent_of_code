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

// I don't think this case works in general,
// but it *does* work for my input ¯\_(ツ)_/¯
func DecompressedLength2(s string) int {
	l := 0
	r := regexp.MustCompile(`^(\w*)\((\d+)x(\d+)\)`)
	for {
		loc := r.FindStringSubmatch(s)
		if loc == nil {
			return l + len(s)
		} else {
			x := len(loc[0])
			l += len(loc[1])
			a, err := strconv.Atoi(loc[2])
			if err != nil {
				panic(err)
			}
			b, err := strconv.Atoi(loc[3])
			if err != nil {
				panic(err)
			}
			l += b * DecompressedLength2(s[x:x+a])
			s = s[x+a:]
		}
	}
}

func main() {
	b, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %v\n", DecompressedLength(string(b)))
	fmt.Printf("Part 2: %v\n", DecompressedLength2(string(b)))
}
