package main

import (
	"fmt"
	"sync"
)

func speakNumber(values map[int][]int, m []int, ok bool, n, turn int) {
	m, ok = values[n]

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
		m, ok := recent[n]
		speakNumber(recent, m, ok, n, turn)
		turn++
	}
	fmt.Println(recent)
	for ; turn <= nth; turn++ {
		m, ok := recent[n]
		//	fmt.Printf("turn: %d  n1: %d m: %v", turn, n, m)
		if ok {
			if m[1] == 0 {
				n = 0
			} else {
				n = m[0] - m[1]
			}
			shift := []int{turn, m[0]}
			recent[n] = shift
		} else {
			n = 0
			recent[n] = []int{turn, 0}
		}

		//speakNumber(recent, m, ok, n, turn)
	}
	return n
}

func main() {
	//var input string = "input.txt"
	//nums := readdata(input)
	//nums := []int{6, 19, 0, 5, 7, 13, 1}
	nums := []int{0, 3, 6}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		t := 2020
		fmt.Printf("The %dth spoken number is: %d\n", t, GetNthNumberFromElvesTalk(nums, t))
	}()
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	t := 30000000
	// 	fmt.Printf("The %dth spoken number is: %d\n", t, GetNthNumberFromElvesTalk(nums, t))
	// }()
	wg.Wait()
}
