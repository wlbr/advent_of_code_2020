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

var testset []*testdata = []*testdata{{"example1.txt", 71, "example2.txt", 1}}

func TestTaskOne(t *testing.T) {
	for _, test := range testset {
		input := ReadData(test.fname1)
		c := task1(input)
		if c != test.expectedtask1 {
			t.Fatalf("Test '%s' failed. Got '%d' -  Wanted: '%d'", test.fname1, c, test.expectedtask1)
		}
	}
}

func TestTaskTwo(t *testing.T) {
	for _, test := range testset {
		input := ReadData(test.fname1)
		c := task2(input)
		if c != test.expectedtask2 {
			t.Fatalf("Test '%s' failed. Got '%d' -  Wanted: '%d'", test.fname2, c, test.expectedtask2)
		}
	}
}
