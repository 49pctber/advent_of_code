package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"

	_ "crypto/MD5"
)

func Nice(s string) bool {

	// look for vowels
	re := regexp.MustCompile(`([aeiou])`)
	r := re.FindAllStringSubmatch(s, -1)
	if len(r) < 3 {
		return false
	}

	// look for double character
	double_char := false
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			double_char = true
			break
		}
	}
	if !double_char {
		return false
	}

	// look for forbidden strings
	re = regexp.MustCompile(`ab|cd|pq|xy`)
	if re.FindStringSubmatch(s) != nil {
		return false
	}

	return true
}

func Nice2(s string) bool {

	pair := false
	idxs := make(map[string]int)

	for i := 0; i < len(s)-1; i++ {
		_, ok := idxs[s[i:i+2]]
		if ok {
			if i-2 >= idxs[s[i:i+2]] {
				pair = true
			}
		} else {
			idxs[s[i:i+2]] = i
		}
	}
	if !pair {
		return false
	}

	double_char := false
	for i := 2; i < len(s); i++ {
		if s[i] == s[i-2] {
			double_char = true
			break
		}
	}
	return double_char
}

func main() {
	fmt.Println("Day 5")

	file, err := os.Open(`input\input5.txt`)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	count := 0
	count2 := 0
	for scanner.Scan() {
		if Nice(scanner.Text()) {
			count++
		}
		if Nice2(scanner.Text()) {
			count2++
		}
	}
	fmt.Println(count)
	fmt.Println(count2)
}
