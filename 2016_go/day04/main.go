package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	_ "embed"
)

type RoomId struct {
	Name     string
	SectorID int
	Checksum string
}

func (rid RoomId) ComputeChecksum() string {
	counts := make(map[rune]int, 27)
	for _, r := range rid.Name {
		counts[r]++
	}
	counts['-'] = -1

	var checksum string

	for i := 0; i < 5; i++ {
		var lr rune
		var hc int = 0

		for r, c := range counts {
			if c > hc {
				lr = r
				hc = c
			} else if c == hc && r < lr {
				lr = r
			}
		}

		checksum = fmt.Sprintf("%s%c", checksum, lr)
		counts[lr] = -1
	}
	return checksum
}

func (rid RoomId) IsValid() bool {
	return rid.ComputeChecksum() == rid.Checksum
}

func (rid RoomId) GetSectorId() int {
	return rid.SectorID
}

func (rid RoomId) GetRealName() string {
	var realname string

	for _, r := range rid.Name {
		var nr rune

		if r == '-' {
			nr = ' '
		} else {
			nr = (r-'a'+rune(rid.SectorID))%26 + 'a'
		}

		realname = fmt.Sprintf("%s%c", realname, nr)
	}

	return realname
}

var re *regexp.Regexp = regexp.MustCompile(`(.*)-(\d+)\[(\w{5})\]`)

func (rid *RoomId) Parse(input string) error {

	d := re.FindStringSubmatch(input)
	if len(d) < 4 {
		return fmt.Errorf("invalid string")
	}

	rid.Name = d[1]
	sid, err := strconv.Atoi(d[2])
	if err != nil {
		panic(err)
	}
	rid.SectorID = sid
	rid.Checksum = d[3]
	return nil
}

//go:embed input
var input string

func Part1() {
	s := strings.Split(input, "\n")
	sum := 0
	rid := RoomId{}

	for _, i := range s {

		err := rid.Parse(i)

		if err != nil {
			continue
		}

		if rid.IsValid() {
			sum += rid.GetSectorId()
			fmt.Println(rid.GetRealName(), rid.GetSectorId())
		}

	}

	fmt.Printf("Part 1: %v\n", sum)
}

func Part2() {
	fmt.Printf("Part 2: %v\n", "x")
}

func main() {
	Part1()
	Part2()
}
