package main

import (
	"fmt"
)

var input string = "input.txt"

func Max(v ...int) int {
	m := v[0]
	for _, e := range v {
		if e > m {
			m = e
		}
	}
	return m
}

func getDevicesJolts(nums []int) int {
	return Max(nums...) + 3
}

func distance(a, b int) int {
	return b - a
}

func getProductOfCountOfDistances(adapters []int) int {
	ones := 0
	threes := 0
	devices := append(adapters, getDevicesJolts(adapters))
	for i := range devices[:len(devices)-1] {
		if distance(devices[i], devices[i+1]) == 1 {
			ones++
		}
		if distance(devices[i], devices[i+1]) == 3 {
			threes++
		}
	}
	return threes * ones
}

func countCombinations(nums []int) int {
	collector := []int{1}
	nums = append(nums, getDevicesJolts(nums))
	for i := 1; i < len(nums); i++ {
		folds := 0
		for j := 0; j < i; j++ {
			if nums[j]+3 >= nums[i] {
				folds += collector[j]
			}
		}
		collector = append(collector, folds)

	}
	return collector[len(collector)-1]
}

func main() {
	nums := readdata(input)

	m := getProductOfCountOfDistances(nums)
	fmt.Printf("Product of differences of 3 and 1 jolts: %d\n", m)

	m = countCombinations(nums)
	fmt.Printf("Valid compbinations of adapters: %d\n", m)
}
