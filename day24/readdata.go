package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func ReadData(fname string) (lines []string) {

	f, err := os.Open(fname)
	if err != nil {
		log.Fatalf("Error opening dataset '%s':  %s", fname, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		lines = append(lines, strings.Trim(strings.ToLower(scanner.Text()), " "))
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return lines
}
