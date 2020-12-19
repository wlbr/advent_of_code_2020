package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"text/scanner"
)

var input string = "input.txt"

func add(nums ...int) (sum int) {
	//fmt.Println("  add", nums)
	for _, n := range nums {
		sum += n
	}
	return sum
}

func mul(nums ...int) (product int) {
	//fmt.Println("  mul", nums)
	product = 1
	for _, n := range nums {
		product *= n
	}
	return product
}

func id(nums ...int) int {
	//fmt.Println("  id", nums)
	return nums[1]
}

func findcorrespondingclosing(set []string) int {
	op := "("
	cl := ")"
	c := 0
	for i := 0; i < len(set); i++ {
		if set[i] == op {
			c++
		}
		if set[i] == cl {
			c--
		}
		if c == 0 {
			return i
		}
	}
	log.Printf("malformat parenthesis")
	return -1
}

func evaluate(expr []string) int {
	ops := id
	if len(expr) == 0 {
		return 0
	}
	result, _ := strconv.Atoi(expr[0])
	for i := 0; i < len(expr); i++ {
		x := expr[i]
		switch x {
		case "+":
			ops = add
			break
		case "*":
			ops = mul
			break
		case "(":
			cpi := findcorrespondingclosing(expr[i:])
			result = ops(result, evaluate(expr[i+1:i+cpi]))
			i += cpi
			break
		default:
			if n, err := strconv.Atoi(x); err == nil {
				result = ops(result, n)
			}
		}
	}
	return result
}

func evaluateExpression(line string) int {
	var s scanner.Scanner
	s.Init(strings.NewReader(line))
	var expr []string
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		expr = append(expr, s.TokenText())
	}
	i := evaluate(expr)
	return i
}

func t1(lines []string) int {
	var results []int
	for _, l := range lines {
		//fmt.Println(l)
		results = append(results, evaluateExpression(l))
	}
	return add(results...)
}

func main() {

	nums := readdata(input)
	fmt.Printf("Task1: %d\n", t1(nums))
	fmt.Printf("Task2: %d\n", t2(nums))

}
