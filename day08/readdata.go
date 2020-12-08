package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func readProgram(fname string) (commands []*Instruction) {

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
		line = strings.ToLower(strings.Trim(line, " "))
		//log.Printf("line: '%s'", line)
		command := ParseInstruction(line)
		//log.Printf("command: '%s'", command)
		commands = append(commands, command)
	}
	//groups = append(groups, group)
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return commands
}
