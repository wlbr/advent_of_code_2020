package main

type Command func(m *VonNeumannMachine, parameter int)

func acc(m *VonNeumannMachine, parameter int) {
	m.program[m.cursor].visited = true
	m.akkumulator += parameter
	m.cursor++
}

func jmp(m *VonNeumannMachine, parameter int) {
	m.program[m.cursor].visited = true
	m.cursor += parameter
}

func nop(m *VonNeumannMachine, parameter int) {
	m.program[m.cursor].visited = true
	m.cursor++
}
