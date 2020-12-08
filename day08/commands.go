package main

type Command func(m *VonNeumannMachine, parameter int)

func acc(m *VonNeumannMachine, parameter int) {
	//log.Printf("acc: %d - %d", m.akkumulator, m.cursor)
	m.program[m.cursor].visited = true
	m.akkumulator += parameter
	m.cursor++
}

func jmp(m *VonNeumannMachine, parameter int) {
	//log.Printf("jmp: %d - %d", m.akkumulator, m.cursor)
	m.program[m.cursor].visited = true
	m.cursor += parameter
}

func nop(m *VonNeumannMachine, parameter int) {
	//log.Printf("nop: %d - %d", m.akkumulator, m.cursor)
	m.program[m.cursor].visited = true
	m.cursor++
}
