package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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
		line := strings.Trim(scanner.Text(), " ")
		lines = append(lines, strings.ToUpper(line))
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return lines
}

func parsedata(lines []string) (cmds []*Command) {
	for _, l := range lines {
		matches := regx1.FindStringSubmatch(l)
		if len(matches) != 3 {
			log.Printf("Warning: Command not in standard format: '%s'", l)
		} else {
			p, err := strconv.Atoi(matches[2])
			if err != nil {
				log.Printf("Warning: Command parameter not a number. cmd:%s param:%v error:%v", l, matches[2], err)
			} else {
				c := NewCommand(matches[1], p)
				if (c.cmd == "L" || c.cmd == "R") && c.param%90 != 0 {
					log.Printf("Warning: degrees to turn not based on 90°. cmd:%s", l)
				}
				cmds = append(cmds, c)
			}
		}
	}
	return cmds
}
