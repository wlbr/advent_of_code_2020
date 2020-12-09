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
	//fmt.Println(last)
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

func checkPossibleSumOfAll(nums []int, target int) (found bool, findings []int) {
	for k := 0; k < len(nums)-1; k++ {
		n := nums[k:]

		for i := 0; i < len(n)-1; i++ {
			m := add(n[:i]...)
			if m == target {
				found = true
				findings = n[:i]
				return found, findings
			}
		}
	}

	return found, findings
}

func lookForSumOfContiguousSum(nums []int, target int) int {
	result := 0

	for i, t := range nums {
		if t == target {
			if r, findings := checkPossibleSumOfAll(nums[:i], target); r {
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
