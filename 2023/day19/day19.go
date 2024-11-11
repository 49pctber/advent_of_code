package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type interval_t struct {
	start int
	end   int
}

func (i interval_t) String() string {
	return fmt.Sprintf("[%d, %d]", i.start, i.end)
}

func (i *interval_t) Contains(x int) bool {
	return i.start <= x || x <= i.end
}

func (i *interval_t) Length() int {
	return i.end - i.start + 1 // endpoint inclusive
}

type intervals_t struct {
	intervals []interval_t
}

func (i intervals_t) String() string {
	s := []string{}
	for _, interval := range i.intervals {
		s = append(s, interval.String())
	}
	return strings.Join(s, " ")
}

func (i *intervals_t) Consolidate() {

	for done := false; !done; {

		done = true

		sort.SliceStable(i.intervals, func(k, j int) bool {
			return i.intervals[k].start < i.intervals[j].start
		})

		for ii := range i.intervals {
			if ii == 0 {
				continue
			}

			if i.intervals[ii-1].end >= i.intervals[ii].start {
				i.intervals[ii-1].end = max(i.intervals[ii-1].end, i.intervals[ii].end)
				i.intervals = append(i.intervals[:ii], i.intervals[ii+1:]...)
				done = false
				break
			}
		}

	}

}

func (i *intervals_t) Inside(x int) bool {

	i.Consolidate()
	for _, interval := range i.intervals {
		if interval.Contains(x) {
			return true
		}
	}

	return false
}

func (i *intervals_t) Length() int {
	i.Consolidate()
	length := 0
	for _, interval := range i.intervals {
		length += interval.Length()
	}
	return length
}

func (i intervals_t) UnionLength(j intervals_t) int {
	k := intervals_t{}
	i.Consolidate()
	j.Consolidate()
	k.intervals = append(k.intervals, i.intervals...)
	k.intervals = append(k.intervals, j.intervals...)
	return k.Length()
}

type rule_t struct {
	category    byte
	operator    byte
	threshold   int
	destination string
}

func (r rule_t) String() string {
	return fmt.Sprintf("{%c%c%d:%v}", r.category, r.operator, r.threshold, r.destination)
}

func (r rule_t) Evaluate(mp machinepart_t) (bool, string) {
	var val int
	switch r.category {
	case 'x':
		val = mp.x
	case 'm':
		val = mp.m
	case 'a':
		val = mp.a
	case 's':
		val = mp.s
	default:
		log.Fatal("shouldn't get here")
	}

	switch r.operator {
	case '<':
		if val < r.threshold {
			return true, r.destination
		} else {
			return false, ""
		}
	case '>':
		if val > r.threshold {
			return true, r.destination
		} else {
			return false, ""
		}
	default:
		log.Fatal("shouldn't get here")
	}

	log.Fatal("shouldn't get here")
	return false, ""
}

type workflow_t struct {
	label               string
	rules               []rule_t
	default_destination string
}

func (w workflow_t) String() string {
	return fmt.Sprintf("%s:{%v:%s}", w.label, w.rules, w.default_destination)
}

func (w workflow_t) Evaluate(mp machinepart_t) string {

	for _, rule := range w.rules {
		if match, dest := rule.Evaluate(mp); match {
			return dest
		}
	}

	return w.default_destination
}

type machinepart_t struct {
	x int
	m int
	a int
	s int
}

func (mp machinepart_t) String() string {
	return fmt.Sprintf("{x=%d m=%d a=%d s=%d}", mp.x, mp.m, mp.a, mp.s)
}

func (mp machinepart_t) SumRatings() int {
	return mp.x + mp.m + mp.a + mp.s
}

func (mp machinepart_t) Accept(wfs map[string]workflow_t) bool {
	for dest := wfs["in"].Evaluate(mp); ; {
		switch dest {
		case "A":
			return true
		case "R":
			return false
		default:
			dest = wfs[dest].Evaluate(mp)
		}
	}
}

func ParseDay19Input(s string) (map[string]workflow_t, []machinepart_t, error) {
	file, err := os.Open(s)
	if err != nil {
		return map[string]workflow_t{}, []machinepart_t{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scan_workflows := true
	workflows := make(map[string]workflow_t)
	machine_parts := []machinepart_t{}
	re_workflow := regexp.MustCompile(`^(\w+)\{(.*)\}$`)
	re_machinepart := regexp.MustCompile(`^\{x=(\d+),m=(\d+),a=(\d+),s=(\d+)\}$`)
	re_rule := regexp.MustCompile(`^(\w)([\<\>])(\d+)\:(\w+)$`)

	for scanner.Scan() {
		if scanner.Text() == "" {
			scan_workflows = false
			continue
		}

		if scan_workflows {
			r := re_workflow.FindStringSubmatch(scanner.Text())
			wf := workflow_t{label: r[1]}
			rs := strings.Split(r[2], ",")
			for _, rr := range rs {
				if r := re_rule.FindStringSubmatch(rr); r != nil {
					rule := rule_t{category: r[1][0], operator: r[2][0], destination: r[4]}
					if rule.threshold, err = strconv.Atoi(r[3]); err != nil {
						return map[string]workflow_t{}, []machinepart_t{}, err
					}
					wf.rules = append(wf.rules, rule)
				} else {
					wf.default_destination = rr
				}
			}
			workflows[r[1]] = wf
		} else {
			r := re_machinepart.FindStringSubmatch(scanner.Text())
			mp := machinepart_t{}
			if mp.x, err = strconv.Atoi(r[1]); err != nil {
				return map[string]workflow_t{}, []machinepart_t{}, err
			}
			if mp.m, err = strconv.Atoi(r[2]); err != nil {
				return map[string]workflow_t{}, []machinepart_t{}, err
			}
			if mp.a, err = strconv.Atoi(r[3]); err != nil {
				return map[string]workflow_t{}, []machinepart_t{}, err
			}
			if mp.s, err = strconv.Atoi(r[4]); err != nil {
				return map[string]workflow_t{}, []machinepart_t{}, err
			}
			machine_parts = append(machine_parts, mp)
		}
	}

	return workflows, machine_parts, nil
}

type mprange_t struct {
	x intervals_t
	m intervals_t
	a intervals_t
	s intervals_t
}

func (r mprange_t) Product() int {
	return r.x.Length() * r.m.Length() * r.a.Length() * r.s.Length()
}

func (i intervals_t) Intersection(j intervals_t) intervals_t {
	// UNTESTED!
	k := intervals_t{}

	for _, ii := range i.intervals {
		for _, jj := range j.intervals {
			mi := max(ii.start, jj.start)
			ma := min(ii.end, jj.end)
			if mi < ma {
				k.intervals = append(k.intervals, interval_t{start: mi, end: ma})
			}
		}
	}

	k.Consolidate()
	return k
}

func (r mprange_t) Split(category byte, threshold int) (mprange_t, mprange_t) {
	low, high := mprange_t{x: r.x, m: r.m, a: r.a, s: r.s}, mprange_t{x: r.x, m: r.m, a: r.a, s: r.s}
	lowint := intervals_t{intervals: []interval_t{{start: 1, end: threshold - 1}}}
	highint := intervals_t{intervals: []interval_t{{start: threshold, end: 4000}}}
	switch category {
	case 'x':
		low.x = r.x.Intersection(lowint)
		high.x = r.x.Intersection(highint)
	case 'm':
		low.m = r.m.Intersection(lowint)
		high.m = r.m.Intersection(highint)
	case 'a':
		low.a = r.a.Intersection(lowint)
		high.a = r.a.Intersection(highint)
	case 's':
		low.s = r.s.Intersection(lowint)
		high.s = r.s.Intersection(highint)
	default:
		log.Fatal("shouldn't be here")
	}
	return low, high
}

func recursiveWorkflows(mprange mprange_t, label string, wfs map[string]workflow_t) int {

	if mprange.Product() == 0 {
		// there are no combinations in this interval
		return 0
	} else if label == "A" {
		// if accepted, return product of lengths of intervals
		return mprange.Product()
	} else if label == "R" {
		// if rejected, return 0
		return 0
	}

	// apply rules
	sum := 0
	for _, rule := range wfs[label].rules {
		var newrange mprange_t

		switch rule.operator {
		case '>':
			mprange, newrange = mprange.Split(rule.category, rule.threshold+1)
			sum += recursiveWorkflows(newrange, rule.destination, wfs)
		case '<':
			newrange, mprange = mprange.Split(rule.category, rule.threshold)
			sum += recursiveWorkflows(newrange, rule.destination, wfs)
		default:
			log.Fatalf("invalid operator %c", rule.operator)
		}

	}

	sum += recursiveWorkflows(mprange, wfs[label].default_destination, wfs)

	return sum
}

func Combinations(wfs map[string]workflow_t) int {
	ranges := mprange_t{
		x: intervals_t{intervals: []interval_t{{start: 1, end: 4000}}},
		m: intervals_t{intervals: []interval_t{{start: 1, end: 4000}}},
		a: intervals_t{intervals: []interval_t{{start: 1, end: 4000}}},
		s: intervals_t{intervals: []interval_t{{start: 1, end: 4000}}},
	}
	return recursiveWorkflows(ranges, "in", wfs)
}

func main() {
	wfs, mps, err := ParseDay19Input(filepath.Join("input", "input19.txt"))
	if err != nil {
		log.Fatal(err)
	}

	part1 := 0
	for _, mp := range mps {
		if mp.Accept(wfs) {
			part1 += mp.SumRatings()
		}
	}
	fmt.Printf("part1: %v\n", part1) // 383682

	part2 := Combinations(wfs)
	fmt.Printf("part2: %v\n", part2)
}
