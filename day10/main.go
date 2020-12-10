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

func getDevicesJolts(nums []int) int {
	return Max(nums...) + 3
}

func getCandidates(nums []int, target int) []int {
	start := -1
	end := -1
	limit := target + 3
	for i, v := range nums {
		if start == -1 && v > target && v <= limit {
			start = i
			end = i
		} else if v > limit {
			end = i - 1
			break
		}
	}
	if start >= 0 && end >= 0 {
		return nums[start : end+1]
	} else {
		return nil
	}
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

func isValidPath(nums []int) bool {

	for i := range nums[:len(nums)-1] {
		d := distance(nums[i], nums[i+1])
		if d != 1 && d != 3 {
			return false
		}
	}
	return true
}

func contains(nums []int, n int) bool {
	for _, i := range nums {
		if n == i {
			return true
		}
	}
	return false
}

func countCombinations(nums []int) int {
	dp := []int{1}
	nums = append(nums, getDevicesJolts(nums))
	for i := 1; i < len(nums); i++ {
		ans := 0
		for j := 0; j < i; j++ {
			if nums[j]+3 >= nums[i] {
				ans += dp[j]
			}
		}
		dp = append(dp, ans)

	}
	return dp[len(dp)-1]
}

func main() {
	nums := readdata(input)

	m := getProductOfCountOfDistances(nums)
	fmt.Printf("Product differences of 3 and 1 jolts: %d\n", m)

	m = countCombinations(nums)
	fmt.Printf("Product differences of 3 and 1 jolts: %d\n", m)

	//n := getASolution(nums)
	//fmt.Printf("Product differences of 3 and 1 jolts: %d\n", n)
	//	fmt.Printf("Sum of min and max of contigues sum: %d\n", lookForSumOfContiguousSum(nums, n))

}
