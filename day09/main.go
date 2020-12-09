package main

import (
	"fmt"
)

var input string = "input.txt"

func Min(v ...int) int {
	m := v[0]
	for _, e := range v {
		if e < m {
			m = e
		}
	}
	return m
}

func Max(v ...int) int {
	m := v[0]
	for _, e := range v {
		if e > m {
			m = e
		}
	}
	return m
}

func add(nums ...int) (sum int) {
	for _, n := range nums {
		sum += n
	}
	return sum
}

func checkPossibleSum(last []int, target int) bool {
	result := false
	for i, m := range last {
		for _, n := range last[i:] {
			if add(m, n) == target {
				result = true
			}
		}
	}
	return result
}

func lookForFirstNotSum(nums []int, preamble int) int {
	result := 0
	for i, target := range nums {
		if !checkPossibleSum(nums[Max(i-preamble, 0):i], target) {
			result = target
		}
	}
	return result
}

func checkPossibleContiguesSum(nums []int, target int) (found bool, findings []int) {
	for k := range nums {
		n := nums[k:]
		for i := range n {
			m := add(n[:i]...)
			if m == target {
				return true, n[:i]
			}
		}
	}
	return false, nil
}

func lookForSumOfContiguousSum(nums []int, target int) int {
	result := 0
	for i, t := range nums {
		if t == target {
			if r, findings := checkPossibleContiguesSum(nums[:i], target); r {
				result = Min(findings...) + Max(findings...)
			}
		}
	}
	return result
}

func main() {
	nums := readdata(input)

	n := lookForFirstNotSum(nums, 25)
	fmt.Printf("First number not being a sum: %d\n", n)
	fmt.Printf("Sum of min and max of contigues sum: %d\n", lookForSumOfContiguousSum(nums, n))

}
