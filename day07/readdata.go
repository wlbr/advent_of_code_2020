package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func getBag(rule string) *bag {
	matches := regx1.FindStringSubmatch(rule)
	if len(matches) != 3 {
		err := fmt.Errorf("Luggage rule not in standard format. Directive: '%s', len(matches):%d", rule, len(matches))
		log.Printf("Error: %s", err)
		return nil
	}
	return NewBag(matches[1], "0")
}

func getPostconditions(pc string) []*bag {
	var bags []*bag

	matches := regx2.FindStringSubmatch(pc)
	if len(matches) == 2 {
		pc = matches[1]
	}
	if pc == "no other bags." {
		return nil
	}
	for _, spc := range strings.Split(pc, ",") {
		//	log.Printf("spc: %s", spc)
		matches := regx3.FindStringSubmatch(spc)
		if len(matches) != 3 {
			err := fmt.Errorf("Luggage rule not in standard format. Directive: '%s', len(matches):%d", spc, len(matches))
			log.Printf("Error: %s", err)
			return nil
		}

		b := NewBag(matches[2], matches[1])
		bags = append(bags, b)
	}

	return bags
}

func getRules(lines []string) map[string]*rule {
	rules := make(map[string]*rule)
	for _, r := range lines {
		r := &rule{bag: getBag(r), carries: getPostconditions(r)}
		rules[r.bag.name] = r
	}
	return rules
}

func readAllRules(fname string) (rules []string) {

	f, err := os.Open(fname)
	if err != nil {
		log.Fatalf("Error opening dataset '%s':  %s", input, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	var line string
	for scanner.Scan() {
		line = scanner.Text()
		line = strings.Trim(line, " ")
		//log.Printf("line: '%s'", line)
		rules = append(rules, line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return rules
}
