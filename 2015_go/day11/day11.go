package main

import (
	"fmt"
	"strings"

	_ "embed"
)

func ValidPassword(s string) bool {

	// invalid characters
	if strings.ContainsAny(s, "iol") {
		return false
	}

	// increasing straight
	straight := false
	for i := 0; i < len(s)-2; i++ {
		if s[i]+1 == s[i+1] && s[i]+2 == s[i+2] {
			straight = true
		}
	}
	if !straight {
		return false
	}

	// two different non-overlapping pairs of letters
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			for j := i + 2; j < len(s)-1; j++ {
				if s[j] == s[j+1] {
					return true
				}
			}
		}
	}

	return false
}

func incrementPassword(s string) string {

	for pos := len(s) - 1; pos >= 0; pos-- {
		var x []string

		if pos > 0 {
			x = append(x, s[:pos])
		}

		var carry bool
		b := s[pos] + 1
		if b > 'z' {
			carry = true
			x = append(x, string('a'))
		} else {
			x = append(x, string(b))
		}

		if pos < len(s)-1 {
			x = append(x, s[pos+1:])
		}

		s = strings.Join(x, "")

		if !carry {
			break
		}
	}

	return s
}

func FindNextPassword(s string) string {
	next := incrementPassword(s)
	for ; !ValidPassword(next); next = incrementPassword(next) {
	}
	return next
}

//go:embed input.txt
var input1 string

func main() {
	new_password := FindNextPassword(input1)
	fmt.Printf("new_password: %v\n", new_password)
	next_password := FindNextPassword(new_password)
	fmt.Printf("next_password: %v\n", next_password)
}
