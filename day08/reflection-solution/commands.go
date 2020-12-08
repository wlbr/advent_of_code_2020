package main

//type Command func(m *VonNeumannMachine, parameter int)

func (m *VonNeumannMachine) Acc(parameter int) {
	//log.Printf("Command acc: %d - %d", m.akkumulator, m.cursor)
	m.program[m.cursor].visited = true
	m.akkumulator += parameter
	m.cursor++
}

func (m *VonNeumannMachine) Jmp(parameter int) {
	//log.Printf("Command jmp: %d - %d", m.akkumulator, m.cursor)
	m.program[m.cursor].visited = true
	m.cursor += parameter
}

func (m *VonNeumannMachine) Nop(parameter int) {
	//log.Printf("Command nop: %d - %d", m.akkumulator, m.cursor)
	m.program[m.cursor].visited = true
	m.cursor++
}
