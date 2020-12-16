package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

func atoi(s string) int {
	if n, err := strconv.Atoi(strings.Trim(s, " ")); err == nil {
		return n
	} else {
		return 0
	}
}

func parseRule(line string, ticketlength int) (r *Rule) {
	parts := strings.Split(line, ":")
	if len(parts) != 2 {
		log.Printf("Warning: Rule not in standard format: '%s'", line)
	} else {
		r = NewRule(parts[0], [][]int{}, ticketlength)
		intervals := strings.Split(parts[1], "or")

		for _, i := range intervals {
			borders := strings.Split(i, "-")
			r.areas = append(r.areas, []int{atoi(borders[0]), atoi(borders[1])})
		}
	}
	return r
}

func parseTicket(t string) (n []int) {
	s := strings.Split(t, ",")
	for _, i := range s {
		n = append(n, atoi(strings.Trim(i, " ")))
	}
	return n
}

const ticketdelimiter = "your ticket:"
const nerbydelimiter = "nearby tickets:"
const ignore = ""

func ParseData(lines []string) (rules []*Rule, ticket []int, nearbies [][]int) {
	var i int

	for j := 0; j < 2; j++ {
		rules = nil
		nearbies = nil
		mode := 'r'
		for i = 0; i < len(lines); i++ {
			l := lines[i]
			switch l {
			case ticketdelimiter:
				mode = 't'
				break
			case nerbydelimiter:
				mode = 'n'
				break
			case ignore:
				break
			default:
				switch mode {
				case 'r':
					rules = append(rules, parseRule(l, len(ticket)))
					break
				case 't':
					ticket = append(ticket, parseTicket(l)...)
					break
				default:
					nearbies = append(nearbies, parseTicket(l))
				}
			}
		}
	}
	return rules, ticket, nearbies
}
