package main

import (
	"testing"
)

type testdata struct {
	fname1        string
	expectedtask1 int
	fname2        string
	expectedtask2 int
}

var testset []*testdata = []*testdata{{"example1.txt", 165, "example2.txt", 208}}

func TestTaskOne(t *testing.T) {
	for _, test := range testset {
		m := readdata(test.fname1)
		p := parsedata(m)
		c := task1(p)
		if c != test.expectedtask1 {
			t.Fatalf("Test '%s' failed. Got '%d' -  Wanted: '%d'", test.fname1, c, test.expectedtask1)
		}
	}
}

func TestTaskTwo(t *testing.T) {
	for _, test := range testset {
		m := readdata(test.fname2)
		p := parsedata(m)
		c := task2(p)
		if c != test.expectedtask2 {
			t.Fatalf("Test '%s' failed. Got '%d' -  Wanted: '%d'", test.fname2, c, test.expectedtask2)
		}
	}
}
