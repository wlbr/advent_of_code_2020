package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var input string = "input.txt"
var regx1 = regexp.MustCompile(`(\d+)-(\d+) (.+): (.+)$`)

func parsePasswordDirective1(buf string) (min, max int, gold, password string, err error) {
	matches := regx1.FindStringSubmatch(buf)
	if len(matches) != 5 {
		err = fmt.Errorf("Password directive not in standard format. Directive: '%s', len(matches):%d", buf, len(matches))
		return min, max, gold, password, err
	}
	min, err = strconv.Atoi(matches[1])
	if err != nil {
		err = fmt.Errorf("Min value not a number. min: '%s'", matches[1])
		return min, max, gold, password, err
	}
	max, err = strconv.Atoi(matches[2])
	if err != nil {
		err = fmt.Errorf("Max value not a number. min: '%s'", matches[1])
		return min, max, gold, password, err
	}
	gold = matches[3]
	password = matches[4]

	return min, max, gold, password, err
}

func checkPassword1(min, max int, gold string, password string, err error) (valid bool) {
	if err == nil {
		c := strings.Count(password, gold)
		if (min <= c) && (max >= c) {
			valid = true
		}
	}
	return valid
}

func checkPassword2(min, max int, gold string, password string, err error) (valid bool) {
	if err == nil {
		//XOR
		if (string((password)[min-1]) == gold) != (string(password[max-1]) == gold) {
			valid = true
		}
	}
	return valid
}

func main() {
	f, err := os.Open(input)
	if err != nil {
		log.Fatalf("Error opening dataset '%s':  %s", input, err)
	}

	scanner := bufio.NewScanner(f)
	count1 := 0
	count2 := 0
	for scanner.Scan() {
		line := scanner.Text()
		if checkPassword1(parsePasswordDirective1(line)) {
			count1++
		}
		if checkPassword2(parsePasswordDirective1(line)) {
			count2++
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	fmt.Printf("Correct passwords count1: %d\n", count1)
	fmt.Printf("Correct passwords count2: %d\n", count2)
}
