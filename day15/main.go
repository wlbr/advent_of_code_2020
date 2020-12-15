package main

import (
	"fmt"
)

func speakNumber(values map[int][]int, n, turn int) {
	m, ok := values[n]
	if !ok {
		values[n] = []int{turn, 0}
	} else {
		shift := []int{turn, m[0]}
		values[n] = shift
	}
}

func wasSpoken(values map[int][]int, n int) (bool, []int) {
	m, ok := values[n]
	if !ok {
		return false, nil
	} else {
		return true, m
	}
}

func GetNthNumberFromElvesTalk(nums []int, nth int) int {
	recent := make(map[int][]int)
	var n, turn int = 0, 1
	for _, n = range nums {
		speakNumber(recent, n, turn)
		turn++
	}

	for ; turn <= nth; turn++ {
		w, s := wasSpoken(recent, n)
		if w {
			if s[1] == 0 {
				n = 0
			} else {
				n = s[0] - s[1]
			}
		} else {
			n = 0
		}
		speakNumber(recent, n, turn)
	}
	return n
}

func main() {
	var input string = "input.txt"
	nums := readdata(input)

	t := 2020
	fmt.Printf("The %dth spoken number is: %d\n", t, GetNthNumberFromElvesTalk(nums, t))
	t = 30000000
	fmt.Printf("The %dth spoken number is: %d\n", t, GetNthNumberFromElvesTalk(nums, t))

}
