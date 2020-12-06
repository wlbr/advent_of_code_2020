package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var input string = "input.txt"

func scanPassport(p string) (passport map[string]string) {

	scanner := bufio.NewScanner(strings.NewReader(p))
	scanner.Split(bufio.ScanWords)

	passport = make(map[string]string)
	for scanner.Scan() {
		p := scanner.Text()

		field := strings.Split(p, ":")
		//log.Printf("f1: '%s', f2: '%s'\n", field[0], field[1])
		passport[field[0]] = field[1]
	}
	return passport
}

func readPassports(fname string) (passports []map[string]string) {

	f, err := os.Open(fname)
	if err != nil {
		log.Fatalf("Error opening dataset '%s':  %s", input, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	var p, line string
	for scanner.Scan() {
		line = scanner.Text()
		line = strings.Trim(line, " ")
		//log.Printf("line: '%s'", line)
		if line == "" {
			passports = append(passports, scanPassport(p))
			p = ""
		} else {
			p = strings.Trim(p+" "+line, " ")
		}
	}
	p = strings.Trim(p+" "+line, " ")
	passports = append(passports, scanPassport(p))
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return passports
}

func checkValidPassport(p map[string]string) (result bool) {
	checks := []check{checkByr, checkIyr, checkEyr, checkHgt, checkEcl, checkHcl, checkPid, checkCid}
	return and(p, checks...)
}

func main() {
	passports := readPassports(input)
	completepassports := 0
	for _, p := range passports {
		//log.Printf("%v", p)
		if checkNeededFields(p) {
			completepassports++
		}
	}
	validpassports := 0
	for _, p := range passports {
		//log.Printf("%v", p)
		if checkValidPassport(p) {
			validpassports++
		}
	}
	fmt.Printf("Found %d candidates, \n %d complete passports, \n %d valid passports.\n", len(passports), completepassports, validpassports)
}
