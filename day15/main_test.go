package main

import (
	"testing"
)

type testdata struct {
	fname         string
	target1       int
	expectedtask1 int
	target2       int
	expectedtask2 int
}

var testset []*testdata = []*testdata{{"example1.txt", 2020, 436, 30000000, 175594},
	{"example2.txt", 2020, 1, 30000000, 2578},
	{"example3.txt", 2020, 10, 30000000, 3544142},
	{"example4.txt", 2020, 27, 30000000, 261214},
	{"example5.txt", 2020, 78, 30000000, 6895259},
	{"example6.txt", 2020, 438, 30000000, 18},
	{"example7.txt", 2020, 1836, 30000000, 362}}

func TestTaskOne(t *testing.T) {
	for _, test := range testset {
		m := readdata(test.fname)
		c := GetNthNumberFromElvesTalk(m, test.target1)
		if c != test.expectedtask1 {
			t.Fatalf("Test '%s' failed. Got '%d' -  Wanted: '%d'", test.fname, c, test.expectedtask1)
		}
	}
}

func TestTaskTwo(t *testing.T) {
	for _, test := range testset {
		m := readdata(test.fname)
		c := GetNthNumberFromElvesTalk(m, test.target2)
		if c != test.expectedtask2 {
			t.Fatalf("Test '%s' failed. Got '%d' -  Wanted: '%d'", test.fname, c, test.expectedtask1)
		}
	}
}
