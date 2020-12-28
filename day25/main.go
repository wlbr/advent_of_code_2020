package main

import (
	"fmt"
)

func main() {
	var input string = "input.txt"

	fmt.Printf("ticket scanning error rate: %d\n", task1(input))
	fmt.Printf("Prouct of the six departure numbers: %d\n", task2(input))
}
