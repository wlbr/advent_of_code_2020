package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// func readAllGroup(p string) (passport map[string]string) {

// 	scanner := bufio.NewScanner(strings.NewReader(p))
// 	scanner.Split(bufio.ScanWords)

// 	passport = make(map[string]string)
// 	for scanner.Scan() {
// 		p := scanner.Text()

// 		field := strings.Split(p, ":")
// 		//log.Printf("f1: '%s', f2: '%s'\n", field[0], field[1])
// 		passport[field[0]] = field[1]
// 	}
// 	return passport
// }

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
