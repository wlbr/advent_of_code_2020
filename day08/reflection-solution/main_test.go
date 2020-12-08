package main

import (
	"testing"
)

type testdata struct {
	fname         string
	expectedtask1 int
	expectedtask2 int
}

var testset []*testdata = []*testdata{{"example.txt", 5, 8}}

func TestTaskOne(t *testing.T) {
	for _, test := range testset {
		m := NewVonNeumannMachine(readProgram(test.fname))
		c, _ := m.mainLoop()
		if c != test.expectedtask1 {
			t.Fatalf("Test '%s' failed. Got '%d' -  Wanted: '%d'", test.fname, c, test.expectedtask1)
		}
	}
}

func TestTaskTwo(t *testing.T) {
	for _, test := range testset {
		m := NewVonNeumannMachine(readProgram(test.fname))
		c, e := m.bruteForceMutator()
		if c != test.expectedtask2 || e != nil {
			t.Fatalf("Test '%s' failed. Got '%d' -  Wanted: '%d'", test.fname, c, test.expectedtask2)
		}
	}
}
