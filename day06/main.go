package main

import (
	"fmt"
)

var input string = "input.txt"

// count counts the occurence of answers given in a group of persons
func count(g []string) map[rune]int {
	m := make(map[rune]int)
	for _, p := range g {
		for _, a := range p {
			m[a] = m[a] + 1
		}
	}
	return m
}

// countAnswersPerGroupAnyone counts all questions anybody answered with yes
func countAnswersPerGroupAnyone(g []string) (c int) {
	m := count(g)
	for range m {
		c++
	}
	return c
}

// countGroupsAnswersEveryone counts all questions everybody answered with yes
func countAnswersPerGroupEveryone(g []string) int {
	m := count(g)
	allsame := 0
	for _, v := range m {
		if v == len(g) {
			allsame++
		}
	}
	return allsame
}

func sumOverAllGroups(groups [][]string, f func([]string) int) (c int) {
	for _, g := range groups {
		c += f(g)
	}
	return c
}

func main() {
	groups := readAllGroups(input)

	fmt.Printf("Counting where ANYONE of a group answered YES  : %d\n", sumOverAllGroups(groups, countAnswersPerGroupAnyone))
	fmt.Printf("Counting where EVERYONE of a group answered YES: %d\n", sumOverAllGroups(groups, countAnswersPerGroupEveryone))
}
