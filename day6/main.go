package main

import (
	"fmt"
)

var input string = "input.txt"

func countGroupsAnswersAnyone(g []string) (c int) {
	m := make(map[rune]int)
	for _, p := range g {
		for _, a := range p {
			m[a] = m[a] + 1
		}
	}
	for range m {
		c++
	}
	return c
}

func countAllGroups(groups [][]string, f func([]string) int) (c int) {
	for _, g := range groups {
		c += f(g)
	}
	return c
}

func countGroupsAnswersEveryone(g []string) int {
	m := make(map[rune]int)
	for _, p := range g {
		for _, a := range p {
			m[a] = m[a] + 1
		}
	}
	allsame := 0
	for _, v := range m {
		if v == len(g) {
			allsame++
		}
	}
	return allsame
}

func main() {
	groups := readAllGroups(input)

	fmt.Printf("Counting where ANYONE of a group answered YES  : %d\n", countAllGroups(groups, countGroupsAnswersAnyone))
	fmt.Printf("Counting where EVERYONE of a group answered YES: %d\n", countAllGroups(groups, countGroupsAnswersEveryone))

	// fmt.Printf("The seat with the highest id is seat no: '%d'\n", max)
	// fmt.Printf("My seat is one of: %v\n", getMySeat(allSeats))

}
