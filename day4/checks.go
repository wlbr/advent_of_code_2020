package main

import (
	"strconv"
	"strings"
)

func checkNeededFields(p map[string]string) (result bool) {
	needed := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	result = true
	for _, f := range needed {
		if p[f] == "" {
			result = false
		}
	}
	return result
}

func checkStrNumber(sn string, min, max int) bool {
	n, err := strconv.Atoi(sn)
	if err == nil && n >= min && n <= max {
		return true
	}
	return false
}

func checkByr(p map[string]string) (result bool) {
	// byr (Birth Year) - four digits; at least 1920 and at most 2002.
	byr := p["byr"]
	if len(byr) == 4 {
		result = checkStrNumber(byr, 1920, 2002)
	}
	return result
}
func checkIyr(p map[string]string) (result bool) {
	//iyr (Issue Year) - four digits; at least 2010 and at most 2020.
	iyr := p["iyr"]
	if len(iyr) == 4 {
		result = checkStrNumber(iyr, 2010, 2020)
	}
	return result
}
func checkEyr(p map[string]string) (result bool) {
	// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
	eyr := p["eyr"]
	if len(eyr) == 4 {
		result = checkStrNumber(eyr, 2020, 2030)
	}
	return result
}
func checkHgt(p map[string]string) (result bool) {
	//hgt (Height) - a number followed by either cm or in:
	//  If cm, the number must be at least 150 and at most 193.
	//  If in, the number must be at least 59 and at most 76.
	hgt := p["hgt"]
	if len(hgt) > 2 {
		if "cm" == hgt[len(hgt)-2:] && checkStrNumber(hgt[:len(hgt)-2], 150, 193) {
			result = true
		}
		if "in" == hgt[len(hgt)-2:] && checkStrNumber(hgt[:len(hgt)-2], 59, 76) {
			result = true
		}
	}
	return result
}

func checkHcl(p map[string]string) (result bool) {
	//hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
	hcl := p["hcl"]
	if len(hcl) == 7 {
		if string(hcl[0]) == "#" {
			result = true
			for _, b := range hcl[1:] {
				if !strings.Contains("0123456789abcdef", string(b)) {
					result = false
				}
			}
		}
	}
	return result
}
func checkEcl(p map[string]string) (result bool) {
	// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
	ecl := p["ecl"]
	colors := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	for _, c := range colors {
		if ecl == c {
			result = true
		}
	}
	return result
}
func checkPid(p map[string]string) (result bool) {
	// pid (Passport ID) - a nine-digit number, including leading zeroes.
	pid := p["pid"]
	if len(pid) == 9 {
		result = true
		for _, b := range pid {
			if !strings.Contains("0123456789", string(b)) {
				result = false
			}
		}
	}
	return result
}

func checkCid(p map[string]string) (result bool) {
	//cid (Country ID) - ignored, missing or not.
	return true
}
