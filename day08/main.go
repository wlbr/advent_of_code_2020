package main

import (
	"fmt"
	"log"
	"strconv"
)

var input string = "input.txt"

type line struct {
	instruction string
	parameter   int
	visited     bool
}

func NewLine(instruction string, parameter string) *line {
	n, e := strconv.Atoi(parameter)
	if e != nil {
		log.Printf("Error in code line, ignored. '%s' '%s' ", instruction, parameter)
		return nil
	}
	return &line{instruction, n, false}
}

func parseInstruction(line string) *line {
	if len(line) < 6 {
		log.Printf("Misformed code line ignored. '%s'", line)
		return nil
	}
	return NewLine(line[:3], line[4:])
}

func (l *line) String() string {
	sign := ""
	if l.parameter >= 0 {
		sign = "+"
	}
	return fmt.Sprintf("'%s' '%s%d'", l.instruction, sign, l.parameter)
}

type vnm struct {
	akkumulator int
	program     []*line
	cursor      int
}

func newVnm(program []*line) *vnm {
	return &vnm{0, program, 0}
}

func (m *vnm) String() string {
	p := ""
	for _, l := range m.program {
		p = fmt.Sprintf("%s\t%s\n", p, l.String())
	}
	return fmt.Sprintf("Akkumulator: %d\nCursor:%d\nProgram:\n%s", m.akkumulator, m.cursor, p)
}

func (m *vnm) reset() {
	for _, ins := range m.program {
		ins.visited = false
	}
	m.cursor = 0
	m.akkumulator = 0
}

func (m *vnm) acc(parameter int) {
	//log.Printf("acc: %d - %d", m.akkumulator, m.cursor)
	m.program[m.cursor].visited = true
	m.akkumulator += parameter
	m.cursor++
}

func (m *vnm) jmp(parameter int) {
	//log.Printf("jmp: %d - %d", m.akkumulator, m.cursor)
	m.program[m.cursor].visited = true
	m.cursor += parameter
}

func (m *vnm) nop(parameter int) {
	//log.Printf("nop: %d - %d", m.akkumulator, m.cursor)
	m.program[m.cursor].visited = true
	m.cursor++
}

func (m *vnm) checkFinished() bool {
	//log.Printf("checkFinished: %d - %d", m.akkumulator, m.cursor)
	if m.cursor < len(m.program) {
		return true
	} else {
		return false
	}
}

func (m *vnm) checkInfiniteLoop() bool {
	//log.Printf("checkInfiniteLoop: %d - %d", m.akkumulator, m.cursor)
	if m.program[m.cursor].visited {
		return true
	} else {
		return false
	}
}

func (m *vnm) mainLoop() (int, error) {
	//log.Printf("mainLoop: %d - %d", m.akkumulator, m.cursor)
	var err error
	m.reset()
	for m.checkFinished() {
		if m.checkInfiniteLoop() {
			err = fmt.Errorf("Infinite Loop, cursor at %d", m.cursor)
			break
		}
		if m.program[m.cursor].instruction == "acc" {
			m.acc(m.program[m.cursor].parameter)
		} else if m.program[m.cursor].instruction == "nop" {
			m.nop(m.program[m.cursor].parameter)
		} else if m.program[m.cursor].instruction == "jmp" {
			m.jmp(m.program[m.cursor].parameter)
		}
	}
	return m.akkumulator, err
}

func (m *vnm) mutate(linenumber int) {
	if m.program[linenumber].instruction == "jmp" {
		m.program[linenumber].instruction = "nop"
		//log.Printf("Changing line %d from %s to %s", linenumber, "jmp", m.program[linenumber])
	} else if m.program[linenumber].instruction == "nop" {
		m.program[linenumber].instruction = "jmp"
		//log.Printf("Changing line %d from %s to %s", linenumber, "nop", m.program[linenumber])
	}

}

func (m *vnm) bruteForceMutator() int {
	res := 0
	for i, _ := range m.program {
		backup := m.program[i].instruction
		//log.Printf("Before try: %s", m.program[i])
		m.mutate(i)
		//log.Printf("  After mutation: %s", m.program[i])
		r, e := m.mainLoop()
		if e == nil {
			res = r
			break
		}
		m.program[i].instruction = backup
		//log.Printf("  After try: %s\n", m.program[i])
	}
	return res
}

func main() {
	m := newVnm(readProgram(input))
	r, e := m.mainLoop()
	if e != nil {
		fmt.Printf("Found Infiniteloop, last value before was %d\n", r)
	}

	fmt.Printf("Found the fix. value before was %d\n", m.bruteForceMutator())
}
