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
		log.Fatalf("Error opening dataset '%s':  %s", fname, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := strings.Split(strings.ToLower(strings.Trim(scanner.Text(), " ")), ",")
		for _, ns := range line {
			if n, e := strconv.Atoi(ns); e == nil {
				nums = append(nums, n)
			} else {
				log.Printf("error: not a integer in input line '%s'. Ignored. error: %s", line, e)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return nums
}
