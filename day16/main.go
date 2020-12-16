package main

import (
	"fmt"
	"strings"
)

type Rule struct {
	kind              string
	areas             [][]int
	possiblePositions []int
	pos               int
}

func NewRule(kind string, areas [][]int, ticketlength int) *Rule {
	r := &Rule{kind: kind, areas: areas}

	for i := 0; i < ticketlength; i++ {
		r.possiblePositions = append(r.possiblePositions, i)
	}
	return r
}

func (r *Rule) hasIn(n int) bool {
	for _, i := range r.areas {
		if n >= i[0] && n <= i[1] {
			return true
		}
	}
	return false
}

func (r *Rule) removeImpossibleNumber(n int) {
	i := find(r.possiblePositions, n)
	if i >= 0 {
		r.possiblePositions = append(r.possiblePositions[:i], r.possiblePositions[i+1:]...)
	}
}

func tickedScanningErrorRate(rules []*Rule, myticket []int, nearbies [][]int) int {
	var errors []int
	var collector []int
	for _, r := range rules {
		for i := 0; i < len(r.areas); i++ {
			a := r.areas[i]
			for j := a[0]; j <= a[1]; j++ {
				collector = append(collector, j)
			}
		}
	}
	maxn := Max(collector...)
	valids := make([]bool, maxn+1)
	for _, c := range collector {
		valids[c] = true
	}

	for _, ticket := range nearbies {
		for _, t := range ticket {
			if t > len(valids)-1 || !valids[t] {
				errors = append(errors, t)
			}
		}
	}

	return add(errors...)
}

func getValidTickets(rules []*Rule, nearbies [][]int) [][]int {
	var collector []int
	for _, r := range rules {
		for i := 0; i < len(r.areas); i++ {
			a := r.areas[i]
			for j := a[0]; j <= a[1]; j++ {
				collector = append(collector, j)
			}
		}
	}

	maxn := Max(collector...)
	valids := make([]bool, maxn+1)
	for _, c := range collector {
		valids[c] = true
	}

	var validtickets [][]int
	for _, ticket := range nearbies {
		ticketIsInvalid := false

		for _, t := range ticket {
			if t > len(valids)-1 || !valids[t] {
				ticketIsInvalid = true
			}
		}
		if !ticketIsInvalid {
			validtickets = append(validtickets, ticket)
		}
	}
	return validtickets
}

func getDepartureRules(rules []*Rule) []*Rule {
	var deps []*Rule
	for _, r := range rules {
		if strings.HasPrefix(strings.ToLower(r.kind), "departure") {
			deps = append(deps, r)
		}
	}
	return deps
}

func removeImpossibleNumbers(rules []*Rule, tickets [][]int) {
	for _, ticket := range tickets {
		for pos, v := range ticket {
			for _, r := range rules {
				if !r.hasIn(v) {
					r.removeImpossibleNumber(pos)
				}
			}
		}
	}
}

func solveLogicEquation(rules []*Rule) {
	for changed := true; changed == true; {
		changed = false
		for _, r := range rules {
			if len(r.possiblePositions) == 1 {
				changed = true
				r.pos = r.possiblePositions[0]
				for _, rr := range rules {
					rr.removeImpossibleNumber(r.pos)
				}
			}
		}
	}
}

func ProductOfDepartureField(rules []*Rule, myticket []int, nearbies [][]int) int {
	departurerules := getDepartureRules(rules)
	validtickets := getValidTickets(rules, nearbies)
	removeImpossibleNumbers(rules, validtickets)
	solveLogicEquation(rules)

	res := 1
	for _, r := range departurerules {
		res = res * myticket[r.pos]
	}
	return res
}

func main() {
	var input string = "input.txt"
	rules, myticket, nearbies := ParseData(ReadData(input))

	fmt.Printf("ticket scanning error rate: %d\n", tickedScanningErrorRate(rules, myticket, nearbies))
	fmt.Printf("Prouct of the six departure numbers: %d\n", ProductOfDepartureField(rules, myticket, nearbies))
}
