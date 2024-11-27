package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const (
	HIGH = true
	LOW  = false
)

type module_i interface {
	Process(p pulse_t) []pulse_t
	GetDestinations() []string
}

type flip_flop_t struct {
	label        string
	state        bool
	destinations []string
}

func (f *flip_flop_t) Process(p pulse_t) []pulse_t {
	switch p.value {
	case LOW:
		f.state = !f.state
		pulses := make([]pulse_t, len(f.destinations))
		for i := range f.destinations {
			pulses[i] = pulse_t{source: f.label, destination: f.destinations[i], value: f.state}
		}
		return pulses
	default:
		return nil
	}
}

func (f flip_flop_t) GetDestinations() []string {
	return f.destinations
}

type conjunction_t struct {
	label        string
	inputs       map[string]bool
	destinations []string
}

func (c *conjunction_t) Process(p pulse_t) []pulse_t {

	c.inputs[p.source] = p.value

	val := LOW
	for _, v := range c.inputs {
		if !v {
			val = HIGH
			break
		}
	}

	pulses := make([]pulse_t, len(c.destinations))
	for i := range c.destinations {
		pulses[i] = pulse_t{source: c.label, destination: c.destinations[i], value: val}
	}
	return pulses
}

func (c conjunction_t) GetDestinations() []string {
	return c.destinations
}

type broadcaster_t struct {
	label        string
	destinations []string
}

func (b *broadcaster_t) Process(p pulse_t) []pulse_t {
	pulses := make([]pulse_t, len(b.destinations))
	for i := range b.destinations {
		pulses[i] = pulse_t{source: b.label, destination: b.destinations[i], value: p.value}
	}
	return pulses
}

func (b broadcaster_t) GetDestinations() []string {
	return b.destinations
}

func ParseSpaceCommaSeparators(s string) []string {
	var ret []string
	for _, t := range strings.Split(strings.TrimSpace(s), ",") {
		ret = append(ret, strings.TrimSpace(t))
	}
	return ret
}

func ParseDay20Input(s string) map[string]module_i {
	file, err := os.Open(s)
	if err != nil {
		log.Fatal(err)
	}
	modules := make(map[string]module_i, 0)

	re := regexp.MustCompile(`([\%\&])*(\w+|broadcaster) ->((?: \w+,*)+)`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		r := re.FindStringSubmatch(scanner.Text())
		var module module_i
		switch {
		case r[2] == "broadcaster":
			module = &broadcaster_t{label: r[2], destinations: ParseSpaceCommaSeparators(r[3])}
		case r[1] == "&":
			module = &conjunction_t{label: r[2], destinations: ParseSpaceCommaSeparators(r[3]), inputs: make(map[string]bool, 0)}
		case r[1] == "%":
			module = &flip_flop_t{label: r[2], destinations: ParseSpaceCommaSeparators(r[3])}
		}
		modules[r[2]] = module
	}

	for src_label := range modules {
		for _, dest_label := range modules[src_label].GetDestinations() {
			switch modules[dest_label].(type) {
			case *conjunction_t:
				modules[dest_label].(*conjunction_t).inputs[src_label] = false
			}
		}
	}
	return modules
}

type pulse_t struct {
	source      string
	destination string
	value       bool
}

func (p pulse_t) String() string {
	var x string
	if p.value {
		x = "high"
	} else {
		x = "low"
	}

	return fmt.Sprintf("%v:%v->%v", p.source, x, p.destination)
}

func PressButton(s string, n_times int) int {

	modules := ParseDay20Input(s)

	queue := make([]pulse_t, 0)
	n_high := 0
	n_low := 0

	for i := 0; i < n_times; i++ {
		queue = append(queue, pulse_t{source: "button", destination: "broadcaster", value: LOW})
		for len(queue) > 0 {
			pulse := queue[0]
			queue = queue[1:]
			switch pulse.value {
			case LOW:
				n_low++
			case HIGH:
				n_high++
			}
			if m, ok := modules[pulse.destination]; ok {
				queue = append(queue, m.Process(pulse)...)
			}
		}
	}

	return n_low * n_high
}

func CountButtonPresses(s string, label string) int {

	modules := ParseDay20Input(s)

	queue := make([]pulse_t, 0)

	for i := 1; true; i++ {
		queue = append(queue, pulse_t{source: "button", destination: "broadcaster", value: LOW})
		for len(queue) > 0 {
			pulse := queue[0]

			if m, ok := modules[pulse.destination]; ok {
				queue = append(queue, m.Process(pulse)...)
			}
			queue = queue[1:]

			if pulse.destination == label && pulse.value == LOW {
				return i
			}
		}
	}

	return math.MaxInt
}

func main() {

	s := filepath.Join("input", "input20.txt")
	part1 := PressButton(s, 1000)
	fmt.Printf("part1: %v\n", part1) // 879834312

	gl := CountButtonPresses(s, "gl")
	bb := CountButtonPresses(s, "bb")
	kk := CountButtonPresses(s, "kk")
	mr := CountButtonPresses(s, "mr")
	part2 := gl * bb * kk * mr

	fmt.Printf("part2: %v\n", part2) // 243037165713371
}
