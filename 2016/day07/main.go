package main

import (
	"fmt"
	"strings"

	_ "embed"
)

func ContainsABBA(s string) bool {
	if len(s) < 4 {
		return false
	}

	for i := 0; i+3 < len(s); i++ {
		if s[i] == s[i+3] && s[i+1] == s[i+2] && s[i] != s[i+1] {
			return true
		}
	}

	return false
}

func ListABA(s string) []string {
	if len(s) < 3 {
		return nil
	}

	abas := make([]string, 0)

	for i := 0; i+2 < len(s); i++ {
		if s[i] == s[i+2] && s[i] != s[i+1] {
			abas = append(abas, s[i:i+2])
		}
	}

	return abas
}

func TlsSupport(ip string) bool {
	x := strings.ReplaceAll(strings.ReplaceAll(ip, "]", "! "), "[", " ")
	strs := strings.Split(x, " ")

	ret := false
	for _, str := range strs {
		cabba := ContainsABBA(str)

		if cabba && str[len(str)-1] == '!' {
			return false
		} else if cabba {
			ret = true
		}
	}

	return ret
}

func SslSupport(ip string) bool {
	x := strings.ReplaceAll(strings.ReplaceAll(ip, "]", "! "), "[", " ")
	strs := strings.Split(x, " ")

	abas := make([]string, 0)
	babs := make([]string, 0)

	for _, str := range strs {
		if str[len(str)-1] == '!' {
			babs = append(babs, ListABA(str)...)
		} else {
			abas = append(abas, ListABA(str)...)
		}
	}

	for _, aba := range abas {
		for _, bab := range babs {
			if aba[0] == bab[1] && aba[1] == bab[0] {
				return true
			}
		}
	}
	return false
}

//go:embed input
var input string

func Part1() {
	ips := strings.Split(input, "\n")
	count := 0
	for _, ip := range ips {
		if TlsSupport(ip) {
			count++
		}
	}
	fmt.Printf("Part 1: %v\n", count)
}

func Part2() {
	ips := strings.Split(input, "\n")
	count := 0
	for _, ip := range ips {
		if SslSupport(ip) {
			count++
		}
	}
	fmt.Printf("Part 2: %v\n", count)
}

func main() {
	Part1()
	Part2()
}
