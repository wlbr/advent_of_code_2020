package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readdata(fname string) (nums []int) {

	f, err := os.Open(fname)
	if err != nil {
		log.Fatalf("Error opening dataset '%s':  %s", input, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), " ")
		if n, e := strconv.Atoi(line); e == nil {
			nums = append(nums, n)
		} else {
			log.Printf("error: not a integer in input line '%s'. Ignored. error: %s", line, e)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return nums
}
