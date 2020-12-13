package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readdata(fname string) (start int, alllines, activelines []int) {

	f, err := os.Open(fname)
	if err != nil {
		log.Fatalf("Error opening dataset '%s':  %s", fname, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	if scanner.Scan() {
		line := scanner.Text()
		if n, e := strconv.Atoi(line); e == nil {
			start = n
		} else {
			log.Printf("error: not a integer in input line '%s'. Ignored. error: %s", line, e)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	if scanner.Scan() {
		line := scanner.Text()
		slines := strings.Split(line, ",")

		for _, l := range slines {
			if n, e := strconv.Atoi(l); e == nil {
				alllines = append(alllines, n)
				activelines = append(activelines, n)
			} else {
				alllines = append(alllines, -1)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return start, alllines, activelines
}
