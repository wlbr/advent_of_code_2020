package main

func Max(v ...int) int {
	if len(v) > 0 {
		m := v[0]
		for _, e := range v {
			if e > m {
				m = e
			}
		}
		return m
	} else {
		return 0
	}
}

func add(nums ...int) (sum int) {
	for _, n := range nums {
		sum += n
	}
	return sum
}

func findn(nums []int, n int) (pos int) {
	for i, v := range nums {
		if v == n {
			return i
		}
	}
	return -1
}
