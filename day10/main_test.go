package main

import (
	"testing"
)

type testdata struct {
	fname         string
	expectedtask1 int
	expectedtask2 int
}

var testset []*testdata = []*testdata{{"example1.txt", 7 * 5, 8}, {"example2.txt", 22 * 10, 19208}}

func TestTaskOne(t *testing.T) {
	for _, test := range testset {
		m := readdata(test.fname)
		c := getProductOfCountOfDistances(m)
		if c != test.expectedtask1 {
			t.Fatalf("Test '%s' failed. Got '%d' -  Wanted: '%d'", test.fname, c, test.expectedtask1)
		}
	}
}

func TestTaskTwo(t *testing.T) {
	for _, test := range testset {
		m := readdata(test.fname)
		c := countCombinations(m)
		if c != test.expectedtask2 {
			t.Fatalf("Test '%s' failed. Got '%d' -  Wanted: '%d'", test.fname, c, test.expectedtask2)
		}
	}
}

type testcandidate struct {
	in  int
	out []int
}
