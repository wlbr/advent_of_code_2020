package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var input string = "input.txt"

func readTorusMap(fname string) (data []string) {

	f, err := os.Open(fname)
	if err != nil {
		log.Fatalf("Error opening dataset '%s':  %s", input, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		strings.TrimRight(line, " ")
		if line != "" {
			data = append(data, line)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return data
}

func countTreesForSlope(torus []string, slopex, slopey int) int {
	j := 0
	treecount := 0
	for i := 0; i < len(torus); i += slopey {
		//log.Printf("len(torus)=%d, len(torus[i]=%d, i=%d, j=%d, torus[%d][j%d])='%s'", len(torus), len(torus[i]), i, j, i, j, string(torus[i][j]))
		if "#" == string(torus[i][j]) {
			treecount++
		}
		j += slopex
		j = j % len(torus[i])
	}
	return treecount
}

func main() {
	torus := readTorusMap(input)
	fmt.Printf("There are %d trees.\n\n", countTreesForSlope(torus, 3, 1))

	slopes := [][]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}

	var results []int
	for _, s := range slopes {
		results = append(results, countTreesForSlope(torus, s[0], s[1]))
		fmt.Printf("There are %d trees.\n", results[len(results)-1])
	}
	product := 1
	for _, r := range results {
		product = product * r
	}
	fmt.Printf("\nThe overall product is %d.\n", product)
}
