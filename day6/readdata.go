package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func readAllGroups(fname string) (groups [][]string) {

	f, err := os.Open(fname)
	if err != nil {
		log.Fatalf("Error opening dataset '%s':  %s", input, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	var line string
	var group []string
	for scanner.Scan() {
		line = scanner.Text()
		line = strings.Trim(line, " ")
		//log.Printf("line: '%s'", line)
		if line == "" {
			groups = append(groups, group)
			group = nil
		} else {
			group = append(group, line)
		}
	}
	groups = append(groups, group)
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return groups
}
