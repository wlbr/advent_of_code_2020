package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func readdata(fname string) (lines []string) {

	f, err := os.Open(fname)
	if err != nil {
		log.Fatalf("Error opening dataset '%s':  %s", fname, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		lines = append(lines, strings.Trim(scanner.Text(), " "))
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return lines
}

func parsedata(lines []string) (cubes map[string]bool) {
	cubes = make(map[string]bool)
	corr := 0
	for y, line := range lines {
		for x, v := range line {
			fmt.Printf("x:%d y:%d ? %s  -  %s\n", x, y, string(v), coord2String([]int{x - corr, y - corr, 0}))
			if string(v) == active {
				cubes[coord2String([]int{x - corr, y - corr, 0})] = true
			} else {
				cubes[coord2String([]int{x - corr, y - corr, 0})] = false
			}
		}
	}
	return cubes
}
