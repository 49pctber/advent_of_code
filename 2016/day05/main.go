package main

import (
	"crypto"
	"fmt"

	_ "crypto/md5"
	_ "embed"
)

func GetPassword(input string) string {
	var password string
	var i int

	for len(password) < 8 {
		h := crypto.MD5.New()
		h.Write([]byte(fmt.Sprintf("%s%d", input, i)))
		s := fmt.Sprintf("%x", h.Sum(nil))
		if s[0:5] == "00000" {
			password = fmt.Sprintf("%s%c", password, s[5])
		}
		i++
	}

	return password
}

func GetPassword2(input string) string {
	password_array := make([]byte, 8)
	i := 0
	mask := 0

	for {
		h := crypto.MD5.New()
		h.Write([]byte(fmt.Sprintf("%s%d", input, i)))
		i++
		sum := h.Sum(nil)
		idx := int(sum[2] & 0x0f)
		if sum[0] == 0 && sum[1] == 0 && sum[2]&0xf0 == 0 {
			if idx >= 0 && idx <= 7 && (mask>>idx)&0x1 == 0 {
				password_array[idx] = (sum[3] & 0xf0) >> 4
				// check if all indexes are populated
				mask |= 1 << idx
				if mask == 0xff {
					break
				}
			} else {
				continue
			}
		}

	}

	password := ""
	for i := 0; i < len(password_array); i++ {
		password = fmt.Sprintf("%s%x", password, password_array[i])
	}
	return password
}

//go:embed input
var input string

func Part1() {
	fmt.Printf("Part 1: %v\n", GetPassword(input))
}

func Part2() {
	fmt.Printf("Part 2: %v\n", GetPassword2(input))
}

func main() {
	Part1()
	Part2()
}
