package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Machine struct {
	registers map[int]int
	Mask      string
}

func NewMachine() *Machine {
	m := &Machine{}
	m.registers = make(map[int]int)
	return m
}

const (
	bx = "x"
	b0 = "0"
	b1 = "1"
)

func encodingV1(in, m string) string {
	switch m {
	case b1:
		return "1"
	case b0:
		return "0"
	case bx:
		return in
	default:
		log.Printf("Unhandled encoding: '%v'", in)
		return ""
	}
}

func encodingV2(in, m string) string {
	switch m {
	case b1:
		return "1"
	case b0:
		return in
	case bx:
		return "x" //floating
	default:
		log.Printf("Unhandled encoding: '%v'", in)
		return ""
	}
}

func applyMaskV1(encodingf func(in, m string) string, value int, mask string) string {
	ns := ""
	format := fmt.Sprintf("%%0%db", len(mask))
	vs := fmt.Sprintf(format, value)
	for i := len(mask) - 1; i >= 0; i-- {
		m := encodingf(string(vs[i]), string(mask[i]))
		ns = fmt.Sprintf("%s%s", m, ns)
	}
	return ns
}

func (m *Machine) SetV1(address, value int) {
	nv := applyMaskV1(encodingV1, value, m.Mask)

	if v, err := strconv.ParseInt(nv, 2, len(m.Mask)+1); err == nil {
		//fmt.Printf("mask:   %s \n value: %s (%d)  result: %d\n", m.Mask, nv, v, value)
		m.registers[address] = int(v)
	} else {
		log.Printf("Error decoding binary number '%s' to decimal.   %s", nv, err)
	}

}

func genFloatingAdresses(initialadress string) []string {
	allAdresses := []string{initialadress}
	for index, char := range strings.ToLower(initialadress) {
		if string(char) == bx {
			var next []string
			for _, address := range allAdresses {
				upper := address[0:index] + b1 + address[index+1:]
				lower := address[0:index] + b0 + address[index+1:]
				next = append(next, upper, lower)
			}
			allAdresses = next
		}
	}
	return allAdresses
}

func (m *Machine) SetV2(address, value int) {
	na := applyMaskV1(encodingV2, address, m.Mask)
	ana := genFloatingAdresses(na)

	for _, a := range ana {
		if v, err := strconv.ParseInt(a, 2, len(m.Mask)+1); err == nil {
			//fmt.Printf("mask:     %s \n address: %s (%d)  result: %d\n", m.Mask, a, v, value)
			m.registers[int(v)] = value
		} else {
			log.Printf("Error decoding binary number '%s' to decimal.   %s", a, err)
		}
	}
}

func (m *Machine) Get(address int) (int, error) {
	if v, ok := m.registers[address]; ok {
		return v, nil
	}
	return 0, fmt.Errorf("address '%d'not found", address)
}

func (m *Machine) Sum() (s int) {
	for _, v := range m.registers {
		s += v
	}
	return s
}

func task1(program []Command) int {
	m := NewMachine()
	for _, c := range program {
		c.HandleV1(m)
	}
	return m.Sum()
}

func task2(program []Command) int {
	m := NewMachine()
	for _, c := range program {
		c.HandleV2(m)
	}
	return m.Sum()
}

func main() {
	var input string = "input.txt"
	prog := parsedata(readdata(input))

	n := task1(prog)
	fmt.Printf("Sum of register values is: %d\n", n)

	m := task2(prog)
	fmt.Printf("Sum of register values after floating algorithm is: %d\n", m)
}
