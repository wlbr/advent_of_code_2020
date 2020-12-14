package main

import (
	"fmt"
	"log"
	"strconv"
)

const (
	MEM  = "mem"
	MASK = "mask"
)

type Command interface {
	HandleV1(m *Machine)
	HandleV2(m *Machine)
}

type Mem struct {
	literal string
	address int
	arg     int
}

func parseMem(params []string) (cmd Command, err error) {
	p := params[1]
	var ar, ad int
	if ad, err = strconv.Atoi(p); err != nil {
		log.Printf("Warning: %s command has non integer address '%+v'", MEM, p)
		return nil, err
	}
	p = params[2]
	if ar, err = strconv.Atoi(p); err != nil {
		log.Printf("Warning: %s command has non integer argument '%+v'", MEM, p)
		return nil, err
	}
	return &Mem{literal: MEM, address: ad, arg: ar}, err
}

func (c *Mem) HandleV1(m *Machine) {
	m.SetV1(c.address, c.arg)
}

func (c *Mem) HandleV2(m *Machine) {
	m.SetV2(c.address, c.arg)
}

type Mask struct {
	literal string
	mask    string
}

func parseMask(params []string) (cmd Command, err error) {
	if len(params) != 3 {
		err := fmt.Errorf("error: %s command has non integer argument '%+v'", MEM, params)
		log.Print(err)
		return nil, err
	}
	return &Mask{literal: MASK, mask: params[2]}, nil
}

func (c *Mask) HandleV1(m *Machine) {
	m.Mask = c.mask
}

func (c *Mask) HandleV2(m *Machine) {
	c.HandleV1(m)
}
