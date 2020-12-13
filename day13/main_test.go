package main

import (
	"testing"
)

type testdata struct {
	fname         string
	expectedtask1 int
	expectedtask2 int
}

var testset []*testdata = []*testdata{{"example.txt", 295, 1068781}}

func TestTaskOne(t *testing.T) {
	for _, test := range testset {
		s, _, busses := readdata(test.fname)
		_, _, r := getNextBus(s, busses)
		if r != test.expectedtask1 {
			t.Fatalf("Test '%s' failed. Got '%d' -  Wanted: '%d'", test.fname, r, test.expectedtask1)
		}
	}
}
func TestTaskTwoST(t *testing.T) {
	for _, test := range testset {
		_, busses, _ := readdata(test.fname)
		r := findBusSequenceST(busses)
		if r != test.expectedtask2 {
			t.Fatalf("Test '%s' failed. Got '%d' -  Wanted: '%d'", test.fname, r, test.expectedtask2)
		}
	}
}

func TestTaskTwoMT(t *testing.T) {
	for _, test := range testset {
		_, busses, _ := readdata(test.fname)

		r := findBusSequenceMT(busses)

		if r != test.expectedtask2 {
			t.Fatalf("Test '%s' failed. Got '%d' -  Wanted: '%d'", test.fname, r, test.expectedtask2)
		}
	}
}
