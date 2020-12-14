package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
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
		lines = append(lines, strings.ToLower(strings.Trim(scanner.Text(), " ")))
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return lines
}

var splitall = regexp.MustCompile(`^([[:alpha:]]+)(\[(.+)\])? = (.+)$`)

func parsedata(lines []string) (program []Command) {
	parsers := make(map[string]func([]string) (Command, error))
	parsers[MEM] = parseMem
	parsers[MASK] = parseMask

	for _, l := range lines {
		matches := splitall.FindStringSubmatch(l)
		if len(matches) != 5 {
			log.Printf("Warning: Command not in standard format: '%s'", l)
		} else {
			if f, ok := parsers[matches[1]]; ok {
				if c, err := f(matches[2:]); err == nil {
					program = append(program, c)
				} else {
					log.Printf("Warning: parseerror '%s' in line '%s'", err, l)
				}
			} else {
				log.Printf("Warning: unhandled command '%s'", l)
			}
		}
	}
	return program
}
