package main

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
)

var input string = "input.txt"

type Instruction struct {
	command   string
	parameter int
	visited   bool
}

func NewInstruction(instruction string, parameter string) *Instruction {
	n, e := strconv.Atoi(parameter)
	if e != nil {
		log.Printf("Error in code line, ignored. '%s' '%s' ", instruction, parameter)
		return nil
	}
	return &Instruction{instruction, n, false}
}

func ParseInstruction(line string) *Instruction {
	if len(line) < 6 {
		log.Printf("Misformed code line ignored. '%s'", line)
		return nil
	}
	return NewInstruction(line[:3], line[4:])
}

func (l *Instruction) String() string {
	sign := ""
	if l.parameter >= 0 {
		sign = "+"
	}
	return fmt.Sprintf("'%s' '%s%d'", l.command, sign, l.parameter)
}

type VonNeumannMachine struct {
	akkumulator int
	program     []*Instruction
	cursor      int
	// commands    map[string]Command
}

func NewVonNeumannMachine(program []*Instruction) *VonNeumannMachine {
	return &VonNeumannMachine{0, program, 0}
}

func (m *VonNeumannMachine) String() string {
	p := ""
	for _, l := range m.program {
		p = fmt.Sprintf("%s\t%s\n", p, l.String())
	}
	return fmt.Sprintf("Akkumulator: %d\nCursor:%d\nProgram:\n%s", m.akkumulator, m.cursor, p)
}

func (m *VonNeumannMachine) Reset() {
	for _, ins := range m.program {
		ins.visited = false
	}
	m.cursor = 0
	m.akkumulator = 0
}

func (m *VonNeumannMachine) checkUnFinished() bool {
	if m.cursor < len(m.program) {
		return true
	} else {
		return false
	}
}

func (m *VonNeumannMachine) checkInfiniteLoop() bool {
	if m.program[m.cursor].visited {
		return true
	} else {
		return false
	}
}

func (m *VonNeumannMachine) callCommand(lit string, param ...int) {
	ifa := reflect.ValueOf(m)
	method := ifa.MethodByName(lit)
	if method.IsValid() {
		var values []reflect.Value
		for _, p := range param {
			values = append(values, reflect.ValueOf(p))
		}
		method.Call(values)
	} else {
		log.Printf("Error: found unkown command: %s", lit)
	}
}

// func (m *VonNeumannMachine) findCommand1(lit string) (Command, error) {
// 	if c, ok := m.commands[lit]; ok {
// 		return c, nil
// 	} else {
// 		return nil, fmt.Errorf("Warning: unknown command, ignored. '%s'", lit)
// 	}
// }

func (m *VonNeumannMachine) mainLoop() (int, error) {
	var err error
	m.Reset()
	for m.checkUnFinished() {
		if m.checkInfiniteLoop() {
			err = fmt.Errorf("Infinite Loop, cursor at %d", m.cursor)
			break
		}
		m.callCommand(m.program[m.cursor].command, m.program[m.cursor].parameter)
		// if f, err := m.findCommand(m.program[m.cursor].command); err == nil {
		// 	f(m, m.program[m.cursor].parameter)
		// } else {
		// 	log.Printf("Warning: unknown command, ignored. '%s'", m.program[m.cursor])
		// 	m.cursor++
		// }
	}
	return m.akkumulator, err
}

func (m *VonNeumannMachine) mutate(linenumber int) {
	if m.program[linenumber].command == "Jmp" {
		m.program[linenumber].command = "Nop"
	} else if m.program[linenumber].command == "Nop" {
		m.program[linenumber].command = "Jmp"
	}

}

func (m *VonNeumannMachine) bruteForceMutator() (int, error) {
	for i, v := range m.program {
		backup := v.command
		m.mutate(i)
		r, e := m.mainLoop()
		if e == nil {
			//res = r
			return r, nil
		}
		m.program[i].command = backup
	}
	return 0, fmt.Errorf("Error: Did not find a fix this program.")
}

func main() {
	m := NewVonNeumannMachine(readProgram(input))

	if r, e := m.mainLoop(); e != nil {
		fmt.Printf("Found Infiniteloop, last value before was:   %5d\n", r)
	}
	if f, e := m.bruteForceMutator(); e == nil {
		fmt.Printf("Found the fix. Last value of akkumulator is: %5d\n", f)
	} else {
		fmt.Println(e)
	}

}
