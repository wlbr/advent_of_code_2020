package main

import (
	"fmt"
	"regexp"
	"strconv"
)

var input string = "input.txt"
var regx1 = regexp.MustCompile(`^(.+) bag.? contain (.+)$`)
var regx2 = regexp.MustCompile(`^.+contain (.+)$`)
var regx3 = regexp.MustCompile(`^ *(\d+) (.+) bag.*$`)
var regx4 = regexp.MustCompile(`no other bags$`)

type bag struct {
	name     string
	quantity int
}

func NewBag(name string, quantity string) *bag {
	b := &bag{name: name}

	n, e := strconv.Atoi(quantity)
	if e != nil {
		b.quantity = 0
	} else {
		b.quantity = n
	}
	return b
}

func (b *bag) String() string {
	return fmt.Sprintf("%d %s", b.quantity, b.name)
}

type rule struct {
	bag     *bag
	carries []*bag
}

func (r *rule) String() string {
	res := "'" + r.bag.name + "'" + " ==>"
	for i, c := range r.carries {
		if i == 0 {
			res += " "
		} else {
			res += ", "
		}
		res += "'" + c.String() + "'"
	}
	return res
}

func contains(set []*bag, s string) (contains bool) {
	for _, i := range set {
		if i.name == s {
			return true
		}
	}
	return false
}

func AppendIfMissing(slice []*bag, bags ...*bag) []*bag {
	out := slice
	for _, b := range bags {
		isin := false
		for _, ele := range slice {
			if ele.name == b.name {
				isin = true
			}
		}
		if !isin {
			out = append(out, b)
		}
	}
	return out
}

func whoCanCarryThisInner(target string, issuer, currentrule *rule, rules map[string]*rule, knowncarriers []*bag) []*bag {
	if contains(currentrule.carries, target) {
		knowncarriers = AppendIfMissing(knowncarriers, issuer.bag)
	}
	for _, c := range currentrule.carries {
		knowncarriers = AppendIfMissing(knowncarriers, whoCanCarryThisInner(target, issuer, rules[c.name], rules, knowncarriers)...)
	}

	return knowncarriers
}

func whoCanCarryThis(target string, rules map[string]*rule) []*bag {
	var knowncarriers []*bag
	i := 0
	for _, r := range rules {
		i++
		//fmt.Println(i)
		knowncarriers = AppendIfMissing(knowncarriers, whoCanCarryThisInner(target, r, r, rules, knowncarriers)...)
	}

	return knowncarriers
}

func howManyBagsInner(target string, rules map[string]*rule, count int) int {
	r, ok := rules[target]
	if ok {
		for _, c := range r.carries {
			count += c.quantity * howManyBagsInner(c.name, rules, 1)
		}
	}
	return count
}

func howManyBags(target string, rules map[string]*rule) int {
	return howManyBagsInner(target, rules, 0)

}

func main() {
	target := "shiny gold"
	rules := getRules(readAllRules(input))

	e := len(whoCanCarryThis(target, rules))
	fmt.Printf("Count of bags that can hold '%s': %d\n", target, e)

	f := howManyBags(target, rules)
	fmt.Printf("The '%s' needs to carry '%d' bags\n", target, f)
}
