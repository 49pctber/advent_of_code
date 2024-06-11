package main

import (
	"fmt"
	"slices"
	"strings"

	_ "embed"
)

type Rule struct {
	lower uint32
	upper uint32
}

func (r Rule) String() string {
	return fmt.Sprintf("%d->%d", r.lower, r.upper)
}

type RuleSet struct {
	Rules []Rule
}

func (rs RuleSet) String() string {
	ret := ""
	for _, rule := range rs.Rules {
		ret = fmt.Sprintf("%s\n%s", ret, rule.String())
	}
	return ret
}

func (rs RuleSet) Sort() {
	slices.SortFunc(rs.Rules, func(a, b Rule) int {
		if a.lower < b.lower {
			return -1
		} else if a.lower == b.lower {
			return 0
		} else {
			return 1
		}
	})
}

func (rs *RuleSet) ConsolidateRules() {
	rs.Sort()
	for i := 0; i+1 < len(rs.Rules); i++ {
		if rs.Rules[i].upper+1 >= rs.Rules[i+1].lower {
			rs.Rules[i].upper = max(rs.Rules[i].upper, rs.Rules[i+1].upper)
			if i+2 < len(rs.Rules) {
				rs.Rules = append(rs.Rules[0:i+1], rs.Rules[i+2:len(rs.Rules)]...)
			} else {
				rs.Rules = rs.Rules[:i]
			}
			i--
		}
	}
}

func LowestAllowed(input string) uint32 {
	strs := strings.Split(input, "\n")
	rules := RuleSet{Rules: make([]Rule, 0)}

	for _, rulestr := range strs {
		var lowerbound, upperbound uint32
		fmt.Sscanf(rulestr, "%d-%d", &lowerbound, &upperbound)
		r := Rule{lower: lowerbound, upper: upperbound}
		rules.Rules = append(rules.Rules, r)
	}
	rules.ConsolidateRules()
	return rules.Rules[0].upper + 1
}

func NumberAllowed(input string) uint32 {
	strs := strings.Split(input, "\n")
	rules := RuleSet{Rules: make([]Rule, 0)}

	for _, rulestr := range strs {
		var lowerbound, upperbound uint32
		fmt.Sscanf(rulestr, "%d-%d", &lowerbound, &upperbound)
		r := Rule{lower: lowerbound, upper: upperbound}
		rules.Rules = append(rules.Rules, r)
	}
	rules.ConsolidateRules()

	var count uint32 = 0
	for i := 1; i < len(rules.Rules); i++ {
		count += rules.Rules[i].lower - rules.Rules[i-1].upper - 1
	}
	// account for max_uint32 somewhere here...

	return count
}

//go:embed input
var input string

func Part1() {
	fmt.Printf("Part 1: %v\n", LowestAllowed(input))
}

func Part2() {
	fmt.Printf("Part 2: %v\n", NumberAllowed(input))
}

func main() {
	Part1()
	Part2()
}
