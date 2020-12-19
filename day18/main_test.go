package main

import (
	"testing"
)

type testdata struct {
	fname         string
	expectedtask1 int
	expectedtask2 int
}

var testset []*testdata = []*testdata{{"example1.txt", 71, 231},
	{"example2.txt", 51, 51},
	{"example3.txt", 26, 46},
	{"example4.txt", 437, 1445},
	{"example5.txt", 12240, 669060},
	{"example6.txt", 13632, 23340}}

func TestTaskOne(t *testing.T) {
	for _, test := range testset {
		m := readdata(test.fname)
		c := t1(m)
		if c != test.expectedtask1 {
			t.Fatalf("Test '%s' failed. Got '%d' -  Wanted: '%d'", test.fname, c, test.expectedtask1)
		}
	}
}
func TestTaskTwo(t *testing.T) {
	for _, test := range testset[:len(testset)-1] {
		m := readdata(test.fname)
		c := t2(m)
		if c != test.expectedtask2 {
			t.Fatalf("Test '%s' failed. Got '%d' -  Wanted: '%d'", test.fname, c, test.expectedtask2)
		}
	}
}
