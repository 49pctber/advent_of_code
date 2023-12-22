package main

import (
	"fmt"
	"strings"
)

func LookAndSay(input string) (output string) {

	count := 1
	b := input[0]

	var v []string

	for i := 1; i < len(input); i++ {

		if input[i] == input[i-1] {
			count++
		} else {
			// output = fmt.Sprintf("%s%d%c", output, count, b)
			v = append(v, fmt.Sprintf("%d%c", count, b))
			count = 1
			b = input[i]
		}
	}

	// return fmt.Sprintf("%s%d%c", output, count, b)
	v = append(v, fmt.Sprintf("%d%c", count, b))
	return strings.Join(v, "")
}

func main() {
	fmt.Println("Day10")

	input := "1321131112"

	for n := 0; n < 50; n++ {
		input = LookAndSay(input)

		if n == 39 || n == 49 {
			fmt.Println(len(input)) // 492982, 6989950
		}
	}

}
