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

	for scanner.Scan() {
		line := strings.ToLower(strings.Trim(scanner.Text(), " "))
		commands = append(commands, ParseInstruction(line))
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return commands
}
